package api

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"time"
)

// ProjectsService handles communication with the project related
// methods of the Asana API.
//
// Asana API docs: https://developers.asana.com/docs/projects
type ProjectsService service

type ProjectResponse struct {
	Data *Project `json:"data,omitempty"`
}

type ProjectsResponse struct {
	Data []*Project `json:"data,omitempty"`
}

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

// ProjectListParams represents the query parameters available when retrieving all projects.
type ProjectListParams struct {
	// The workspace or organization to filter projects on.
	Workspace string `url:"workspace"`

	// The team to filter projects on.
	Team string `url:"team"`

	// Only return projects whose archived field takes on the value of this parameter.
	Archived bool `url:"archived"`
}

// ProjectRequestOpts represents the options available when creating or updating a Project.
type ProjectRequestOpts struct {
	Archived      bool                   `json:"archived"`
	Color         *string                `json:"color"`
	CurrentStatus *ProjectStatus         `json:"current_status,omitempty"`
	CustomFields  map[string]interface{} `json:"custom_fields,omitempty"`
	DefaultView   string                 `json:"default_view,omitempty"`
	DueOn         *string                `json:"due_on"`
	Followers     string                 `json:"followers,omitempty"`
	HTMLNotes     string                 `json:"html_notes,omitempty"`
	IsTemplate    *bool                  `json:"is_template,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Notes         string                 `json:"notes,omitempty"`
	Owner         *string                `json:"owner"`
	Public        bool                   `json:"public"`
	StartOn       *string                `json:"start_on"`
	Team          string                 `json:"team,omitempty"`
	Workspace     string                 `json:"workspace,omitempty"`
}

// List returns the compact project records for some filtered set of projects. Use one or more of the parameters
// provided to filter the projects returned.
//
// Asana API docs: https://developers.asana.com/docs/get-multiple-projects
func (p *ProjectsService) List(params ...interface{}) (*ProjectsResponse, *simpleresty.Response, error) {
	result := new(ProjectsResponse)
	urlStr, urlStrErr := p.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/projects"), params...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := p.client.http.Get(urlStr, result, nil)

	return result, response, err
}

// Create a new project in a workspace or team.
//
// Every project is required to be created in a specific workspace or organization, and this cannot be changed once set.
// Note that you can use the workspace parameter regardless of whether or not it is an organization.
// If the workspace for your project is an organization, you must also supply a team to share the project with.
// Returns the full record of the newly created project.
//
// Asana API docs: https://developers.asana.com/docs/create-a-project
func (p *ProjectsService) Create(createOpts *ProjectRequestOpts, ioOpts *InputOutputOpts) (*ProjectResponse, *simpleresty.Response, error) {
	result := new(ProjectResponse)
	urlStr := p.client.http.RequestURL("/projects")

	body := struct {
		Data    *ProjectRequestOpts `json:"data"`
		Options *InputOutputOpts    `json:"options,omitempty"`
	}{Data: createOpts, Options: ioOpts}

	response, err := p.client.http.Post(urlStr, result, body)
	return result, response, err
}

// Get returns the complete project record for a single project.
//
// Asana API docs: https://developers.asana.com/docs/get-a-project
func (p *ProjectsService) Get(id string, params ...interface{}) (*ProjectResponse, *simpleresty.Response, error) {
	result := new(ProjectResponse)
	urlStr, urlStrErr := p.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/projects/%s", id), params...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := p.client.http.Get(urlStr, result, nil)

	return result, response, err
}

// Update a project.
//
// A specific, existing project can be updated by making a PUT request on the URL for that project.
// Only the fields provided in the data block will be updated; any unspecified fields will remain unchanged.
// When using this method, it is best to specify only those fields you wish to change, or else you may overwrite changes made by another user since you last retrieved the task.
//
// Returns the complete updated project record.
//
// Asana API docs: https://developers.asana.com/docs/update-a-project
func (p *ProjectsService) Update(id string, updateOpts *ProjectRequestOpts, ioOpts *InputOutputOpts) (*ProjectResponse, *simpleresty.Response, error) {
	result := new(ProjectResponse)
	urlStr := p.client.http.RequestURL("/projects/%s", id)

	body := struct {
		Data    *ProjectRequestOpts `json:"data"`
		Options *InputOutputOpts    `json:"options,omitempty"`
	}{Data: updateOpts, Options: ioOpts}

	response, err := p.client.http.Put(urlStr, result, body)
	return result, response, err
}

// Delete a project.
//
// Asana API docs: https://developers.asana.com/docs/delete-a-project
func (p *ProjectsService) Delete(id string) (*simpleresty.Response, error) {
	urlStr := p.client.http.RequestURL("/projects/%s", id)

	response, err := p.client.http.Delete(urlStr, nil, nil)
	return response, err
}
