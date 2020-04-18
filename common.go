package goemt

/*
DataInterface is the interface used when dealing with structs that embeds Common struct. (i.e. busemtmad.stopDetail)
*/
type DataInterface interface {
	GetAPIReturnCode() string
	GetAPIReturnDescription() string
}

/*
Common is a struct that holds the common returned values from the EMT Rest API
*/
type Common struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Datetime    string `json:"datetime"` //In the future I'd like to play with json.Unmarshaler interface to have a time.Time struct
}

/*
GetAPIReturnCode method returns the code from the rest API call returned by the EMT Rest API (i.e. "00" means that the API manage to do the job well)
*/
func (c *Common) GetAPIReturnCode() string {
	return c.Code
}

/*
GetAPIReturnDescription method returns the description from the rest API call returned by the EMT Rest API
*/
func (c *Common) GetAPIReturnDescription() string {
	return c.Description
}
