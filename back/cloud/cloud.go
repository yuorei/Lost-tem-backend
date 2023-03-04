package cloud

type Cloud interface {
	UploadImage([]byte, string) error
	ObjectRecognition(string) ([]string, error)
	Close()
}
