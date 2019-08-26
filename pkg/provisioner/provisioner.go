package provisioner

//go:generate go run ../../gen/genoai/main.go --input ./specs/provisioner.json --name provisioner --package provisioner --output provisioner-impl.go --tpl ./gen/tpl/gen-impl.tpl --imports fmt --imports github.com/spf13/cobra
//go:generate gofmt -w provisioner-impl.go
