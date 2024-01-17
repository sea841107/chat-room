package user

import (
	"chatserver/json"
	"log"
	"sync"
)

var Service = &service{
	id:      "UserService",
	models:  make(map[string]userModel),
	room:    make(map[string]userRoom),
	poll:    make(map[string]userPoll),
	rwMutex: &sync.RWMutex{},
}

type service struct {
	id      string
	models  map[string]userModel
	room    map[string]userRoom
	poll    map[string]userPoll
	rwMutex *sync.RWMutex
}

func (s *service) Start() {
	s.models[LoginCall] = newLoginModel()
	s.models[HistoryCall] = newHistoryModel()
	s.models[MessageCall] = newMessageModel()
	s.models[SearchCall] = newSearchModel()
	s.models[UpdateCall] = newUpdateModel()
}

func (s *service) ProcessMsg(msg []byte, response chan []byte) {
	header, err := json.DecodeHeader(msg)
	if err != nil {
		log.Println("Json Decode Header Error", err)
	}

	if model, ok := s.models[header]; ok {
		model.process(s, msg, response)
	}
}

func (s *service) addMessage(roomName string, message json.Message) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	room := userRoom{}
	if r, ok := s.room[roomName]; ok {
		room = r
	}

	room.userMessage = append(room.userMessage, message)
	if len(room.userMessage) > MaxMessageCount {
		room.userMessage = room.userMessage[1:]
	}
	s.room[roomName] = room
}

func (s *service) getMessage(roomName string) (messages []json.Message) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	if r, ok := s.room[roomName]; ok {
		return r.userMessage
	}

	return nil
}

func (s *service) replaceMessages(roomName string, messages []json.Message) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	room := userRoom{}

	for i := len(messages) - 1; i >= 0; i-- {
		room.userMessage = append(room.userMessage, messages[i])
	}

	s.room[roomName] = room
}

func (s *service) subscribe(from, to string, onNotify func(from, to string), onTimeout func()) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	if poll, ok := s.poll[from]; ok {
		poll.onTimeout()
		delete(s.poll, from)
	}

	s.poll[from] = userPoll{
		from:      from,
		to:        to,
		onNotify:  onNotify,
		onTimeout: onTimeout,
	}
}

func (s *service) updateSubscribe(from, to string) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	if poll, ok := s.poll[from]; ok {
		poll.from = from
		poll.to = to
		s.poll[from] = poll
	}
}

func (s *service) notify(from, to string) {
	s.rwMutex.Lock()

	var pollList []userPoll
	if to == Lobby {
		for k, poll := range s.poll {
			if k != from && poll.to == to {
				pollList = append(pollList, poll)
				delete(s.poll, k)
			}
		}
	} else {
		if poll, ok := s.poll[to]; ok {
			pollList = append(pollList, poll)
			delete(s.poll, to)
		}
	}

	s.rwMutex.Unlock()

	for _, poll := range pollList {
		poll.onNotify(poll.from, poll.to)
	}
}

func getRoomName(name1, name2 string) string {
	if name1 == Lobby || name2 == Lobby {
		return "room_" + Lobby
	} else if name1 < name2 {
		return "room_" + name1 + "_" + name2
	} else {
		return "room_" + name2 + "_" + name1
	}
}
