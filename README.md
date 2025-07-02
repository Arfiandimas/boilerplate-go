# {{PROJECT_NAME}}
This boilerplate support Golang version 1.13.x or latest

## Query Builder Documentation
Using depiq Query builder please see documentation [depiq](https://github.com/orn-id/depiq/tree/master/docs)

## Migrate Command
### Create migrate file
``` bash
go run src/main.go db:migrate create table_schema_or_any sql
```
### Run Migrate UP
``` bash
go run src/main.go db:migrate up
```

### Run Migrate Down
``` bash
go run src/main.go db:migrate down
```



## Getting Started
For local development install golang latest version or min version 1.13.x. after install golang next install [air](https://github.com/cosmtrek/air).
```bash
sudo curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
```
Install service dependecy

package dependecy
```
go mod tidy
```

service dependecy
```
docker-compose up -d
```
Running air for up the server
```
air
```
Start your enhancement

# Documentation
## Install Godoc
```
go install -v golang.org/x/tools/cmd/godoc
```

## Run Server
```
godoc -http=:3000
```

## Documentation
After running doc server please open this Link [Documentation](http://localhost:3000/pkg/github.com/oni-kit/{{PROJECT_NAME}}/)
