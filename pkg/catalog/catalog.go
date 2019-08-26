package catalog

//go:generate go run ../../gen/genoai/main.go --input ./specs/catalog.json --name catalog --package catalog --output catalog-impl.go --tpl ./gen/tpl/gen-impl.tpl --imports fmt --imports github.com/spf13/cobra
//go:generate gofmt -w catalog-impl.go
