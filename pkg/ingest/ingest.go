package ingest

//go:generate go run ../../gen/genoai/main.go --input ./specs/ingest.json --name ingest --package ingest --output ingest-impl.go --tpl ./gen/tpl/gen-impl.tpl --imports fmt --imports github.com/spf13/cobra
//go:generate goimports -w ingest-impl.go
