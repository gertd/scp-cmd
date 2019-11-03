package collect

//go:generate go run ../../gen/genoai/main.go --input ./specs/collect.json --name collect --package collect --output collect-impl.go --tpl ./gen/tpl/gen-impl.tpl --imports fmt --imports github.com/spf13/cobra
//go:generate goimports -w collect-impl.go
