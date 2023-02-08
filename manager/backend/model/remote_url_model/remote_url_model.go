package remote_url_model

type PageUrl struct {
	Id      uint32 `json:"id"`
	Title   string `json:"title" validate:"required,max=32"`
	PageUrl string `json:"page_url" validate:"required,max=512"`
}
