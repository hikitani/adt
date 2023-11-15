package adt

type f[T any] func() T

func V[T any](v T) f[T] {
	return func() T {
		return v
	}
}

type union[T1, T2 any] interface {
	f[T1] | f[T2]
}

type sumType[T1, T2 any] interface{}

func SumType[T1, T2 any]() sumType[T1, T2] {
	return nil
}

type OneOf[T1, T2 any] struct {
	v1 *T1
	v2 *T2
}

func (o *OneOf[T1, T2]) T1() (v1 T1, ok bool) {
	if o.v1 == nil {
		return
	}

	return *o.v1, true
}

func (o *OneOf[T1, T2]) T2() (v2 T2, ok bool) {
	if o.v2 == nil {
		return
	}

	return *o.v2, true
}

func (o *OneOf[T1, T2]) SetT1(v T1) {
	var zeroVal OneOf[T1, T2]
	*o = zeroVal
	o.v1 = new(T1)
	*o.v1 = v
}

func (o *OneOf[T1, T2]) SetT2(v T2) {
	var zeroVal OneOf[T1, T2]
	*o = zeroVal
	o.v2 = new(T2)
	*o.v2 = v
}
