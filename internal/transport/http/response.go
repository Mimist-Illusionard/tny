package http

type CreateShortLinkResponse struct {
	ID          string `json:"id"`
	ShortCode   string `json:"shortCode"`
	OriginalURL string `json:"originalUrl"`
}

type GetOriginalLinkResponse struct {
	OriginalURL string `json:"originalUrl"`
}
