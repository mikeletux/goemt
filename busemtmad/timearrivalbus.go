package busemtmad

import (
	"fmt"
	"github.com/mikeletux/goemt"
	"strings"
)

/*
PostInfoTimeArrival struct implements the post information needed to send to the server
*/
type PostInfoTimeArrival struct {
	//Could be EN for english or ES for spanish
	CultureInfo string `json:"cultureInfo"`

	//Y(es) for getting name stop or N(ot)
	TextStopRequired string `json:"Text_StopRequired_YN"`

	//Y(es) for data estimations to arrival Bus or N(ot)
	TextEstimationsRequired string `json:"Text_EstimationsRequired_YN"`

	//Y(es) for getting incidents related to lines in this stop s or N(ot)
	TextIncidencesRequired string `json:"Text_IncidencesRequired_YN"`

	//????????", year-month-day to reference of incidents
	DateTimeReferencedIncidencies string `json:"DateTime_Referenced_Incidencies_YYYYMMDD"`
}

/*
TimeArrivalBus struct is where all information regarding Time Arrival from Buses will be stored
*/
type TimeArrivalBus struct {
	Arrive []struct {
		DistanceBus     int    `json:"DistanceBus"`
		Bus             int    `json:"bus"`
		Destination     string `json:"destination"`
		Deviation       int    `json:"deviation"`
		Stop            string `json:"stop"`
		PositionTypeBus string `json:"positionTypeBus"`
		IsHead          string `json:"isHead"`
		Line            string `json:"line"`
		EstimateArrive  int    `json:"estimateArrive"`
		Geometry        struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"Arrive"`
	StopInfo []struct {
		Lines []struct {
			Label            string `json:"label"`
			Line             string `json:"line"`
			NameA            string `json:"nameA"`
			NameB            string `json:"nameB"`
			MetersFromHeader int    `json:"metersFromHeader"`
			To               string `json:"to"`
		} `json:"lines"`
		StopID    string `json:"stopId"`
		StopName  string `json:"stopName"`
		Direction string `json:"Direction"`
		Geometry  struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		//ExtraInfo -> Not clear from documentation. Need to figure out its structure
		Incident struct {
			ListaIncident struct {
				Data []struct {
					Title       string `json:"title"`
					GUID        string `json:"guid"`
					Description string `json:"description"`
					PubDate     string `json:"pubDate"`
					RssFrom     string `json:"rssFrom"`
					RssTo       string `json:"rssTo"`
					Cause       string `json:"cause"`
					Effect      string `json:"effect"`
					MoreInfo    struct {
						URL    string `json:"@url"`
						Type   string `json:"@type"`
						Length string `json:"@length"`
					} `json:"moreInfo"`
				} `json:"data"`
			} `json:"ListaIncident"`
		} `json:"Incident"`
	} `json:"StopLines"`
}

type timeArrivalBusAux struct {
	goemt.Common
	Data []TimeArrivalBus `json:"data"`
}

/*
GetTimeArrivalBus returns all datails regarding time arrivals in buses
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	stopID -> stop id
	lineArrive -> returns info only from that specific line. If 0, returns data from all lines that stops there
	postData -> struct used in the body of the POST HTTP method
Returns:
	returnedData -> struct returned from EMT Rest API ready to consume
	err -> if an error occurs, this var will be different from zero
*/
func GetTimeArrivalBus(api goemt.IAPI, stopID int, lineArrive int, postData PostInfoTimeArrival) (returnedData []TimeArrivalBus, err error) {
	//ADD ERROR HANDLING postData STRUCT!!!!!
	var data timeArrivalBusAux
	if stopID == 0 {
		return returnedData, fmt.Errorf("stopID and lineArrive must be different from 0")
	}
	finalServiceEnpoint := strings.ReplaceAll(endpointTimeArrivalBus, "<stopId>", fmt.Sprintf("%d", stopID))
	if lineArrive == 0 {
		finalServiceEnpoint = strings.ReplaceAll(finalServiceEnpoint, "<lineArrive>/", fmt.Sprintf(""))
	} else {
		finalServiceEnpoint = strings.ReplaceAll(finalServiceEnpoint, "<lineArrive>", fmt.Sprintf("%d", lineArrive))
	}
	err = postInfoToPlatform(api, finalServiceEnpoint, postData, &data)
	if err != nil {
		return returnedData, err
	}
	return data.Data, nil
}
