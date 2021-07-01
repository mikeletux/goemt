package busemtmad

import (
	"os"
	"testing"

	"github.com/mikeletux/goemt"
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
		XClientID: os.Getenv("EMT_XCLIENTID"),
		PassKey:   os.Getenv("EMT_PASSKEY"),
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
		{2, c, 999999999, true},
		{3, nil, 2537, true},
	}

	for _, v := range testData {
		data, err := GetStopDetail(v.api, v.stopID)
		if v.hasError {
			if err == nil {
				t.Errorf("FAIL: Test was supposed to fail but suceed in test no %d. Error %s", v.testNo, err)
			} else {
				t.Logf("SUCESS: Test was supposed to fail. Test no %d", v.testNo)
			}
		} else {
			if err != nil {
				t.Errorf("FAIL: Test no %d was supposed to succeed. Error %s", v.testNo, err)
			} else {
				t.Logf("SUCCESS: Test was successful. Data %v", data)
			}
		}
	}
}
