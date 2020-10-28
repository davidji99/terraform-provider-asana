package api

// InputOutputOpts represents the options available for POST or PUT request
// or when the `application/json` content type.
type InputOutputOpts struct {
	// Provides “pretty” output.
	Pretty bool `json:"opt_pretty"`

	// Defines fields to return.
	Fields []string `json:"opt_fields"`
}
