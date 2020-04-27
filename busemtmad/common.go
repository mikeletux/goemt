package busemtmad

import (
	"encoding/json"
	"fmt"
	"github.com/mikeletux/goemt"
)

const (
	endpointStopDetail                   = "/transport/busemtmad/stops/<stopId>/detail/"
	endpointStopsAroundStop              = "/transport/busemtmad/stops/arroundstop/<stopId>/<radius>/"
	endpointStopsAroundGeographicalPoint = "/transport/busemtmad/stops/arroundxy/<longitude>/<latitude>/<radius>/"
	endpointTimeArrivalBus               = "/transport/busemtmad/stops/<stopId>/arrives/<lineArrive>/" //POST METHOD
	endpointCalendar                     = "/transport/busemtmad/calendar/<startdate>/<enddate>/"
	endpointIncidents                    = "/transport/busemtmad/lines/incidents/<lineid>/"
	endpointInfoLineDetail               = "/transport/busemtmad/lines/<lineId>/info/<dateref>/"
	endpointInfoLineGeneral              = "/transport/busemtmad/lines/info/<dateref>/"
	endpointInfoStopsGeneral             = "/transport/busemtmad/stops/list/" //POST METHOD
	endpointLineStops                    = "/transport/busemtmad/lines/<lineId>/stops/<direction>/"
	endpointListStops                    = "/transport/busemtmad/stops/list/" //POST METHOD
	endpointOperationGroups              = "/transport/busemtmad/lines/groups/"
	endpointRouteOfLine                  = "/transport/busemtmad/lines/<labelId>/route/"
	endpointStopsAroundPlaces            = "/transport/busemtmad/stops/arroundstreet/<namePlace>/<number_or_zero_street_number>/<radius>/"
	endpointTimeTableStartStop           = "/transport/busemtmad/lines/<lineId>/timetable/"
	endpointTimeTableTrips               = "/transport/busemtmad/lines/<lineId>/trips/<dateRef>/"
)

/*
GetInfoFromPlatform function fills a give structure with the data from the EMT Rest API
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	serviceEndpoint -> endpoint to queri (i.e. "/transport/busemtmad/stops/arroundstop/<stopId>/<radius>/")
	dataStructure -> struct where all the json info will be Unmarshalled
*/
func getInfoFromPlatform(api goemt.IAPI, serviceEndpoint string, dataStructure goemt.DataInterface) (err error) {
	if api == nil {
		return fmt.Errorf("api being used is nil")
	}
	data, err := api.Get(serviceEndpoint)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &dataStructure)
	if err != nil {
		return err
	}
	if dataStructure.GetAPIReturnCode() != "00" {
		return fmt.Errorf(dataStructure.GetAPIReturnDescription())
	}
	return nil
}

func postInfoToPlatform(api goemt.IAPI, serviceEndpoint string, postDataStructure interface{}, returnDataStructure goemt.DataInterface) (err error) {
	if api == nil {
		return fmt.Errorf("api being used is nil")
	}
	data, err := api.Post(serviceEndpoint, postDataStructure)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &returnDataStructure)
	if err != nil {
		return err
	}
	if returnDataStructure.GetAPIReturnCode() != "00" {
		return fmt.Errorf(returnDataStructure.GetAPIReturnDescription())
	}
	return nil
}
