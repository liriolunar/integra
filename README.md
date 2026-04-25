# Integra

Jogo em terminal feito em Go para apoiar o ensino de Cálculo 1 com perguntas de limites, derivadas e integrais.

## Como rodar o jogo

```bash
go run ./cmd/calculus-challenge
```

Os quizzes são carregados automaticamente a partir de `/questions`.

## Formato dos quizzes

O formato escolhido foi **YAML**, porque ele fica mais legível para editar quizzes com várias perguntas e alternativas.

Cada arquivo em `/questions` deve ser um `.yaml` ou `.yml` com nome do quiz na raiz e uma lista de perguntas:

```yaml
name: Nome do quiz
description: Resumo opcional
questions:
  - topic: Limites
    prompt: Pergunta
    options:
      - A
      - B
      - C
      - D
    answer: 0
    explanation: Explicação
```

## Apresentação

Slides em Presenterm:

```bash
presenterm docs/apresentacao-integra.md
```
