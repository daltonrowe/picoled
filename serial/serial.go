package serial

import (
	"machine"
	"errors"
	"github.com/buger/jsonparser"
)

func Read() (data string, err error) {
	bytes := [256]byte{}
	next := 0

	for machine.Serial.Buffered() > 0 {
		c, err := machine.Serial.ReadByte()

		if err == nil {
			bytes[next] = c
			next += 1
		}
	}

	if next > 0 {
		s := bytes[0:next]
		name, err := jsonparser.GetString(s, "name")

		if err == nil {
			data = name
		}
	}

	if len(data) == 0 {
		err = errors.New("Effect not found")
	}
	
	return
}

