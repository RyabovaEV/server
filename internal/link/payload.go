package link

// LinkCreateRequest структора запроса на создание ссылки
type LinkCreateRequest struct {
	URL string `json:"url" validate:"required,url"`
}

// LinkUpdateRequest структура запроса на обновление ссылки
type LinkUpdateRequest struct {
	URL  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

// GetAllLinksResponce структура запроса на получение ссылок
type GetAllLinksResponce struct {
	Links []Link `json:"links"`
	Count int64  `json:"count"`
}
