package busemtmad

import (
	"os"
	"testing"

	"github.com/mikeletux/goemt"
)

func TestGetTimeArrivalBus(t *testing.T) {
	type TestDataItem struct {
		testNo     int
		api        goemt.IAPI
		stopID     int
		lineArrive int
		postData   PostInfoTimeArrival
		hasError   bool
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

	postData := PostInfoTimeArrival{
		CultureInfo:             "ES",
		TextStopRequired:        "Y",
		TextEstimationsRequired: "Y",
		TextIncidencesRequired:  "Y",
	}

	testData := []TestDataItem{
		{1, c, 2537, 0, postData, false},
		{2, c, 2537, 60, postData, false}, //Line 60 stops at 2537 stop
		{3, c, 0, 0, postData, true},
		{4, nil, 2537, 0, postData, true},
	}

	for _, v := range testData {
		data, err := GetTimeArrivalBus(v.api, v.stopID, v.lineArrive, v.postData)
		if v.hasError {
			if err != nil {
				t.Logf("SUCESS: Test no %d was supposed to fail and failed with error %v", v.testNo, err)
			} else {
				t.Errorf("FAIL: Test no %d was supposed to fail and succeed", v.testNo)
			}
		} else {
			if err != nil {
				t.Errorf("FAIL: Test no %d was supposed to succeed. Error %v", v.testNo, err)
			} else {
				t.Logf("SUCCESS: Test no %d was successful. Data retreived %v", v.testNo, data)
			}
		}
	}

}
