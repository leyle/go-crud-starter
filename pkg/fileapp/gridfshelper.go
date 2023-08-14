package fileapp

import (
	"fmt"
	"github.com/leyle/crud-objectid/pkg/objectid"
	"github.com/leyle/go-crud-starter/configandcontext"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"mime/multipart"
	"time"
)

func saveFileToGridFS(ctx *configandcontext.APIContext, file multipart.File, header *multipart.FileHeader) (*UploadedFile, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("read uploaded file body failed")
		return nil, err
	}

	// parse file metadata
	originalFilename := header.Filename
	fileSize := header.Size
	ext := GetFileExt(originalFilename)
	fileHash := GetFileSha256(data)
	newFilename := fmt.Sprintf("%s.%s", fileHash, ext)

	ctx.Logger.Debug().
		Str("userId", ctx.KYCUser.Address).
		Str("originalFilename", originalFilename).
		Str("fileHash", fileHash).
		Str("filename", newFilename).
		Msg("try to parse and save file")

	// before really save to db, we need to check if this is a repeated upload
	dbFile, err := checkFileRepeatUploadedByFilename(ctx, newFilename)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("check file repeat status failed")
		return nil, err
	}
	if dbFile != nil {
		return dbFile, nil
	}

	fileMetaInfo := &FileMetadata{
		Id:               objectid.GetObjectId(),
		UserId:           ctx.KYCUser.Address,
		OriginalFilename: originalFilename,
		Filename:         newFilename,
		Size:             fileSize,
		Ext:              ext,
		Sha256:           fileHash,
		Created:          time.Now().Unix(),
	}

	uploadOpts := options.GridFSUpload()
	uploadOpts.SetMetadata(fileMetaInfo)

	bucket, err := getGridFSBucket(ctx, defaultBucketName)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("get gridFS bucket failed")
		return nil, err
	}

	stream, err := bucket.OpenUploadStreamWithID(fileMetaInfo.Id, fileMetaInfo.Filename, uploadOpts)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("open gridFS upload stream failed")
		return nil, err
	}
	defer stream.Close()

	_, err = stream.Write(data)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("write file data into gridFS failed")
		return nil, err
	}

	result := &UploadedFile{
		Metadata: fileMetaInfo,
	}

	return result, nil
}

func getFileFromGridFSById(ctx *configandcontext.APIContext, id string) (*UploadedFile, error) {
	bucket, err := getGridFSBucket(ctx, defaultBucketName)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("get gridFS bucket failed")
		return nil, err
	}

	stream, err := bucket.OpenDownloadStream(id)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("open gridFS download stream failed")
		return nil, err
	}
	defer stream.Close()

	// read metadata
	var metadata FileMetadata
	md := stream.GetFile().Metadata
	err = bson.Unmarshal(md, &metadata)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("unmarshal gridFS file metadata failed")
		return nil, err
	}

	// read file content
	fileData, err := io.ReadAll(stream)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("read gridFS file data failed")
		return nil, err
	}

	upFile := &UploadedFile{
		Data:     fileData,
		Metadata: &metadata,
	}

	return upFile, nil
}

func checkFileRepeatUploadedByFilename(ctx *configandcontext.APIContext, filename string) (*UploadedFile, error) {
	// filename means sha256 hash plus file extension
	bucket, err := getGridFSBucket(ctx, defaultBucketName)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("get gridFS bucket failed")
		return nil, err
	}

	stream, err := bucket.OpenDownloadStreamByName(filename)
	if err != nil {
		if err == gridfs.ErrFileNotFound {
			return nil, nil
		}
		ctx.Logger.Error().Err(err).Msg("open gridFS download stream failed")
		return nil, err
	}
	defer stream.Close()

	// read metadata
	var metadata FileMetadata
	md := stream.GetFile().Metadata
	ctx.Logger.Warn().Interface("metadata", md).Send()
	err = bson.Unmarshal(md, &metadata)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("unmarshal gridFS file metadata failed")
		return nil, err
	}

	metadata.Repeated = true
	upFile := &UploadedFile{
		Metadata: &metadata,
	}

	return upFile, nil
}
