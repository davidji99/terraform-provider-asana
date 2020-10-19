package api

// Enum represents the common fields available for EnumOption and EnumValue.
type Enum struct {
	CommonResourceFields

	// The color of the enum option. Defaults to ‘none’.
	Color *string `json:"color,omitempty"`

	// Whether or not the enum option is a selectable value for the custom field.
	Enabled *bool `json:"enabled"`

	// The name of the enum option/value.
	Name *string `json:"name,omitempty"`
}
