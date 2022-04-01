# ipatser
---

## Deployment : Container
---
```bash
docker-compose down; docker image prune -f; docker rmi vcs_ipatser; docker-compose up;
```

```
name: vcs_ipatser
size: 13.1MB
base: alpine
```

## GO MOD
---

```bash
go mod init main
```
```
go: creating new go.mod: module main
go: to add module requirements and sums:
	go mod tidy
```
```bash
go test
```
```
main.go:10:5: no required module provides package github.com/gorilla/mux; to add it:
	go get github.com/gorilla/mux
main.go:11:5: no required module provides package github.com/lib/pq; to add it:
	go get github.com/lib/pq
```
```bash
go get github.com/gorilla/mux
```
```
go: downloading github.com/gorilla/mux v1.8.0
go: added github.com/gorilla/mux v1.8.0
```
```bash
go get github.com/lib/pq
```
```
go: downloading github.com/lib/pq v1.10.4
go: added github.com/lib/pq v1.10.4
```
```bash
ls
```
```
go.mod  go.sum  main.go  script
```

## LIST ALL REST API ENDPOINTS
---

```bash
sudo apt install -y jq
```

### LIST ALL ENTRIES
```bash
curl -X GET 'http://localhost:8000/movies' | jq
```
```bash
{
  "type": "success",
  "data": [
    {
      "movieid": "1",
      "moviename": "movie3"
    },
    {
      "movieid": "3",
      "moviename": "movie1"
    },
    {
      "movieid": "2",
      "moviename": "movie2"
    }
  ],
  "message": ""
}
```

### INSERT ENTRY
```bash
curl -X POST 'http://localhost:8000/movies?movieid=4&moviename=movie4'
```
```bash
{"type":"success","data":null,"message":"The movie has been inserted successfully!"}
```

```bash
curl -X GET 'http://localhost:8000/movies' | jq
```
```bash
{
  "type": "success",
  "data": [
    {
      "movieid": "1",
      "moviename": "movie3"
    },
    {
      "movieid": "3",
      "moviename": "movie1"
    },
    {
      "movieid": "2",
      "moviename": "movie2"
    },
    {
      "movieid": "4",
      "moviename": "movie4"
    }
  ],
  "message": ""
}
```

### DELETE BY ID
```bash
curl -X DELETE 'http://localhost:8000/movies/1'
```
```bash
{"type":"success","data":null,"message":"The movie has been deleted successfully!"}
```
```bash
curl -X GET 'http://localhost:8000/movies' | jq
```
```bash
{
  "type": "success",
  "data": [
    {
      "movieid": "3",
      "moviename": "movie1"
    },
    {
      "movieid": "2",
      "moviename": "movie2"
    },
    {
      "movieid": "4",
      "moviename": "movie4"
    }
  ],
  "message": ""
}
```

### DELETE ALL ENTRIES
```bash
curl -X DELETE 'http://localhost:8000/movies'
```
```bash
{"type":"success","data":null,"message":"All movies have been deleted successfully!"}
```
```bash
curl -X GET 'http://localhost:8000/movies' | jq
```
```bash
{
  "type": "success",
  "data": null,
  "message": ""
}
```
