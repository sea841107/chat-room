package user

import (
	"chatserver/database"
	"chatserver/json"
	"log"
)

type historyModel struct{}

func newHistoryModel() historyModel {
	return historyModel{}
}

func (model historyModel) process(service *service, msg []byte, serverChan chan []byte) {
	var result []byte

	call, err := json.DecodeHistoryCall(msg)
	if err != nil {
		log.Println("Json Decode Call Error", HistoryCall, err)
	} else {
		var messageList []json.Message

		roomName := getRoomName(call.From, call.To)
		if history := database.Service.FindHistory(roomName, MaxMessageCount+call.Count); history != nil {
			database.Service.DeleteHint(call.From, call.To)
			service.replaceMessages(roomName, history)
			messageList = service.getMessage(roomName)
		} else {
			// add friend if there is no history
			if call.To != Lobby {
				database.Service.AddFriend(call.From, call.To)
			}
		}

		service.updateSubscribe(call.From, call.To)

		result, err = json.EncodeHistoryRecall(HistoryRecall, call.From, call.To, messageList)
		if err != nil {
			log.Println("Json Encode Recall Error", HistoryRecall, err)
		}
	}

	serverChan <- result
}
