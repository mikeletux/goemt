package busemtmad

import (
	"fmt"
	"github.com/mikeletux/goemt"
	"strings"
)

/*
Calendar struct is where info regarding each calendar day will be stored
*/
type Calendar struct {
	//Calendar date
	Date string `json:"date"`

	//in this day or not (depending on this attribute, data planification could have change).
	Strike string `json:"strike"`

	//LA.- Working day, SA.- Saturday, FE.- Holiday
	DayType string `json:"dayType"`
}

type auxCalendar struct {
	goemt.Common
	Data []Calendar `json:"data"`
}

/*
GetCalendar function returns the EMT transport bus calendar
Params:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	startDate -> Date start on YYYYMMDD formnat
	endDate -> Date end on YYYYMMDD format
Returns:
	[]Calendar -> slice with all the calendars
	err -> if err, err will be set, otherwise will be nil
*/
func GetCalendar(api goemt.IAPI, startDate, endDate int) (calendar []Calendar, err error) {
	var data auxCalendar
	if startDate == 0 || endDate == 0 {
		return calendar, fmt.Errorf("startDate and endDate must be different from 0")
	}
	finalServiceEnpoint := strings.ReplaceAll(endpointCalendar, "<startdate>", fmt.Sprintf("%v", startDate))
	finalServiceEnpoint = strings.ReplaceAll(finalServiceEnpoint, "<enddate>", fmt.Sprintf("%v", endDate))
	err = getInfoFromPlatform(api, finalServiceEnpoint, &data)
	if err != nil {
		return calendar, err
	}
	return data.Data, nil
}
