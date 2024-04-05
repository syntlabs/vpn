package Antares_vpn

import (
	"fmt"
	"net/http"
	"time"
)

func req(
	daemon NetworkDaemon, url string, headers, payload map[string]string,
) (*http.Response, error) {

	daemonClient := &http.Client{}
	request, req_err := http.NewRequest(daemon.method, url, nil)

	if req_err != nil {
		raise(ErrorsGlobal.network.brokenRequest)
	}

	if headers != nil {
		for key, val := range headers {
			request.Header.Add(key, val)
		}
	}
	if payload != nil {
		for key, val := range payload {
			request.PostForm.Add(key, val)
		}
	}

	response, res_err := daemonClient.Do(request)

	if res_err != nil {
		raise(ErrorsGlobal.network.brokenResponse)
	}

	return response, nil
}

func serviceAvailable(d NetworkDaemon) bool {

	go func() (avail bool, code int) {
		for {
			resp, _ := req(d, MAIN_URL_PATH, nil, nil)

			time.Sleep(d.updateTime * time.Second)
			fmt.Print("Network daemon is running...")

			if resp.StatusCode != 200 {
				avail = false
			}

			return avail, resp.StatusCode
		}
	}()
}
