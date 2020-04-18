package busemtmad

import (
	"encoding/json"
	"fmt"
	"github.com/mikeletux/goemt"
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
