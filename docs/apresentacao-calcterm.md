---
title: "CalcTerm"
sub_title: "Jogo em terminal para apoio ao ensino de Cálculo 1"
theme:
  name: light
  override:
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

- Objetivo: desenvolver um jogo aplicável ao ensino de Cálculo 1.
- Conteúdos trabalhados: **limites**, **derivadas** e **integrais**.
- Produto final: um quiz em terminal simples, visual e acessível.

<!-- end_slide -->

Problema
========

- Exercícios tradicionais nem sempre mantêm o estudante engajado.
- Em Cálculo 1, a prática frequente é importante para consolidar conceitos.
- O projeto busca oferecer uma forma mais interativa de revisão.

<!-- end_slide -->

Solução Proposta
================

- Foi desenvolvido o **CalcTerm**, um quiz em TUI (*Terminal User Interface*).
- O jogo apresenta perguntas com alternativas **A, B, C e D**.
- A proposta principal é apoiar o **aprendizado individual** do estudante.
- O aluno responde no próprio ritmo e recebe feedback imediato.
- O fluxo atual possui **7 perguntas**:
  - 3 de limites,
  - 3 de derivadas,
  - 1 integral final como pegadinha.

<!-- end_slide -->

Por que Terminal?
=================

- O terminal é leve, rápido e funciona bem em laboratório ou notebook simples.
- A interface TUI evita distrações e mantém o foco no conteúdo matemático.
- O uso de Go permite gerar uma aplicação simples de executar e manter.

<!-- end_slide -->

Tecnologias Utilizadas
======================

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

- As perguntas foram simplificadas para uso em aula.
- A ordem foi definida para acompanhar a progressão do conteúdo:
  - primeiro **limites**,
  - depois **derivadas**,
  - por último uma **integral-surpresa**.
- A última questão funciona como um momento de atenção extra para a turma.

Exemplo da pegadinha final:

```typst +render
$ integral e^(2x) dif x = 1/2 e^(2x) + C $
```

<!-- end_slide -->

Diferenciais do Projeto
=======================

- Interface visual simples e objetiva.
- Fácil adaptação do banco de perguntas.
- Pode ser expandido com:
  - níveis de dificuldade,
  - cronômetro,
  - novas listas de exercícios.

<!-- end_slide -->

Demonstração
============

Para executar o jogo:

```bash
go run ./cmd/calculus-challenge
```

Para executar esta apresentação no Presenterm:

```bash
presenterm docs/apresentacao-calcterm.md
```

<!-- end_slide -->

Conclusão
=========

- O **CalcTerm** atende ao objetivo do trabalho ao propor um jogo para apoio ao ensino de Cálculo 1.
- A solução combina:
  - conteúdo matemático,
  - prática individual,
  - implementação simples e acessível.
- O projeto pode ser usado em revisão de conteúdo e atividades de apoio ao estudo.

Obrigado!
