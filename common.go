package goemt

/*
Common struct for all rest API responses
*/
type Common struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Datetime    string `json:"datetime"` //In the future I'd like to play with json.Unmarshaler interface to have a time.Time struct
}
