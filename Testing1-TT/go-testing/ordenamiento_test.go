package go_testing

import "testing"

func TestOrdenar(t *testing.T) {
	slice1 := []int{5, 4, 3, 2, 1}
	desiredOrder := []int{1, 2, 3, 4, 5}

	result := Ordenar(slice1)

	for i, _ := range desiredOrder {
		if result[i] != desiredOrder[i] {
			t.Errorf("La funcion Ordenar() retorno %v y se esperaba %v", result, desiredOrder)
			return
		}
	}
}
