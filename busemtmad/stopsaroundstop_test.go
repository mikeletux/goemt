package busemtmad

import (
	"os"
	"testing"

	"github.com/mikeletux/goemt"
)

func TestGetStopsAroundStop(t *testing.T) {
	type TestDataItem struct {
		testNo   int
		api      goemt.IAPI
		stopID   int
		radius   int
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

	testData := []TestDataItem{
		{1, c, 2537, 500, false},
		{2, c, 9999999, 1243123, true},
		{3, nil, 123123, 12, true},
	}

	for _, v := range testData {
		data, err := GetStopsAroundStop(v.api, v.stopID, v.radius)
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

func TestGetStopsAroundGeographicalPoint(t *testing.T) {
	type TestDataItem struct {
		testNo    int
		api       goemt.IAPI
		longitude float64
		latitude  float64
		radius    int
		hasError  bool
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

	testData := []TestDataItem{
		{1, c, -3.703525, 40.417013, 500, false},
		{2, c, -1.926158, 41.814420, 50, false}, //Coordinates from Soria LOL
		{3, nil, -3.703525, 40.417013, 12, true},
		{4, c, 0.0, 0.0, 500, true},
	}

	for _, v := range testData {
		data, err := GetStopsAroundGeographicalPoint(v.api, v.longitude, v.latitude, v.radius)
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
