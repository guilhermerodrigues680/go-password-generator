# GO Password Generator

## Instalação

Vá na página de [releases](https://github.com/guilhermerodrigues680/go-password-generator/releases) e faça o download da versão mais recente.

## Como usar

Gerando uma senha com **passwordgenerator** com as configurações padrão:

```sh
$ passwordgenerator <flags> <comprimento_da_senha>
# ou
$ passwordgenerator <comprimento_da_senha>
```

Exemplo:

```sh
# Gera uma senha com caracteres Maiusculos, Minusculos, Numeros e Simbolos
$ passwordgenerator 8

# Gera uma senha somente com caracteres Minusculos e Numeros
$ passwordgenerator -ln 8
```

### Command-Line Options

```console
Password Generator é um gerador de senha de linha de comando

Usage:
  passwordgenerator [flags]

Flags:
  -h, --help        help for passwordgenerator
  -l, --lowecase    usar caracteres minusculos
  -n, --numbers     usar numeros
  -s, --symbols     usar simbolos
  -u, --uppercase   usar caracteres maiusculos
  -v, --verbose     verboso
      --version     version for passwordgenerator
```