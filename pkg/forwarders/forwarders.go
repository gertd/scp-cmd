package forwarders

//go:generate go run ../../gen/genoai/main.go --input ./specs/forwarders.json --name forwarders --package forwarders --output forwarders-impl.go --tpl ./gen/tpl/gen-impl.tpl --imports fmt --imports github.com/spf13/cobra
//go:generate goimports -w forwarders-impl.go
