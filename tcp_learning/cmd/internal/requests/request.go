package requests

import (
	"bytes"
	"fmt"
	"io"
)

// defining enums
type parserState string

const (
	StateInit parserState = "init"
	StateDone parserState = "done"
)

var MALFORMED_REQ_LINE = fmt.Errorf("malformed request line")
var UNSUPPORTED_HTTP_VER = fmt.Errorf("unsupported http version")

var SEPARATOR = []byte("\r\n")

// creating  a struct that  will be used to create a HTTP request format.
type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
	state       parserState
}

func newRequest() *Request {
	return &Request{
		state: StateInit,
	}
}

func ParseRequestLine(b []byte) (*RequestLine, int, error) {
	idx := bytes.Index(b, SEPARATOR)
	if idx == -1 {
		return nil, 0, nil
	}

	startLine := b[:idx]
	read := idx + len(SEPARATOR)

	parts := bytes.Split(startLine, []byte(" ")) // we get an array of [Method, path, http-version]
	if len(parts) != 3 {
		return nil, 0, MALFORMED_REQ_LINE
	}
	httpParts := bytes.Split(parts[2], []byte("/"))

	if len(httpParts) != 2 || string(httpParts[0]) != "HTTP" || string(httpParts[1]) != "1.1" {
		return nil, 0, UNSUPPORTED_HTTP_VER
	}

	// Since as per RFC9110 message-parsing protocol, the startline should only have single-space as separators
	// length  of parts must be 3.
	if len(parts) != 3 {
		return nil, 0, MALFORMED_REQ_LINE
	}
	rl := &RequestLine{
		RequestTarget: string(parts[1]),
		HttpVersion:   string(httpParts[1]),
		Method:        string(parts[0]),
	}

	return rl, read, nil

}

func (r *Request) parse(data []byte) (int, error) {
	read := 0
outer:
	for {
		switch r.state {
		case StateInit:
			rl, n, err := ParseRequestLine(data[read:])
			if err != nil {
				return 0, err
			}
			if n == 0 {
				break outer
			}
			r.RequestLine = *rl
			read += n
			r.state = StateDone
		case StateDone:
			break outer //nothing new to parse here
		}
	}
	return read, nil
}

func (r *Request) done() bool {
	return r.state == StateDone
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	request := newRequest()
	buf := make([]byte, 1024)
	bufLen := 0
	for !request.done() {
		n, err := reader.Read(buf[bufLen:])
		if err != nil {
			return nil, err
		}
		readN, err := request.parse(buf[:bufLen+n])
		if err != nil {
			return nil, err
		}

		// moving  data to beginning , and readjusting buffer length
		copy(buf, buf[readN:bufLen])
		bufLen -= readN
	}
	return request, nil
}
