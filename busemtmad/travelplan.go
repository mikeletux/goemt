package busemtmad

import (
	"github.com/mikeletux/goemt"
)

/*
PostTravelPlan struct is used for posting info to the EMT Rest API
*/
type PostTravelPlan struct {
	//P (Public transport) C (by Car from/to Parking) W (Walking) M (Mixed)
	RouteType string `json:"routeType"`

	//longitude origin
	CoordinateXFrom float64 `json:"coordinateXFrom"`

	//latitude origin
	CoordinateYFrom float64 `json:"coordinateYFrom"`

	//longitude destination
	CoordinateXTo float64 `json:"coordinateXTo"`

	//latitude destination
	CoordinateYTo float64 `json:"coordinateYTo"`

	//Symbolic name of origin
	OriginName string `json:"originName"`

	//Symbolic name of destination
	DestinationName string `json:"destinationName"`

	//Optional Array on JSON format contains coordinates of polygon. (If input, represents one exclusion area and route surround it).
	Polygon string `json:"polygon"`

	//optional day of routeplan
	Day int `json:"day"`

	//optional month of routeplan
	Month int `json:"month"`

	//optional year of routeplan
	Year int `json:"year"`

	//optional hour of routeplan
	Hour int `json:"hour"`

	//optional minute of routeplan
	Minute int `json:"minute"`

	//EN (english) ES (spanish)
	Culture string `json:"culture"`

	//(if true returns the full multiline string for mapping representations)
	Itinerary bool `json:"itinerary"`

	//(if true and routeType is "M" or "P" recover routes by bus)
	AllowBus bool `json:"allowBus"`

	//(if true and routeType is "M" or "P" recover routes by public bike)
	AllowBike bool `json:"allowBike"`

	//(reserved for future use)
	PreferPublic bool `json:"preferPublic"`

	//(reserved for future use)
	IsResidentOrInvited bool `json:"isResidentOrInvited"`

	//(reserved for future use)
	IsEnvFriendly bool `json:"isEnvFriendly"`

	//(reserved for future use)
	UsingTaxi bool `json:"usingTaxi"`

	//(reserved for future use)
	UsingRentedCar bool `json:"usingRentedCar"`
}

/*
TravelPlan struct is where calculate route between two points will be stored
*/
type TravelPlan struct {
	Description   string  `json:"description"`
	DepartureTime string  `json:"departureTime"`
	ArrivalTime   string  `json:"arrivalTime"`
	Duration      float64 `json:"duration"`
	Distance      float64 `json:"distance"`
	Sections      []struct {
		Order    int     `json:"order"`
		Type     string  `json:"type"`
		Duration float64 `json:"duration"`
		Distance float64 `json:"distance"`
		Source   struct {
			Type     string `json:"type"`
			Geometry struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
			Properties struct {
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"properties"`
		} `json:"source"`
		Destination struct {
			Type     string `json:"type"`
			Geometry struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
			Properties struct {
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"properties"`
		} `json:"destination"`
		Route struct {
			Type     string `json:"type"`
			Features []struct {
				Type     string `json:"type"`
				Geometry struct {
					Type        string    `json:"type"`
					Coordinates []float64 `json:"coordinates"`
				} `json:"geometry"`
				Properties struct {
					Description string  `json:"description"`
					Duration    float64 `json:"duration"`
					Distance    float64 `json:"distance"`
				} `json:"properties"`
			} `json:"features"`
		} `json:"route"`
		Itinerary struct {
			Type        string      `json:"type"`
			Coordinates [][]float64 `json:"coordinates"`
		} `json:"itinerary"`
		IDLine string `json:"idLine,omitempty"`
	} `json:"sections"`
}

type auxTravelPlan struct {
	goemt.Common
	Data TravelPlan `json:"data"`
}

/*
GetTravelPlan func returns calculated route between two points
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	postData -> struct used in the body of the POST HTTP method
Return:
	travelPlan -> struct with all requested data
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetTravelPlan(api goemt.IAPI, postData PostTravelPlan) (travelPlan TravelPlan, err error) {
	var data auxTravelPlan
	err = postInfoToPlatform(api, endpointTravelPlan, postData, &data)
	if err != nil {
		return travelPlan, err
	}
	return data.Data, nil
}
