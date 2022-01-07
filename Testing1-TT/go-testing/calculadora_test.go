package go_testing

import "testing"

func TestRestart(t *testing.T) {
	num1 := 10
	num2 := 4
	desiredResult := 6

	result := Restar(num1, num2)

	if result != desiredResult {
		t.Errorf("Funcion Restar() arrojo el resultado %v pero se esperaba %v", result, desiredResult)
	}
}
