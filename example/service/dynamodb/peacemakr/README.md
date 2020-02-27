# Example

[Peacemakr](https://peacemakr.io/) is Secure Data Platform that abstracts away all the hard problems of data security behind a simple interface. This example shows how the [Go SDK](https://github.com/peacemakr-io/peacemakr-go-sdk) integrates with Dynamodb. 

## Getting Started

Make sure you have docker installed and pull a local dynamodb from Dockerhub.

```sh
docker pull amazon/dynamodb-local 
# Start dynamodb at 8000
docker run -p 8000:8000 amazon/dynamodb-local    
```

## Running the example

To run the example, make sure you have the dep installed.
```sh
# get deps
go mod vendor 

# run create_encrypted example
go run create_encrypted_item.go

# run read_encrypted example
go run read_encrypted_item.go
# Example output:
#  Found item:
#    Year:   2020
#    Title:  Encrypted Text
#    Plot:   Encrypted Nothing happens at all.
#    Rating: 2.0
```


## Encrypt Semi-Structured Data
Peacemakr Go SDK provide tools that makes encrypting semi-structured data easy. 

To Encrypt a field, simply add a `encrypt:"true"` tag next to the field. As shown in the example:
```Go
type ItemInfo struct {
	Plot   string `json:"plot" encrypt:"true"`
	Rating string `json:"rating"`
}

type Item struct {
	Year  int      `json:"year"`
	Title string   `json:"title" encrypt:"true"`
	Info  ItemInfo `json:"info"`
}
```
The tool will automatically detect and encrypt the data.


