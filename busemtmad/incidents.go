package busemtmad

import (
	"fmt"
	"github.com/mikeletux/goemt"
	"strings"
)

/*
Incidents struct is where all data from EMT Rest API regarding incidents will be stored
*/
type Incidents struct {
	LastBuildDate string `json:"lastBuildDate"`
	Description   string `json:"description"`
	Copyright     string `json:"copyright"`

	//Title of feed data
	Title string `json:"title"`

	//Ignored
	Generator string `json:"generator"`

	//Ignored
	Language string `json:"language"`

	//feed location
	Link      string `json:"link"`
	TTL       string `json:"ttl"`
	WebMaster string `json:"webMaster"`

	//General image of feed and data asociated
	Image struct {
		URL         string `json:"url"`
		Link        string `json:"link"`
		Description string `json:"description"`
		Title       string `json:"title"`
	} `json:"image"`
	Item []struct {
		//In GoogleTransit, incident of this context (effect)
		GoogleTransitEffect string `json:"GoogleTransitEffect"`

		//Start time estimated of incident
		RssAfectaDesde string `json:"rssAfectaDesde"`

		//Description of what happened
		Description string `json:"description"`

		//date of publication
		PubDate string `json:"pubDate"`

		//author-department of publish
		Author string `json:"author"`

		//Title of what happened
		Title string `json:"title"`

		//end of incident (estimate)
		RssAfectaHasta string `json:"rssAfectaHasta"`

		//feed location
		Link string `json:"link"`

		//publisher of incident
		Rsspublisher string `json:"rsspublisher"`

		//end of publish
		Rssfinaldate string `json:"rssfinaldate"`

		//unique id for reference each incident
		GUID string `json:"guid"`

		//In GoogleTransit, incident of this context (Cause)
		GoogleTransitCause string `json:"GoogleTransitCause"`

		//lines related to incident belongs to "item"
		Category []string `json:"category"`

		//Ignore
		Media struct {
			Media struct {
				AltTitle       string `json:"@AltTitle"`
				AltDescription string `json:"@AltDescription"`
				Type           string `json:"@Type"`
				Format         string `json:"@Format"`
			} `json:"MEDIA"`
		} `json:"MEDIA"`

		//data link in external document
		Enclosure struct {
			Length string `json:"@length"`
			Type   string `json:"@type"`
			URL    string `json:"@url"`
		} `json:"enclosure"`
	} `json:"item"`
}

type auxIncidents struct {
	goemt.Common
	Data []Incidents
}

/*
GetIncidents function returns all incidents related to an specific bus line
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	lineId -> line from where you want to get the incidents from
Returns:
	incidents -> slice of associated incidents
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetIncidents(api goemt.IAPI, lineID int) (incidents []Incidents, err error) {
	var data auxIncidents
	if lineID == 0 {
		return incidents, fmt.Errorf("lineId must be different from 0")
	}
	finalServiceEnpoint := strings.ReplaceAll(endpointIncidents, "<lineid>", fmt.Sprintf("%v", lineID))
	err = getInfoFromPlatform(api, finalServiceEnpoint, &data)
	if err != nil {
		return incidents, err
	}
	return data.Data, nil
}
