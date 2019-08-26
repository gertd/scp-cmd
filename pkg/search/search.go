package search

//go:generate go run ../../gen/genoai/main.go --input ./specs/search.json --name search --package search --output search-impl.go --tpl ./gen/tpl/gen-impl.tpl --imports fmt --imports github.com/spf13/cobra
//go:generate gofmt -w search-impl.go
