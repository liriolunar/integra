---
title: "Integra"
sub_title: "Jogo em terminal para apoio ao ensino de Cálculo 1"
theme:
  name: light
  override:
    palette:
      colors:
        solarized_yellow: "b58900"
        solarized_orange: "cb4b16"
        solarized_red: "dc322f"
        solarized_magenta: "d33682"
        solarized_violet: "6c71c4"
        solarized_blue: "268bd2"
        solarized_cyan: "2aa198"
        solarized_green: "859900"
      classes:
        accent:
          foreground: "2aa198"
        concept:
          foreground: "268bd2"
        warm:
          foreground: "cb4b16"
        strong:
          foreground: "dc322f"
        architecture:
          foreground: "6c71c4"
        positive:
          foreground: "859900"
        title_accent:
          foreground: "b58900"
    default:
      margin:
        percent: 8
      colors:
        foreground: "657b83"
        background: "fdf6e3"
    slide_title:
      colors:
        foreground: "b58900"
    headings:
      h1:
        colors:
          foreground: "268bd2"
      h2:
        colors:
          foreground: "2aa198"
    intro_slide:
      title:
        colors:
          foreground: "268bd2"
      subtitle:
        colors:
          foreground: "586e75"
      author:
        colors:
          foreground: "6c71c4"
    code:
      theme_name: "Solarized (light)"
    block_quote:
      prefix: "▍ "
    footer:
      style: progress_bar
authors:
  - Luiz Gustavo Silva Carvalho
  - Karoline Rodrigues Costa
  - Leonardo Brito Gomes
  - Raquel Pereira Barros Conceição
  - André Rocha Vital
  - Breno Amorim Da Silva
  - Enzo Portella Zorzin
---

Motivação
=========

<!-- font_size: 6 -->

- Revisar Cálculo 1 com listas de exercícios tradicionais é monótono.
- Ferramentas web ou mobile exigem instalação, login ou conexão.
- A proposta foi um quiz leve e rápido, rodando direto no terminal.

<!-- end_slide -->

O que é
=======

<!-- font_size: 6 -->

- Quiz em terminal com perguntas de múltipla escolha.
- Temas: limites, derivadas e integrais.
- O estudante escolhe uma alternativa e recebe feedback imediato.
- Respostas erradas vêm acompanhadas da explicação.

<!-- end_slide -->

Fluxo do Jogo
=============

<!-- font_size: 6 -->

1. O usuário escolhe um quiz no menu inicial.
2. Cada pergunta é exibida com até quatro alternativas.
3. O estudante navega entre as opções e confirma a resposta.
4. O jogo indica se acertou ou errou e mostra a explicação.
5. Ao final, exibe a pontuação e a porcentagem de acertos.

<!-- end_slide -->

Conteúdo Externo
================

<!-- font_size: 4 -->

As perguntas ficam no diretório `/questions` em arquivos YAML. Isso permite adicionar novos quizzes sem recompilar o código.

```yaml
name: Cálculo 1 — Revisão Guiada
questions:
  - topic: Limites
    prompt: Qual é o limite de x² quando x tende a 0?
    options: ["0", "1", "2", "não existe"]
    answer: 0
    explanation: x² tende a 0 quando x se aproxima de 0.
```

<!-- end_slide -->

Menu de Quizzes
===============

<!-- font_size: 6 -->

Na tela inicial, o usuário vê a lista de quizzes disponíveis e seleciona um para começar. Novos arquivos YAML adicionados ao diretório `/questions` aparecem automaticamente no menu.

<!-- end_slide -->

Arquitetura
===========

<!-- font_size: 4 -->

```text
󰉋 cmd/
└── 󰈔 calculus-challenge/
    └──  main.go

󰉋 internal/
└── 󰈔 game/
    ├──  model.go
    ├──  view.go
    └──  theme.go

󰉋 pkg/
└── 󰈔 quiz/
    ├──  loader.go
    └──  types.go

󰉋 questions/
└──  integra-calculo-1.yaml
```

- `cmd/`: ponto de entrada
- `internal/game/`: lógica, interface e tema
- `pkg/quiz/`: parser YAML e tipos
- `questions/`: banco de perguntas

<!-- end_slide -->

Tecnologias
===========

<!-- font_size: 6 -->

- **Go** — linguagem principal
- **Bubble Tea** — framework para TUI
- **Lip Gloss** — estilização de componentes
- **YAML** — formato dos quizzes
- **Presenterm** — geração dos slides

<!-- end_slide -->

Características
===============

<!-- font_size: 6 -->

- Quizzes são independentes do código-fonte.
- Interface simples, navegável pelo teclado.
- Útil para revisão rápida e estudo individual.

<!-- end_slide -->

<!-- font_size: 6 -->

Obrigado!

<!-- end_slide -->
