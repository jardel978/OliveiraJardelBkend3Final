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
	Logradouro  string        `json:"logradouro"`
	Numero      string        `json:"numero"`
	Bairro      string        `json:"bairro"`
	CEP         string        `json:"cep"`
	Complemento string        `json:"complemento,omitempty"`
	Cidade      string        `json:"cidade"`
	Estado      domain.Estado `json:"estado"`
}
