package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"log"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/iowar/kecmatch/keccak256"
	"github.com/iowar/kecmatch/models"
)

const (
	pow10 = 10000000000
)

var (
	// atomic iteration.
	iteration int64
)

func main() {
	name := flag.String("fn_name", "DefaultFunc", "# solidity function name")
	args := flag.String("fn_args", "(address[],uint256)", "# solidity function arguments")
	sig := flag.String("fn_sig", "00badfee", "# solidity function method signature")

	search_goroutines := flag.Int("goroutines", 8, "# goroutines ")
	logger_interval := flag.Int("interval", 3, "# logger interval (second)")

	flag.Parse()

	funcName := *name
	funcArgs := *args
	funcSignature := *sig

	interval := time.Duration(*logger_interval) * time.Second
	goroutines := *search_goroutines

	if len(funcSignature) != 8 {
		log.Fatal("signature should be 4 byte!")
		return
	}

	keccak := keccak256.New()

	signature, err := hex.DecodeString(funcSignature)
	if err != nil {
		log.Fatal(err)
	}

	go metrics(interval)

	sms := models.NewSolidityMethodStatus(goroutines)
	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {

		start := i * pow10
		wg.Add(1)

		go func(start int, signature []byte, funcName, funcArgs string, wg *sync.WaitGroup) {
			defer wg.Done()

			for j := start; j < start+pow10; j++ {
				if sms.IsFound() {
					break
				}

				data := funcName + strconv.Itoa(j) + funcArgs

				if bytes.Compare(keccak.Hash([]byte(data)), signature) == 0 {
					sms.SetInfo(data, atomic.LoadInt64(&iteration), hex.EncodeToString(signature))

					break
				}

				atomic.AddInt64(&iteration, 1)
			}
		}(start, signature, funcName, funcArgs, &wg)
	}

	wg.Wait()

	sms.Display()
	//d, i, s := sms.GetInfo()
	//log.Printf("\033[31m found %s signature using %s in %d iteration \033[0m", s, d, i)
}

func metrics(interval time.Duration) {
	var (
		// duplication rate based on by number of iterations.
		// margin of error per iteration.
		perErrorCost = math.Pow(256, 28) / ((math.Pow(256, 32)) * 2.)
		// cum vars. for geo. distribution family.
		p = 1. / math.Pow(16, 8)
		q = 1 - p
	)

	for {
		time.Sleep(interval)

		totalIteration := atomic.LoadInt64(&iteration)
		errorCost := perErrorCost * float64(totalIteration)

		probability := (1. - math.Pow(q, float64(totalIteration))) * 100.
		log.Printf(
			"\033[32m iteration: %d \033[33m margin of error: %.5f \033[34m probability of approaching desired sig: %.3f%%  \033[0m",
			totalIteration,
			errorCost,
			probability,
		)
	}
}
