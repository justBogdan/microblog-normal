package models

type Post struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	PostID int32  `json:"post_id"`
	Likes  []Like `json:"likes"`
}
