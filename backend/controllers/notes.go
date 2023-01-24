package controllers

type NoteInput struct {
	Title       string `json:"title" binding:"required,gte=1"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required,gte=1"`
}
