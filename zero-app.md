### 1. "Sign up"

1. route definition

- Url: /auth/signUp
- Method: POST
- Request: `signUpReq`
- Response: `messageRes`

2. request definition



```golang
type SignUpReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age int `json:"age"`
}
```


3. response definition



```golang
type MessageRes struct {
	Message string `json:"message"`
}
```

### 2. "Sign In"

1. route definition

- Url: /auth/signIn
- Method: GET
- Request: `signInReq`
- Response: `signInRes`

2. request definition



```golang
type SignInReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type SignInRes struct {
	Token string `json:"token"`
}
```

### 3. "Get profile by token"

1. route definition

- Url: /user/profile
- Method: GET
- Request: `-`
- Response: `getProfileRes`

2. request definition



3. response definition



```golang
type GetProfileRes struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age int `json:"age"`
}
```

### 4. "Update profile"

1. route definition

- Url: /user/profile
- Method: PUT
- Request: `updateProfileReq`
- Response: `messageRes`

2. request definition



```golang
type UpdateProfileReq struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age int `json:"age"`
}
```


3. response definition



```golang
type MessageRes struct {
	Message string `json:"message"`
}
```

### 5. "Create article"

1. route definition

- Url: /articles
- Method: POST
- Request: `createArticleReq`
- Response: `messageRes`

2. request definition



```golang
type CreateArticleReq struct {
	Title string `json:"title"`
	Body string `json:"body"`
}
```


3. response definition



```golang
type MessageRes struct {
	Message string `json:"message"`
}
```

### 6. "Get article by id"

1. route definition

- Url: /articles/:article_id
- Method: GET
- Request: `getArticleReq`
- Response: `getArticleRes`

2. request definition



```golang
type GetArticleReq struct {
	ArticleId int `path:"article_id"`
}
```


3. response definition



```golang
type GetArticleRes struct {
	Title string `json:"title"`
	Body string `json:"body"`
}
```

### 7. "Get all articles"

1. route definition

- Url: /articles
- Method: GET
- Request: `-`
- Response: `getArticleRes`

2. request definition



3. response definition



```golang
type GetArticleRes struct {
	Title string `json:"title"`
	Body string `json:"body"`
}
```

### 8. "Update article"

1. route definition

- Url: /articles/:article_id
- Method: PUT
- Request: `updateArticleReq`
- Response: `messageRes`

2. request definition



```golang
type UpdateArticleReq struct {
	Id int `path:"article_id"`
	Title string `json:"title"`
	Body string `json:"body"`
}
```


3. response definition



```golang
type MessageRes struct {
	Message string `json:"message"`
}
```

### 9. "Delete article"

1. route definition

- Url: /articles/:article_id
- Method: DELETE
- Request: `deleteArticleReq`
- Response: `messageRes`

2. request definition



```golang
type DeleteArticleReq struct {
	Id int `path:"article_id"`
}
```


3. response definition



```golang
type MessageRes struct {
	Message string `json:"message"`
}
```

