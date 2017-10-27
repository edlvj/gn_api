package controllers

import (
	"github.com/revel/revel"
	"github.com/edlvj/gn-api/app/database"
	"github.com/edlvj/gn-api/app/models"
    "github.com/edlvj/gn-api/app/forms"
	"gopkg.in/mgo.v2/bson"
)

type Likes struct {
	*revel.Controller
}

func (c Likes) Create(post_id, like_type string) revel.Result {
    post := models.Post{}
    likeForm := forms.LikeForm{post_id, like_type}
    change_type := bson.M{}

    likeForm.Validate(c.Validation)

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        return c.RenderJSON(c.Validation);
    }  

    id := bson.M{"_id": bson.ObjectIdHex(likeForm.PostId)}; 

    if err := database.Posts.Find(id).One(&post); err != nil {
        return c.RenderJSON(err);
    } 

    if likeForm.LikeType == "like" {
        change_type = bson.M{"like.likes": 1}
    } else if likeForm.LikeType == "dislike" {
        change_type = bson.M{"like.dislikes": 1}
    }; 

    if err := database.Posts.Update(bson.M{"_id": post.ID}, bson.M{"$inc": change_type }); err != nil {
        return c.RenderJSON(err);
    }
    return c.RenderJSON(post);
}