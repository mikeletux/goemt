package busemtmad

import (
	"github.com/mikeletux/goemt"
	"os"
	"testing"
)

func TestGetCalendar(t *testing.T) {
	type TestDataItem struct {
		testNo    int
		api       goemt.IAPI
		startDate int
		endDate   int
		hasError  bool
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

	testData := []TestDataItem{
		{1, c, 20200601, 20200607, false},
		{2, c, 123, 123, true},
		{3, c, 20200605, 20200601, true}, //start date is after the enddate
	}

	for _, v := range testData {
		//Execute func
		data, err := GetCalendar(c, v.startDate, v.endDate)
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
