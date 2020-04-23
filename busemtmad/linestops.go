package busemtmad

import (
	"fmt"
	"github.com/mikeletux/goemt"
	"strings"
)

/*
LineStops struct is where the list of stops of a line-direction and every line that coincides with some stop will be stored
*/
type LineStops struct {
	//Line number
	Line string `json:"line"`

	//Array of timetables of every line (see infoline/detail)
	Timetable []struct {
		IDLine     string `json:"IdLine"`
		NameA      string `json:"nameA"`
		TypeOfDays []struct {
			DayType    string `json:"DayType"`
			Direction1 struct {
				MaximumFrequency string `json:"MaximumFrequency"`
				MinimunFrequency string `json:"MinimunFrequency"`
				StopTime         string `json:"StopTime"`
				Frequency        string `json:"Frequency"`
				StartTime        string `json:"StartTime"`
			} `json:"Direction1"`
			Direction2 struct {
				MaximumFrequency string `json:"MaximumFrequency"`
				MinimunFrequency string `json:"MinimunFrequency"`
				StopTime         string `json:"StopTime"`
				Frequency        string `json:"Frequency"`
				StartTime        string `json:"StartTime"`
			} `json:"Direction2"`
		} `json:"typeOfDays"`
	} `json:"timetable"`

	//Array contains detail of bus stops
	Stops []struct {
		//If exists electronic panel, the idNumber
		Pmv string `json:"pmv"`

		//Stop name
		Name string `json:"name"`

		//Number of stop
		Stop string `json:"stop"`

		//Lines using this stop
		DataLine []string `json:"dataLine"`
		//Direction
		PostalAddress string `json:"postalAddress"`

		//GEO-JSON projection
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"stops"`
}

type auxLineStops struct {
	goemt.Common
	Data []LineStops `json:"data"`
}

/*
GetLineStops func will get the list of stops of a line-direction and every line that coincides with some stop
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	lineID -> Line EMT
	direction ->  [1] - From A to B [2] - From B to A
Returns:
	lineStops -> slice with all requested data
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetLineStops(api goemt.IAPI, lineID, direction int) (lineStops []LineStops, err error) {
	if lineID == 0 || direction == 0 {
		return lineStops, fmt.Errorf("lineID and direction bust be different from 0")
	}
	finalServiceEnpoint := strings.ReplaceAll(endpointLineStops, "<lineId>", fmt.Sprintf("%v", lineID))
	finalServiceEnpoint = strings.ReplaceAll(finalServiceEnpoint, "<direction>", fmt.Sprintf("%v", direction))
	var data auxLineStops
	err = getInfoFromPlatform(api, finalServiceEnpoint, &data)
	if err != nil {
		return lineStops, err
	}
	return data.Data, nil
}
