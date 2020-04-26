package busemtmad

import (
	"github.com/mikeletux/goemt"
	"os"
	"testing"
)

func TestGetListStops(t *testing.T) {
	type TestDataItem struct {
		testNo   int
		api      goemt.IAPI
		postData PostListStops
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
		{1, c, PostListStops{2537}, false},
		// {2, c, nil, false},
		{3, c, PostListStops{2537, 72}, false},
		//{4, c, PostListStops{"ccc"}, true},
		{5, nil, nil, true},
	}

	for _, v := range testData {
		//Execute func
		data, err := GetListStops(v.api, v.postData)
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
