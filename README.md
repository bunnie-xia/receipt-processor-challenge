# Project Requirements

### Development Requirements Summary
- Data does not need to persist and cache information in memory.
- Shall assume docker is already installed.
- Use the described APIs.

### POST
**Endpoint:** "/receipts/{id}/points"

**Input:** Receipt object in JSON
- Calculate rewarded points.
- Cache the unmarshalled receipt, and the rewarded points with the same id through different maps for easy access. (May combine, depending on preferences.)
 - All cache-related files are placed within the same module as preferred by Golang. (eg, points, cache and llm).

**Return:** the generated "id" for that object in JSON

### GET
**Endpoint:** "/receipts/{id}/points"

**Input:** JSON "id"
- Look up the receipt with the extracted id.

**Return:** the rewarded total points in JSON

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

## Installation info
Golang installation: https://go.dev/doc/install) <br/>
Docker installation: https://www.docker.com/get-started










