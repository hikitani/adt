package adt

type f[T any] func() T

func V[T any](v T) f[T] {
	return func() T {
		return v
	}
}

//go:generate go run github.com/hikitani/adt/cmd/gen -out oneof2.go -num 2
//go:generate go run github.com/hikitani/adt/cmd/gen -out oneof3.go -num 3
//go:generate go run github.com/hikitani/adt/cmd/gen -out oneof4.go -num 4
//go:generate go run github.com/hikitani/adt/cmd/gen -out oneof5.go -num 5
//go:generate go run github.com/hikitani/adt/cmd/gen -out oneof6.go -num 6
//go:generate go run github.com/hikitani/adt/cmd/gen -out oneof7.go -num 7
//go:generate go run github.com/hikitani/adt/cmd/gen -out oneof8.go -num 8
//go:generate go run github.com/hikitani/adt/cmd/gen -out oneof9.go -num 9
