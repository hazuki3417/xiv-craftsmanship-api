package schema

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
 * marker struct
 */
type PrimaryKey struct {
	Id string `bson:"_id,omitempty"`
}

/*
 * marker interface
 */
type PrimaryKeyInteface interface {
	GetId() string
	ToStrId(id *primitive.ObjectID) (string, error)
	ToObjId(id string) (primitive.ObjectID, error)
}

func (p PrimaryKey) GetId() string {
	return p.Id
}

func (p PrimaryKey) ToStrId(id *primitive.ObjectID) (string, error) {
	if id == nil {
		return "", errors.New("id is nil")
	}
	return id.Hex(), nil
}

func (p PrimaryKey) ToObjId(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

type Content struct {
	// 構造体を埋め込んだ場合はinlineタグを指定すること
	PrimaryKey    `bson:",inline"`
	WorkspaceId   string    `bson:"workspaceId" validate:"required"`
	Tags          []string  `bson:"tags"`
	S3DirPath     string    `bson:"s3DirPath" validate:"required"`
	ThumbnailFile *File     `bson:"thumbnailFile"`
	OriginalFile  File      `bson:"originalFile"`
	CreatedAt     time.Time `bson:"createdAt" validate:"required"`
	UpdatedAt     time.Time `bson:"updatedAt" validate:"required"`
}

type File struct {
	Name      string `bson:"name" validate:"required"`
	Extension string `bson:"extension" validate:"required"`
	Type      string `bson:"type" validate:"required"`
	Size      int    `bson:"size" validate:"gte=0"`
}
