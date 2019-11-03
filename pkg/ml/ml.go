package ml

//go:generate go run ../../gen/genoai/main.go --input ./specs/ml.json --name ml --package ml --output ml-impl.go --tpl ./gen/tpl/gen-impl.tpl --imports fmt --imports github.com/spf13/cobra
//go:generate goimports -w ml-impl.go
