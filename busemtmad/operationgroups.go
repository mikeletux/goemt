package busemtmad

import (
	"github.com/mikeletux/goemt"
)

/*
OperationGroup struct is where the list of groups of operations related to EMT lines will be stored
*/
type OperationGroup struct {
	//Group that belongs each line.
	Group string `json:"group"`

	//Simple description
	Description string `json:"description"`

	//Subgroup of line
	SubGroup string `json:"subGroup"`
}

type auxOperationGroup struct {
	goemt.Common
	Data []OperationGroup `json:"data"`
}

/*
GetOperationGroups func v returns the list of groups of operations related to EMT lines
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
returns:
	operationGroups -> slice with all requested data
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetOperationGroups(api goemt.IAPI) (operationGroups []OperationGroup, err error) {
	var data auxOperationGroup
	err = getInfoFromPlatform(api, endpointOperationGroups, &data)
	if err != nil {
		return operationGroups, err
	}
	return data.Data, nil
}
