package error_messages

type ErrorMsg struct {
	DbConnectionError string "Failed to connect DB"
}

var Error *ErrorMsg
