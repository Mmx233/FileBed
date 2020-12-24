package Public

type body struct {
	Status string `json:"status"`
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func CallError(Code int,Msg string)body{
	return body{Status: "error",Code: Code,Msg: Msg,Data: map[string]interface{}{}}
}

func CallSuccess(Code int,Msg string,Data interface{})body{
	return body{Status: "success",Code: Code,Msg: Msg,Data: Data}
}

func CallErrorWithCode(Code int)body{//统一error code
	codes:=map[int]string{
		1:"您的访问超出频次限制",
	}
	return CallError(Code,codes[Code])
}
