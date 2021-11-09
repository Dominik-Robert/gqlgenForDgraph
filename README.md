# gqlgenForDgraph

generates for the [gqlgen](https://github.com/99designs/gqlgen) project the relvant code for the dgraph database structures.

## How to use it?
Install the binary:
```bash
go install github.com/dominik-robert/gqlgenForDgraph@latest
```

Run the binary (the GOPATH/bin folder should be in your PATH)
```bash
gqlgenForDgraph
```
## Structures
The binary adds for the json-Field dType the right format:

```golang
UID           string       `json:"uid"`
DType         []string     `json:"dgraph.type,omitempty`
```
