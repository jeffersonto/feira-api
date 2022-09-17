# 1. Atributos de Estruturas de Dados e Documentação em Português

Data: 17/09/2022

## Status

Aceito

## Contexto

Sistema com única utilização para a Prefeitura de São Paulo, não sendo necessário a internacionalização.

## Decisão

As estruturas de dados e documentações deverão ser escritas em português, observando o conceito da Linguagem Ubíqua do DDD (Domain-Driven Design), facilitando, assim, a comunicação entre os especialistas de domínio e os desenvolvedores.

  - Exemplo
```go
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
```

Contudo, ainda, observando a padronização de escrita de software em inglês, a sua estruturação ao nível de arquitetura de software e organização de pasta deverá observar a língua inglesa.

## Consequências

- Em caso de internacionalização haverá a necessidade de refatoração;
- Difícil compreensão de trechos de códigos por membros da equipe de desenvolvimento de outros países.
