// Nome do pacote
package main

// Importa o pacote testing
import "testing"

// Funçaõ de teste
func TestReturn(t *testing.T) {

	resultado := greeting("Code.education Rocks!!")

	if resultado != "<b>Code.education Rocks!!/b>" {
		t.Errorf("Erro")
	}

	
}