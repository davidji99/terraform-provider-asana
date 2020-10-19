package api

// UsersService handles communication with the user related
// methods of the Asana API.
//
// Asana API docs: https://developers.asana.com/docs/user
type UsersService service

// User represents an account in Asana that can be given access to various workspaces, projects, and tasks.
type User struct {
	CommonResourceFields

	// Read-only except when same user as requester. The user’s name.
	Name *string `json:"name,omitempty"`

	// The user's email address.
	Email *string `json:"email,omitempty"`

	// A map of the user’s profile photo in various sizes, or null if no photo is set.
	// Sizes provided are 21, 27, 36, 60, and 128. Images are in PNG format.
	Photo *UserPhoto `json:"photo,omitempty"`

	// Workspaces and organizations this user may access.
	// Note: The API will only return workspaces and organizations that also contain the authenticated user.
	Workspaces []*Workspace `json:"workspaces,omitempty"`
}

// UserPhoto represents a user's profile photo in various sizes, or null if no photo is set.
// Sizes provided are 21, 27, 36, 60, and 128. Images are in PNG format.
type UserPhoto struct {
	Image128 *string `json:"image_128x128,omitempty"`
	Image21  *string `json:"image_21x21,omitempty"`
	Image27  *string `json:"image_27x27,omitempty"`
	Image36  *string `json:"image_36x36,omitempty"`
	Image60  *string `json:"image_60x60,omitempty"`
}
