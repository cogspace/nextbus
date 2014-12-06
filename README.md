NextBus API for Go
==================

This is an API library for reading data from the NextBus public JSON feed in Go.

Example
-------

```go
package main

import (
    "nextbus"
    "fmt"
    "log"
)

agencies, err := nextbus.GetAgencies()
    if err != nil {
        log.Fatal(err)
    }

for _, agency := range agencies {
    fmt.Println(agency.Tag)
}
```

Further Reading
---------------
The NextBus feed [documentation][1] only covers the XML API, but the JSON API is very similar.

[1]: http://www.nextbus.com/xmlFeedDocs/NextBusXMLFeed.pdf
