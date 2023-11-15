/*
Generated automatically. DO NOT EDIT.
*/
package adt

import "reflect"

type union7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any] interface {
	f[T1] | f[T2] | f[T3] | f[T4] | f[T5] | f[T6] | f[T7]
}
type sumType7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any] interface{}

func SumType7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any]() sumType7[T1, T2, T3, T4, T5, T6, T7] {
	return nil
}

type OneOf7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any] struct {
	v1 *T1
	v2 *T2
	v3 *T3
	v4 *T4
	v5 *T5
	v6 *T6
	v7 *T7
}

func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) T1() (v1 T1, ok bool) {
	if o.v1 == nil {
		return
	}
	return *o.v1, true
}
func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) T2() (v2 T2, ok bool) {
	if o.v2 == nil {
		return
	}
	return *o.v2, true
}
func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) T3() (v3 T3, ok bool) {
	if o.v3 == nil {
		return
	}
	return *o.v3, true
}
func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) T4() (v4 T4, ok bool) {
	if o.v4 == nil {
		return
	}
	return *o.v4, true
}
func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) T5() (v5 T5, ok bool) {
	if o.v5 == nil {
		return
	}
	return *o.v5, true
}
func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) T6() (v6 T6, ok bool) {
	if o.v6 == nil {
		return
	}
	return *o.v6, true
}
func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) T7() (v7 T7, ok bool) {
	if o.v7 == nil {
		return
	}
	return *o.v7, true
}
func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) SetT1(v T1) {
	var zeroVal OneOf7[T1, T2, T3, T4, T5, T6, T7]
	*o = zeroVal
	o.v1 = new(T1)
	*o.v1 = v
}
func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) SetT2(v T2) {
	var zeroVal OneOf7[T1, T2, T3, T4, T5, T6, T7]
	*o = zeroVal
	o.v2 = new(T2)
	*o.v2 = v
}
func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) SetT3(v T3) {
	var zeroVal OneOf7[T1, T2, T3, T4, T5, T6, T7]
	*o = zeroVal
	o.v3 = new(T3)
	*o.v3 = v
}
func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) SetT4(v T4) {
	var zeroVal OneOf7[T1, T2, T3, T4, T5, T6, T7]
	*o = zeroVal
	o.v4 = new(T4)
	*o.v4 = v
}
func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) SetT5(v T5) {
	var zeroVal OneOf7[T1, T2, T3, T4, T5, T6, T7]
	*o = zeroVal
	o.v5 = new(T5)
	*o.v5 = v
}
func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) SetT6(v T6) {
	var zeroVal OneOf7[T1, T2, T3, T4, T5, T6, T7]
	*o = zeroVal
	o.v6 = new(T6)
	*o.v6 = v
}
func (o *OneOf7[T1, T2, T3, T4, T5, T6, T7]) SetT7(v T7) {
	var zeroVal OneOf7[T1, T2, T3, T4, T5, T6, T7]
	*o = zeroVal
	o.v7 = new(T7)
	*o.v7 = v
}
func Cases7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](v OneOf7[T1, T2, T3, T4, T5, T6, T7], case1 func(v T1), case2 func(v T2), case3 func(v T3), case4 func(v T4), case5 func(v T5), case6 func(v T6), case7 func(v T7), caseNull func()) {
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
	if v, ok := v.T5(); ok {
		case5(v)
		return
	}
	if v, ok := v.T6(); ok {
		case6(v)
		return
	}
	if v, ok := v.T7(); ok {
		case7(v)
		return
	}
	caseNull()
}
func NewOneOf7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, U union7[T1, T2, T3, T4, T5, T6, T7]](typ sumType7[T1, T2, T3, T4, T5, T6, T7], v U) OneOf7[T1, T2, T3, T4, T5, T6, T7] {
	var res OneOf7[T1, T2, T3, T4, T5, T6, T7]
	assign7(&res, v)
	return res
}
func assign7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, U union7[T1, T2, T3, T4, T5, T6, T7]](val *OneOf7[T1, T2, T3, T4, T5, T6, T7], v U) {
	funcVal := reflect.ValueOf(v)
	outVal := funcVal.Call(nil)[0]
	var (
		v1 T1
		v2 T2
		v3 T3
		v4 T4
		v5 T5
		v6 T6
		v7 T7
	)
	v1Typ := reflect.TypeOf(v1)
	v2Typ := reflect.TypeOf(v2)
	v3Typ := reflect.TypeOf(v3)
	v4Typ := reflect.TypeOf(v4)
	v5Typ := reflect.TypeOf(v5)
	v6Typ := reflect.TypeOf(v6)
	v7Typ := reflect.TypeOf(v7)
	switch outVal.Type() {
	case v1Typ:
		val.SetT1(outVal.Interface().(T1))
	case v2Typ:
		val.SetT2(outVal.Interface().(T2))
	case v3Typ:
		val.SetT3(outVal.Interface().(T3))
	case v4Typ:
		val.SetT4(outVal.Interface().(T4))
	case v5Typ:
		val.SetT5(outVal.Interface().(T5))
	case v6Typ:
		val.SetT6(outVal.Interface().(T6))
	case v7Typ:
		val.SetT7(outVal.Interface().(T7))
	default:
		panic("unreachable code")
	}
}
