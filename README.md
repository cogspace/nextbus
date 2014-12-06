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

Further Reading
---------------
The NextBus feed [documentation][1] only covers the XML API, but the JSON API is very similar.

[1]: http://www.nextbus.com/xmlFeedDocs/NextBusXMLFeed.pdf
