package api

// CustomFieldsService handles communication with the custom field related
// methods of the Asana API.
//
// Asana API docs: https://developers.asana.com/docs/custom-fields
type CustomFieldsService service

// Custom field store the metadata that is used in order to add user-specified information to tasks in Asana.
type CustomField struct {
	CommonResourceFields

	// ISO 4217 currency code to format this custom field. This will be null if the format is not currency.
	CurrencyCode *string `json:"currency_code,omitempty"`

	// This is the string that appears next to the custom field value. This will be null if the format is not custom.
	CustomLabel *string `json:"custom_label,omitempty"`

	// Only relevant for custom fields with custom format. This depicts where to place the custom label.
	// This will be null if the format is not custom.
	CustomLabelPosition *string `json:"custom_label_position,omitempty"`

	// (Opt In) The description of the custom field.
	Description *string `json:"description,omitempty"`

	// Conditional. Determines if the custom field is enabled or not.
	Enabled *bool `json:"enabled"`

	// Conditional. Only relevant for custom fields of type enum. This array specifies the possible values which an
	// enum custom field can adopt. To modify the enum options, refer to working with enum options.
	EnumOptions []*Enum `json:"enum_options,omitempty"`

	// Conditional. Only relevant for custom fields of type enum. This object is the chosen value of an enum custom field.
	EnumValues *Enum `json:"enum_value,omitempty"`

	// 	The format of this custom field.
	Format *string `json:"format,omitempty"`

	// Conditional. This flag describes whether a follower of a task with this field should receive inbox notifications
	// from changes to this field.
	HasNotificationEnabled *bool `json:"has_notifications_enabled,omitempty"`

	// This flag describes whether this custom field is available to every container in the workspace.
	// Before project-specific custom fields, this field was always true.
	IsGlobalToWorkspace *bool `json:"is_global_to_workspace,omitempty"`

	// The name of the custom field.
	Name *string `json:"name,omitempty"`

	// Conditional. This number is the value of a number custom field.
	NumberValue *int `json:"number_value,omitempty"`

	// Only relevant for custom fields of type ‘Number’. This field dictates the number of places after the decimal
	// to round to, i.e. 0 is integer values, 1 rounds to the nearest tenth, and so on. Must be between 0 and 6, inclusive.
	// For percentage format, this may be unintuitive, as a value of 0.25 has a precision of 0, while a value of 0.251
	// has a precision of 1. This is due to 0.25 being displayed as 25%.
	// The identifier format will always have a precision of 0.
	Precision *int `json:"precision,omitempty"`

	// The type of the custom field. Must be one of the given values.
	ResourceSubtype *string `json:"resource_subtype,omitempty"`

	// 	Conditional. This string is the value of a text custom field.
	TextValue *string `json:"text_value,omitempty"`

	// Deprecated: new integrations should prefer the resource_subtype field.
	// The type of the custom field. Must be one of the given values.
	Type *string `json:"type,omitempty"`
}
