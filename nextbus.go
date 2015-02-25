package nextbus

import (
	"encoding/json"
	"net/http"
)

// Stop combines a route and stop tag for input to GetPredictionsMulti.
type Stop struct {
	Route string
	Stop  string
}

func (s Stop) format() string {
	return s.Route + "|" + s.Stop
}

// fetch GETs url and unmarshals the response body as JSON into v.
func fetch(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

// GetAgencies returns a list of all agencies known to NextBus.
func GetAgencies() ([]Agency, error) {
	var wrapper struct{ Agency []Agency }
	err := fetch(args{}.url("agencyList"), &wrapper)
	return wrapper.Agency, err
}

// GetRoutes returns a list of all routes for a given agency (see GetAgencies).
func GetRoutes(agency string) ([]Route, error) {
	var wrapper struct{ Route []Route }
	err := fetch(args{{"a", agency}}.url("routeList"), &wrapper)
	return wrapper.Route, err
}

// GetRouteConfig returns detailed information for the specified route (see GetRoutes).
func GetRouteConfig(agency string, route string, verbose, terse bool) (RouteConfig, error) {
	a := args{{"a", agency}, {"r", route}}
	if verbose {
		a.add("verbose", "")
	}
	if terse {
		a.add("terse", "")
	}
	var wrapper struct{ Route RouteConfig }
	err := fetch(a.url("routeConfig"), &wrapper)
	return wrapper.Route, err
}

// GetPredictionsStopId returns predicted arrival times for a stop given by stopId.
func GetPredictionsStopId(agency string, stopId string) (Predictions, error) {
	return getPredictions(args{{"a", agency}, {"stopId", stopId}})
}

// GetPredictions returns predicted arrival times for a stop given by route and stop tag.
func GetPredictions(agency string, route string, stop string) (Predictions, error) {
	return getPredictions(args{{"a", agency}, {"r", route}, {"s", stop}})
}

func getPredictions(args args) (Predictions, error) {
	var wrapper struct{ Predictions Predictions }
	err := fetch(args.url("predictions"), &wrapper)
	return wrapper.Predictions, err
}

// GetPredictionsMulti returns predicted arrival times for multiple stops.
func GetPredictionsMulti(agency string, stops []Stop) ([]Predictions, error) {
	a := args{{"a", agency}}
	for _, stop := range stops {
		a.add("stops", stop.format())
	}
	var wrapper struct{ Predictions []Predictions }
	err := fetch(a.url("predictionsForMultiStops"), &wrapper)
	return wrapper.Predictions, err
}

// GetSchedules returns schedule information for the given route (see GetRoutes).
func GetSchedules(agency string, route string) ([]Schedule, error) {
	var wrapper struct{ Route []Schedule }
	err := fetch(args{{"a", agency}, {"r", route}}.url("schedule"), &wrapper)
	return wrapper.Route, err
}
