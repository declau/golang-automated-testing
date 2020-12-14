package enderecos

import "testing"

type cenarioDeteste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	//Nome da Função começa obrigatóriamente com a palavra Test

	cenariosDeTeste := []cenarioDeteste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo inválido"},
		{"Estrada Trem Baum", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA JOÃO PINHEIRO", "Avenida"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("Tipo de endereço recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

	/*
		//Usado para função básica
		if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
			t.Errorf("Tipo diferente do esperado! Esperava o tipo %s e recebeu %s",
				tipoDeEnderecoEsperado,
				tipoDeEnderecoRecebido,
			)
		} */

}
