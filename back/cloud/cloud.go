package cloud

type Cloud interface {
	UploadImage([]byte, string) error
	ObjectRecognition(string) ([]string, error)
	LabelRecognition(string) ([]string, error)
	GetURL(string) (string, error)
	Close()
}
