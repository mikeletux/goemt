package busemtmad

import (
	//"encoding/json"
	"fmt"
	"github.com/mikeletux/goemt"
	"strings"
)

const (
	endpointStopDetail = "/transport/busemtmad/stops/<stopId>/detail/"
)

/*
Stop struct holds info regarding one bus stop
*/
type Stop struct {
	//If the stop contaions an electronic panel, contains the number (or empty).
	Pmv string `json:"pmv"`
	//Stop name.
	Name string `json:"name"`
	//id Stop.
	Stop string `json:"stop"`
	//geographical position
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
	//Array of information about the lines using this stop.
	Dataline []struct {
		//Name of line A
		HeaderA string `json:"headerA"`

		//Name of line B
		HeaderB string `json:"headerB"`

		//B mean from A to B, A mean from B to A.
		Direction string `json:"direction"`

		//Public name
		Label string `json:"label"`

		//time of start the service line
		StartTime string `json:"startTime"`

		//time of end the service line
		StopTime string `json:"stopTime"`

		//minimun frequency of line
		MinFreq string `json:"minFreq"`

		//Maximun frequency of line
		MaxFreq string `json:"maxFreq"`

		//related to current query (LA.- Working day, SA.- Saturday, FE.- Festive)
		DayType string `json:"dayType"`

		//code line
		Line string `json:"line"`
	} `json:"dataLine"`
}

/*
stopDetail struct holds data regarding an EMT bus stop when read from the API
*/
type stopDetail struct {
	goemt.Common
	Data []struct {
		Stops []Stop `json:"stops"`
	} `json:"data"`
}

/*
GetStopDetail function returns details of the stop request from EMTMADRID.
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	stopID -> Stop number.
*/
func GetStopDetail(api goemt.IAPI, stopID int) (stopDetails []Stop, err error) {
	var stopData stopDetail
	finalServiceEnpoint := strings.ReplaceAll(endpointStopDetail, "<stopId>", fmt.Sprintf("%d", stopID))
	err = getInfoFromPlatform(api, finalServiceEnpoint, &stopData)
	if err != nil {
		return stopDetails, err
	}
	return stopData.Data[0].Stops, nil
}
