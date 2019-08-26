package action

//go:generate go run ../../gen/genoai/main.go --input ./specs/action.json --name action --package action --output action-impl.go --tpl ./gen/tpl/gen-impl.tpl --imports fmt --imports github.com/spf13/cobra
//go:generate gofmt -w action-impl.go
