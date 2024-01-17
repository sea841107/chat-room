package user

import (
	"chatserver/database"
	"chatserver/json"
	"log"
	"sort"
)

type updateModel struct{}

func newUpdateModel() updateModel {
	return updateModel{}
}

func (model updateModel) process(service *service, msg []byte, response chan []byte) {
	var result []byte

	call, err := json.DecodeUpdateCall(msg)
	if err != nil {
		log.Println("Json Decode Call Error", UpdateCall, err)
		response <- result
	} else {
		onNotify := func(from, to string) {
			roomName := getRoomName(from, to)
			messageList := service.getMessage(roomName)

			user, _ := database.Service.FindUser(from)
			userList := user.Friend
			hintList := user.Hint
			sort.Strings(userList)
			sort.Strings(hintList)

			result, err = json.EncodeUpdateRecall(UpdateRecall, from, to, userList, hintList, messageList)
			if err != nil {
				log.Println("Json Encode Recall Error", UpdateRecall, err)
			}

			response <- result
		}

		onTimeout := func() {
			close(response)
		}

		service.subscribe(call.From, call.To, onNotify, onTimeout)
	}
}
