package hospitalQueue

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type Patient struct {
	Name string
	Age  uint8
}

type QueuePosition struct {
	data Patient
	next *QueuePosition
}

type HospitalQueue struct {
	first *QueuePosition
	last  *QueuePosition
	size  uint16
}

func (hq *HospitalQueue) add(patient Patient) {
	if hq.first == nil || hq.last == nil {
		hq.first = &QueuePosition{
			data: patient,
			next: nil,
		}
		hq.last = hq.first
		hq.size = 1
	} else {
		hq.last.next = &QueuePosition{
			data: patient,
			next: nil,
		}

		hq.last = hq.last.next
		hq.size = hq.size + 1
	}
}

func (hq *HospitalQueue) unroll() []Patient {
	next := hq.first
	unwoundQueue := make([]Patient, hq.size)

	for i, _ := range unwoundQueue {
		if next != nil {
			unwoundQueue[i] = next.data
		} else {
			lastValidIndex := i - 1

			hq.size = uint16(lastValidIndex + 1)
			unwoundQueue = unwoundQueue[0:(lastValidIndex + 1)]
			break
		}

		next = next.next
	}

	return unwoundQueue
}

type QueuePageData struct {
	Data []Patient
}

func GinServer() {
	port := os.Getenv("PORT")
	htmlPath := os.Getenv("HTML_PATH")

	if port != "" && htmlPath != "" {
		sourceQueue := HospitalQueue{
			first: nil,
			last:  nil,
			size:  0,
		}
		sourceQueue.add(Patient{
			Name: "Lucas",
			Age:  37,
		})

		sourceQueue.add(Patient{
			Name: "Gabriel",
			Age:  73,
		})

		testData := QueuePageData{
			Data: sourceQueue.unroll(),
		}

		router := gin.Default()
		router.LoadHTMLGlob(htmlPath)

		router.GET("/test", func(context *gin.Context) {
			context.HTML(http.StatusOK, "index.html", testData)
		})

		router.Run(":" + port)
	} else {
		log.Fatal("No port defined!")
	}
}
