package main

import "testing"

func TestGreeting(t *testing.T) {
	result := greeting("Hello World")

	if len(result) == 0 {
		t.Error("O retorno da mensagem falhou")
	}
}
