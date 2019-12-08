# mockbob
A CLI tool for generating Spongebob meme text.

Because sometimes manually typing out that format takes too many braincells.

## Usage
To use `mockbob`, simply call it in your terminal.

```shell
$ mockbob herpaderp
hErPaDeRp
```

For multiple words, use quotes:

```shell
$ mockbob "do you even lift bro"
dO yOu EvEn LiFt BrO
```

`mockbob` will preserve punctuation:

```shell
mockbob "do you even lift bro?"
dO yOu EvEn LiFt BrO?
```

And if you want to start off with capital letters, pass the `-c` flag.

```shell
$ mockbob -c "do you even lift bro?"
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