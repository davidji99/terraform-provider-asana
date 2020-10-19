package api

import "time"

// ProjectStatusesService handles communication with the project status related
// methods of the Asana API.
//
// Asana API docs: https://developers.asana.com/docs/project-status
type ProjectStatusesService service

// ProjectStatus is an update on the progress of a particular project, and is sent out to all project followers
// when created. These updates include both text describing the update and a color code intended to represent the
// overall state of the project: "green" for projects that are on track, "yellow" for projects at risk, and "red"
// for projects that are behind.
type ProjectStatus struct {
	CommonResourceFields

	// The title of the project status update.
	Title *string `json:"title,omitempty"`

	// A user object represents an account in Asana that can be given access to various workspaces, projects, and tasks.
	Author *User `json:"author,omitempty"`

	// The color associated with the status update.
	Color *string `json:"color,omitempty"`

	// The text content of the status update with formatting as HTML.
	HtmlText *string `json:"html_text,omitempty"`

	// The time at which this project status was last modified.
	// Note: This does not currently reflect any changes in associations such as comments that may have been added
	// or removed from the project status.
	ModifiedAt *string

	// The text content of the status update.
	Text *string `json:"text,omitempty"`

	// The time at which this resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// A user object represents an account in Asana that can be given access to various workspaces, projects, and tasks.
	CreatedBy *User `json:"created_by,omitempty"`
}
