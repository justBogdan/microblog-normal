package models

type Like struct {
	LikeId          int32  `json:"like_id"`
	FromWhoNickName string `json:"from_id"`
}
