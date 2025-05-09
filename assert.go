package assert

func Error(err error) {
	if err != nil {
		panic(err)
	}
}

func Error1[T any](a T, err error) T {
	if err != nil {
		panic(err)
	}

	return a
}

func Error2[T1 any, T2 any](a1 T1, a2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}

	return a1, a2
}

func Error3[T1 any, T2 any, T3 any](a1 T1, a2 T2, a3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}

	return a1, a2, a3
}
