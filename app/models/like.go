package models

/*
Like model
*/

type Like struct {
	Likes int        `json:"likes" bson:"likes"`
	Dislikes int     `json:"dislikes" bson:"dislikes"`
}
