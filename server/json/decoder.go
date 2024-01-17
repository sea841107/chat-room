package json

import (
	"encoding/json"
)

func DecodeHeader(msg []byte) (string, error) {
	header := &Header{}
	err := json.Unmarshal(msg, &header)
	return header.Header, err
}

func DecodeLoginCall(msg []byte) (*LoginCall, error) {
	call := &LoginCall{}
	err := json.Unmarshal(msg, &call)
	return call, err
}

func DecodeMessageCall(msg []byte) (*MessageCall, error) {
	call := &MessageCall{}
	err := json.Unmarshal(msg, &call)
	return call, err
}

func DecodeHistoryCall(msg []byte) (*HistoryCall, error) {
	call := &HistoryCall{}
	err := json.Unmarshal(msg, &call)
	return call, err
}

func DecodeSearchCall(msg []byte) (*SearchCall, error) {
	call := &SearchCall{}
	err := json.Unmarshal(msg, &call)
	return call, err
}

func DecodeUpdateCall(msg []byte) (*UpdateCall, error) {
	call := &UpdateCall{}
	err := json.Unmarshal(msg, &call)
	return call, err
}
