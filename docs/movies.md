# Movies
---

```bash
sudo apt install -y jq
```

### LIST ALL ENTRIES
---

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
---

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
---

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
---

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
