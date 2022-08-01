package files

type S3File struct {
	Name         string `json:"name"`
	LastModified string `json:"last_modified"`
	Size         string `json:"size"`
	Type         string `json:"type"`
}

type S3PutUrl struct {
	Url string `json:"url"`
}
