# deferlint

This package warns when defer is used within loops.

## Example

```bash
$ go install github.com/eatonphil/deferlint/cmd/deferlint
$ go vet -vettool=$(which deferlint) ./tests
# github.com/eatonphil/deferlint/tests
tests/deferloop.go:7:3: defer in loop found "defer f()"
```

## References

https://github.com/fatih/addlint
