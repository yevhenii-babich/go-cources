Welcome to the chi/_examples/rest generated docs.

## Routes

<details>
<summary>`/`</summary>

- [RequestID](/middleware/request_id.go#L67)
- [Logger](/middleware/logger.go#L41)
- [Recoverer](/middleware/recoverer.go#L21)
- [URLFormat](/middleware/url_format.go#L47)
- [SetContentType.func1](https://github.com/yevhenii-babich/go-cources/rtrpackaged/vendor/github.com/go-chi/render/content_type.go#L49)
- **/**
  - _GET_
  - [main.main.func1](https://github.com/yevhenii-babich/go-cources/rtrpackaged/main.go#L67)

</details>
<details>
<summary>`/admin`</summary>

- [RequestID](/middleware/request_id.go#L67)
- [Logger](/middleware/logger.go#L41)
- [Recoverer](/middleware/recoverer.go#L21)
- [URLFormat](/middleware/url_format.go#L47)
- [SetContentType.func1](https://github.com/yevhenii-babich/go-cources/rtrpackaged/vendor/github.com/go-chi/render/content_type.go#L49)
- **/admin**
  - [main.AdminOnly](https://github.com/yevhenii-babich/go-cources/rtrpackaged/main.go#L134)
  - **/**
  - _GET_
  - [main.adminRouter.func1](https://github.com/yevhenii-babich/go-cources/rtrpackaged/main.go#L121)

</details>
<details>
<summary>`/admin/accounts`</summary>

- [RequestID](/middleware/request_id.go#L67)
- [Logger](/middleware/logger.go#L41)
- [Recoverer](/middleware/recoverer.go#L21)
- [URLFormat](/middleware/url_format.go#L47)
- [SetContentType.func1](https://github.com/yevhenii-babich/go-cources/rtrpackaged/vendor/github.com/go-chi/render/content_type.go#L49)
- **/admin**
  - [main.AdminOnly](https://github.com/yevhenii-babich/go-cources/rtrpackaged/main.go#L134)
  - **/accounts**
  - _GET_
  - [main.adminRouter.func2](https://github.com/yevhenii-babich/go-cources/rtrpackaged/main.go#L124)

</details>
<details>
<summary>`/admin/users/{userId}`</summary>

- [RequestID](/middleware/request_id.go#L67)
- [Logger](/middleware/logger.go#L41)
- [Recoverer](/middleware/recoverer.go#L21)
- [URLFormat](/middleware/url_format.go#L47)
- [SetContentType.func1](https://github.com/yevhenii-babich/go-cources/rtrpackaged/vendor/github.com/go-chi/render/content_type.go#L49)
- **/admin**
  - [main.AdminOnly](https://github.com/yevhenii-babich/go-cources/rtrpackaged/main.go#L134)
  - **/users/{userId}**
  - _GET_
  - [main.adminRouter.func3](https://github.com/yevhenii-babich/go-cources/rtrpackaged/main.go#L127)

</details>
<details>
<summary>`/articles`</summary>

- [RequestID](/middleware/request_id.go#L67)
- [Logger](/middleware/logger.go#L41)
- [Recoverer](/middleware/recoverer.go#L21)
- [URLFormat](/middleware/url_format.go#L47)
- [SetContentType.func1](https://github.com/yevhenii-babich/go-cources/rtrpackaged/vendor/github.com/go-chi/render/content_type.go#L49)
- **/articles**
  - **/**
  - _GET_
  - [main.paginate](https://github.com/yevhenii-babich/go-cources/rtrpackaged/main.go#L147)
  - [ListArticles](https://github.com/yevhenii-babich/go-cources/rtrpackaged/handler/articles.go#L15)
  - _POST_
  - [CreateArticle](https://github.com/yevhenii-babich/go-cources/rtrpackaged/handler/articles.go#L30)

</details>
<details>
<summary>`/articles/search`</summary>

- [RequestID](/middleware/request_id.go#L67)
- [Logger](/middleware/logger.go#L41)
- [Recoverer](/middleware/recoverer.go#L21)
- [URLFormat](/middleware/url_format.go#L47)
- [SetContentType.func1](https://github.com/yevhenii-babich/go-cources/rtrpackaged/vendor/github.com/go-chi/render/content_type.go#L49)
- **/articles**
  - **/search**
  - _GET_
  - [SearchArticles](https://github.com/yevhenii-babich/go-cources/rtrpackaged/handler/articles.go#L24)

</details>
<details>
<summary>`/articles/{articleID}`</summary>

- [RequestID](/middleware/request_id.go#L67)
- [Logger](/middleware/logger.go#L41)
- [Recoverer](/middleware/recoverer.go#L21)
- [URLFormat](/middleware/url_format.go#L47)
- [SetContentType.func1](https://github.com/yevhenii-babich/go-cources/rtrpackaged/vendor/github.com/go-chi/render/content_type.go#L49)
- **/articles**
  - **/{articleID}**
  - [ArticleCtx](https://github.com/yevhenii-babich/go-cources/rtrpackaged/mw/middleware.go#L16)
  - **/**
  - _GET_
  - [GetArticle](https://github.com/yevhenii-babich/go-cources/rtrpackaged/handler/articles.go#L50)
  - _PUT_
  - [UpdateArticle](https://github.com/yevhenii-babich/go-cources/rtrpackaged/handler/articles.go#L65)
  - _DELETE_
  - [DeleteArticle](https://github.com/yevhenii-babich/go-cources/rtrpackaged/handler/articles.go#L80)

</details>
<details>
<summary>`/articles/{articleSlug:[a-z-]+}`</summary>

- [RequestID](/middleware/request_id.go#L67)
- [Logger](/middleware/logger.go#L41)
- [Recoverer](/middleware/recoverer.go#L21)
- [URLFormat](/middleware/url_format.go#L47)
- [SetContentType.func1](https://github.com/yevhenii-babich/go-cources/rtrpackaged/vendor/github.com/go-chi/render/content_type.go#L49)
- **/articles**
  - **/{articleSlug:[a-z-]+}**
  - _GET_
  - [ArticleCtx](https://github.com/yevhenii-babich/go-cources/rtrpackaged/mw/middleware.go#L16)
  - [GetArticle](https://github.com/yevhenii-babich/go-cources/rtrpackaged/handler/articles.go#L50)

</details>
<details>
<summary>`/panic`</summary>

- [RequestID](/middleware/request_id.go#L67)
- [Logger](/middleware/logger.go#L41)
- [Recoverer](/middleware/recoverer.go#L21)
- [URLFormat](/middleware/url_format.go#L47)
- [SetContentType.func1](https://github.com/yevhenii-babich/go-cources/rtrpackaged/vendor/github.com/go-chi/render/content_type.go#L49)
- **/panic**
  - _GET_
  - [main.main.func3](https://github.com/yevhenii-babich/go-cources/rtrpackaged/main.go#L75)

</details>
<details>
<summary>`/ping`</summary>

- [RequestID](/middleware/request_id.go#L67)
- [Logger](/middleware/logger.go#L41)
- [Recoverer](/middleware/recoverer.go#L21)
- [URLFormat](/middleware/url_format.go#L47)
- [SetContentType.func1](https://github.com/yevhenii-babich/go-cources/rtrpackaged/vendor/github.com/go-chi/render/content_type.go#L49)
- **/ping**
  - _GET_
  - [main.main.func2](https://github.com/yevhenii-babich/go-cources/rtrpackaged/main.go#L71)

</details>