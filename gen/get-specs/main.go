package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	scp "gitlab.com/d5s/scp-cmd/gen/pkg/scp"
	yaml3 "gopkg.in/yaml.v3"
)

const (
	envClientID     = "SCP_CLIENT_ID"
	envClientSecret = "SCP_CLIENT_SECRET"
)

func main() {

	var (
		formatArg string
		format    = scp.FormatUnknown
	)

	flag.StringVar(&formatArg, "format", "", "format [json | yaml]")
	flag.Parse()

	switch strings.ToLower(formatArg) {
	case scp.FormatJSON.String():
		format = scp.FormatJSON
	case scp.FormatYAML.String():
		format = scp.FormatYAML
	default:
		log.Fatalf("unknown format value %s", formatArg)
	}

	log.Printf("getspecs --format %s", format)

	// get client id and secret of registered client
	// fail hard when not exist
	clientID := os.Getenv(envClientID)
	if clientID == "" {
		log.Fatalf("environment variable %s not set", envClientID)
	}
	clientSecret := os.Getenv(envClientSecret)
	if clientSecret == "" {
		log.Fatalf("environment variable %s not set", envClientSecret)
	}

	clnt := scp.NewClient(clientID, clientSecret)
	if err := clnt.Authenticate(); err != nil {
		log.Fatal(err)
	}

	info, err := clnt.GetInfo()
	if err != nil {
		log.Fatal(err)
	}

	for serviceName, service := range info.Services {

		// ignore gateway-canary entry
		if serviceName == "gateway-canary" {
			continue
		}

		versions := make(map[string]int)
		for k, v := range service.Versions {
			versions[v] = k
		}

		version := service.RecommendedVersion

		if _, ok := versions[version]; !ok {
			log.Fatalf("service %s version %s not found", serviceName, version)
		}

		spec, err := clnt.GetOpenAPISpec(serviceName, version, format)
		if err != nil {
			log.Fatal(err)
		}

		var writer *os.File
		fileName := fmt.Sprintf("./%s.%s", serviceName, format)
		writer, err = os.Create(fileName)
		// writer = os.Stdout
		defer writer.Close()

		if format == scp.FormatJSON {
			writeJSON(spec, writer)
		} else if format == scp.FormatYAML {
			writeYAML(spec, writer)
		}
	}
}

func writeJSON(v interface{}, w *os.File) error {

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	if _, err := w.Write(b); err != nil {
		return err
	}
	w.Sync()
	return nil
}

func writeYAML(v interface{}, w *os.File) error {
	b, err := yaml3.Marshal(v)
	if err != nil {
		return err
	}
	if _, err := w.Write(b); err != nil {
		return err
	}
	w.Sync()
	return nil
}
