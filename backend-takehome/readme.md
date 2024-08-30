# Migration 
install goose : 
```go install github.com/pressly/goose/v3/cmd/goose@latest```

run : 
goose -dir migration/ mysql "{{database_username}}:{{database_username}}@tcp(localhost:3306)/{{database_name}}?parseTime=true" up
