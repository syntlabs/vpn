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
	prom     string
	timeLeft time.Time
}

const prSize = 8
const checkdelay = 60 // SECONDS
const prefix = "Synt"

var port = os.Getenv("promPort")

var hashVaultPath = ""
var salt string = os.Getenv("Salt")

func promgen(timeEnd string) string {

	prom := randStr(prSize)
	hash := sha256.Sum256(append([]byte(prom), []byte(salt)...))
	hhex := hex.EncodeToString(hash[:])

	saveHash(hhex)

	switch timeEnd {
	case "1 month":
		return fmt.Sprintf("%s_1m_%s", prefix, prom)
	case "3 months":
		return fmt.Sprintf("%s_3m_%s", prefix, prom)
	case "6 month":
		return fmt.Sprintf("%s_6m_%s", prefix, prom)
	case "1 year":
		return fmt.Sprintf("%s_1y_%s", prefix, prom)
	default:
		return fmt.Sprintf("%s%s", prefix, prom)
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
			} else {
				v <- false
			}
		}
		close(v)
	}()

	return <-v
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

	if r.Method != "GET" {

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//cookie, err := r.Cookie(fmt.Sprintf("session_%s", data))

		if err != nil {
			http.Error(w, "Session not found", http.StatusUnauthorized)
			return
		} else {
			fmt.Print("Session is found")
		}
	}

	if !time.Time.Equal(data.timeLeft, time.Now()) {

		ch := make(chan struct{})

		go func() {
			for time.Now().Before(data.timeLeft) {
				time.Sleep(checkdelay * time.Second)
			}
			close(ch)
		}()

		<-ch

		http.SetCookie(w, &http.Cookie{
			Name:   "session_id",
			Value:  "",
			MaxAge: -1,
		})
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
