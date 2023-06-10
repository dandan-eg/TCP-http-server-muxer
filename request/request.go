package request

import "strings"

type Request struct {
	Method string
	URI    string
	Param  map[string]string
}

/*
New creates a new instance of the Request structure from a request line.
If the request line contains parameters, they will be parsed and stored in the request's parameter map.
The function returns an initialized instance of the Request structure with the values extracted from the request line.
Example : New("GET /home?name=dandan HTTP/1.1") will return an instance of Request {Method: "GET", URI: "/home", Param: {"name" : "dandan"}}
*/
func New(req string) Request {
	fs := strings.Fields(req)
	method := fs[0]
	uri, query := getURIAndQuery(fs[1])

	param := extractParam(query)

	return Request{
		Method: method,
		URI:    uri,
		Param:  param,
	}
}

func getURIAndQuery(uri string) (string, string) {
	parts := strings.SplitN(uri, "?", 2)
	if len(parts) < 2 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

func extractParam(query string) map[string]string {
	if query == "" {
		return nil
	}

	param := make(map[string]string)

	options := strings.Split(query, "&")

	for _, otp := range options {
		keyValue := strings.SplitN(otp, "=", 2)
		if len(keyValue) == 2 {
			k, v := keyValue[0], keyValue[1]
			param[k] = v
		}
	}

	if len(param) == 0 {
		return nil
	}

	return param
}
