# Monkey Interpreter

A small Go implementation of the Monkey programming language.

Built while following *Writing An Interpreter In Go* by Thorsten Ball.

It supports integers, booleans, strings, arrays, hashes, conditionals, functions,
closures, and a few built-in functions.

## Run

```bash
go run .
```

Example:

```monkey
let add = fn(a, b) { a + b; };
add(2, 3);
```

## Test

```bash
go test ./...
```
