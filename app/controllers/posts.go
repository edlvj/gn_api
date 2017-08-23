package controllers

import (
	"net/http"

	"github.com/revel/revel"
	"github.com/edlvj/gn-api/app/database"
	"github.com/edlvj/gn-api/app/models"
	"gopkg.in/mgo.v2/bson"
)

type Posts struct {
	*revel.Controller
}

func (c Posts) Index() revel.Result {
	results := []models.Post{}
	
	if err := database.Posts.Find(bson.M{}).All(&results); err != nil {
		return c.RenderJSON(err);
	} 
	return c.RenderJSON(results)
}

func (c Posts) Create(title, text, post_type, image_url string) revel.Result {
	post := models.Post{ Like: models.Like{0, 0}}

	post.Title = title
	post.Text = text
	post.PostType = post_type
	post.ImageUrl = image_url

	post.Validate(c.Validation)

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        return c.RenderJSON(c.Validation);
    }

    if err := database.Posts.Insert(post); err != nil {
	 	return c.RenderJSON(err);
	}

	c.Response.Status = http.StatusCreated
	return c.RenderJSON(post) 
}