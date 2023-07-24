package hospitalQueue

import (
	"fmt"
	"github.com/bxcodec/faker/v4"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

type QueuePageData struct {
	Data []Patient
}

type CounterChan struct {
	setZero   chan bool
	increment chan bool
	decrement chan bool
	value     int
}

type DataDistribution struct {
	enqueueOdd int
	dequeueOdd int

	sampleSpace int
	mean        int
	deviation   int
	counter     chan int
}

var wg sync.WaitGroup

var threadProfile = pprof.Lookup("threadProfile")

func init() {
	runtime.GOMAXPROCS(int(math.Max(
		1,
		float64(runtime.NumCPU()/2),
	)))
}

func randPatientArray(size int) []Patient {
	array := make([]Patient, size)

	for i := range array {
		array[i] = randPatient()
	}

	return array
}

func randPatient() Patient {
	ageLimit := 100

	return Patient{
		Name: faker.Name(),
		Age:  uint8(rand.Intn(ageLimit)),
	}
}

func (dd *DataDistribution) new(enqOdd int, deqOdd int, mean int) DataDistribution {
	sampleSpace := enqOdd + deqOdd
	p := float64(enqOdd) / float64(sampleSpace)
	n := float64(mean) / p

	return DataDistribution{
		enqueueOdd: enqOdd,
		dequeueOdd: deqOdd,

		sampleSpace: sampleSpace,
		mean:        mean,
		deviation:   int(math.Sqrt(n * p * (1 - p))),
		counter:     make(chan int),
	}
}

func fillQueue(sourceQueue *HospitalQueue, dataDist *DataDistribution) {

	if (<-dataDist.counter) < dataDist.deviation {
		randSample := rand.Intn(dataDist.sampleSpace+1) + 1
		if randSample < dataDist.enqueueOdd {
			pat := randPatient()
			sourceQueue.Enqueue(pat)
			fmt.Println("\t |enq|" + pat.getFmtData())

			dataDist.counter <- (<-dataDist.counter) + 1
		} else {
			pat, err := sourceQueue.Dequeue()

			if err == nil {
				fmt.Println("\t |deq|" + pat.getFmtData())

				dataDist.counter <- (<-dataDist.counter) - 1
			} else {
				fmt.Println("\t |err|" + err.Error())
			}
		}
	} else {
		sourceQueue.Reset(
			randPatientArray(dataDist.deviation)...)
		dataDist.counter <- 0
		fmt.Println("\t |Deviation compensation| QUEUE RESET")
	}

	time.Sleep(time.Second)
	defer wg.Done()
}

func parallelFill(sourceQueue *HospitalQueue) {
	dd := (&DataDistribution{}).new(8, 2, 20)
	for {
		wg.Add(1)
		println("Test")
		fillQueue(sourceQueue, &dd)
	}
}

func GinServer() {
	port := os.Getenv("PORT")
	htmlPath := os.Getenv("HTML_PATH")
	first := Patient{
		Name: "Pedro Vinicius",
		Age:  19,
	}

	sourceQueue, _ := (&HospitalQueue{}).New(first)

	wg.Add(1)
	go parallelFill(sourceQueue)

	if port != "" && htmlPath != "" {
		router := gin.Default()
		router.LoadHTMLGlob(htmlPath)

		router.GET("/view", func(context *gin.Context) {
			context.HTML(
				http.StatusOK,
				"index.html",
				QueuePageData{
					Data: sourceQueue.Unroll(),
				})
		})

		timeOut := 100 * time.Millisecond
		maxHeaderBytes := 1 << 20

		server := &http.Server{
			Addr:              ":" + port,
			Handler:           router,
			ReadTimeout:       timeOut,
			ReadHeaderTimeout: timeOut,
			WriteTimeout:      timeOut,
			MaxHeaderBytes:    maxHeaderBytes,
		}

		err := server.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("No port defined!")
	}
}
