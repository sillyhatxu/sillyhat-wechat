package common

type ResponseEntity struct {
	Ret  int         `json:"ret"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const RESPONSE_SUCCESS int = 0

func Success(data interface{}) *ResponseEntity {
	return &ResponseEntity{Ret: RESPONSE_SUCCESS, Data: data, Msg: "success"}
}

func ErrorData(ret int, data interface{}, msg string) *ResponseEntity {
	return &ResponseEntity{Ret: ret, Data: data, Msg: msg}
}

func Error(ret int, msg string) *ResponseEntity {
	return &ResponseEntity{Ret: ret, Msg: msg}
}
