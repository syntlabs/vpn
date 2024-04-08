package main

import (
	"Client/Network"
	"encoding/json"
	"fmt"
	"time"
)

const SUBS_PATH = ""

func subscriptionTime(d NetworkDaemon, client Client.UserClient) bool {

	uniqSubPath := fmt.Sprintf(
		"%s%s%s%s%s", TRANSFER_PROTOCOL, MAIN_ROUTE, MAIN_PORT, SUBS_PATH, client.fingerprint,
	)

	//var payload map[string]string
	//var headers map[string]string

	response, reserr := req(uniqSubPath, nil, nil)
	defer response.Body.Close()

	if reserr != nil {
		raise(Err.net.BrokenResponse)
	}

	rbytes, errjson := json.Marshal(response.Body)

	if errjson != nil {
		raise(Err.net.JsonConversion)
	}

	var data map[string]interface{}

	err := json.Unmarshal(rbytes, &data)

	if poi_ty {
		if time.Now().Minute() <= data["subscription_time_left"]) == true {
			client.subscription := 0
			viewSubJustEnded()
		}
	}
}
