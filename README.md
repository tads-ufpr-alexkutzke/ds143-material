Tecnologia em Análise e Desenvolvimento de Sistemas

Setor de Educação Profissional e Tecnológica - SEPT

Universidade Federal do Paraná - UFPR

---

*DS143 - Estruturas de Dados II*

Prof. Alexander Robert Kutzke

# Material da Disciplina DS143 - Estruturas de Dados II

Material publicado em
**<https://tads-ufpr-alexkutzke.github.io/ds143-material/>**.

O conteúdo é escrito em Markdown, em `src/`, e publicado com
[mdBook](https://rust-lang.github.io/mdBook/). Cada push na `main` dispara a
publicação por GitHub Actions.

## Trabalhar no material localmente

```bash
cargo install mdbook mdbook-katex     # uma vez
mdbook serve --open                   # recarrega a cada alteração
mdbook build                          # gera o site em book/
```

`src/SUMMARY.md` define o índice e a ordem das páginas: um arquivo que não
estiver listado ali não aparece no livro.

## Espelhos

O repositório é espelhado em
[GitLab](https://gitlab.com/ds143-alexkutzke/material) e em
[Codeberg](https://codeberg.org/ds143/material). Todo push vai para os três
destinos:

```bash
git push origin main && git push gitlab main && git push codeberg main
```
