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

Projeto Aplicado 6
==================

<!-- font_size: 6 -->

- Problema: revisar **<span class="concept">Cálculo 1</span>** de forma mais atraente.
- Solução: um quiz em terminal chamado **<span class="accent">Integra</span>**.
- Foco: **<span class="concept">limites</span>**, **<span class="warm">derivadas</span>** e **<span class="positive">integrais</span>**.

<!-- end_slide -->

Problema
========

<!-- font_size: 6 -->

- Exercícios repetitivos podem reduzir o **<span class="strong">engajamento</span>**.
- Cálculo exige **<span class="warm">prática frequente</span>**.
- Queríamos uma revisão:
  - <span class="concept">rápida</span>,
  - <span class="accent">simples</span>,
  - <span class="positive">interativa</span>.

<!-- end_slide -->

Solução
=======

<!-- font_size: 6 -->

- Criamos o **<span class="accent">Integra</span>**.
- Um jogo em **<span class="architecture">terminal</span>** com perguntas de múltipla escolha.
- O estudante responde e recebe **<span class="positive">feedback imediato</span>**.

```typst +render
$ integral e^(2x) dif x = 1/2 e^(2x) + C $
```

<!-- end_slide -->

Como Funciona
=============

<!-- font_size: 6 -->

1. Pergunta no topo.
2. Alternativas em destaque.
3. Escolha do estudante.
4. Correção + explicação.

Resultado:

- <span class="positive">prática individual</span>;
- <span class="concept">revisão rápida</span>;
- <span class="warm">reforço imediato</span>.

<!-- end_slide -->

Quizzes Externos
================

<!-- font_size: 4 -->

- Os conteúdos ficam em **<span class="accent">/questions</span>**.
- O formato escolhido foi **<span class="architecture">YAML</span>**.

```yaml
name: Cálculo 1 — Revisão Guiada
questions:
  - topic: Limites
    prompt: Qual é o limite lim x→0 de x² ?
    options: ["0", "1", "2", "não existe"]
    answer: 0
    explanation: x² tende a 0.
```

<!-- end_slide -->

Menu de Quizzes
===============

<!-- font_size: 6 -->

- Ao abrir o jogo, o usuário escolhe um **<span class="concept">quiz</span>**.
- Novos quizzes podem ser adicionados sem mudar o código.
- Isso deixa o projeto **<span class="architecture">modular</span>** e **<span class="positive">reaproveitável</span>**.

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

<!-- end_slide -->

Tecnologias
===========

<!-- font_size: 6 -->

- **<span class="concept">Go</span>**
- **<span class="warm">Bubble Tea</span>**
- **<span class="accent">Lip Gloss</span>**
- **<span class="architecture">YAML</span>**
- **<span class="positive">Presenterm</span>**

<!-- end_slide -->

Diferenciais
============

<!-- font_size: 6 -->

- <span class="positive">leve e rápido</span>;
- <span class="concept">fácil de executar</span>;
- <span class="architecture">fácil de expandir</span>;
- <span class="warm">bom para revisão e estudo individual</span>.

<!-- end_slide -->

Demonstração
============

<!-- font_size: 4 -->

```bash
go run ./cmd/calculus-challenge
```

```bash
presenterm docs/apresentacao-integra.md
```

<!-- end_slide -->

Conclusão
=========

<!-- font_size: 6 -->

- O **<span class="accent">Integra</span>** transforma revisão em interação.
- Une conteúdo matemático com uma interface simples.
- É uma base pronta para **<span class="positive">novos quizzes</span>** e **<span class="architecture">novas funcionalidades</span>**.

Obrigado!
