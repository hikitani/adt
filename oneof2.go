/*
Generated automatically. DO NOT EDIT.
*/
package adt

import "reflect"

type union2[T1 any, T2 any] interface {
	f[T1] | f[T2]
}
type sumType2[T1 any, T2 any] interface{}

func SumType2[T1 any, T2 any]() sumType2[T1, T2] {
	return nil
}

type OneOf2[T1 any, T2 any] struct {
	v1 *T1
	v2 *T2
}

func (o *OneOf2[T1, T2]) T1() (v1 T1, ok bool) {
	if o.v1 == nil {
		return
	}
	return *o.v1, true
}
func (o *OneOf2[T1, T2]) T2() (v2 T2, ok bool) {
	if o.v2 == nil {
		return
	}
	return *o.v2, true
}
func (o *OneOf2[T1, T2]) SetT1(v T1) {
	var zeroVal OneOf2[T1, T2]
	*o = zeroVal
	o.v1 = new(T1)
	*o.v1 = v
}
func (o *OneOf2[T1, T2]) SetT2(v T2) {
	var zeroVal OneOf2[T1, T2]
	*o = zeroVal
	o.v2 = new(T2)
	*o.v2 = v
}
func Cases2[T1 any, T2 any](v OneOf2[T1, T2], case1 func(v T1), case2 func(v T2), caseNull func()) {
	if v, ok := v.T1(); ok {
		case1(v)
		return
	}
	if v, ok := v.T2(); ok {
		case2(v)
		return
	}
	caseNull()
}
func NewOneOf2[T1 any, T2 any, U union2[T1, T2]](typ sumType2[T1, T2], v U) OneOf2[T1, T2] {
	var res OneOf2[T1, T2]
	assign2(&res, v)
	return res
}
func assign2[T1 any, T2 any, U union2[T1, T2]](val *OneOf2[T1, T2], v U) {
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
