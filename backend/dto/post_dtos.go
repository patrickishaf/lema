package dtos

type CreatePostReqBody struct {
	Title  string `json:"title" validate:"required"`
	Body   string `json:"body" validate:"required"`
	UserID uint   `json:"user_id" validate:"required"`
}
