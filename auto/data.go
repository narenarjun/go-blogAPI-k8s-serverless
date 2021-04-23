package auto

import "blogapi/api/models"

var users = []models.User{
	models.User{Nickname: "redpanda", Email: "redpanda@chilling.com", Password: "skdhaklhfsd"},
}

var posts = []models.Post{
	models.Post{
		Title:   "Post 1",
		Content: "Welcome,  Hello world !!!",
	},
}
