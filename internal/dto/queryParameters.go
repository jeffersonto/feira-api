package dto

import (
	"github.com/jeffersonto/feira-api/internal/entity"
)

type QueryParameters struct {
	Distrito  string `form:"distrito"`
	Regiao5   string `form:"regiao5"`
	NomeFeira string `form:"nomeFeira"`
	Bairro    string `form:"bairro"`
}

func (dto *QueryParameters) ToEntity() entity.Filter {
	return entity.Filter{
		Distrito:  dto.Distrito,
		Regiao5:   dto.Regiao5,
		Bairro:    dto.Bairro,
		NomeFeira: dto.NomeFeira,
	}
}
