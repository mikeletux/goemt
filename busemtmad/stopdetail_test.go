package busemtmad

import (
	"github.com/mikeletux/goemt"
	"os"
	"testing"
)

func TestGetStopDetail(t *testing.T) {
	type TestDataItem struct {
		testNo   int
		api      goemt.IAPI
		stopID   int
		hasError bool
	}

	configProtected := goemt.ClientConfig{
		Enpoint:   os.Getenv("EMT_ENDPOINT"),
		Email:     os.Getenv("EMT_EMAIL"),
		Password:  os.Getenv("EMT_PASSWORD"),
		XAPIKey:   os.Getenv("EMT_XAPIKEY"),
		XClientID: os.Getenv("EMT_XCLIENTID"),
	}

	c, err := goemt.Connect(configProtected)
	if err != nil {
		t.Error(err)
		return
	}
	defer c.Logout()

	//NEED TO USE FAKE API TO TEST MORE!!!
	testData := []TestDataItem{
		{1, c, 2537, false},
	}

	for _, v := range testData {
		data, err := GetStopDetail(v.api, v.stopID)
		if v.hasError {
			// DO CHECKS!
		} else {
			if err != nil {
				t.Errorf("FAIL: Test no %d was supposed to succeed. Error %s", v.testNo, err)
			} else {
				t.Logf("SUCCESS: Test was successful. Data %v", data)
			}
		}
	}
}
