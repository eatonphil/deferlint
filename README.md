# deferlint

This package warns when defer is used within loops.

## Building

```
go get github.com/eatonphil/deferlint
go vet -vettool=$GOPATH/cmd/deferlint ./pkg/...
```
