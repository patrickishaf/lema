package dtos

type CreatePostReqBody struct {
	Title  string `json:"title" validate:"required,min=3"`
	Body   string `json:"body" validate:"required,min=10"`
	UserID string `json:"user_id" validate:"required,min=1"`
}
