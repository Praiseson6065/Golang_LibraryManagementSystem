
## To Run This Server
```
go mod download
go mod tidy
go run ./cmd/
```


```bash
cd cart
swag init --generalInfo ../cmd/user.go -o ../docs/usercart --instanceName usercart --parseDependency --parseInternal --parseDepth 5 --dir ../cmd,../cart
```