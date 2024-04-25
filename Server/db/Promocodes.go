package Promocodes

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
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

var globVault = os.Getenv("hashes")

var salt string = os.Getenv("Salt")

func promgen(promSize int) (string, [sha256.Size]byte) {

	var ps, pb = RandStringBytesMaskImprSrcSB(promSize)

	return ps, sha256.Sum256(append(pb, []byte(salt)...))
}

func isValid(prom string, salt string) {

	checkBytes := bytes.Join([][]byte{[]byte(prom), []byte(salt)}, nil)

	checkhash := sha256.Sum256(checkBytes)

	file, err := os.Open(globVault)
	if err != nil {
		panic(runtime.PanicNilError{})
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {

		line := sc.Text()

		if line == base64.StdEncoding.EncodeToString(checkhash[:]) {
			//uniq promocode found
		} else {
			panic(runtime.PanicNilError{})
		}
	}
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

func createProms(thr int) [][]interface{} {

	var promocodes = make([][]interface{}, nPr)
	var wg sync.WaitGroup

	job := func() {
		defer log.Printf("Work (promgen) is done")

		for i := 0; i < nPr/thr; i++ {
			slice := []interface{}{promgen(prSize)}

			promocodes = slices.Insert(promocodes, 0, slice)
		}
	}

	if thr > 1 {
		for _ = range thr {
			go job()
		}
	} else {
		go job()
	}
	wg.Wait()

	return promocodes
}

func saveData(data interface{}) {

	const vault = ""

	file, err := os.Open(vault)
	defer file.Close()

	if err != nil {
		panic(runtime.PanicNilError{})
	}

	//writer := bufio.NewWriter()
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
