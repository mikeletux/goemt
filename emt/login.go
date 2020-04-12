package emt

import "encoding/json"

/*
UserLogin struct holds the values needed for a user to authenticate against the EMT Open APÎ
*/
type UserLogin struct {
	// Email verified that user has registered using https://mobilitylabs.emtmadrid.es (mandatory if not put the X-ClientId and passKey params)
	Email string `json:email`

	// Personal password (mandatory if not put the X-ClientId and passKey params)
	Password string `json:password`

	// Optional when email and password are inserted, if not input, MobilityLabs openapi is asumed
	XAPIKey string `json:X-ApiKey`

	// Optional when email and password are inserted, MobilityLabs openapi is asumed. Mandatory when passKey is inserted
	XClientID string `json:X-ClientId`

	// Optional. Mandatory if not exists email and password.
	PassKey string `json:passKey`
}

// LoginProtected gets a token to query the rest of endpoints.
// Protected: Same functionality as Advanced but allows to protect your portal credentials and increase time session up to 86400 seconds. Mandatory X-ClientId and passKey.
func LoginProtected(xClientID, passKey string) string {
	//xClientID and passKey ARE HEADERS, NOT JSON!
}
