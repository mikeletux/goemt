package busemtmad

import (
	"os"
	"testing"

	"github.com/mikeletux/goemt"
)

func TestGetStopsAroundPlaces(t *testing.T) {
	type TestDataItem struct {
		testNo    int
		api       goemt.IAPI
		namePlace string
		number    int
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
		{1, c, "paseo de santa maria de la cabeza", 112, 300, false},
		{2, c, "calle delicias", 0, 100, false},
		{3, nil, "paseo de santa maria de la cabeza", 112, 100, true},
	}

	for _, v := range testData {
		//Execute func
		data, err := GetStopsAroundPlaces(v.api, v.namePlace, v.number, v.radius)
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
