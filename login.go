package goemt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const loginEnpoint = "/mobilitylabs/user/login/"

/*
LoginResponse is the data structure where the response from the login endpoint is recorded
*/
type LoginResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Datetime    string `json:"datetime"`
	Data        []struct {
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

// LoginProtected gets a token to query the rest of endpoints.
// Protected: Same functionality as Advanced but allows to protect your portal credentials and increase time session up to 86400 seconds. Mandatory X-ClientId and passKey.
func LoginProtected(c *http.Client, config ClientConfig) (s string, err error) {
	req, err := http.NewRequest("GET", config.Enpoint+loginEnpoint, nil)
	if err != nil {
		return s, err
	}
	req.Header.Add("X-ClientId", config.XClientID)
	req.Header.Add("passKey", config.PassKey)
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
	return data.Data[0].AccessToken, nil

}
