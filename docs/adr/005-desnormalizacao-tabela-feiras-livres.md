# 5. Desnormalização da tabela de feiras livres

Data: 17/09/2022

## Status

Aceito

## Contexto

Aplicação simples de Gerenciamento de Feira, observando melhor performance e simplicidade de carga de dados

## Decisão

Optou-se por manter a tabela desnormalizada, com intuito de evitar querys com joins, entregando, neste caso, uma microperformance (pois trata de tabela extremamente pequena).

Porém, nota-se, a possibilidade de normalização em 5 tabelas distintas, sendo:

- feiras_livres
- distritos
- subprefeituras
- regiao5
- regiao8

## Consequências

- Microperformance, por ausência de joins;
- A desnormalização, porém, ocasiona uma estrutura de linhas de tabela grande, podendo ocasionar `chained row`.

