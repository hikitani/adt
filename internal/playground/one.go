package adt

import "reflect"

func Cases[T1, T2 any](v OneOf[T1, T2], case1 func(v T1), case2 func(v T2), caseNull func()) {
	if v, ok := v.T1(); ok {
		case1(v)
	}

	if v, ok := v.T2(); ok {
		case2(v)
	}

	caseNull()
}

func Create[T1, T2 any, U union[T1, T2]](typ sumType[T1, T2], v U) OneOf[T1, T2] {
	var res OneOf[T1, T2]
	assign(&res, v)
	return res
}

func assign[T1, T2 any, U union[T1, T2]](val *OneOf[T1, T2], v U) {
	funcVal := reflect.ValueOf(v)
	outVal := funcVal.Call(nil)[0]
	var (
		v1 T1
		v2 T2
	)
	v1Typ := reflect.TypeOf(v1)
	v2Typ := reflect.TypeOf(v2)

	switch outVal.Type() {
	case v1Typ:
		val.SetT1(outVal.Interface().(T1))
	case v2Typ:
		val.SetT2(outVal.Interface().(T2))
	default:
		panic("unreachable code")
	}
}

func Is[T1, T2 any, U union[T1, T2]](v OneOf[T1, T2], typ U) bool {
	targetTyp := reflect.TypeOf(typ).Out(0)
	switch {
	case v.v1 != nil:
		return reflect.TypeOf(v.v1).Elem() == targetTyp
	case v.v2 != nil:
		return reflect.TypeOf(v.v2).Elem() == targetTyp
	default:
		return false
	}
}
