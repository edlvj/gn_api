package models

import "github.com/revel/revel"
import "gopkg.in/mgo.v2/bson"

/*
Post model
*/
type Post struct {
	ID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string        `json:"title" bson:"title"`
	Text  string        `json:"text" bson:"text"`
	PostType string         `json:"post_type" bson:"post_type"`
	ImageUrl string        `json:"image_url" bson:"image_url"`
	Like                `json:"like" bson:"like"`
}


func (post *Post) Validate(v *revel.Validation) {
	v.Required(post.Title)
	v.Required(post.PostType == "url" || post.PostType == "news").Key("post_type").
		Message("The incorrect post type.")
}
