---
title: "Integra"
sub_title: "Jogo em terminal para apoio ao ensino de Cálculo 1"
theme:
  name: light
  override:
    palette:
      colors:
        solarized_cyan: "2aa198"
      classes:
        accent:
          foreground: "2aa198"
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

- Objetivo: desenvolver um jogo aplicável ao ensino de Cálculo 1.
- Conteúdos trabalhados: **<span class="accent">limites</span>**, **<span class="accent">derivadas</span>** e **<span class="accent">integrais</span>**.
- Produto final: um quiz em terminal simples, visual e acessível.

<!-- end_slide -->

Problema
========

<!-- font_size: 6 -->

- Exercícios tradicionais nem sempre mantêm o estudante engajado.
- Em Cálculo 1, a prática frequente é importante para consolidar conceitos.
- O projeto busca oferecer uma forma mais interativa de revisão.

<!-- end_slide -->

Solução Proposta
================

<!-- font_size: 6 -->

- Foi desenvolvido o **<span class="accent">Integra</span>**, um quiz em TUI (*Terminal User Interface*).
- O jogo apresenta perguntas com alternativas **A, B, C e D**.
- A proposta principal é apoiar o **<span class="accent">aprendizado individual</span>** do estudante.
- O aluno responde no próprio ritmo e recebe feedback imediato.
- O fluxo atual possui **7 perguntas**:
  - 3 de <span class="accent">limites</span>,
  - 3 de <span class="accent">derivadas</span>,
  - 1 <span class="accent">integral</span> final.

<!-- end_slide -->

Por que Terminal?
=================

<!-- font_size: 6 -->

- O terminal é leve, rápido e funciona bem em laboratório ou notebook simples.
- A interface TUI evita distrações e mantém o foco no conteúdo matemático.
- O uso de Go permite gerar uma aplicação simples de executar e manter.

<!-- end_slide -->

Tecnologias Utilizadas
======================

<!-- font_size: 6 -->

- **Go**: linguagem principal do projeto.
- **Bubble Tea**: gerenciamento de estado e navegação da interface.
- **Lip Gloss**: estilização visual no terminal.
- **Presenterm**: criação desta apresentação em Markdown para terminal.

Estrutura principal do projeto:

```text
cmd/calculus-challenge/main.go
internal/game/model.go
internal/game/view.go
internal/game/theme.go
pkg/quiz/questions.go
```

<!-- end_slide -->

Como o Jogo Funciona
====================

<!-- font_size: 6 -->

1. O jogo mostra a pergunta no topo.
2. As alternativas aparecem em destaque no centro.
3. O estudante escolhe uma opção.
4. O sistema mostra a resposta correta e uma explicação curta.

Resultado esperado:

- revisão rápida do conteúdo;
- autonomia no estudo;
- feedback imediato durante a prática.

<!-- end_slide -->

Decisões Pedagógicas
====================

<!-- font_size: 6 -->

- As perguntas foram simplificadas para uso em aula.
- A ordem foi definida para acompanhar a progressão do conteúdo:
  - primeiro **<span class="accent">limites</span>**,
  - depois **<span class="accent">derivadas</span>**,
  - por último uma **<span class="accent">integral</span>**.
- A última questão funciona como um momento de atenção extra para a turma.

Exemplo da questão final:

```typst +render
$ integral e^(2x) dif x = 1/2 e^(2x) + C $
```

<!-- end_slide -->

Diferenciais do Projeto
=======================

<!-- font_size: 6 -->

- Interface visual simples e objetiva.
- Fácil adaptação do banco de perguntas.
- Pode ser expandido com:
  - níveis de dificuldade,
  - cronômetro,
  - novas listas de exercícios.

<!-- end_slide -->

Demonstração
============

<!-- font_size: 6 -->

Para executar o jogo:

```bash
go run ./cmd/calculus-challenge
```

Para executar esta apresentação no Presenterm:

```bash
presenterm docs/apresentacao-integra.md
```

<!-- end_slide -->

Conclusão
=========

<!-- font_size: 2 -->

- O **<span class="accent">Integra</span>** atende ao objetivo do trabalho ao propor um jogo para apoio ao ensino de Cálculo 1.
- A solução combina:
  - conteúdo matemático,
  - prática individual,
  - implementação simples e acessível.
- O projeto pode ser usado em revisão de conteúdo e atividades de apoio ao estudo.

Obrigado!
