package busemtmad

import (
	"os"
	"testing"

	"github.com/mikeletux/goemt"
)

func TestGetTravelPlan(t *testing.T) {
	type TestDataItem struct {
		testNo   int
		api      goemt.IAPI
		postData PostTravelPlan
		hasError bool
	}

	configProtected := goemt.ClientConfig{
		Enpoint:   os.Getenv("EMT_ENDPOINT"),
		XClientID: os.Getenv("EMT_XCLIENTID"),
		PassKey:   os.Getenv("EMT_PASSKEY"),
	}

	c, err := goemt.Connect(configProtected)
	if err != nil {
		t.Error(err)
		return
	}
	defer c.Logout()

	postData := PostTravelPlan{
		RouteType:       "P",
		Itinerary:       true,
		CoordinateXFrom: -3.701077,
		CoordinateYFrom: 40.4469,
		CoordinateXTo:   -3.674902,
		CoordinateYTo:   40.400149,
		OriginName:      "Calle Maudes 6",
		DestinationName: "Calle Cerro de la Plata 4",
		Day:             2,
		Month:           5,
		Year:            2020,
		Hour:            18,
		Minute:          18,
		Culture:         "es",
		AllowBus:        true,
		AllowBike:       false,
	}

	testData := []TestDataItem{
		{1, c, postData, false},
		{2, nil, postData, true},
	}

	for _, v := range testData {
		//Execute func
		data, err := GetTravelPlan(v.api, v.postData)
		if v.hasError {
			if err != nil {
				t.Logf("SUCCESS: Test no %d has supposed to fail. Error: %v", v.testNo, err)
			} else {
				t.Errorf("FAIL: Test no %d was supposed to fail and it succeed.", v.testNo)
			}
		} else {
			if err != nil {
				t.Errorf("FAIL: Test no %d was supposed to succeed but failed. Error %v", v.testNo, err)
			} else {
				t.Logf("SUCCESS: Test no %d succeed with data %v", v.testNo, data)
			}
		}
	}
}
