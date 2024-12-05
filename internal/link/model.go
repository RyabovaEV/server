package link

import (
	"math/rand/v2"

	"gorm.io/gorm"
)

// Link модель ссылки
type Link struct {
	gorm.Model
	URL  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

// NewLink функция конструктор для Link
func NewLink(url string) *Link {
	link := &Link{
		URL:  url,
		Hash: RandStringRunes(10),
	}
	link.GenerateHash()
	return link
}

// GenerateHash генерация HAsh
func (link *Link) GenerateHash() {
	link.Hash = RandStringRunes(10)
}

var lettersRunes = []rune("abcdifghigklmnopqrstuvwxyzABCDIFGHIGKLMNOPQRSTUVWXYZ")

// RandStringRunes рандомное создание HAsh: n - кол во символов в hash
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = lettersRunes[rand.IntN(len(lettersRunes))]
	}
	return string(b)
}
