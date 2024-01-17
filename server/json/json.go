package json

type Message struct {
	Name string `json:"name"`
	Text string `json:"text"`
	Time int64  `json:"time"`
}

type Header struct {
	Header string `json:"header"`
}

type LoginCall struct {
	Header   string `json:"header"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRecall struct {
	Header string `json:"header"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

type MessageCall struct {
	Header  string  `json:"header"`
	To      string  `json:"to"`
	Message Message `json:"message"`
}

type MessageRecall struct {
	Header  string  `json:"header"`
	To      string  `json:"to"`
	Message Message `json:"message"`
}

type HistoryCall struct {
	Header string `json:"header"`
	From   string `json:"from"`
	To     string `json:"to"`
	Count  int    `json:"count"`
}

type HistoryRecall struct {
	Header      string    `json:"header"`
	From        string    `json:"from"`
	To          string    `json:"to"`
	MessageList []Message `json:"messageList"`
}

type SearchCall struct {
	Header string `json:"header"`
	From   string `json:"from"`
	Name   string `json:"name"`
}

type SearchRecall struct {
	Header   string   `json:"header"`
	UserList []string `json:"userList"`
}

type UpdateCall struct {
	Header string `json:"header"`
	From   string `json:"from"`
	To     string `json:"to"`
}

type UpdateRecall struct {
	Header      string    `json:"header"`
	From        string    `json:"from"`
	To          string    `json:"to"`
	UserList    []string  `json:"userList"`
	HintList    []string  `json:"hintList"`
	MessageList []Message `json:"messageList"`
}
