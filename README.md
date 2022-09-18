<div align="center">

# FEIRA-API

![golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![sqllite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white)

</div>

## Menu

- [Visão Geral](#visão-geral)
- [Architectural Decision Record (ADR)](#architectural-decision-record-adr)
- [Requisitos](#requisitos)
- [Como executar](#como-executar)
  - [Testes](#testes)
  - [Principais End-Points e Retornos](#principais-end-points-e-retornos)
- [Collection Postman](#collection-postman)
- [Débitos Técnicos](#debitos-tecnicos)
- [Licença](#licença)
- [Autor](#autor)

## Visão Geral

Este projeto visa carregar informações das feiras públicas disponibilizadas no formato CSV pela Prefeitura de São Paulo, e alimentar um banco de dados em memória (SQL Lite) para simular um fluxo de CRUD com os recursos: GET, POST, PUT e DELETE.

## Architectural Decision Record (ADR)

[1. Atributos de Estruturas de Dados e Documentação em Português](docs/adr/0001-atributos-de-estruturas-de-dados-e-documentacao-em-portugues.md)

[2. Banco de Dados em Memória (SQL Lite) para Gerenciamento de Feiras](docs/adr/0002-banco-de-dados-em-memoria-para-gerenciamento-de-feiras.md)

[3. Layout de Estrutura de Pasta e Arquitetura de Software](docs/adr/0003-layout-de-estrutura-de-pasta-e-arquitetura-de-software.md)

[4. Testes focados no core da aplicação e percentual de cobertura de testes](docs/adr/004-testes-focados-no-core-da-aplicacao-e-percentual-de-cobertura-de-testes.md)

[5. Desnormalização da tabela de feiras livres](docs/adr/005-desnormalizacao-tabela-feiras-livres.md)

## Requisitos
- [Git](https://git-scm.com/downloads)
- [Golang 1.17 ou superior](https://go.dev/doc/install)

## Como executar

Seguir os comandos abaixo:

1. Clone do projeto através do Git

```
git clone https://github.com/jeffersonto/feira-api.git
```

2. Na pasta raiz do projeto, executar a integração das dependências:
```
go mod tidy
```

3. Executar o run do GO
```
go run github.com/jeffersonto/feira-api/cmd
```

Ou abrir o projeto em sua IDE preferida e executá-lo através de atalhos disponíveis.

O arquivo de feira será importado automaticamente a cada execução da aplicação, não sendo necessário quaisquer passos adicionais.

### Testes

Para execução dos testes com cobertura, na pasta raiz do projeto, seguir os passos:

1. Executar o go test
```
go test -v -coverprofile cover.out ./...
```

2. Executar o go tool cover
```
go tool cover -html cover.out -o cover.html
```
3. Abrir o `cover.html` no seu navegador de preferência.

### Principais End-Points e Retornos

- Verifica se a aplicação está em execução (Ping)
```
curl --location --request GET 'http://localhost:8080/ping'
```
> > 200 - Ok: busca realizada com registros

- Busca uma feira por ID
```
curl --location --request GET 'http://localhost:8080/v1/feiras/1'
```
> > 200 - Ok: busca realizada com registros
>
> > 204 - No Content: busca realizada sem registros
>
> > 500 - Internal Server Error: erro desconhecido na busca dos dados

- Busca feiras Por Query Params
```
curl --location --request GET 'http://localhost:8080/v1/feiras?bairro=VL FORMOSA'
```


Parameters | Type
------------- | -------------
distrito | string
regiao5 | string
nomeFeira | string
bairro | string

> > 200 - Ok: busca realizada com registros
>
> > 204 - No Content: busca realizada sem registros
>
> > 500 - Internal Server Error: erro desconhecido na busca dos dados

- Busca todas as feiras
```
curl --location --request GET 'http://localhost:8080/v1/feiras''
```
> > 200 - Ok: busca realizada com registros
>
> > 204 - No Content: busca realizada sem registros
>
> > 500 - Internal Server Error: erro desconhecido na busca dos dados

- Cria uma Nova Feira
```
curl --location --request POST 'http://localhost:8080/v1/feiras' \
--header 'Content-Type: application/json' \
--data-raw '{
    "longitude": -46550164,
    "latitude": -23558733,
    "setor_censitario": 355030885000091,
    "area_ponderacao": 3550308005040,
    "codigo_ibge": "87",
    "distrito": "VILA FORMOSA",
    "codigo_subprefeitura": 26,
    "subprefeitura": "ARICANDUVA-FORMOSA-CARRAO",
    "regiao5": "Leste",
    "regiao8": "Leste 1",
    "nome_feira": "VILA FORMOSA",
    "registro": "4041-0",
    "logradouro": "RUA MARAGOJIPE",
    "numero": "S/N",
    "bairro": "VL FORMOSA",
    "referencia": "TV RUA PRETORIA"
}'
```
> > 201 - Created: novo registro criado com sucesso
>
> > 400 - Bad Request: body passado incorretamente
>
> > 500 - Internal Server Error: erro desconhecido na busca dos dados

- Atualiza uma Feira
```
curl --location --request PUT 'http://localhost:8080/v1/feiras/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "longitude": -46550164,
    "latitude": -23558733,
    "setor_censitario": 355030885000091,
    "area_ponderacao": 3550308005040,
    "codigo_ibge": "87",
    "distrito": "VILA FORMOSA",
    "codigo_subprefeitura": 26,
    "subprefeitura": "ARICANDUVA-FORMOSA-CARRAO",
    "regiao5": "Leste",
    "regiao8": "Leste 1",
    "nome_feira": "VILA FORMOSA",
    "registro": "4041-0",
    "logradouro": "RUA MARAGOJIPE",
    "numero": "S/N",
    "bairro": "VL FORMOSA",
    "referencia": "TV RUA PRETORIA - 3"
}'
```
> > 200 - Ok
>
> > 204 - No Content: não foi encontrado registro para atualizar
>
> > 400 - Bad Request: body ou path parameter passado incorretamente
>
> > 500 - Internal Server Error: erro desconhecido na busca dos dados

- Deleta uma Feira por ID
```
curl --location --request DELETE 'http://localhost:8080/v1/feiras/1'
```
> > 204 - No Content: deleção executada corretamente ou não houve afetação por ausência do dado
>
> > 500 - Internal Server Error: erro desconhecido na busca dos dado

## Collection Postman

No link abaixo será possível fazer o download da collection do Postman para facilitar as chamadas dos end-points:

[Feita-Api Collection Postman](resources/collection/Feira-API.postman_collection.json)

## Débitos Técnicos

- [ ] Implementação de cache
- [ ] Atualização de feiras por Patch

## Licença

The [MIT License]() (MIT)

## Autor

- Jefferson de Almeida Costa:

  - Linkedin: https://www.linkedin.com/in/jeffersonacosta/
  - Email: jefferson.acosta@hotmail.com
