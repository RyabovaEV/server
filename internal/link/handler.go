// Package link обработка HTTP запросов для работы со ссылками
package link

import (
	"net/http"
	"server/pkg/req"
	"server/pkg/res"
	"strconv"

	"gorm.io/gorm"
)

// LinkHendlerDeps связь с конфигурацией
type LinkHendlerDeps struct {
	LinkRepository *LinkRepository
}

// LinkHendler структура хендлера для ссылок
type LinkHendler struct {
	LinkRepository *LinkRepository
}

// NewLinkHendler создание нового хендлера
func NewLinkHendler(router *http.ServeMux, deps LinkHendlerDeps) {
	handler := &LinkHendler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

// GoTo получение ссылки
func (handler *LinkHendler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, link.URL, http.StatusTemporaryRedirect)
	}
}

// Create вставка ссылки
func (handler *LinkHendler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}

		link := NewLink(body.URL)
		for {

			existedLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}

		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdLink, 201)
	}

}

// UpdateLink изменение ссылки
func (handler *LinkHendler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// парсим боди
		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		// проверяем id что мы его модем прочитать
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// создаем изменение (это относится к бизнес логике)
		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			URL:   body.URL,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, link, 201)
	}
}

// Delete удаление ссылки
func (handler *LinkHendler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = handler.LinkRepository.GetByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, nil, 200)
	}
}
