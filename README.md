# Project Requirements
***[Take-home backend challenge for](https://github.com/fetch-rewards/receipt-processor-challenge)***

### Development Requirements
- Use local cache.
- Shall assume docker is already installed.
- Use the described APIs.

### POST
***"/receipts/{id}/points"***\
***Input: receipt object in JSON**
- Calculate rewarded points.
- Store the unmarshalled receipt, and the rewarded points with the same id through a different map for easy access. (May combine, depending on preferences.)[^1]

[^1]All cache-related files are placed within the same module as preferred by Golang. (eg, points, cache and llm).\
***Return: the generated "id" for that object in JSON***

## GET
***"/receipts/{id}/points"***
***Input: JSON "id"***
- Look up receipt with the extracted id.\
***Return: the rewarded total points in JSON***

# Instructions Running Application
### Golang
```
cd project
go run main.go
```

### Docker
```
cd project
docker build -t fetch-reward .
docker run -it --rm --name fetch-reward fetch-reward
```
### Defalt Port
http://localhost:8080

### Installation info
[Golang installation](https://go.dev/doc/install)
[Docker installation](https://www.docker.com/get-started)










