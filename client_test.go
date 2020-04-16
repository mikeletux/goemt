package goemt

import (
	"os"
	"testing"
)

type TestDataItem struct {
	testNo   int
	config   ClientConfig
	hasError bool
}

//test login function

func TestLogin(t *testing.T) {
	configBasic := ClientConfig{
		Enpoint:  os.Getenv("EMT_ENDPOINT"),
		Email:    os.Getenv("EMT_EMAIL"),
		Password: os.Getenv("EMT_PASSWORD"),
	}

	configAdvanced := ClientConfig{
		Enpoint:   os.Getenv("EMT_ENDPOINT"),
		Email:     os.Getenv("EMT_EMAIL"),
		Password:  os.Getenv("EMT_PASSWORD"),
		XAPIKey:   os.Getenv("EMT_XAPIKEY"),
		XClientID: os.Getenv("EMT_XCLIENTID"),
	}

	configProtected := ClientConfig{
		Enpoint:   os.Getenv("EMT_ENDPOINT"),
		XClientID: os.Getenv("EMT_XCLIENTID"),
		PassKey:   os.Getenv("EMT_PASSKEY"),
	}

	configOverloaderd := ClientConfig{
		Enpoint:   os.Getenv("EMT_ENDPOINT"),
		XClientID: os.Getenv("EMT_XCLIENTID"),
		PassKey:   os.Getenv("EMT_PASSKEY"),
		Email:     os.Getenv("EMT_EMAIL"),
		Password:  os.Getenv("EMT_PASSWORD"),
		XAPIKey:   os.Getenv("EMT_XAPIKEY"),
	}

	// input-result data items
	dataItems := []TestDataItem{
		{1, configBasic, false},
		{2, configAdvanced, false},
		{3, configProtected, false},
		{4, configOverloaderd, true},
		{5, ClientConfig{}, true},
	}

	for _, v := range dataItems {
		api, err := Connect(v.config)
		if v.hasError {
			//We expect to have an error
			if err == nil {
				t.Errorf("FAILED: Connect() was supposed to failed but suceed in test number %d.", v.testNo)
			} else {
				t.Logf("SUCCEED: Connect() fail in test number %d with error %v", v.testNo, err)
			}
		} else {
			//We expect to succeed
			if err != nil {
				t.Errorf("FAILED: Connect() was supposed to succeed but failed in test number %d with error %v", v.testNo, err)
			} else {
				if len(api.auth) > 0 {
					t.Logf("SUCCEED: Connect() succeed in test number %d with token %v", v.testNo, api.auth)
					api.Logout()
				} else {
					t.Errorf("FAILED: Connect() was supposed to succeed but failed in test number %d", v.testNo)
				}
			}

		}
	}

}

func TestLogout(t *testing.T) {
	type TestDataItem struct {
		testNo   int
		config   ClientConfig
		hasError bool
	}

	configBasic := ClientConfig{
		Enpoint:  os.Getenv("EMT_ENDPOINT"),
		Email:    os.Getenv("EMT_EMAIL"),
		Password: os.Getenv("EMT_PASSWORD"),
	}

	configWrong := ClientConfig{
		Enpoint:  os.Getenv("EMT_ENDPOINT"),
		Email:    os.Getenv("EMT_EMAIL"),
		Password: "wrong_pass",
	}

	configEmpty := ClientConfig{}

	dataItems := []TestDataItem{
		{1, configBasic, false},
		{2, configWrong, true},
		{3, configEmpty, true},
	}

	for _, v := range dataItems {
		api, err := Connect(v.config)
		if v.hasError {
			//Suppose to fail
			if err != nil {
				t.Logf("SUCCEED: Logout() was supposed to fail in test %d Error: %v", v.testNo, err)
			}
		} else {
			//Suppose to suceed
			err = api.Logout()
			if err != nil {
				t.Errorf("FAILED: Logout() was suppose to succeed in test %d. Error: %v", v.testNo, err)
			} else {
				t.Logf("SUCCEED: Logout() suceed in test %d", v.testNo)
			}
		}
	}
}

func TestGet(t *testing.T) {
	type TestDataItem struct {
		testNo   int
		url      string
		hasError bool
	}

	dataTest := []TestDataItem{
		{1, "/mobilitylabs/user/whoami/", false},
		{2, "/mobilitylabs/user/fakeurl/", true},
	}

	configBasic := ClientConfig{
		Enpoint:  os.Getenv("EMT_ENDPOINT"),
		Email:    os.Getenv("EMT_EMAIL"),
		Password: os.Getenv("EMT_PASSWORD"),
	}

	c, err := Connect(configBasic)
	if err != nil {
		t.Errorf("couldn't connect to the service")
		return
	}
	defer c.Logout()
	for _, v := range dataTest {
		data, err := c.Get(v.url)
		if v.hasError {
			if err != nil {
				t.Logf("SUCCEED: Get() was supposed to fail in test %d Error: %v", v.testNo, err)
			} else {
				t.Errorf("FAILED: Get() succeeed, was supposed to fail in test %d Error: %v", v.testNo, err)
			}
		} else {
			if err != nil {
				t.Errorf("FAILED: Get() failed, was supposed to succeed in test %d Error: %v", v.testNo, err)
			} else {
				t.Logf("SUCCEED: Get() succeed in test %d with data %v", v.testNo, data)
			}
		}
	}
}

//NEED TO TEST POST!!!!!!
