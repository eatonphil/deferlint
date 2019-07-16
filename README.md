# deferlint

This package warns when defer is used within loops.

## Example

```bash
$ go get github.com/eatonphil/deferlint
$ go vet -vettool=/Users/philipeaton/go/bin/deferlint ./tests
# github.com/eatonphil/deferlint/tests
tests/deferloop.go:7:3: defer in loop found "defer f()"
```
