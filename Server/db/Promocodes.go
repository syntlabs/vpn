package Promocodes

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type userStruct struct {
	prom  string `json:"prom"`
	sesid string `json:"sesid"`
	paid  bool   `json:"paid"`
}

const prSize = 8
const checkdelay = 60 // SECONDS
const prefix = "Synt"

var hashVaultPath = ""
var port = os.Getenv("promPort")
var salt string = os.Getenv("Salt")

func promgen(timeEnd string) string {

	prom := randStr(prSize)
	hash := sha256.Sum256(append([]byte(prom), []byte(salt)...))
	hhex := hex.EncodeToString(hash[:])

	saveHash(hhex)

	switch timeEnd {
	case "1 month":
		return fmt.Sprintf("%s_%s_1", prefix, prom)
	case "3 months":
		return fmt.Sprintf("%s_%s_3", prefix, prom)
	case "6 month":
		return fmt.Sprintf("%s_%s_6", prefix, prom)
	default:
		return fmt.Sprintf("")
	}
}

func isValid(prom string) bool {

	hash := sha256.Sum256(append([]byte(prom), []byte(salt)...))

	file, err := os.Open(hashVaultPath)
	if err != nil {
		panic(runtime.PanicNilError{})
	}

	sc := bufio.NewScanner(file)

	v := make(chan bool)

	go func() {
		for sc.Scan() {
			line := sc.Text()
			if line == hex.EncodeToString(hash[:]) {
				v <- true
				close(v)
				return
			}
			v <- false
		}
		close(v)
	}()

	select {
	case result, ok := <-v:
		if !ok {
			return false
		}
		return result
	}
}

func main() {

	http.HandleFunc("/", netHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func netHandler(w http.ResponseWriter, r *http.Request) {

	var data userStruct

	if r.Method == "GET" {

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if isValid(data.prom) && !data.paid {

			data.paid = true

			timer := newTimerOnce()
			timer.run(time.Hour*24*time.Duration(int(data.prom[len(data.prom)-1])), func() {
				//
			})
		}
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}

func saveHash(h string) {

	file, err := os.Open(hashVaultPath)
	defer file.Close()

	if err != nil {
		panic(runtime.PanicNilError{})
	}

	_, err = file.WriteString(h)

	if err != nil {
		panic(runtime.PanicNilError{})
	}
}

func randStr(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)

	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return sb.String()
}
