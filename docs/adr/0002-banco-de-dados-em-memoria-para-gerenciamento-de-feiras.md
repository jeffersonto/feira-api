# 2. Banco de Dados em Memória (SQL Lite) para Gerenciamento de Feiras

Data: 17/09/2022

## Status

Aceito

## Contexto

Gerenciamento de feiras carregadas a partir de um arquivo texto fornecido pela Prefeitura de São Paulo.

## Decisão

Com intuito de salvar temporariamente os registros, em momento de execução somente, foi escolhida o banco de dados Sql Lite e configurado em forma de cache, haja vista que se trata de um desafio técnico para apresentação da organização de código em uma dinâmica controlada de um CRUD.

## Consequências

- A cada execução do app o arquivo (DEINFO_AB_FEIRASLIVRES_2014.csv) de feiras é recarregado;
- Alterações realizadas em execuções anteriores (inserções, atualizações e deleções) são desfeitas a cada execução do app;

## Decisões Para Outro Contextos

Caso esta aplicação deva ser utilizada em ambiente produtivo, com intuito de disponibilizar e permitir a gestão das informações pelos feirantes, deve-se observar:

- Utilizar um banco de dados relacional como MySql ou Postgresql, pois se aproveita de baixos custos ou nenhum de licenciamento;

  - Em caso de necessidade de escalabilidade, utilizar inicialmente um cluster master (escrita) e dois workers (somente leitura);

- Em caso de alta variabilidade de concorrência de acessos pode-ser migrar a um banco NoSQL se favorecendo de escalabilidade horizontal.
