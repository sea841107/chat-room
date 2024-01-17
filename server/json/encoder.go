package json

import "encoding/json"

func EncodeLoginRecall(header, status, name string) ([]byte, error) {
	data := &LoginRecall{
		Header: header,
		Status: status,
		Name:   name,
	}
	msg, err := json.Marshal(data)
	return msg, err
}

func EncodeMessageRecall(header, to string, message Message) ([]byte, error) {
	data := &MessageRecall{
		Header:  header,
		To:      to,
		Message: message,
	}
	msg, err := json.Marshal(data)
	return msg, err
}

func EncodeHistoryRecall(header string, from, to string, messageList []Message) ([]byte, error) {
	data := &HistoryRecall{
		Header:      header,
		From:        from,
		To:          to,
		MessageList: messageList,
	}
	msg, err := json.Marshal(data)
	return msg, err
}

func EncodeSearchRecall(header string, userList []string) ([]byte, error) {
	data := &SearchRecall{
		Header:   header,
		UserList: userList,
	}
	msg, err := json.Marshal(data)
	return msg, err
}

func EncodeUpdateRecall(header string, from, to string, userList, hintList []string, messageList []Message) ([]byte, error) {
	data := &UpdateRecall{
		Header:      header,
		From:        from,
		To:          to,
		UserList:    userList,
		HintList:    hintList,
		MessageList: messageList,
	}
	msg, err := json.Marshal(data)
	return msg, err
}
