package main

import "fmt"

func main() {
	puntajes := []int{4, 5, 5, 5, 5, 4, 1, 3, 5, 4}

	suma := 0
	min, max := 5, 1
	conteo := [5]int{}

	for _, p := range puntajes {
		if p < 1 || p > 5 {
			fmt.Printf("Puntaje inválido: %d\n", p)
			continue
		}

		suma += p
		if p < min {
			min = p
		}
		if p > max {
			max = p
		}
		conteo[p-1]++
	}

	fmt.Printf(`
Resultados de la encuesta:
Número total de respuestas: %d
Promedio de satisfacción: %.2f
Puntaje más alto: %d
Puntaje más bajo: %d

Distribución de puntajes:
1 punto: %d clientes
2 puntos: %d clientes
3 puntos: %d clientes
4 puntos: %d clientes
5 puntos: %d clientesS
`, len(puntajes), float64(suma)/float64(len(puntajes)), max, min,
		conteo[0], conteo[1], conteo[2], conteo[3], conteo[4])
}
