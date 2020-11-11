package main

import "testing"

func TestSquareRoot(t *testing.T) {
	result := squareRoot(25)

	if result <= 0 {
		t.Error("Retorno inesperado. Deveria ser maior que 0")
	}
}
