[![CircleCI](https://circleci.com/gh/tlkamp/mockbob.svg?style=svg)](https://circleci.com/gh/tlkamp/mockbob)

# mockbob

A CLI tool for generating Spongebob meme text.

Because sometimes manually typing out that format takes too many braincells.

## Usage

> Note: Stdin takes precedence over argument provided values

To use `mockbob`, simply call it in your terminal.

```shell
$ mockbob herpaderp
hErPaDeRp
```

For multiple words, use quotes:

```shell
# Arg
$ mockbob "do you even lift bro"
dO yOu EvEn LiFt BrO

# Stdin
$ echo "do you even lift bro" | mockbob
dO yOu EvEn LiFt BrO
```

`mockbob` will preserve punctuation:

```shell
# Arg
$ mockbob "do you even lift bro?"
dO yOu EvEn LiFt BrO?

# Stdin
$ echo "do you even lift bro?" | mockbob
dO yOu EvEn LiFt Bro?
```

And if you want to start off with capital letters, pass the `-c` flag.

```shell
# Arg
$ mockbob -c "do you even lift bro?"
Do YoU eVeN lIfT bRo?

# Stdin
$ echo "do you even lift bro?" | mockbob -c
Do YoU eVeN lIfT bRo?
```

## Installation

To install this module, simply run:

```shell
$ go get github.com/tlkamp/mockbob

# mockbob is now available in your terminal
$ which mockbob
/Users/traci/go/bin/mockbob

$ mockbob -h
mockbob will take any set of input text, and return it in a Spongebob meme mocking format.
        Examples:
        mockbob "do you even lift bro" -> dO yOu EvEn LiFt BrO
        mockbob -c "do you even lift bro" -> Do YoU eVeN lIfT bRo
        mockbob herpderp -> hErPdErP
        mockbob -c herpderp -> HeRpDeRp

Usage:
  mockbob [word or sentence (quoted)] [flags]

Flags:
  -h, --help         help for mockbob
  -c, --start-caps   start the text with a capital letter
```
