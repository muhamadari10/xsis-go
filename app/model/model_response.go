package model

type StatusData int

type StatusMessagePair struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

const (
	Success  StatusData = 2001
	InvInput StatusData = 3041
	NotFound StatusData = 4041
)

func (sD StatusData) GetMessage() string {
	return statusMessage[sD].Message
}

func (sD StatusData) GetCode() int {
	return statusMessage[sD].Code
}

func (sD StatusData) GetStatus() string {
	return statusMessage[sD].Status
}

var statusMessage = map[StatusData]StatusMessagePair{
	Success: {
		Status:  "success",
		Message: "Success %s",
		Code:    200,
	},
	InvInput: {
		Status:  "error",
		Message: "Error %s",
		Code:    422,
	},
	NotFound: {
		Status:  "error",
		Message: "Movie not found",
		Code:    404,
	},
}
