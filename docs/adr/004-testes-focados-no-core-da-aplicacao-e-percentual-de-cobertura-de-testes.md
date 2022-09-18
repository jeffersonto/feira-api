# 4. Testes focados no core da aplicação e percentual de cobertura de testes

Data: 17/09/2022

## Status

Aceito

## Contexto

Aplicação simples de Gerenciamento de Feira, para simulação da dinâmica de CRUD

## Decisão

Os testes foram realizados observando as funções que representam as regras de negócios.

Desta maneira, a cobertura de teste tem maior incidência na pasta `internal`, bem como em `pkg` para o arquivo `commons.go`

> **Outro acordo estabelecido é que novas PRs não podem conter menos de 80% nas pastas supracitadas.**

## Consequências

- Alta coberturua de testes nas pastas referente a regras de negócio;
- Acordo estabelecido sobre o mínimo percentual de testes;

