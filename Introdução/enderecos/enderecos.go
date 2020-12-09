package enderecos

import "strings"

//TipoDeEndereco verifica se um endereço tem um tipo válido e retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecotemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecotemUmTipoValido = true
		}
	}

	if enderecotemUmTipoValido {
		return strings.Title(primeiraPalavraEndereco)
	}

	return "Tipo inválido"
}
