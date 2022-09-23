package lib

import (
	"bitbucket.org/taubyte/go-sdk/event"

	"bytes"
	"strings"
	"time"

	"github.com/o1egl/govatar"

	"image/png"

	"runtime/debug"
)

//export avatar
func avatar(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	gender, _ := h.Query().Get("gender")
	if gender == "" {
		if time.Now().UnixNano()%2 == 0 {
			gender = "female"
		} else {
			gender = "male"
		}
	}

	username, _ := h.Query().Get("username")
	if username == "" {
		username = time.Now().String()
	}

	var _gender govatar.Gender
	switch strings.ToLower(gender) {
	case "male", "m":
		_gender = govatar.MALE
	case "female", "f":
		_gender = govatar.FEMALE
	}

	var b bytes.Buffer

	img, err := govatar.GenerateForUsername(_gender, username)
	if err != nil {
		h.Write([]byte("generate failed with " + err.Error()))
		h.Write(debug.Stack())
		return 1
	}

	err = png.Encode(&b, img)
	if err != nil {
		h.Write([]byte("png encoding failed with " + err.Error()))
		h.Write(debug.Stack())
		return 1
	}

	h.Headers().Set("Content-Type", "image/png")
	h.Write(b.Bytes())

	return 0
}

