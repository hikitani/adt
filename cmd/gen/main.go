package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	. "github.com/dave/jennifer/jen"
)

func main() {
	out := flag.String("out", "", "Set file path when code will be generated")
	numTypes := flag.Uint("num", 2, "Num of types")
	flag.Parse()

	log.SetFlags(0)

	if *out == "" {
		log.Fatal("invalid -out: expected path, got nothing")
	}

	if *numTypes < 2 {
		log.Fatal("invalid -num: minimum number of types is 2")
	}

	file, err := os.Create(*out)
	if err != nil {
		log.Fatalf("Cannot create file %s: %s", *out, err)
	}
	defer file.Close()

	f := NewFile("adt")
	f.PackageComment("Generated automatically. DO NOT EDIT.\n")

	f.Add()

	for _, code := range renderOneOf(int(*numTypes)) {
		f.Add(code)
	}

	for _, code := range renderFuncs(int(*numTypes)) {
		f.Add(code)
	}

	if err := f.Render(file); err != nil {
		log.Fatalf("Failed to code generate: %s", err)
	}
}

func genTypes(numTypes int) []Code {
	genTypes := []Code{}
	for i := 1; i <= numTypes; i++ {
		genTypes = append(genTypes, genTypId(i).Any())
	}
	return genTypes
}

func sumTypeNames(numTypes int) []Code {
	sumTypeNames := []Code{}
	for i := 1; i <= numTypes; i++ {
		sumTypeNames = append(sumTypeNames, genTypId(i))
	}
	return sumTypeNames
}

func renderOneOf(numTypes int) (res []Code) {
	genTypes := genTypes(numTypes)

	unionTypes := []Code{}
	for i := 1; i <= numTypes; i++ {
		unionTypes = append(unionTypes, Id("f").Types(genTypId(i)))
	}

	name := func(s string) string {
		return fmt.Sprintf("%s%d", s, numTypes)
	}

	unionType := Type().
		Id(name("union")).
		Types(genTypes...).
		Interface(Union(unionTypes...))

	sumType := Type().
		Id(name("sumType")).
		Types(genTypes...).
		Interface()

	sumTypeNames := sumTypeNames(numTypes)

	sumTypeFunc := Func().
		Id(name("SumType")).
		Types(genTypes...).
		Params().
		Id(name("sumType")).Types(sumTypeNames...).
		Block(Return(Nil()))

	eitherFields := []Code{}
	for i := 1; i <= numTypes; i++ {
		f := valId(i).Op("*").Add(genTypId(i))
		eitherFields = append(eitherFields, f)
	}

	eitherType := Type().
		Id(name("OneOf")).
		Types(genTypes...).
		Struct(eitherFields...)

	res = append(res,
		unionType,
		sumType,
		sumTypeFunc,
		eitherType,
	)

	for i := 1; i <= numTypes; i++ {
		f := Func().
			Params(Id("o").Op("*").Id(name("OneOf")).Types(sumTypeNames...)).
			Add(genTypId(i)).
			Params().
			Params(
				valId(i).Add(genTypId(i)),
				Id("ok").Bool(),
			).
			Block(
				If(
					Id("o").Dot(valId(i).GoString()).Op("==").Nil(),
				).Block(
					Return(),
				),
				Return(
					Op("*").Id("o").Dot(valId(i).GoString()),
					True(),
				),
			)
		res = append(res, f)
	}

	for i := 1; i <= numTypes; i++ {
		f := Func().
			Params(Id("o").Op("*").Id(name("OneOf")).Types(sumTypeNames...)).
			Id(fmt.Sprintf("SetT%d", i)).
			Params(Id("v").Add(genTypId(i))).
			Block(
				Var().Id("zeroVal").Id(name("OneOf")).Types(sumTypeNames...),
				Op("*").Id("o").Op("=").Id("zeroVal"),
				Id("o").Dot(valId(i).GoString()).Op("=").New(genTypId(i)),
				Op("*").Id("o").Dot(valId(i).GoString()).Op("=").Id("v"),
			)
		res = append(res, f)
	}

	return
}

func valId(id int) *Statement {
	return Id(fmt.Sprintf("v%d", id))
}

func genTypId(id int) *Statement {
	return Id(fmt.Sprintf("T%d", id))
}

func renderFuncs(numTypes int) []Code {
	name := func(s string) string {
		return fmt.Sprintf("%s%d", s, numTypes)
	}

	genTypes := genTypes(numTypes)
	sumTypeNames := sumTypeNames(numTypes)

	caseArgs := []Code{
		Id("v").Add(Id(name("OneOf")).Types(sumTypeNames...)),
	}
	for i := 1; i <= numTypes; i++ {
		caseArg := Id(fmt.Sprintf("case%d", i)).Func().Params(Id("v").Add(genTypId(i)))
		caseArgs = append(caseArgs, caseArg)
	}

	caseCheckers := []Code{}
	for i := 1; i <= numTypes; i++ {
		caseChecker := If(
			List(Id("v"), Id("ok")).Op(":=").Id("v").Dot(fmt.Sprintf("T%d", i)).Call(),
			Id("ok"),
		).Block(
			Id(fmt.Sprintf("case%d", i)).Call(Id("v")),
			Return(),
		)
		caseCheckers = append(caseCheckers, caseChecker)
	}
	casesFunc := Func().
		Id(name("Cases")).
		Types(genTypes...).
		Params(merge(caseArgs, []Code{
			Id("caseNull").Func().Params(),
		})...).
		Block(merge(caseCheckers, []Code{
			Id("caseNull").Call(),
		})...)

	createGenTypes := make([]Code, len(genTypes))
	copy(createGenTypes, genTypes)
	createGenTypes = append(createGenTypes,
		Id("U").Id(name("union")).Types(sumTypeNames...))
	createFunc := Func().
		Id(name("NewOneOf")).
		Types(createGenTypes...).
		Params(
			Id("typ").Id(name("sumType")).Types(sumTypeNames...),
			Id("v").Id("U"),
		).
		Params(Id(name("OneOf")).Types(sumTypeNames...)).
		Block(
			Var().Id("res").Id(name("OneOf")).Types(sumTypeNames...),
			Id(name("assign")).Call(Op("&").Id("res"), Id("v")),
			Return(Id("res")),
		)

	valDefs := []Code{}
	typDefs := []Code{}
	typCases := []Code{}
	for i := 1; i <= numTypes; i++ {
		vTypVal := fmt.Sprintf("v%dTyp", i)

		valDefs = append(valDefs, valId(i).Add(genTypId(i)))
		typDefs = append(typDefs,
			Id(vTypVal).Op(":=").Qual("reflect", "TypeOf").Call(valId(i)))
		typCases = append(typCases,
			Case(Id(vTypVal)).Block(
				Id("val").Dot(fmt.Sprintf("SetT%d", i)).Call(
					Id("outVal").Dot("Interface").Call().Dot("").Call(genTypId(i)),
				),
			),
		)
	}
	typCases = append(typCases, Default().Block(Panic(Lit("unreachable code"))))
	assignFunc := Func().
		Id(name("assign")).
		Types(createGenTypes...).
		Params(
			Id("val").Op("*").Id(name("OneOf")).Types(sumTypeNames...),
			Id("v").Id("U"),
		).
		Block(merge(
			[]Code{
				Id("funcVal").Op(":=").Qual("reflect", "ValueOf").Call(Id("v")),
				Id("outVal").Op(":=").Id("funcVal").Dot("Call").Call(Nil()).Index(Id("0")),
				Var().Defs(valDefs...),
			},
			typDefs,
			[]Code{
				Switch(Id("outVal").Dot("Type").Call()).Block(typCases...),
			},
		)...)

	return []Code{
		casesFunc,
		createFunc,
		assignFunc,
	}
}

func merge[T any](slices ...[]T) []T {
	l := 0
	for _, s := range slices {
		l += len(s)
	}
	res := make([]T, l)
	resSlice := res
	for _, s := range slices {
		copy(resSlice, s)
		resSlice = resSlice[len(s):]
	}
	return res
}
