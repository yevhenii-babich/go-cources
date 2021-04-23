package mw

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/yevhenii-babich/go-cources/rtrpackaged/db"
	"github.com/yevhenii-babich/go-cources/rtrpackaged/models"
)

// ArticleCtx middleware is used to load an Article object from
// the URL parameters passed through as the request. In case
// the Article could not be found, we stop here and return a 404.
func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var article *models.Article
		var err error

		if articleID := chi.URLParam(r, "articleID"); articleID != "" {
			article, err = db.GetArticle(articleID)
		} else if articleSlug := chi.URLParam(r, "articleSlug"); articleSlug != "" {
			article, err = db.GetArticleBySlug(articleSlug)
		} else {
			_ = render.Render(w, r, ErrNotFound)
			return
		}

		if err != nil {
			_ = render.Render(w, r, ErrNotFound)
			return
		}

		//goland:noinspection ALL
		ctx := context.WithValue(r.Context(), "article", article) //nolint:revive,staticcheck
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
