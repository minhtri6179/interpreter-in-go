# Lexing

## Transformation

Source code -> Tokens -> Abstract syntax tree

- Token is the smallest unit, the source code will be parsed to this unit before the second state: build the abstract syntax tree

To build own interpreter, we need to define own tokens. The tokenize process can simple describe like this: take source code as input and return the tokens
