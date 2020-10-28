package api

// InputOutputParams specify query parameters to control how your request is interpreted and how the response is generated.
// This struct is used for any GET requests or any PUT/POST requests that use the
// `application/x-www-form-urlencoded` content type.
//
// Reference: https://developers.asana.com/docs/input-output-options
type InputOutputParams struct {
	// Provides the response in "pretty" output. In the case of JSON this means doing proper line breaking
	// and indentation to make it readable. This will take extra time and increase the response size
	// so it is advisable only to use this during debugging.
	Pretty bool `url:"opt_pretty"`

	// Some requests return compact representations of objects, to conserve resources and complete
	// the request more efficiently. Other times requests return more information than you may need.
	// This option allows you to list the exact set of fields that the API should be sure to return for the objects.
	// The field names should be provided as paths, described below.
	Fields []string `url:"opt_fields"`
}

// ListParams represents the query parameters available to most (if not all) list methods.
type ListParams struct {
	// Results per page.
	Limit int `url:"opt_fields"`

	// Offset token.
	Offset string `url:"offset"`
}
