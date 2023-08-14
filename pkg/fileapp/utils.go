package fileapp

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/leyle/go-crud-starter/configandcontext"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

func GetFileExt(filename string) string {
	if !strings.Contains(filename, ".") {
		return ""
	}

	infos := strings.Split(filename, ".")
	ext := infos[len(infos)-1]
	return strings.ToLower(ext)
}

func GetFileSha256(data []byte) string {
	h := sha256.New()
	h.Write(data)
	val := h.Sum(nil)
	return hex.EncodeToString(val)
}

func convertStructToBsonM(data interface{}) (bson.M, error) {
	d1, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}

	dm := bson.M{}
	err = bson.Unmarshal(d1, &dm)
	if err != nil {
		return nil, err
	}

	return dm, nil
}

func getGridFSBucket(ctx *configandcontext.APIContext, prefix string) (*gridfs.Bucket, error) {
	db := ctx.Ds.Client().Database(ctx.Cfg.Mongodb.Database)
	opts := options.GridFSBucket()
	opts.SetName(prefix)
	bucket, err := gridfs.NewBucket(db, opts)
	return bucket, err
}
