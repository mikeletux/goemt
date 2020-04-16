package busemtmad

import (
	"encoding/json"
	"fmt"
	"github.com/mikeletux/goemt"
	"strings"
)

const (
	serviceEndpoint = "/transport/busemtmad/stops/<stopId>/detail/"
)

/*
StopDetail struct holds data regarding an EMT bus stop
*/
type StopDetail struct {
	goemt.Common
	Data []struct {
		Stops []struct {
			Pmv      string `json:"pmv"`
			Name     string `json:"name"`
			Geometry string `json:"geometry"`
			Stop     string `json:"stop"`
			Dataline []struct {
				HeaderA   string `json:"headerA"`
				HeaderB   string `json:"headerB"`
				Direction string `json:"direction"`
				Label     string `json:"label"`
				StartTime string `json:"startTime"`
				StopTime  string `json:"stopTime"`
				MinFreq   string `json:"minFreq"`
				MaxFreq   string `json:"maxFreq"`
				DayType   string `json:"dayType"`
				Line      string `json:"line"`
			} `json:"dataLine"`
		} `json:"stops"`
	} `json:"data"`
}

/*
GetStopDetail function returns details of the stop request from EMTMADRID.
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	stopID -> Stop number.
*/
func GetStopDetail(api goemt.IAPI, stopID int) (stopDetails StopDetail, err error) {
	finalServiceEnpoint := strings.ReplaceAll(serviceEndpoint, "<stopId>", fmt.Sprintf("%d", stopID))
	data, err := api.Get(finalServiceEnpoint)
	if err != nil {
		return stopDetails, err
	}
	err = json.Unmarshal(data, &stopDetails)
	if err != nil {
		return stopDetails, err
	}
	return stopDetails, nil
}
