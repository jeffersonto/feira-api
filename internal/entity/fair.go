package entity

type Fair struct {
	ID                  int64  `json:"id" db:"id" csv:"ID"`
	Longitude           int64  `json:"longitude" db:"longitude" csv:"LONG"`
	Latitude            int64  `json:"latitude" db:"latitude" csv:"LAT"`
	SetorCensitario     int64  `json:"setor_censitario" db:"setor_censitario" csv:"SETCENS"`
	AreaPonderacao      int64  `json:"area_ponderacao" db:"area_ponderacao" csv:"AREAP"`
	CodigoIBGE          string `json:"codigo_ibge" db:"codigo_ibge" csv:"CODDIST"`
	Distrito            string `json:"distrito" db:"distrito" csv:"DISTRITO"`
	CodigoSubPrefeitura int64  `json:"codigo_subprefeitura" db:"codigo_subprefeitura" csv:"CODSUBPREF"`
	SubPrefeitura       string `json:"subprefeitura" db:"subprefeitura" csv:"SUBPREFE"`
	Regiao5             string `json:"regiao5" db:"regiao5" csv:"REGIAO5"`
	Regiao8             string `json:"regiao8" db:"regiao8" csv:"REGIAO8"`
	NomeFeira           string `json:"nome_feira" db:"nome_feira" csv:"NOME_FEIRA"`
	Registro            string `json:"registro" db:"registro" csv:"REGISTRO"`
	Logradouro          string `json:"logradouro" db:"logradouro" csv:"LOGRADOURO"`
	Numero              string `json:"numero" db:"numero" csv:"NUMERO"`
	Bairro              string `json:"bairro" db:"bairro" csv:"BAIRRO"`
	Referencia          string `json:"referencia" db:"referencia" csv:"REFERENCIA"`
}
