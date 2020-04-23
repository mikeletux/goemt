package busemtmad

import (
	"fmt"
	"github.com/mikeletux/goemt"
	"strings"
)

/*
InfolineDetail Struct where information regarding detail of one EMT line will be stored
*/
type InfolineDetail struct {
	//date reference of line data
	DateRef string `json:"dateRef"`

	//Name of Header A
	NameA string `json:"nameA"`

	//Name of Header B
	NameB string `json:"nameB"`

	//Group of data
	Line string `json:"line"`

	//Public line code
	Label string `json:"label"`

	//Timetable of line in all typeDays
	TimeTable []struct {
		Direction1 struct {
			MaximumFrequency string `json:"MaximumFrequency"`
			MinimunFrequency string `json:"MinimunFrequency"`
			FrequencyText    string `json:"FrequencyText"`
			StopTime         string `json:"StopTime"`
			StartTime        string `json:"StartTime"`
		} `json:"Direction1"`
		Direction2 struct {
			MaximumFrequency string `json:"MaximumFrequency"`
			MinimunFrequency string `json:"MinimunFrequency"`
			FrequencyText    string `json:"FrequencyText"`
			StopTime         string `json:"StopTime"`
			StartTime        string `json:"StartTime"`
		} `json:"Direction2"`
		IDDayType string `json:"idDayType"`
	} `json:"timeTable"`
}

type auxInfolineDetail struct {
	goemt.Common
	Data []InfolineDetail `json:"data"`
}

/*
GetInfolineDetail gets detailed info regarding a EMT line
Params:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	lineID -> Line (or label) for getting data
	dateref -> date reference on YYYYMMDD format
Return:
	infolineDetails -> slice with the wanted info
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetInfolineDetail(api goemt.IAPI, lineID, dateref int) (infolineDetails []InfolineDetail, err error) {
	if lineID == 0 || dateref == 0 {
		return infolineDetails, fmt.Errorf("lineID and dataref must be different from 0")
	}
	finalServiceEnpoint := strings.ReplaceAll(endpointInfoLineDetail, "<lineId>", fmt.Sprintf("%v", lineID))
	finalServiceEnpoint = strings.ReplaceAll(finalServiceEnpoint, "<dateref>", fmt.Sprintf("%v", dateref))
	var data auxInfolineDetail
	err = getInfoFromPlatform(api, finalServiceEnpoint, &data)
	if err != nil {
		return infolineDetails, err
	}
	return data.Data, nil
}
