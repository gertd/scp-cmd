package main

import (
	"context"

	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/gertd/scp-cmd/gen/pkg/gen"
	"github.com/getkin/kin-openapi/openapi3"
	flag "github.com/spf13/pflag"
)

type config struct {
	rootdir  string
	input    string
	svcname  string
	pkgname  string
	output   string
	template string
	imports  []string
}

func main() {

	cfg := config{}
	var err error

	if cfg.rootdir, err = rootDir(); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&cfg.input, "input", "", "filepath to OpenAPI model file")
	flag.StringVar(&cfg.svcname, "name", "", "service name")
	flag.StringVar(&cfg.pkgname, "package", "", "package name")
	flag.StringVar(&cfg.output, "output", "", "output filename")
	flag.StringVar(&cfg.template, "tpl", "", "template filepath")
	flag.StringArrayVar(&cfg.imports, "imports", nil, "imports")
	flag.Parse()

	log.Printf("genoai -- OpenAPI generator")
	log.Printf("  root    %s", cfg.rootdir)
	log.Printf("--input   %s", cfg.input)
	log.Printf("--name    %s", cfg.svcname)
	log.Printf("--package %s", cfg.pkgname)
	log.Printf("--output  %s", cfg.output)
	log.Printf("--tpl     %s", cfg.template)
	log.Printf("--imports %s", cfg.imports)

	if err := validate(cfg); err != nil {
		log.Fatal(err)
	}

	if err := generate(cfg); err != nil {
		log.Fatal(err)
	}
}

func generate(c config) error {

	spec, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(path.Join(c.rootdir, c.input))
	if err != nil {
		return err
	}

	service := gen.Service{Name: c.svcname}

	if err := transform(spec, &service); err != nil {
		return err
	}

	g := gen.Generator{
		Generator:   "genoai",
		PackageName: c.pkgname,
		Imports:     c.imports,
		Service:     service,
		Model:       spec,
	}

	b, err := g.Generate(path.Join(c.rootdir, c.template))
	if err != nil {
		return err
	}

	var f *os.File
	if c.output == "" {
		f = os.Stdout
	} else {
		f, err = os.Create(c.output)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	fmt.Fprintf(f, string(b))

	log.Printf("%s", c.output)

	return nil
}

func transform(spec *openapi3.Swagger, service *gen.Service) error {

	ctx := context.Background()

	service.APIMap = make(map[string]*gen.API)

	for _, path := range spec.Paths {

		for _, op := range path.Operations() {

			service.APIMap[op.OperationID] = &gen.API{
				Descr:        or(op.Description, op.Summary),
				ParameterMap: make(map[string]*gen.Parameter),
			}

			for _, parm := range path.Parameters {

				if parm.Ref != "" {
					parm.Validate(ctx)
				}
				if parm.Value.Schema.Ref != "" {
					parm.Value.Schema.Validate(ctx)
				}

				service.APIMap[op.OperationID].ParameterMap[parm.Value.Name] = &gen.Parameter{
					Descr:    parm.Value.Description,
					Datatype: parm.Value.Schema.Value.Type,
				}
			}

			for _, parm := range op.Parameters {

				if parm.Ref != "" {
					parm.Validate(ctx)
				}
				if parm.Value.Schema.Ref != "" {
					parm.Value.Schema.Validate(ctx)
				}

				service.APIMap[op.OperationID].ParameterMap[parm.Value.Name] = &gen.Parameter{
					Descr:    parm.Value.Description,
					Datatype: parm.Value.Schema.Value.Type,
				}
			}
		}
	}

	return nil
}

func validate(c config) error {

	if c.input == "" {
		return fmt.Errorf("input argument not specified")
	}
	if c.svcname == "" {
		return fmt.Errorf("name argument not specified")
	}
	if c.pkgname == "" {
		return fmt.Errorf("package argument not specified")
	}
	if !dirExists(c.rootdir) {
		return fmt.Errorf("root directory %s not found", c.rootdir)
	}
	if !fileExists(path.Join(c.rootdir, c.input)) {
		return fmt.Errorf("input file %s not found", c.input)
	}
	if !fileExists(path.Join(c.rootdir, c.template)) {
		return fmt.Errorf("template file %s not found", c.template)
	}
	return nil
}

func or(s1, s2 string) string {
	if s1 == "" {
		return s2
	}
	return s1
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func dirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// rootDir -- return path of .git directory
func rootDir() (string, error) {

	cmd := exec.Command("git", "rev-parse", "--show-toplevel")

	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git rev-parse --show-toplevel %s", stderr.String())
	}

	dirPath := strings.TrimSuffix(stdout.String(), "\n")
	if !dirExists(dirPath) {
		return "", fmt.Errorf("directory %s not found", dirPath)
	}

	return dirPath, nil
}
