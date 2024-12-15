package Schemas


import (
"go.mongodb.org/mongo-driver/bson/primitive"
)

type Video struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	VideoName      string             `json:"video_name" bson:"video_name"`
	Uploader       string             `json:"uploader_username" bson:"uploader_username"`
	Description    string             `json:"description" bson:"description"`
	Tags           []string           `json:"tags" bson:"tags"`
	Comments       []Comment          `json:"comments" bson:"comments"`
	VideoID        string             `json:"video_id" bson:"video_id"`
}
