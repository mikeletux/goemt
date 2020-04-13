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
	apiVersion      = "v2"
)

/*
ClientConfig struct is used for user login
*/
type ClientConfig struct {
	//Enpoint is where the EMT openapi is located
	Enpoint string

	// Email verified that user has registered using https://mobilitylabs.emtmadrid.es (mandatory if not put the X-ClientId and passKey params)
	Email string

	// Personal password (mandatory if not put the X-ClientId and passKey params)
	Password string

	// Optional when email and password are inserted, if not input, MobilityLabs openapi is asumed
	XAPIKey string

	// Optional when email and password are inserted, MobilityLabs openapi is asumed. Mandatory when passKey is inserted
	XClientID string

	// Optional. Mandatory if not exists email and password.
	PassKey string

	//Insecure enforces validation of SSL certificate
	Insecure bool

	//TLSHandshakeTimeout controls TLS handshake timeout
	TLSHandshakeTimeout int

	//HTTPClient is the optional client to connect with.
	HTTPClient *http.Client
}

/*
getLoginMethod Checks three different kind of login from https://apidocs.emtmadrid.es/#api-Block_1_User_identity-login
	- Basic: Allows to use the API on basic level (up to 25k hits/day). Mandatory request params are email and password
	- Advanced: Allows to use the API on advanced level (up to 250k/day). Mandatory register your application in MobilityLabs and including in the request params are email, password, X-ApiKey and X-ClientId.
	- Protected: Same functionality as Advanced but allows to protect your portal credentials and increase time session up to 86400 seconds. Mandatory X-ClientId and passKey.
*/
func (c ClientConfig) getLoginMethod() (m string, err error) {
	if c.Email != "" && c.Password != "" && c.XAPIKey == "" && c.XClientID == "" && c.PassKey == "" {
		return "basic", nil
	}
	if c.Email != "" && c.Password != "" && c.XAPIKey != "" && c.XClientID != "" && c.PassKey == "" {
		return "advanced", nil
	}
	if c.Email == "" && c.Password == "" && c.XAPIKey == "" && c.XClientID != "" && c.PassKey != "" {
		return "protected", nil
	}
	return m, fmt.Errorf("login parameters are ambiguous")
}

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
Connect returns a APIClient usable for getting info from the EMT Rest API.
It needs a ClientConfig struct to be able to log in.
*/
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
	//Check kind of login mode
	loginMode, err := config.getLoginMethod()
	if err != nil {
		return c, err
	}
	//Need to authenticate
	token, err := Login(client.HTTPClient, config, loginMode)
	if err != nil {
		return c, err
	}
	client.auth = token

	return client, nil
}

/*
Logout method closes the session against the EMT rest API
*/
func (c *APIClient) Logout() error {
	err := Logout(c)
	if err != nil {
		return err
	}
	return nil
}
