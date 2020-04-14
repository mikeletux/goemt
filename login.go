package goemt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	userIdentity   = "/mobilitylabs/user"
	loginMethod    = "/login/"
	logoutMethod   = "/logout/"
	loginEnpoint   = userIdentity + loginMethod
	logoutEndpoint = userIdentity + logoutMethod
)

/*
LoginResponse is the data structure where the response from the login endpoint is recorded
*/
type LoginResponse struct {
	Common
	Data []struct {
		UpdatedAt          string `json:"updatedAt"`
		UserName           string `json:"userName"`
		AccessToken        string `json:"accessToken"`
		TokenSecExpiration int    `json:"tokenSecExpiration"`
		Email              string `json:"email"`
		IDUser             string `json:"idUser"`
		APICounter         struct {
			Current    int    `json:"current"`
			DailyUse   int    `json:"dailyUse"`
			Owner      int    `json:"owner"`
			LicenceUse string `json:"licenceUse"`
			AboutUses  string `json:"aboutUses"`
		} `json:"apiCounter"`
	} `json:"data"`
}

// Login gets a token to query the rest of endpoints.
func Login(c *http.Client, config ClientConfig, mode string) (s string, err error) {
	req, err := http.NewRequest("GET", config.Enpoint+loginEnpoint, nil)
	if err != nil {
		return s, err
	}
	switch mode {
	case "basic":
		req.Header.Add("email", config.Email)
		req.Header.Add("password", config.Password)
	case "advanced":
		req.Header.Add("email", config.Email)
		req.Header.Add("password", config.Password)
		req.Header.Add("X-ApiKey", config.XAPIKey)
		req.Header.Add("X-ClientId", config.XClientID)
	case "protected":
		req.Header.Add("X-ClientId", config.XClientID)
		req.Header.Add("passKey", config.PassKey)
	}
	res, err := c.Do(req)
	if err != nil {
		return s, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return s, err
	}
	//Defined the variable where the json data is going to be stored
	var data LoginResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return s, err
	}
	if data.Code != "00" {
		return s, fmt.Errorf("emt server error - code %s - description %s", data.Code, data.Description)
	}
	return data.Data[0].AccessToken, nil
}

/*
Logout function finishes a session with the EMT rest API server
Needs a APIClient struct initialized to sucessfully close the connection
*/
func Logout(c *APIClient) error {
	if c == nil {
		return fmt.Errorf("APIClient cannot be nil")
	}
	req, err := http.NewRequest("GET", c.endpoint+logoutEndpoint, nil)
	if err != nil {
		return err
	}
	if c.auth == "" {
		return fmt.Errorf("cannot log out, there's no session created")
	}
	req.Header.Add("accessToken", c.auth)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var data Common
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	if data.Code != "03" {
		return fmt.Errorf("emt server error - code %s - description %s", data.Code, data.Description)
	}
	return nil
}
