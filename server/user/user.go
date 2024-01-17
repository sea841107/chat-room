package user

import (
	"chatserver/json"
)

const (
	MaxMessageCount = 10
	Lobby           = "Lobby"

	LoginCall     = "LoginCall"
	LoginRecall   = "LoginRecall"
	HistoryCall   = "HistoryCall"
	HistoryRecall = "HistoryRecall"
	MessageCall   = "MessageCall"
	MessageRecall = "MessageRecall"
	SearchCall    = "SearchCall"
	SearchRecall  = "SearchRecall"
	UpdateCall    = "UpdateCall"
	UpdateRecall  = "UpdateRecall"
)

type userModel interface {
	process(service *service, msg []byte, response chan []byte)
}

type userRoom struct {
	userMessage []json.Message
}

type userPoll struct {
	from      string
	to        string
	onNotify  func(from, to string)
	onTimeout func()
}
