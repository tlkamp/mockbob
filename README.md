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

