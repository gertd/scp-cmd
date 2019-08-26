package streams

//go:generate go run ../../gen/genoai/main.go --input ./specs/streams.json --name streams --package streams --output streams-impl.go --tpl ./gen/tpl/gen-impl.tpl --imports fmt --imports github.com/spf13/cobra
//go:generate gofmt -w streams-impl.go
