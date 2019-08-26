package identity

//go:generate go run ../../gen/genoai/main.go --input ./specs/identity.json --name identity --package identity --output identity-impl.go --tpl ./gen/tpl/gen-impl.tpl --imports fmt --imports github.com/spf13/cobra
//go:generate gofmt -w identity-impl.go
