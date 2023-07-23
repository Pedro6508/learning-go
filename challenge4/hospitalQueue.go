package hospitalQueue

import (
	"errors"
	"strconv"
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

func (pat *Patient) getFmtData() string {
	return "Name: " + pat.Name + " Age: " + strconv.FormatInt(int64(pat.Age), 10)
}

func (hq *HospitalQueue) Reset(patients ...Patient) {
	if len(patients) > 0 {
		first := &QueuePosition{
			data: patients[0],
			next: nil,
		}

		hq.first = first
		hq.last = first
		hq.size = 1

		for _, pat := range patients[1:] {
			hq.Enqueue(pat)
		}
	}

	hq.first = nil
	hq.last = nil
	hq.size = 0
}

func (hq *HospitalQueue) Dequeue() (Patient, error) {
	first := hq.first

	if first == nil {
		return Patient{}, errors.New("empty Queue")
	} else if first == hq.last {
		hq.last = nil
	}

	hq.first = hq.first.next
	hq.size -= 1

	return first.data, nil
}

func (hq *HospitalQueue) Enqueue(patient Patient) {
	patientPosition := QueuePosition{
		data: patient,
		next: nil,
	}

	if hq.size == 0 {

		hq.first = &patientPosition
		hq.last = &patientPosition
		hq.size = 1

		return
	}

	hq.last.next = &patientPosition
	hq.last = &patientPosition
	hq.size += 1
}

func (hq *HospitalQueue) New(patients ...Patient) (*HospitalQueue, error) {
	if len(patients) > 0 {
		first := &QueuePosition{
			data: patients[0],
			next: nil,
		}

		queue := HospitalQueue{
			first: first,
			last:  first,
			size:  1,
		}

		for _, pat := range patients[1:] {
			queue.Enqueue(pat)
		}

		return &queue, nil
	}

	return &HospitalQueue{
		first: nil,
		last:  nil,
		size:  0,
	}, nil
}

func (hq *HospitalQueue) Unroll() []Patient {
	next := hq.first
	unwoundQueue := make([]Patient, hq.size)

	for i := range unwoundQueue {
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
