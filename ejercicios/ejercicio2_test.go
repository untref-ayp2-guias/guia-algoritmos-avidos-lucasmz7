package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEjercicio2(t *testing.T) {
	tareas := []Tarea{
		{"t1", 34},
		{"t2", 13},
		{"t3", 16},
		{"t4", 7},
		{"t5", 9},
		{"t6", 19},
		{"t7", 10},
		{"t8", 21},
		{"t9", 42},
		{"t10", 14},
		{"t11", 25},
	}
	Ejercicio2(tareas)
	for i := 0; i < len(tareas)-1; i++ {
		assert.LessOrEqualf(t, tareas[i].Tiempo, tareas[i+1].Tiempo, "Las tareas no están ordenadas en la posición %d", i)
	}
}
