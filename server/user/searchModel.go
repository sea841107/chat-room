package user

import (
	"chatserver/database"
	"chatserver/json"
	"log"
	"sort"
)

type searchModel struct{}

func newSearchModel() searchModel {
	return searchModel{}
}

func (model searchModel) process(service *service, msg []byte, response chan []byte) {
	var result []byte
	call, err := json.DecodeSearchCall(msg)
	if err != nil {
		log.Println("Json Decode Call Error", SearchCall, err)
	} else {
		var userList []string
		if call.Name == "" {
			// no query, search friend
			if user, ok := database.Service.FindUser(call.From); ok {
				userList = user.Friend
				sort.Strings(userList)
			}
		} else {
			userList = database.Service.FindUserList(call.Name)
		}

		result, err = json.EncodeSearchRecall(SearchRecall, userList)
		if err != nil {
			log.Println("Json Encode Recall Error", SearchRecall, err)
		}
	}

	response <- result
}
