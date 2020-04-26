package busemtmad

import (
	"github.com/mikeletux/goemt"
)

/*
PostListStops is used by the HTTP Post method.
It contains an array contains a list of stops for getting.
*/
type PostListStops []int

/*
ListStops struct is where the list of stops from EMTMADRID will be stored
*/
type ListStops struct {
	//Stop number.
	Node string `json:"node"`

	//Indicates if stop offer public wifi.
	Wifi string `json:"wifi"`

	//Stop name.
	Name string `json:"name"`

	//geographical position.
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`

	//Array of line/direction used by node.
	Lines []string `json:"lines"`
}

type auxListStops struct {
	goemt.Common
	Data []ListStops `json:"data"`
}

/*
GetListStops fuction returns the list of stops from EMTMADRID.
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	postData -> struct used in the body of the POST HTTP method
Returns:
	listStops -> slice with all requested data
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetListStops(api goemt.IAPI, postData PostListStops) (listStops []ListStops, err error) {
	if postData == nil {
		postData = PostListStops{}
	}
	var data auxListStops
	err = postInfoToPlatform(api, endpointListStops, postData, &data)
	if err != nil {
		return listStops, err
	}
	return data.Data, nil
}
