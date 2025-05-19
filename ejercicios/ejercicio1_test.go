package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectorActividades(t *testing.T) {
	actividades := []Actividad{
		{"a1", 1, 4},
		{"a2", 3, 5},
		{"a3", 0, 6},
		{"a4", 5, 7},
		{"a5", 3, 9},
		{"a6", 5, 9},
		{"a7", 6, 10},
		{"a8", 8, 11},
		{"a9", 8, 12},
		{"a10", 2, 14},
		{"a11", 12, 15},
	}
	seleccionadas := SelectorRecursivo(actividades)
	assert.Len(t, seleccionadas, 4, "Se esperaban 4 actividades seleccionadas")

	assert.Contains(t, seleccionadas, Actividad{"a1", 1, 4}, "Se esperaba la actividad a1")
	assert.Contains(t, seleccionadas, Actividad{"a4", 5, 7}, "Se esperaba la actividad a4")
	assert.Contains(t, seleccionadas, Actividad{"a8", 8, 11}, "Se esperaba la actividad a8")
	assert.Contains(t, seleccionadas, Actividad{"a11", 12, 15}, "Se esperaba la actividad a11")
}
