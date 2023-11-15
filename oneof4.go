/*
Generated automatically. DO NOT EDIT.
*/
package adt

import "reflect"

type union4[T1 any, T2 any, T3 any, T4 any] interface {
	f[T1] | f[T2] | f[T3] | f[T4]
}
type sumType4[T1 any, T2 any, T3 any, T4 any] interface{}

func SumType4[T1 any, T2 any, T3 any, T4 any]() sumType4[T1, T2, T3, T4] {
	return nil
}

type OneOf4[T1 any, T2 any, T3 any, T4 any] struct {
	v1 *T1
	v2 *T2
	v3 *T3
	v4 *T4
}

func (o *OneOf4[T1, T2, T3, T4]) T1() (v1 T1, ok bool) {
	if o.v1 == nil {
		return
	}
	return *o.v1, true
}
func (o *OneOf4[T1, T2, T3, T4]) T2() (v2 T2, ok bool) {
	if o.v2 == nil {
		return
	}
	return *o.v2, true
}
func (o *OneOf4[T1, T2, T3, T4]) T3() (v3 T3, ok bool) {
	if o.v3 == nil {
		return
	}
	return *o.v3, true
}
func (o *OneOf4[T1, T2, T3, T4]) T4() (v4 T4, ok bool) {
	if o.v4 == nil {
		return
	}
	return *o.v4, true
}
func (o *OneOf4[T1, T2, T3, T4]) SetT1(v T1) {
	var zeroVal OneOf4[T1, T2, T3, T4]
	*o = zeroVal
	o.v1 = new(T1)
	*o.v1 = v
}
func (o *OneOf4[T1, T2, T3, T4]) SetT2(v T2) {
	var zeroVal OneOf4[T1, T2, T3, T4]
	*o = zeroVal
	o.v2 = new(T2)
	*o.v2 = v
}
func (o *OneOf4[T1, T2, T3, T4]) SetT3(v T3) {
	var zeroVal OneOf4[T1, T2, T3, T4]
	*o = zeroVal
	o.v3 = new(T3)
	*o.v3 = v
}
func (o *OneOf4[T1, T2, T3, T4]) SetT4(v T4) {
	var zeroVal OneOf4[T1, T2, T3, T4]
	*o = zeroVal
	o.v4 = new(T4)
	*o.v4 = v
}
func Cases4[T1 any, T2 any, T3 any, T4 any](v OneOf4[T1, T2, T3, T4], case1 func(v T1), case2 func(v T2), case3 func(v T3), case4 func(v T4), caseNull func()) {
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
	if v, ok := v.T4(); ok {
		case4(v)
		return
	}
	caseNull()
}
func NewOneOf4[T1 any, T2 any, T3 any, T4 any, U union4[T1, T2, T3, T4]](typ sumType4[T1, T2, T3, T4], v U) OneOf4[T1, T2, T3, T4] {
	var res OneOf4[T1, T2, T3, T4]
	assign4(&res, v)
	return res
}
func assign4[T1 any, T2 any, T3 any, T4 any, U union4[T1, T2, T3, T4]](val *OneOf4[T1, T2, T3, T4], v U) {
	funcVal := reflect.ValueOf(v)
	outVal := funcVal.Call(nil)[0]
	var (
		v1 T1
		v2 T2
		v3 T3
		v4 T4
	)
	v1Typ := reflect.TypeOf(v1)
	v2Typ := reflect.TypeOf(v2)
	v3Typ := reflect.TypeOf(v3)
	v4Typ := reflect.TypeOf(v4)
	switch outVal.Type() {
	case v1Typ:
		val.SetT1(outVal.Interface().(T1))
	case v2Typ:
		val.SetT2(outVal.Interface().(T2))
	case v3Typ:
		val.SetT3(outVal.Interface().(T3))
	case v4Typ:
		val.SetT4(outVal.Interface().(T4))
	default:
		panic("unreachable code")
	}
}
