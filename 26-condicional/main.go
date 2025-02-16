package main

func main() {
	a := 1
	b := 2
	c := 3

	if a > b && c > a {
		println("entrou no if")
	}

	if a > b || c > a {
		println("entrou no if")
	}

	if a > b {
		println(a)
	} else {
		println(b)
	}

	switch a {
	case 1:
		println("entrou no case 1")
	case 2:
		println("entrou no case 2")
	case 3:
		println("entrou no case 3")
	default:
		println("entrou no case default")
	}
}
