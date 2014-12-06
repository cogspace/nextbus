package nextbus

type Agency struct {
	Title       string
	RegionTitle string
	Tag         string
}

type Route struct {
	Title      string
	ShortTitle string
	Tag        string
}

type RouteConfig struct {
	Title         string
	ShortTitle    string
	Tag           string
	LatMin        string
	LonMin        string
	LatMax        string
	LonMax        string
	Color         string
	OppositeColor string
	Stop          []struct {
		Tag        string
		Title      string
		ShortTitle string
		Lat        string
		Lon        string
		StopId     string
	}
	Direction []struct {
		Stop []struct {
			Title  string
			StopId string
			Tag    string
			Lat    string
			Lon    string
		}
		Title    string
		UseForUI string
		Tag      string
		Name     string
	}
	Path []struct {
		Point []struct {
			Lat string
			Lon string
		}
	}
}

type Predictions struct {
	AgencyTitle                  string
	RouteTag                     string
	RouteCode                    string
	RouteTitle                   string
	StopTitle                    string
	DirTitleBecauseNoPredictions string
	Direction                    []struct {
		Title      string
		Prediction []struct {
			Seconds           string
			Minutes           string
			EpochTime         string
			IsDeparture       bool
			DirTag            string
			Block             string
			TripTag           string
			Branch            string
			AffectedByLayover bool
			IsScheduleBased   bool
			Delayed           bool
		}
	}
	Message []struct {
		Text string
	}
}

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
