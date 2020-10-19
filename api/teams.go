package api

// TeamsService handles communication with the team related
// methods of the Asana API.
//
// Asana API docs: https://developers.asana.com/docs/teams
type TeamsService service

// Team is used to group related projects and people together within an organization.
// Each project in an organization is associated with a team.
type Team struct {
	CommonResourceFields

	// The name of the team.
	Name *string `json:"name,omitempty"`

	// (Opt-in) The description of the team.
	Description *string `json:"description,omitempty"`

	// (Opt-in) The description of the team with formatting as HTML.
	HtmlDescription *string `json:"html_description,omitempty"`

	// The organization/workspace the team belongs to.
	Organization *Workspace `json:"organization,omitempty"`

	// A url that points directly to the object within Asana.
	PermalinkURL *string `json:"permalink_url,omitempty"`
}
