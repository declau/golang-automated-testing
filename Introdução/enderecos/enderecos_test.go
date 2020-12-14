package enderecos

import "testing"

// Teste de unidade

func TestTipoDeEndereco(t *testing.T) {
	//Nome da Função começa obrigatóriamente com a palavra Test
	enderecoParateste := "Avenida Paulista"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParateste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("Tipo diferente do esperado! Esperava o tipo %s e recebeu %s",
			tipoDeEnderecoEsperado,
			tipoDeEnderecoRecebido,
		)
	}

}
