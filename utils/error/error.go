package error

type ErrorType uint

func HandleError() {

}

type CustomError struct {
	ErrorType     ErrorType
	OriginalError error
	contextInfo   map[string]string
}
