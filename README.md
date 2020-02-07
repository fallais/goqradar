# goqradar

**goqradar** is a library written in Golang that helps you with the **IBM QRadar REST API**.

## Usage

```go
import "github.com/fallais/goqradar"
```

Construct a new QRadar client, then use the various services on the client to
access different parts of the QRadar API. For example:

```go
client := qradar.NewClient(nil)

// List offenses
offenses, err := client.SIEM.ListOffenses(context.Background(), opt)
```

If you want to provide your own `http.Client`, or change the version, you can do it :

```go
httpClient := &http.Client{}

client := qradar.NewClient(httpClient)

client.Version = "7.0"
```
