package utilerr

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrChanTimeOut    = errors.New("channel timeout")
)

func IsErrRecordNotFound(errStr string) bool {
	if errStr == "record not found" {
		return true
	}
	return false
}
