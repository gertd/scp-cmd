package commerce

//go:generate go run ../../gen/genoai/main.go --input ./specs/commerce.json --name commerce --package commerce --output commerce-impl.go --tpl ./gen/tpl/gen-impl.tpl --imports fmt --imports github.com/spf13/cobra
//go:generate goimports -w commerce-impl.go
