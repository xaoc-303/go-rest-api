# go-rest-api

```
go run *.go
```

| explanation | method | uri | additional |
| --- | --- | --- | --- |
| home page | GET | http://localhost:8080/ | |
| say hello | GET | http://localhost:8080/hello/ | |
| print headers | GET | http://localhost:8080/headers/ | |
| get all articles | GET | http://localhost:8080/articles/ | |
| get article | GET | http://localhost:8080/article/1/ | |
| new article | POST | http://localhost:8080/article/ | {"title":"newTitle","desc":"newDesc","content":"newContent"} |
| delete article | DELETE | http://localhost:8080/article/1/ | |
| update article | PUT | http://localhost:8080/article/1/ | {"title":"updTitle","desc":"updDesc","content":"updContent"} |
