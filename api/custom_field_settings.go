package api

// CustomFieldSetting represent the many-to-many join of the Custom Field and Project
// as well as stores information that is relevant to that particular pairing.
type CustomFieldSetting struct {
	CommonResourceFields

	// The custom field that is applied to the parent.
	CustomField *CustomField `json:"custom_field,omitempty"`

	// is_important is used in the Asana web application to determine if this custom field
	// is displayed in the list/grid view of a project or portfolio.
	IsImportant *bool `json:"is_important,omitempty"`

	// The parent to which the custom field is applied. This can be a project or portfolio and indicates
	// that the tasks or projects that the parent contains may be given custom field values for this custom field.
	Parent *Project `json:"parent,omitempty"`

	// Deprecated: new integrations should prefer the parent field.
	// The id of the project that this custom field settings refers to.
	Project *Project `json:"project,omitempty"`
}
