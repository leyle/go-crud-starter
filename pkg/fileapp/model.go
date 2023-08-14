package fileapp

const defaultBucketName = "cbdc-kyc"

type UploadedFile struct {
	Data     []byte        `json:"-" bson:"-"` // return field
	Metadata *FileMetadata `json:"-" bson:"-"`
}

type FileMetadata struct {
	Id               string `json:"id" bson:"id"`
	UserId           string `json:"userId" bson:"userId"` // uploaded user, maybe ethereum address
	OriginalFilename string `json:"originalFilename" bson:"originalFilename"`
	Filename         string `json:"filename" bson:"filename"`
	Size             int64  `json:"size" bson:"size"`
	Ext              string `json:"ext" bson:"ext"`
	Sha256           string `json:"sha256" bson:"sha256"` // sha256
	Created          int64  `json:"created" bson:"created"`
	Repeated         bool   `json:"repeated" bson:"-"`
}
