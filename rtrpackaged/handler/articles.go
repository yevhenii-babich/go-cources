package handler

import (
	"fmt"
	"net/http"

	"github.com/yevhenii-babich/go-cources/rtrpackaged/mw"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yevhenii-babich/go-cources/rtrpackaged/db"
	"github.com/yevhenii-babich/go-cources/rtrpackaged/models"
)

func ListArticles(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderList(w, r, NewArticleListResponse(db.List())); err != nil {
		_ = render.Render(w, r, mw.ErrRender(err))
		return
	}
}

// SearchArticles searches the Articles data for a matching article.
// It's just a stub, but you get the idea.
func SearchArticles(w http.ResponseWriter, r *http.Request) {
	list := db.List()
	list[0].ID = "10"
	_ = render.RenderList(w, r, NewArticleListResponse(list))
}

// CreateArticle persists the posted Article and returns it
// back to the client as an acknowledgement.
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var data ArticleRequest
	if err := render.Bind(r, &data); err != nil {
		_ = render.Render(w, r, mw.ErrInvalidRequest(err))
		return
	}

	article := data.Article
	if _, err := db.NewArticle(article); err != nil {
		_ = render.Render(w, r, mw.ErrInvalidRequest(err))
	}

	render.Status(r, http.StatusCreated)
	_ = render.Render(w, r, NewArticleResponse(article))
}

// GetArticle returns the specific Article. You'll notice it just
// fetches the Article right off the context, as its understood that
// if we made it this far, the Article must be on the context. In case
// its not due to a bug, then it will panic, and our Recoverer will save us.
func GetArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(chi.URLParam(r, "articleID"))

	// Assume if we've reach this far, we can access the article
	// context because this handler is a child of the ArticleCtx
	// middleware. The worst case, the recoverer middleware will save us.
	article := r.Context().Value("article").(*models.Article)

	if err := render.Render(w, r, NewArticleResponse(article)); err != nil {
		_ = render.Render(w, r, mw.ErrRender(err))
		return
	}
}

// UpdateArticle updates an existing Article in our persistent store.
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*models.Article)

	data := &ArticleRequest{Article: article}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, mw.ErrInvalidRequest(err))
		return
	}
	article = data.Article
	_, _ = db.UpdateArticle(article.ID, article)

	_ = render.Render(w, r, NewArticleResponse(article))
}

// DeleteArticle removes an existing Article from our persistent store.
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	var err error

	// Assume if we've reach this far, we can access the article
	// context because this handler is a child of the ArticleCtx
	// middleware. The worst case, the recoverer middleware will save us.
	article := r.Context().Value("article").(*models.Article)

	article, err = db.RemoveArticle(article.ID)

	if err != nil {
		_ = render.Render(w, r, mw.ErrInvalidRequest(err))
		return
	}

	_ = render.Render(w, r, NewArticleResponse(article))
}
