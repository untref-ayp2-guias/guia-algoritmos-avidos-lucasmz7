package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFarolas(t *testing.T) {
	entrada := []Farolas{
		{1, 3, false},
		{4, 5, false},
		{7, 7, false},
		{11, 6, false},
		{15, 8, false},
		{26, 9, false},
		{37, 10, false},
	}
	x := []int{1, 4, 13, 16, 21, 25, 37, 45}
	//Debe encender las farolas de las posiciones 7, 15, 26 y 37
	salida, err := EncenderFarolas(entrada, x)
	assert.NoError(t, err, "EncenderFarolas no debería retornar error")
	assert.Len(t, salida, 4, "Error al encender farolas")
}
