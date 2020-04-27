package busemtmad

import (
	"fmt"
	"github.com/mikeletux/goemt"
	"strings"
)

/*
TimeTableTrip struct is where the list of trips of operations related to EMT lines will be stored.
*/
type TimeTableTrip struct {
	//Unique internal id in simualtaneous service
	LogicBus string `json:"logicBus"`

	//From A to B or from B to A
	Direction string `json:"direction"`

	//Number of trip
	TripNum string `json:"tripNum"`

	//Theoretical end trip
	EndTimeTrip string `json:"endTimeTrip"`

	//Theoretical init trip
	StartTimeTrip string `json:"startTimeTrip"`

	//Date reference
	Date string `json:"date"`

	//Related to line timetable and calendar
	DayType string `json:"dayType"`

	//EMT bus line
	Line string `json:"line"`
}

type auxTimeTableTrip struct {
	goemt.Common
	Data []TimeTableTrip `json:"data"`
}

/*
GetTimeTableTrips func returns the list of trips of operations related to EMT lines
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	lineID -> idLine EMT
	dataref -> date reference on YYYYMMDD format
Returns:
	timeTableTrips -> slice with all requested data
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetTimeTableTrips(api goemt.IAPI, lineID, dataref int) (timeTableTrips []TimeTableTrip, err error) {
	if lineID == 0 || dataref == 0 {
		return timeTableTrips, fmt.Errorf("lineID and dataref must be different from 0")
	}
	finalServiceEnpoint := strings.ReplaceAll(endpointTimeTableTrips, "<lineId>", fmt.Sprintf("%v", lineID))
	finalServiceEnpoint = strings.ReplaceAll(finalServiceEnpoint, "<dateRef>", fmt.Sprintf("%v", dataref))
	var data auxTimeTableTrip
	err = getInfoFromPlatform(api, finalServiceEnpoint, &data)
	if err != nil {
		return timeTableTrips, err
	}
	return data.Data, nil
}
