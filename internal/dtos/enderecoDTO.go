package dtos

import "OliveiraJardelBkend3Final/internal/domain"

type EnderecoRequestBody struct {
	ID          uint          `json:"id,omitempty"`
	Logradouro  string        `json:"logradouro"`
	Numero      string        `json:"numero"`
	Bairro      string        `json:"bairro"`
	CEP         string        `json:"cep"`
	Complemento string        `json:"complemento,omitempty"`
	Cidade      string        `json:"cidade"`
	Estado      domain.Estado `json:"estado"`
}

type EnderecoResponseBody struct {
	ID          uint          `json:"id,omitempty"`
	Logradouro  string        `json:"logradouro,omitempty"`
	Numero      string        `json:"numero,omitempty"`
	Bairro      string        `json:"bairro,omitempty"`
	CEP         string        `json:"cep,omitempty"`
	Complemento string        `json:"complemento,omitempty"`
	Cidade      string        `json:"cidade,omitempty"`
	Estado      domain.Estado `json:"estado,omitempty"`
}
