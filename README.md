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
```

If you want to provide your own `http.Client`, you can do it :

```go
httpClient := &http.Client{}
client := qradar.NewClient(httpClient)
```

If you want to downgrade the default version (which is 12.0), you do it as follow :

```go
client.Version = "7.0"
```

Then you can start using it.

```go
fields := "id,description,status,assigned_to,magnitude,start_time,last_updated_time,follow_up,offense_source,offense_type"
filter := "status = \"OPEN\""

// List the offenses
offenses, total, err := s.client.SIEM.ListOffenses(ctx, fields, filter, "", 0, 40)
```