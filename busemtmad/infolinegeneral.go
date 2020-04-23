package busemtmad

import (
	"fmt"
	"github.com/mikeletux/goemt"
	"strings"
)

/*
InfolineGeneral struct will store  the list of lines actives
*/
type InfolineGeneral struct {
	//Date init of current line configuration
	StartDate string `json:"startDate"`

	//Group that belongs each line.
	Group string `json:"group"`

	//Name of Header A
	NameA string `json:"nameA"`

	//Name of Header B
	NameB string `json:"nameB"`

	//Date end of current line configuration
	EndDate string `json:"endDate"`

	//Public code of line
	Label string `json:"label"`

	//Line of EMT.
	Line string `json:"line"`
}

type auxInfolineGeneral struct {
	goemt.Common
	Data []InfolineGeneral `json:"data"`
}

/*
GetInfolineGeneral func will return the list of lines actives in the reference date
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	dateref -> date reference on YYYYMMDD format
Returns:
	infolinesGeneral -> slice with all requested data
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetInfolineGeneral(api goemt.IAPI, dateref int) (infolinesGeneral []InfolineGeneral, err error) {
	if dateref == 0 {
		return infolinesGeneral, fmt.Errorf("dateref must be different from 0 (format YYYYMMDD)")
	}
	finalServiceEnpoint := strings.ReplaceAll(endpointInfoLineGeneral, "<dateref>", fmt.Sprintf("%v", dateref))
	var data auxInfolineGeneral
	err = getInfoFromPlatform(api, finalServiceEnpoint, &data)
	if err != nil {
		return infolinesGeneral, err
	}
	return infolinesGeneral, nil
}
