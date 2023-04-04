package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type ProjectDocument interface {
	GetDocument() interface{}
}
type ProjectDoc struct {
	Name           string    `json:"name"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
}

func (pd *ProjectDoc) GetDocument() interface{} {
	return bson.D{{"name", pd.Name}, {"lastModifiedAt", pd.LastModifiedAt}}
}
