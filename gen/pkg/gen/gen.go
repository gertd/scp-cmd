package gen

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"
	"unicode"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
)

// Generator --
type Generator struct {
	Generator   string
	PackageName string
	Imports     []string
	Service     Service
	Model       *openapi3.Swagger
}

// Service --
type Service struct {
	Name   string
	APIMap map[string]*API
}

// API --
type API struct {
	Descr        string
	ParameterMap map[string]*Parameter
}

// Parameter --
type Parameter struct {
	Datatype string
	Descr    string
}

var (
	funcMap = template.FuncMap{
		"pubIdent":       publicIdentifier,
		"privIdent":      privateIdentifier,
		"first":          first,
		"toLower":        strings.ToLower,
		"toUpper":        strings.ToUpper,
		"captilize":      strings.Title,
		"titleCase":      strings.ToTitle,
		"camelCase":      strcase.ToCamel,
		"lowerCamelCase": strcase.ToLowerCamel,
		"snakeCase":      snakeCase,
		"kebabCase":      strcase.ToKebab,
		"lint":           lint,
		"trim":           trim,
		"args":           args,
	}
)

// Generate --
func (g *Generator) Generate(templateFile string) ([]byte, error) {

	b, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return nil, err
	}

	t := template.Must(template.New("").Funcs(funcMap).Parse(string(b)))

	var buf bytes.Buffer

	err = t.Execute(&buf, g)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func publicIdentifier(s string) string {
	return lint(strcase.ToCamel(s))
}

func privateIdentifier(s string) string {
	return lint(strcase.ToLowerCamel(s))
}

func first(s string) string {
	return strings.ToLower(string(s[0]))
}

func snakeCase(s string) string {
	offset := strings.Index(s, "IPv")
	if offset > 0 {
		return toSnakeCase(s[:offset]) + "_" + strings.ToLower(s[offset:])
	}
	return toSnakeCase(s)
}

func toSnakeCase(in string) string {
	runes := []rune(in)

	var out []rune
	for i := 0; i < len(runes); i++ {
		if i > 0 && (unicode.IsUpper(runes[i]) || unicode.IsNumber(runes[i])) && ((i+1 < len(runes) && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

func trim(s string) string {
	return strings.TrimRight(s, "\r\n")
}

func args(kvs ...interface{}) (map[string]interface{}, error) {

	if len(kvs)%2 != 0 {
		return nil, fmt.Errorf("fnArgs requires even number of arguments")
	}

	m := make(map[string]interface{})
	for i := 0; i < len(kvs); i += 2 {
		s, ok := kvs[i].(string)
		if !ok {
			return nil, errors.New("even args to args must be strings")
		}
		m[s] = kvs[i+1]
	}
	return m, nil
}
