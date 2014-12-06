package nextbus

type args []struct{ key, value string }

func (a args) url(cmd string) string {
	url := "http://webservices.nextbus.com/service/publicJSONFeed?command=" + cmd
	if len(a) > 0 {
		url += "&" + a.String()
	}
	return url
}

func (a args) String() string {
	out := ""
	for i, arg := range a {
		if i > 0 {
			out += "&"
		}
		out += arg.key
		if len(arg.value) > 0 {
			out += "=" + arg.value
		}
	}
	return out
}

func (a *args) add(key, value string) {
	*a = append(*a, struct{ key, value string }{key, value})
}
