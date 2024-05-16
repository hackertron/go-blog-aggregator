# go-blog-aggregator

## build instructions
1. goose <postgres-url> up for migrations
2. sqlc generate
3. go build && ./blog-agg

## packages used
1. echo : https://echo.labstack.com/
2. goose : https://github.com/pressly/goose
3. sqlc : https://sqlc.dev/