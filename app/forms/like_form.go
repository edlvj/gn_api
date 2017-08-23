package forms

import "github.com/revel/revel"

/*
LikeForm form
*/

type LikeForm struct {
	PostId string    `bson:"post_id"`
	LikeType string  `bson:"like_type"`
}

func (like *LikeForm) Validate(v *revel.Validation) {
	v.Required(like.PostId)
	v.Required(like.LikeType == "like" || like.LikeType == "dislike").Key("like_type").
		Message("The incorrect like type.")
}