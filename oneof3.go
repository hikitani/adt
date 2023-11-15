/*
Generated automatically. DO NOT EDIT.
*/
package adt

import "reflect"

type union3[T1 any, T2 any, T3 any] interface {
	f[T1] | f[T2] | f[T3]
}
type sumType3[T1 any, T2 any, T3 any] interface{}

func SumType3[T1 any, T2 any, T3 any]() sumType3[T1, T2, T3] {
	return nil
}

type OneOf3[T1 any, T2 any, T3 any] struct {
	v1 *T1
	v2 *T2
	v3 *T3
}

func (o *OneOf3[T1, T2, T3]) T1() (v1 T1, ok bool) {
	if o.v1 == nil {
		return
	}
	return *o.v1, true
}
func (o *OneOf3[T1, T2, T3]) T2() (v2 T2, ok bool) {
	if o.v2 == nil {
		return
	}
	return *o.v2, true
}
func (o *OneOf3[T1, T2, T3]) T3() (v3 T3, ok bool) {
	if o.v3 == nil {
		return
	}
	return *o.v3, true
}
func (o *OneOf3[T1, T2, T3]) SetT1(v T1) {
	var zeroVal OneOf3[T1, T2, T3]
	*o = zeroVal
	o.v1 = new(T1)
	*o.v1 = v
}
func (o *OneOf3[T1, T2, T3]) SetT2(v T2) {
	var zeroVal OneOf3[T1, T2, T3]
	*o = zeroVal
	o.v2 = new(T2)
	*o.v2 = v
}
func (o *OneOf3[T1, T2, T3]) SetT3(v T3) {
	var zeroVal OneOf3[T1, T2, T3]
	*o = zeroVal
	o.v3 = new(T3)
	*o.v3 = v
}
func Cases3[T1 any, T2 any, T3 any](v OneOf3[T1, T2, T3], case1 func(v T1), case2 func(v T2), case3 func(v T3), caseNull func()) {
	if v, ok := v.T1(); ok {
		case1(v)
		return
	}
	if v, ok := v.T2(); ok {
		case2(v)
		return
	}
	if v, ok := v.T3(); ok {
		case3(v)
		return
	}
	caseNull()
}
func NewOneOf3[T1 any, T2 any, T3 any, U union3[T1, T2, T3]](typ sumType3[T1, T2, T3], v U) OneOf3[T1, T2, T3] {
	var res OneOf3[T1, T2, T3]
	assign3(&res, v)
	return res
}
func assign3[T1 any, T2 any, T3 any, U union3[T1, T2, T3]](val *OneOf3[T1, T2, T3], v U) {
	funcVal := reflect.ValueOf(v)
	outVal := funcVal.Call(nil)[0]
	var (
		v1 T1
		v2 T2
		v3 T3
	)
	v1Typ := reflect.TypeOf(v1)
	v2Typ := reflect.TypeOf(v2)
	v3Typ := reflect.TypeOf(v3)
	switch outVal.Type() {
	case v1Typ:
		val.SetT1(outVal.Interface().(T1))
	case v2Typ:
		val.SetT2(outVal.Interface().(T2))
	case v3Typ:
		val.SetT3(outVal.Interface().(T3))
	default:
		panic("unreachable code")
	}
}
