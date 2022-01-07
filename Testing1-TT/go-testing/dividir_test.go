package go_testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	num := 10
	den := 2
	desiredResult := 5
	result, err := Dividir(num, den)

	if err != nil {
		assert.Equal(t, 0, result)
		assert.Equal(t, fmt.Errorf("El denominador no puede ser 0"), err)
	} else {
		assert.Equal(t, desiredResult, result)
		assert.Equal(t, nil, err)
	}

}
