package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T Number](a T, b T) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func main() {
	m := map[string]int{"Wesley": 100, "João": 2000, "Maria": 3000}
	m2 := map[string]MyNumber{"Wesley": 100, "João": 2000, "Maria": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Compara(10, 10.0))
}
