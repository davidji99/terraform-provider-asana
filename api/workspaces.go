package api

// WorkspacesService handles communication with the workspace related
// methods of the Asana API.
//
// Asana API docs: https://developers.asana.com/docs/workspace
type WorkspacesService service

// Workspace is the highest-level organizational unit in Asana. All projects and tasks have an associated workspace.
type Workspace struct {
	CommonResourceFields

	// The name of the workspace.
	Name *string `json:"name,omitempty"`

	// The email domains that are associated with this workspace.
	EmailDomains []string `json:"email_domains,omitempty"`

	// Whether the workspace is an organization.
	IsOrganization *bool `json:"is_organization"`
}
