package main

import (
	"fmt"
	"github.com/mikeletux/goemt"
	"github.com/mikeletux/goemt/busemtmad"
	"os"
)

func main() {

	const stopToQuery = 2537 //Number of the stop to be query

	//By setting XClientID and PassKey on ClientConfig struct, we enforce the protected mode when log in
	configProtected := goemt.ClientConfig{
		Enpoint:   os.Getenv("EMT_ENDPOINT"), // in this case will be https://openapi.emtmadrid.es/v2
		XClientID: os.Getenv("EMT_XCLIENTID"),
		PassKey:   os.Getenv("EMT_PASSKEY"),
	}

	//We create the connect struct that we'll use to query the platform
	c, err := goemt.Connect(configProtected)
	if err != nil {
		panic(err)
	}
	defer c.Logout()

	//GetTimeArrivalBus func needs a struct to use it when post
	postData := busemtmad.PostInfoTimeArrival{
		CultureInfo:             "ES",
		TextStopRequired:        "Y",
		TextEstimationsRequired: "Y",
		TextIncidencesRequired:  "N",
	}

	busTimes, err := busemtmad.GetTimeArrivalBus(c, stopToQuery, 0, postData)

	for _, v := range busTimes {
		for _, arrive := range v.Arrive {
			fmt.Printf("Bus number %v will arrive in %v seconds and heads to %v\n", arrive.Line, arrive.EstimateArrive, arrive.Destination)
		}

	}

}
