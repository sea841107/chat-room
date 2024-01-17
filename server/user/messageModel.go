package user

import (
	"chatserver/database"
	"chatserver/json"
	"log"
)

type messageModel struct{}

func newMessageModel() messageModel {
	return messageModel{}
}

func (model messageModel) process(service *service, msg []byte, response chan []byte) {
	var result []byte

	call, err := json.DecodeMessageCall(msg)
	if err != nil {
		log.Println("Json Decode Call Error", MessageCall, err)
	} else {
		roomName := getRoomName(call.Message.Name, call.To)
		if database.Service.AddMessage(roomName, call.Message.Name, call.Message.Text, call.Message.Time) {
			database.Service.AddHint(call.To, call.Message.Name)
			database.Service.DeleteHint(call.Message.Name, call.To)
			service.addMessage(roomName, call.Message)

			result, err = json.EncodeMessageRecall(MessageRecall, call.To, call.Message)
			if err != nil {
				log.Println("Json Encode Recall Error", MessageRecall, err)
			}
		}
	}

	response <- result

	service.notify(call.Message.Name, call.To)
}
