package db

import (
	"bytes"
	"crypto/sha256"
	"log"
	"math/rand"
	"os"
	"runtime"
	"slices"
	"strings"
	"sync"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)
const prSize = 8
const nPr = 100000
const threads = 4
const grandMode = "one"

var salt string = os.Getenv("Salt")

func promgen(promSize int) (string, [sha256.Size]byte) {

	var ps, pb = RandStringBytesMaskImprSrcSB(promSize)

	return ps, sha256.Sum256(append(pb, []byte(salt)...))
}

func isValid(prom string, salt string) {

	checkBytes := bytes.Join([][]byte{[]byte(prom), []byte(salt)}, nil)

	checkhash := sha256.Sum256(checkBytes)
}

func main() {

	var data *any

	switch grandMode {
	case "one":
		*data = [2]interface{}{promgen(prSize)}
	default:
		*data = []interface{}{createProms(threads)}
	}

	saveData(&data)
}

func createProms(thr int) [][2]interface{} {

	var promocodes = make([][2]interface{}, nPr)
	var wg sync.WaitGroup

	job := func() {
		defer log.Printf("Work (promgen) is done")

		for _ = range nPr / thr {

			slice := []interface{}{promgen(prSize)}

			promocodes = slices.Insert(promocodes, 0, promgen(prSize))
		}
	}

	if thr > 1 {
		_ = runtime.GOMAXPROCS(threads)
		for _ = range thr {
			go job()
		}
	} else {
		go job()
	}
	wg.Wait()

	return promocodes
}

func saveData(data ...interface{}) {

	const vault = ""

	file, err := os.Open(vault)
	defer file.Close()

	if err != nil {
		panic(runtime.PanicNilError{})
	}

	file.Write(data)
}
func RandStringBytesMaskImprSrcSB(n int) (string, []byte) {
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

	return sb.String(), []byte(sb.String())
}
