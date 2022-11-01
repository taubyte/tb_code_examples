package lib

import (
    "bitbucket.org/taubyte/go-sdk/event"
	"time"
)

//export ping
func ping(e event.Event) uint32 {
    h, err := e.HTTP()
	if err != nil {
		return 1
	}


	h.Write([]byte(fmt.Sprintf("{\"pong\": %d}", time.Now().Unix())))

    return 0
}
