// Package {{ .PackageName }} -- generated by  {{ .Generator }} 
// !! DO NOT EDIT !! 
// 
package {{ .PackageName }}

import (
	{{ with .Imports -}}
    {{ range . -}} 
	{{ . | printf "%q" }}
	{{ end -}}
	{{ end -}}
)

{{ range $apiName, $api := .Service.APIMap }}
// {{ $apiName | captilize | lint }} -- {{ $api.Descr | trim }} (implementation)
func {{ $apiName | captilize | lint }}(cmd *cobra.Command, args []string) error {
	fmt.Printf("called {{ $apiName | captilize | lint }}\n")
	return nil
}
{{ end }}