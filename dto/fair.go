package dto

import "github.com/jeffersonto/feira-api/entity"

type Fair struct {
	ID                  int64  `json:"id" form:"id"`
	Longitude           int64  `json:"longitude" form:"longitude" binding:"required"`
	Latitude            int64  `json:"latitude" form:"latitude" binding:"required"`
	SetorCensitario     int64  `json:"setor_censitario" form:"setor_censitario" binding:"required"`
	AreaPonderacao      int64  `json:"area_ponderacao" form:"area_ponderacao" binding:"required"`
	CodigoIBGE          string `json:"codigo_ibge" form:"codigo_ibge" binding:"required"`
	Distrito            string `json:"distrito" form:"distrito" binding:"required"`
	CodigoSubPrefeitura int64  `json:"codigo_subprefeitura" form:"codigo_subprefeitura" binding:"required"`
	SubPrefeitura       string `json:"subprefeitura" form:"subprefeitura" binding:"required"`
	Regiao5             string `json:"regiao5" form:"regiao5" binding:"required"`
	Regiao8             string `json:"regiao8" form:"regiao8" binding:"required"`
	NomeFeira           string `json:"nome_feira" form:"nome_feira" binding:"required"`
	Registro            string `json:"registro" form:"registro" binding:"required"`
	Logradouro          string `json:"logradouro" form:"logradouro" binding:"required"`
	Numero              string `json:"numero" form:"numero"`
	Bairro              string `json:"bairro" form:"bairro"`
	Referencia          string `json:"referencia" form:"referencia"`
}

func (dto *Fair) ToEntity() entity.Fair {
	return entity.Fair{
		ID:                  dto.ID,
		Longitude:           dto.Longitude,
		Latitude:            dto.Latitude,
		SetorCensitario:     dto.SetorCensitario,
		AreaPonderacao:      dto.AreaPonderacao,
		CodigoIBGE:          dto.CodigoIBGE,
		Distrito:            dto.Distrito,
		CodigoSubPrefeitura: dto.CodigoSubPrefeitura,
		SubPrefeitura:       dto.SubPrefeitura,
		Regiao5:             dto.Regiao5,
		Regiao8:             dto.Regiao8,
		NomeFeira:           dto.NomeFeira,
		Registro:            dto.Registro,
		Logradouro:          dto.Logradouro,
		Numero:              dto.Numero,
		Bairro:              dto.Bairro,
		Referencia:          dto.Referencia,
	}
}
