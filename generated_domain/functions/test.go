package lib

import (
	"bitbucket.org/taubyte/go-sdk/event"
)

//export test
func test(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	h.Write([]byte("TEST"))

	return 0
}
