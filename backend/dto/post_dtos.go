package dtos

type CreatePostReqBody struct {
	Title  string `json:"title" validate:"required,min=3,endsnotwith=."`
	Body   string `json:"body" validate:"required,min=10"`
	UserID uint   `json:"user_id" validate:"required,min=1"`
}
