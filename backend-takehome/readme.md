# Migration 
install goose : 
```go install github.com/pressly/goose/v3/cmd/goose@latest```

run : 
goose -dir migration/ mysql "root:Root_123@tcp(localhost:3306)/bythenapp?parseTime=true" up
