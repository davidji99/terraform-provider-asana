package api

import "time"

// ProjectsService handles communication with the project related
// methods of the Asana API.
//
// Asana API docs: https://developers.asana.com/docs/projects
type ProjectsService service

// Project represents a prioritized list of tasks in Asana or a board with columns of tasks represented as cards.
type Project struct {
	CommonResourceFields

	// Name of the project. This is generally a short sentence fragment that fits on a line in the UI
	// for maximum readability. However, it can be longer.
	Name *string `json:"name,omitempty"`

	// True if the project is archived, false if not. Archived projects do not show in the UI by default
	// and may be treated differently for queries.
	Archived *bool `json:"archived,omitempty"`

	// Color of the project.
	Color *string `json:"color,omitempty"`

	// The time at which this resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// A project status is an update on the progress of a particular project, and is sent out to all project followers
	// when created. These updates include both text describing the update and a color code intended to represent the
	// overall state of the project: "green" for projects that are on track, "yellow" for projects at risk,
	// and "red" for projects that are behind.
	CurrentStatus *ProjectStatus `json:"current_status,omitempty"`

	// Array of Custom Field Settings (in compact form).
	CustomFieldSettings []*CustomFieldSetting `json:"custom_field_settings,omitempty"`

	// The default view (list, board, calendar, or timeline) of a project.
	DefaultView *string `json:"default_view,omitempty"`

	// Deprecated: new integrations should prefer the due_on field.
	DueDate *string `json:"due_date,omitempty"`

	// The day on which this project is due. This takes a date with format YYYY-MM-DD.
	DueOn *string `json:"due_on,omitempty"`

	// Opt In. The notes of the project with formatting as HTML.
	HtmlNotes *string `json:"html_notes,omitempty"`

	// Opt In. Determines if the project is a template.
	IsTemplate *bool `json:"is_template"`

	// Array of users who are members of this project.
	Members []*User `json:"members,omitempty"`

	// The time at which this project was last modified.
	// Note: This does not currently reflect any changes in associations such as tasks or comments
	// that may have been added or removed from the project.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`

	// More detailed, free-form textual information associated with the project.
	Notes *string `json:"notes,omitempty"`

	// True if the project is public to the organization. If false, do not share this project with other users in
	// this organization without explicitly checking to see if they have access.
	Public *bool `json:"public"`

	// The day on which work for this project begins, or null if the project has no start date.
	// This takes a date with YYYY-MM-DD format. Note: due_on or due_at must be present in the request when setting
	// or unsetting the start_on parameter. Additionally, start_on and due_on cannot be the same date
	StartOn *string `json:"start_on,omitempty"`

	// Create-only. The workspace or organization this project is associated with.Once created,
	// projects cannot be moved to a different workspace. This attribute can only be specified at creation time.
	Workspace *Workspace `json:"workspace,omitempty"`

	// Array of Custom Fields.
	CustomFields []*CustomField `json:"custom_fields,omitempty"`

	// Array of users following this project. Followers are a subset of members who receive all notifications
	// for a project, the default notification setting when adding members to a project in-product.
	Followers []*User `json:"followers,omitempty"`

	// The icon for a project.
	Icon *string `json:"icon,omitempty"`

	// A user object represents an account in Asana that can be given access to various workspaces, projects, and tasks.
	Owner *User `json:"owner,omitempty"`

	// A url that points directly to the object within Asana.
	PermalinkURL *string `json:"permalink_url,omitempty"`

	// Create-only. The team that this project is shared with. This field only exists for projects in organizations.
	Team *Team `json:"team,omitempty"`
}
