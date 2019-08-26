package appreg

//go:generate go run ../../gen/genoai/main.go --input ./specs/app-registry.json --name app-registry --package appreg --output appreg-impl.go --tpl ./gen/tpl/gen-impl.tpl --imports fmt --imports github.com/spf13/cobra
//go:generate gofmt -w appreg-impl.go
