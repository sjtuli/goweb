package article

import (
	"github.com/gorilla/mux"
	"goblog/app/models"
	"goblog/pkg/route"
	"strconv"
)

var router = mux.NewRouter()

// Article 文章模型
type Article struct {
	models.BaseModel

	ID    uint64
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}
