package domain

import "gorm.io/gorm"

type Estado int

const (
	ACRE_AC Estado = iota
	ALAGOAS_AL
	AMAPA_AP
	AMAZONAS_AM
	BAHIA_BA
	CEARA_CE
	DISTRITO_FEDERAL_DF
	ESPIRITO_SANTO_ES
	GOIAS_GO
	MARANHAO_MA
	MATO_GROSSO_MT
	MATO_GROSSO_DO_SUL_MS
	MINAS_GERAIS_MG
	PARA_PA
	PARAIBA_PB
	PARANA_PR
	PERNAMBUCO_PE
	PIAUI_PI
	RIO_DE_JANEIRO_RJ
	RIO_GRANDE_DO_NORTE_RN
	RIO_GRANDE_DO_SUL_RS
	RONDONIA_RO
	RORAIMA_RR
	SANTA_CATARINA_SC
	SAO_PAULO_SP
	SERGIPE_SE
	TOCANTINS_TO
)

var estados = [...]string{
	"Acre (AC)",
	"Alagoas (AL)",
	"Amapá (AP)",
	"Amazonas (AM)",
	"Bahia (BA)",
	"Ceará (CE)",
	"DistritoFederal (DF)",
	"EspíritoSanto (ES)",
	"Goiás (GO)",
	"Maranhão (MA)",
	"MatoGrosso (MT)",
	"MatoGrossoDoSul (MS)",
	"MinasGerais (MG)",
	"Pará (PA)",
	"Paraíba (PB)",
	"Paraná (PR)",
	"Pernambuco (PE)",
	"Piauí (PI)",
	"RioDeJaneiro (RJ)",
	"RioGrandeDoNorte (RN)",
	"RioGrandeDoSul (RS)",
	"Rondônia (RO)",
	"Roraima (RR)",
	"SantaCatarina (SC)",
	"SãoPaulo (SP)",
	"Sergipe (SE)",
	"Tocantins (TO)"}

func (e Estado) String() string {
	return estados[e]
}

func (e Estado) IndexEnum() int {
	return int(e)
}

type Bairro struct {
	gorm.Model
	Nome      string     `gorm:"column:nome"`
	Enderecos []Endereco `gorm:"column:enderecos"`
}

type Cidade struct {
	gorm.Model
	Nome      string     `gorm:"column:nome"`
	Estado    Estado     `gorm:"column:estado"`
	Enderecos []Endereco `gorm:"column:enderecos"`
}

type Endereco struct {
	gorm.Model
	Logradouro  string `gorm:"column:logradouro;size=60;not null"`
	Numero      string `gorm:"column:numero;size=7;not null"`
	BairroID    uint   `gorm:"column:bairro_id;not null"`
	CEP         string `gorm:"column:cep;size=8;not null"`
	CidadeID    uint   `gorm:"column:cidade_id;not null"`
	Complemento string `gorm:"column:complemento;size=250"`
}
