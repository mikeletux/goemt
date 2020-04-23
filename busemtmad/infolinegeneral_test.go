package busemtmad

import (
	"github.com/mikeletux/goemt"
	"os"
	"testing"
)

func TestGetInfolineGeneral(t *testing.T) {
	type TestDataItem struct {
		testNo   int
		api      goemt.IAPI
		dateref  int
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

	testData := []TestDataItem{
		{1, c, 20200501, false},
		{2, c, 0, true},
		{3, nil, 20200501, true},
	}

	for _, v := range testData {
		//Execute func
		data, err := GetInfolineGeneral(v.api, v.dateref)
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
