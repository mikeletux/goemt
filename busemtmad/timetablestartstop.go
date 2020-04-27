package busemtmad

import (
	"fmt"
	"github.com/mikeletux/goemt"
	"strings"
)

/*
TimeTable struct is where start and stop time from one EMT line will be stored
*/
type TimeTable struct {
	//start of current planification
	DateIni string `json:"dateIni"`

	//end of current planification
	DateEnd string `json:"dateEnd"`

	//First time in reference date from A to B
	FirstTimeServiceA string `json:"firstTimeServiceA"`

	//First time in reference date from B to A
	FirstTimeServiceB string `json:"firstTimeServiceB"`

	//Last time in reference date from A to B
	EndTimeServiceA string `json:"endTimeServiceA"`

	//Last time in reference date from B to A
	EndTimeServiceB string `json:"endTimeServiceB"`

	//Related to line timetable and calendar
	DayType string `json:"dayType"`

	//idLine EMT
	Line string `json:"line"`
}

type auxTimeTable struct {
	goemt.Common
	Data []TimeTable `json:"data"`
}

/*
GetTimeTableStartStop func returns start and stop time from one EMT line
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	lineID -> idLine EMT
Returns:
	timeTables -> slice with all requested data
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetTimeTableStartStop(api goemt.IAPI, lineID int) (timeTables []TimeTable, err error) {
	if lineID == 0 {
		return timeTables, fmt.Errorf("lineID must be different from 0")
	}
	finalServiceEnpoint := strings.ReplaceAll(endpointTimeTableStartStop, "<lineId>", fmt.Sprintf("%v", lineID))
	var data auxTimeTable
	err = getInfoFromPlatform(api, finalServiceEnpoint, &data)
	if err != nil {
		return timeTables, err
	}
	return data.Data, nil
}
