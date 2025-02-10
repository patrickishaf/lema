package dtos

type CreatePostReqBody struct {
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
	AuthorId uint   `json:"author_id" validate:"required"`
}
