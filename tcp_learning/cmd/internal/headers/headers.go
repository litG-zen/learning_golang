package headers

import (
	"bytes"
	"fmt"
)

var rn = []byte("\r\n")

type Headers map[string]string

func NewHeaders() Headers {
	return map[string]string{}
}

func parserHeader(fieldLine []byte) (string, string, error) {
	parts := bytes.SplitN(fieldLine, []byte(":"), 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("Malformed field line")
	}

	name := parts[0]
	value := bytes.TrimSpace(parts[1])

	if bytes.HasSuffix(name, []byte(" ")) {
		return "", "", fmt.Errorf("Malformed field name")
	}

	return string(name), string(value), nil
}

func (h Headers) Parse(data []byte) (int, bool, error) {
	read := 0
	done := false
	for {
		idx := bytes.Index(data[read:], rn)
		if idx == -1 {
			break
		}
		if idx == 0 {
			done = true
			break
		}

		name, value, err := parserHeader(data[:idx])
		if err != nil {
			return 0, false, err
		}
		read += idx + len(rn)
		h[name] = value
	}

	return read, done, nil
}
