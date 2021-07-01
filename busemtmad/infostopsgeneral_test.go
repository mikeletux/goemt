package busemtmad

import (
	"os"
	"testing"

	"github.com/mikeletux/goemt"
)

func TestGetInfostopsGeneral(t *testing.T) {
	type TestDataItem struct {
		testNo   int
		api      goemt.IAPI
		postData PostInfoStops
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

	postDataTest := []int{62, 1234, 2537}

	testData := []TestDataItem{
		{1, c, postDataTest, false},
		{2, c, nil, false},
		{3, nil, postDataTest, true},
	}

	for _, v := range testData {
		//Execute func
		data, err := GetInfostopsGeneral(v.api, v.postData)
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
