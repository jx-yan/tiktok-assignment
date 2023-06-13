package main

type Message struct {
	Sender    string `json:"sender"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type DeleteAllRequest struct {
	Chat string `thrift:"1,required"`
}

type DeleteAllResponse struct {
	Code int32  `thrift:"1,required"`
	Msg  string `thrift:"2,required"`
}
