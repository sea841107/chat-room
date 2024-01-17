package user

import (
	"chatserver/database"
	"chatserver/json"
	"log"
)

type loginModel struct{}

func newLoginModel() loginModel {
	return loginModel{}
}

func (model loginModel) process(service *service, msg []byte, response chan []byte) {
	var result []byte

	call, err := json.DecodeLoginCall(msg)
	if err != nil {
		log.Println("Json Decode Call Error", LoginCall, err)
	} else {
		var status string
		if user, ok := database.Service.FindUser(call.Name); ok {
			if user.Password == call.Password {
				status = "Success"
			} else {
				status = "Fail"
			}
		} else if database.Service.AddUser(call.Name, call.Password) {
			status = "New"
		} else {
			status = "Error"
		}

		result, err = json.EncodeLoginRecall(LoginRecall, status, call.Name)
		if err != nil {
			log.Println("Json Encode Recall Error", LoginRecall, err)
		}
	}

	response <- result
}
