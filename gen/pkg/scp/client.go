package scp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"gopkg.in/yaml.v3"
)

const (
	tokenURL = "https://auth.scp.splunk.com/token" //nolint:gosec
	baseURL  = "https://api.scp.splunk.com"
)

// Client --
type Client struct {
	clientID     string
	clientSecret string
	httpClient   *http.Client
	baseURL      *url.URL
	tokenSource  oauth2.TokenSource
}

// Info -- response payload from /info call, to discover all exposed services and versions
type Info struct {
	InfoVersion string             `json:"infoVersion"`
	Services    map[string]Service `json:"services"`
}

// Service --
type Service struct {
	RecommendedVersion string   `json:"recommendedVersion"`
	Title              string   `json:"title"`
	Versions           []string `json:"versions"`
}

// NewClient --
func NewClient(clientID, clientSecret string) *Client {

	base, _ := url.Parse(baseURL)
	c := Client{
		clientID:     clientID,
		clientSecret: clientSecret,
		baseURL:      base,
		httpClient:   http.DefaultClient,
	}
	return &c
}

// Authenticate --
func (c *Client) Authenticate() error {

	clientConfig := clientcredentials.Config{
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
		TokenURL:     tokenURL,
		Scopes:       []string{"client_credentials"},
	}

	c.tokenSource = clientConfig.TokenSource(context.Background())

	token, err := c.tokenSource.Token()
	if err != nil {
		return err
	}
	if !token.Valid() {
		return fmt.Errorf("token not valid")
	}

	return nil
}

// GetInfo --
func (c *Client) GetInfo() (*Info, error) {

	req, err := c.newRequest("GET", "/info", nil)
	if err != nil {
		return nil, err
	}

	var info Info
	_, err = c.do(req, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetOpenAPISpec -- Retrieve json or yaml formated OpenAPI document
// https://api.scp.splunk.com/system/{service}/specs/{version}/openapi.[json|yaml]
func (c *Client) GetOpenAPISpec(service, version string, format Format) (interface{}, error) {

	if format == FormatUnknown {
		return nil, fmt.Errorf("invalid format unknown")
	}

	path := fmt.Sprintf("/system/%s/specs/%s/openapi.%s", service, version, format.String())

	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var data interface{}
	_, err = c.do(req, &data)

	return data, err
}

// accessToken --
func (c *Client) accessToken() string {
	token, err := c.tokenSource.Token()
	if err != nil {
		return ""
	}
	return token.AccessToken
}

// newRequest --
func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.baseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept", "application/yaml")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken()))
	return req, nil
}

// do --
func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.Header["Content-Type"][0] == "application/yaml" {
		err = yaml.NewDecoder(resp.Body).Decode(v)
	} else {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}
