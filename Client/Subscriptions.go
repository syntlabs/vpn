package main

import (
	main2 "awesomeProject/Client/Network"
	"encoding/json"
	"fmt"
	"time"
)

const SUBS_PATH = ""
const TIME_POINTER = ""

func accessIsAllowed(d NetworkDaemon, client Client.UserClient) bool {

	uniqSubPath := fmt.Sprintf(
		"%s%s%s%s%s", TRANSFER_PROTOCOL, MAIN_ROUTE, MAIN_PORT, SUBS_PATH, client.fingerprint,
	)

	var payload map[string]string
	var headers map[string]string

	response, reserr := main2.req(d, uniqSubPath, headers, payload)
	defer response.Body.Close()

	if reserr != nil {
		raise(ErrorsGlobal.network.brokenResponse)
	}

	rbytes, errjson := json.Marshal(response.Body)

	if errjson != nil {
		raise(ErrorsGlobal.network.jsonConversionError)
	}

	var data map[string]any

	err := json.Unmarshal(rbytes, &data)

	var poi_ty bool

	switch data[TIME_POINTER].(type) {
	case float64:
		poi_ty = true
	default:
		poi_ty = false
		raise(ErrorsGlobal.value.timepointer)
	}

	if poi_ty {
		if TimeEnd(float64(time.Now().Minute()), data[TIME_POINTER]) == true {
			client.subscription := 0
		}
	}
}

func TimeEnd(start float64, end float64) bool {

	if (end/60)-(start/60) <= 0 {
		return false
	} else {
		return true
	}
}
