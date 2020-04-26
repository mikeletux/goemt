package busemtmad

import (
	"fmt"
	"github.com/mikeletux/goemt"
	"strings"
)

/*
RouteOfLine struct will store the full itinerary for a line
*/
type RouteOfLine struct {
	//label of line
	Label string `json:"label"`

	//number of line
	Line string `json:"line"`

	//direction name A
	NameSectionA string `json:"nameSectionA"`

	//direction name B
	NameSectionB string `json:"nameSectionB"`

	//Structure of arrays contains multiline strings of itinerary
	Itinerary struct {
		ToA struct {
			Type     string `json:"type"`
			Features []struct {
				Type       string `json:"type"`
				Properties struct {
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Distance    int    `json:"distance"`
					StrokeWidth int    `json:"stroke-width"`
					Stroke      string `json:"stroke"`
				} `json:"properties"`
				Geometry struct {
					Type        string        `json:"type"`
					Coordinates [][][]float64 `json:"coordinates"`
				} `json:"geometry"`
			} `json:"features"`
		} `json:"toA"`

		ToB struct {
			Type     string `json:"type"`
			Features []struct {
				Type       string `json:"type"`
				Properties struct {
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Distance    int    `json:"distance"`
					StrokeWidth int    `json:"stroke-width"`
					Stroke      string `json:"stroke"`
				} `json:"properties"`
				Geometry struct {
					Type        string        `json:"type"`
					Coordinates [][][]float64 `json:"coordinates"`
				} `json:"geometry"`
			} `json:"features"`
		} `json:"toB"`
	} `json:"itinerary"`

	//Structure of arrays contains stops locations and asociates values
	Stops struct {
		ToA struct {
			Type     string `json:"type"`
			Features []struct {
				Type       string `json:"type"`
				Properties struct {
					StopNum  int    `json:"stopNum"`
					StopName string `json:"stopName"`
					Distance int    `json:"distance"`
				} `json:"properties"`
				Geometry struct {
					Type        string    `json:"type"`
					Coordinates []float64 `json:"coordinates"`
				} `json:"geometry"`
			} `json:"features"`
		} `json:"toA"`

		ToB struct {
			Type     string `json:"type"`
			Features []struct {
				Type       string `json:"type"`
				Properties struct {
					StopNum  int    `json:"stopNum"`
					StopName string `json:"stopName"`
					Distance int    `json:"distance"`
				} `json:"properties"`
				Geometry struct {
					Type        string    `json:"type"`
					Coordinates []float64 `json:"coordinates"`
				} `json:"geometry"`
			} `json:"features"`
		} `json:"toB"`
	} `json:"stops"`
}

type auxRouteOfLine struct {
	goemt.Common
	Data RouteOfLine `json:"data"`
}

/*
GetRouteOfLine func returns the full itinerary for a line
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	labelID -> Public id of line (or line number)
Returns:
	routeOfLine -> struct with all requested data
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetRouteOfLine(api goemt.IAPI, labelID int) (routeOfLine RouteOfLine, err error) {
	if labelID == 0 {
		return routeOfLine, fmt.Errorf("labelID must be different from zero")
	}
	finalServiceEnpoint := strings.ReplaceAll(endpointRouteOfLine, "<labelId>", fmt.Sprintf("%v", labelID))
	var data auxRouteOfLine
	err = getInfoFromPlatform(api, finalServiceEnpoint, &data)
	if err != nil {
		return routeOfLine, err
	}
	return data.Data, nil
}
