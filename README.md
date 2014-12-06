NextBus API for Go
==================

This is an API library for reading data from the NextBus public JSON feed in Go.

```go
import "github.com/cogspace/nextbus"
```

Things you can do
-----------------

### Get all the agencies NextBus knows about.


```go
agencies, err := nextbus.GetAgencies()
```

### Get all the routes for a given agency.

```go
routes, err := nextbus.GetRoutes("glendale")
```

### Retrieve route information (terse).

```go
routeConfig, err := nextbus.GetRouteConfig("glendale", "12", false, true)
```

### Retrieve schedule information for a route.

```go
schedules, err := nextbus.GetSchedules("glendale", "12")
```

### Get bus predictions for one stop.

```go
predictions, err := nextbus.GetPredictions("glendale", "12", "gtc_d")
```

### Get predictions for multiple stops.

```go
predictions, err := nextbus.GetPredictionsMulti(
    "glendale",
    []nextbus.Stop{
        { Route: "12", Stop: "gtc_d" },
        { Route: "12", Stop: "sanchev" },
    },
)
```

Nested Types
------------
In the interest of concise, sequential code, some of the types are composed
of rather deeply nested anonymous subtypes, for example:

```go
type Schedule struct {
	ServiceClass string
	Title        string
	Tr           []struct {
		Stop []struct {
			Content   string
			Tag       string
			EpochTime string
		}
		BlockId string
	}
	Direction string
	Tag       string
	Header    struct {
		Stop []struct {
			Content string
			Tag     string
		}
	}
	ScheduleClass string
}
```

The only real problem with this is declaring a type for, say, `Tr` can be...
cumbersome.

```go
func getBlockId(tr struct{ Stop []struct{ Content, Tag, EpochTime string}; BlockId string }) string {
    return tr.BlockId
}

func getStops(tr struct{ Stop []struct{ Content, Tag, EpochTime string }; BlockId string }) []struct{ Content, Tag, EpochTime string } {
    // Now that's just silly...
    return tr.Stop
}
```

To alleviate this problem, simply name whichever types you need to use.

```go
type Stop struct{ Content, Tag, EpochTime string }
type Tr struct{ Stop []struct{ Content, Tag, EpochTime string }; BlockId string }

func getBlockId(tr Tr) string {
    return tr.BlockId
}

func getStops(tr Tr) []Stop {
    return tr.Stop
}
```

In practice, naming theses types is not usually necessary in Go client code.

```go
ch := make(chan string)
go func(ch chan string) {
    schedules, err := nextbus.GetSchedules("glendale", "12")
    if err != nil {
        log.Fatal(err)
    }

    for _, schedule := range schedules {
        for _, tr := range schedule.Tr {
            ch <- tr.BlockId
        }
    }
}(ch)
...
```

Further Reading
---------------
The NextBus feed [documentation][1] only covers the XML API, but the JSON API is very similar.

[1]: http://www.nextbus.com/xmlFeedDocs/NextBusXMLFeed.pdf
