// Package {{ .PackageName }} -- generated by {{ .Generator }}
// !! DO NOT EDIT !! 
// 
package {{ .PackageName }}

import (
	{{ with .Imports -}}
    {{ range . -}} 
	{{ . | printf "%q" }}
	{{ end -}}
	{{ end -}}
    impl "github.com/gertd/scp-cmd/pkg/{{ .PackageName }}"
)

{{ range $apiName, $api := .Service.APIMap }}
// {{ $apiName | lint }} -- {{ $api.Descr | trim }}
var {{ $apiName | lint }}Cmd = &cobra.Command{
	Use:   "{{ $apiName | kebabCase }}",
	Short: "{{ $api.Descr | trim }}",
	RunE:  impl.{{ $apiName | captilize | lint }},
}
{{ end }}

func init() {

    {{ range $apiName, $api := .Service.APIMap -}}
    {{ $.PackageName }}Cmd.AddCommand({{ $apiName | lint }}Cmd)
    {{ end }}

}
