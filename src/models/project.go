package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type ProjectDocument interface {
	GetFullDocument() interface{}
	GetBaseDocument() interface{}
	GetSeedPaperDocument() interface{}
}

type Paper struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type ProjectDoc struct {
	Name           string    `json:"name"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	Description    string    `json:"description"`
	Tags           []string  `json:"tags"`
	User           string    `json:"user"`
	SeedPapers     []Paper   `json:"seedPapers"`
	ReadingList    []Paper   `json:"readingList"`
}

func (pd *ProjectDoc) GetFullDocument() interface{} {
	return bson.D{{"name", pd.Name}, {"lastModifiedAt", pd.LastModifiedAt}, {"description", pd.Description}, {"tags", pd.Tags}, {"user", pd.User}, {"seedPapers", pd.SeedPapers}, {"readingList", pd.ReadingList}}
}

func (pd *ProjectDoc) GetBaseDocument() interface{} {
	return bson.D{{"name", pd.Name}, {"lastModifiedAt", pd.LastModifiedAt}, {"description", pd.Description}, {"tags", pd.Tags}, {"user", pd.User}}
}

func (pd *ProjectDoc) GetSeedPaperDocument() interface{} {
	return bson.D{{"name", pd.Name}, {"seedPapers", pd.SeedPapers}}
}
