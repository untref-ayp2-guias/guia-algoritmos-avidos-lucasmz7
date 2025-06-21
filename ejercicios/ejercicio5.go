package ejercicios

var billetes = []int{10000, 2000, 1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}

func Cambiar(cantidad int) map[int]int {
	cambio := make(map[int]int)
	return helper(cantidad, billetes, cambio)
}

func helper(cantidad int, monedas []int, cambio map[int]int) map[int]int {
	if cantidad == 0 {
		return cambio
	}
	for _, moneda := range monedas {
		if moneda <= cantidad {
			cambio[moneda]++
			return helper(cantidad-moneda, monedas, cambio)
		}
	}
	return cambio
}
