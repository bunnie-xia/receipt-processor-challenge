
### Requirements
# Development Requirements
- Use local cache.
- Shall assume already have docker installed on testing end.
- Use described API 

# POST
Input: takes in receipt in JSON

- Calculate rewarded points
- Unmarshal the receipt to the data that's ready to use
- Store the unmarshaled receipt, and rewarded-points with the same generated id in different map for easy access. (May combine, depends on preferences.)
  - All related files: points, cache and llm are in same module as prefered by Golang

Return: the generated "id" for that object in JSON (that's ready to pass in path "/receipts/{id}/points")

# GET
Input: JSON "id".
- looks up receipt with the extracted id
Return: the rewarded total points in JSON

Take-home backend challenge for (https://github.com/fetch-rewards/receipt-processor-challenge)

### Instructions Running Application
# Golang
`cd project`
`go run main.go`

localhost:8080

# Docker
`cd project`
`docker build -t fetch-reward .`
`docker run -it --rm --name fetch-reward fetch-reward`


### Installation info
- Golang installation: https://go.dev/doc/install
- Docker installation: https://www.docker.com/get-started

[Defalt Port](http://localhost:8080)

(https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax#links)
.. may add test..
