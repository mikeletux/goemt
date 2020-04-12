package goemt

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	userAgent       = "goemt/1.0"
	applicationJSON = "application/json"
)

/*
APIClient struct where all connection data is stored
*/
type APIClient struct {
	// Endpoint is the URL of the EMT service
	endpoint string

	// HTTPClient is of direct HTTP actions
	HTTPClient *http.Client

	// Auth is where the token for auth will be hold
	auth string
}

/*
ClientConfig struct is used for user login
*/
type ClientConfig struct {
	//Enpoint is where the EMT openapi is located
	Enpoint string

	//xClientID value given by emtmadrid
	XClientID string

	//passKey value given by emtmadrid
	PassKey string

	//Insecure enforces validation of SSL certificate
	Insecure bool

	//TLSHandshakeTimeout controls TLS handshake timeout
	TLSHandshakeTimeout int

	//HTTPClient is the optional client to connect with.
	HTTPClient *http.Client
}

func Connect(config ClientConfig) (c *APIClient, err error) {

	if !strings.HasPrefix(config.Enpoint, "http") {
		return c, fmt.Errorf("endpoint must start with http or https")
	}

	client := &APIClient{
		endpoint: config.Enpoint,
	}

	if config.TLSHandshakeTimeout == 0 {
		config.TLSHandshakeTimeout = 10
	}

	if config.HTTPClient == nil {
		defaultTransport := http.DefaultTransport.(*http.Transport)
		transport := &http.Transport{
			Proxy:                 defaultTransport.Proxy,
			DialContext:           defaultTransport.DialContext,
			ForceAttemptHTTP2:     defaultTransport.ForceAttemptHTTP2,
			MaxIdleConns:          defaultTransport.MaxIdleConns,
			IdleConnTimeout:       defaultTransport.IdleConnTimeout,
			TLSHandshakeTimeout:   time.Second * time.Duration(config.TLSHandshakeTimeout),
			ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: config.Insecure,
			},
		}
		client.HTTPClient = &http.Client{Transport: transport}
	} else {
		client.HTTPClient = config.HTTPClient
	}

	//Need to autenticate

}
