package busemtmad

import (
	"github.com/mikeletux/goemt"
)

/*
PostInfoStops struct implements the post information needed to send to the server
ints are the numbers of stop
*/
type PostInfoStops []int

/*
InfostopsGeneral struct is where the list of stops actives in the commercial operation will be stored
*/
type InfostopsGeneral struct {
	//number of stop
	Node string `json:"node"`

	//Name of bus stop
	Name string `json:"name"`

	//1.- this stop contains a wifi AP
	Wifi string `json:"wifi"`

	//GEO-JSON geometry of stop point
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`

	//array of lines belong to stop
	Lines []string `json:"lines"`
}

type auxInfostopsGeneral struct {
	goemt.Common
	Data []InfostopsGeneral `json:"data"`
}

/*
GetInfostopsGeneral func will return the list of stops actives in the commercial operation
Params:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	postData -> struct used in the body of the POST HTTP method
Returns:
	infoStopsGeneral -> slice with all requested data
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetInfostopsGeneral(api goemt.IAPI, postData PostInfoStops) (infoStopsGeneral []InfostopsGeneral, err error) {
	if postData == nil {
		postData = PostInfoStops{}
	}
	var data auxInfostopsGeneral
	err = postInfoToPlatform(api, endpointInfoStopsGeneral, postData, &data)
	if err != nil {
		return infoStopsGeneral, err
	}
	return data.Data, nil
}
