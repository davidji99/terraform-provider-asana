// Copyright 2020
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Code generated by gen-accessors; DO NOT EDIT.
package api

import (
	"time"
)

// GetGID returns the GID field if it's non-nil, zero value otherwise.
func (c *CommonResourceFields) GetGID() string {
	if c == nil || c.GID == nil {
		return ""
	}
	return *c.GID
}

// GetResourceType returns the ResourceType field if it's non-nil, zero value otherwise.
func (c *CommonResourceFields) GetResourceType() string {
	if c == nil || c.ResourceType == nil {
		return ""
	}
	return *c.ResourceType
}

// GetCurrencyCode returns the CurrencyCode field if it's non-nil, zero value otherwise.
func (c *CustomField) GetCurrencyCode() string {
	if c == nil || c.CurrencyCode == nil {
		return ""
	}
	return *c.CurrencyCode
}

// GetCustomLabel returns the CustomLabel field if it's non-nil, zero value otherwise.
func (c *CustomField) GetCustomLabel() string {
	if c == nil || c.CustomLabel == nil {
		return ""
	}
	return *c.CustomLabel
}

// GetCustomLabelPosition returns the CustomLabelPosition field if it's non-nil, zero value otherwise.
func (c *CustomField) GetCustomLabelPosition() string {
	if c == nil || c.CustomLabelPosition == nil {
		return ""
	}
	return *c.CustomLabelPosition
}

// GetDescription returns the Description field if it's non-nil, zero value otherwise.
func (c *CustomField) GetDescription() string {
	if c == nil || c.Description == nil {
		return ""
	}
	return *c.Description
}

// GetEnabled returns the Enabled field if it's non-nil, zero value otherwise.
func (c *CustomField) GetEnabled() bool {
	if c == nil || c.Enabled == nil {
		return false
	}
	return *c.Enabled
}

// HasEnumOptions checks if CustomField has any EnumOptions.
func (c *CustomField) HasEnumOptions() bool {
	if c == nil || c.EnumOptions == nil {
		return false
	}
	if len(c.EnumOptions) == 0 {
		return false
	}
	return true
}

// GetEnumValues returns the EnumValues field.
func (c *CustomField) GetEnumValues() *Enum {
	if c == nil {
		return nil
	}
	return c.EnumValues
}

// GetFormat returns the Format field if it's non-nil, zero value otherwise.
func (c *CustomField) GetFormat() string {
	if c == nil || c.Format == nil {
		return ""
	}
	return *c.Format
}

// GetHasNotificationEnabled returns the HasNotificationEnabled field if it's non-nil, zero value otherwise.
func (c *CustomField) GetHasNotificationEnabled() bool {
	if c == nil || c.HasNotificationEnabled == nil {
		return false
	}
	return *c.HasNotificationEnabled
}

// GetIsGlobalToWorkspace returns the IsGlobalToWorkspace field if it's non-nil, zero value otherwise.
func (c *CustomField) GetIsGlobalToWorkspace() bool {
	if c == nil || c.IsGlobalToWorkspace == nil {
		return false
	}
	return *c.IsGlobalToWorkspace
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (c *CustomField) GetName() string {
	if c == nil || c.Name == nil {
		return ""
	}
	return *c.Name
}

// GetNumberValue returns the NumberValue field if it's non-nil, zero value otherwise.
func (c *CustomField) GetNumberValue() int {
	if c == nil || c.NumberValue == nil {
		return 0
	}
	return *c.NumberValue
}

// GetPrecision returns the Precision field if it's non-nil, zero value otherwise.
func (c *CustomField) GetPrecision() int {
	if c == nil || c.Precision == nil {
		return 0
	}
	return *c.Precision
}

// GetResourceSubtype returns the ResourceSubtype field if it's non-nil, zero value otherwise.
func (c *CustomField) GetResourceSubtype() string {
	if c == nil || c.ResourceSubtype == nil {
		return ""
	}
	return *c.ResourceSubtype
}

// GetTextValue returns the TextValue field if it's non-nil, zero value otherwise.
func (c *CustomField) GetTextValue() string {
	if c == nil || c.TextValue == nil {
		return ""
	}
	return *c.TextValue
}

// GetType returns the Type field if it's non-nil, zero value otherwise.
func (c *CustomField) GetType() string {
	if c == nil || c.Type == nil {
		return ""
	}
	return *c.Type
}

// GetCustomField returns the CustomField field.
func (c *CustomFieldSetting) GetCustomField() *CustomField {
	if c == nil {
		return nil
	}
	return c.CustomField
}

// GetIsImportant returns the IsImportant field if it's non-nil, zero value otherwise.
func (c *CustomFieldSetting) GetIsImportant() bool {
	if c == nil || c.IsImportant == nil {
		return false
	}
	return *c.IsImportant
}

// GetParent returns the Parent field.
func (c *CustomFieldSetting) GetParent() *Project {
	if c == nil {
		return nil
	}
	return c.Parent
}

// GetProject returns the Project field.
func (c *CustomFieldSetting) GetProject() *Project {
	if c == nil {
		return nil
	}
	return c.Project
}

// GetColor returns the Color field if it's non-nil, zero value otherwise.
func (e *Enum) GetColor() string {
	if e == nil || e.Color == nil {
		return ""
	}
	return *e.Color
}

// GetEnabled returns the Enabled field if it's non-nil, zero value otherwise.
func (e *Enum) GetEnabled() bool {
	if e == nil || e.Enabled == nil {
		return false
	}
	return *e.Enabled
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (e *Enum) GetName() string {
	if e == nil || e.Name == nil {
		return ""
	}
	return *e.Name
}

// HasFields checks if InputOutputOpts has any Fields.
func (i *InputOutputOpts) HasFields() bool {
	if i == nil || i.Fields == nil {
		return false
	}
	if len(i.Fields) == 0 {
		return false
	}
	return true
}

// HasFields checks if InputOutputParams has any Fields.
func (i *InputOutputParams) HasFields() bool {
	if i == nil || i.Fields == nil {
		return false
	}
	if len(i.Fields) == 0 {
		return false
	}
	return true
}

// GetArchived returns the Archived field if it's non-nil, zero value otherwise.
func (p *Project) GetArchived() bool {
	if p == nil || p.Archived == nil {
		return false
	}
	return *p.Archived
}

// GetColor returns the Color field if it's non-nil, zero value otherwise.
func (p *Project) GetColor() string {
	if p == nil || p.Color == nil {
		return ""
	}
	return *p.Color
}

// GetCreatedAt returns the CreatedAt field if it's non-nil, zero value otherwise.
func (p *Project) GetCreatedAt() time.Time {
	if p == nil || p.CreatedAt == nil {
		return time.Time{}
	}
	return *p.CreatedAt
}

// GetCurrentStatus returns the CurrentStatus field.
func (p *Project) GetCurrentStatus() *ProjectStatus {
	if p == nil {
		return nil
	}
	return p.CurrentStatus
}

// HasCustomFields checks if Project has any CustomFields.
func (p *Project) HasCustomFields() bool {
	if p == nil || p.CustomFields == nil {
		return false
	}
	if len(p.CustomFields) == 0 {
		return false
	}
	return true
}

// HasCustomFieldSettings checks if Project has any CustomFieldSettings.
func (p *Project) HasCustomFieldSettings() bool {
	if p == nil || p.CustomFieldSettings == nil {
		return false
	}
	if len(p.CustomFieldSettings) == 0 {
		return false
	}
	return true
}

// GetDefaultView returns the DefaultView field if it's non-nil, zero value otherwise.
func (p *Project) GetDefaultView() string {
	if p == nil || p.DefaultView == nil {
		return ""
	}
	return *p.DefaultView
}

// GetDueDate returns the DueDate field if it's non-nil, zero value otherwise.
func (p *Project) GetDueDate() string {
	if p == nil || p.DueDate == nil {
		return ""
	}
	return *p.DueDate
}

// GetDueOn returns the DueOn field if it's non-nil, zero value otherwise.
func (p *Project) GetDueOn() string {
	if p == nil || p.DueOn == nil {
		return ""
	}
	return *p.DueOn
}

// HasFollowers checks if Project has any Followers.
func (p *Project) HasFollowers() bool {
	if p == nil || p.Followers == nil {
		return false
	}
	if len(p.Followers) == 0 {
		return false
	}
	return true
}

// GetHtmlNotes returns the HtmlNotes field if it's non-nil, zero value otherwise.
func (p *Project) GetHtmlNotes() string {
	if p == nil || p.HtmlNotes == nil {
		return ""
	}
	return *p.HtmlNotes
}

// GetIcon returns the Icon field if it's non-nil, zero value otherwise.
func (p *Project) GetIcon() string {
	if p == nil || p.Icon == nil {
		return ""
	}
	return *p.Icon
}

// GetIsTemplate returns the IsTemplate field if it's non-nil, zero value otherwise.
func (p *Project) GetIsTemplate() bool {
	if p == nil || p.IsTemplate == nil {
		return false
	}
	return *p.IsTemplate
}

// HasMembers checks if Project has any Members.
func (p *Project) HasMembers() bool {
	if p == nil || p.Members == nil {
		return false
	}
	if len(p.Members) == 0 {
		return false
	}
	return true
}

// GetModifiedAt returns the ModifiedAt field if it's non-nil, zero value otherwise.
func (p *Project) GetModifiedAt() time.Time {
	if p == nil || p.ModifiedAt == nil {
		return time.Time{}
	}
	return *p.ModifiedAt
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (p *Project) GetName() string {
	if p == nil || p.Name == nil {
		return ""
	}
	return *p.Name
}

// GetNotes returns the Notes field if it's non-nil, zero value otherwise.
func (p *Project) GetNotes() string {
	if p == nil || p.Notes == nil {
		return ""
	}
	return *p.Notes
}

// GetOwner returns the Owner field.
func (p *Project) GetOwner() *User {
	if p == nil {
		return nil
	}
	return p.Owner
}

// GetPermalinkURL returns the PermalinkURL field if it's non-nil, zero value otherwise.
func (p *Project) GetPermalinkURL() string {
	if p == nil || p.PermalinkURL == nil {
		return ""
	}
	return *p.PermalinkURL
}

// GetPublic returns the Public field if it's non-nil, zero value otherwise.
func (p *Project) GetPublic() bool {
	if p == nil || p.Public == nil {
		return false
	}
	return *p.Public
}

// GetStartOn returns the StartOn field if it's non-nil, zero value otherwise.
func (p *Project) GetStartOn() string {
	if p == nil || p.StartOn == nil {
		return ""
	}
	return *p.StartOn
}

// GetTeam returns the Team field.
func (p *Project) GetTeam() *Team {
	if p == nil {
		return nil
	}
	return p.Team
}

// GetWorkspace returns the Workspace field.
func (p *Project) GetWorkspace() *Workspace {
	if p == nil {
		return nil
	}
	return p.Workspace
}

// GetColor returns the Color field if it's non-nil, zero value otherwise.
func (p *ProjectRequest) GetColor() string {
	if p == nil || p.Color == nil {
		return ""
	}
	return *p.Color
}

// GetCurrentStatus returns the CurrentStatus field.
func (p *ProjectRequest) GetCurrentStatus() *ProjectStatus {
	if p == nil {
		return nil
	}
	return p.CurrentStatus
}

// GetDueOn returns the DueOn field if it's non-nil, zero value otherwise.
func (p *ProjectRequest) GetDueOn() string {
	if p == nil || p.DueOn == nil {
		return ""
	}
	return *p.DueOn
}

// GetIsTemplate returns the IsTemplate field if it's non-nil, zero value otherwise.
func (p *ProjectRequest) GetIsTemplate() bool {
	if p == nil || p.IsTemplate == nil {
		return false
	}
	return *p.IsTemplate
}

// GetOwner returns the Owner field if it's non-nil, zero value otherwise.
func (p *ProjectRequest) GetOwner() string {
	if p == nil || p.Owner == nil {
		return ""
	}
	return *p.Owner
}

// GetStartOn returns the StartOn field if it's non-nil, zero value otherwise.
func (p *ProjectRequest) GetStartOn() string {
	if p == nil || p.StartOn == nil {
		return ""
	}
	return *p.StartOn
}

// GetData returns the Data field.
func (p *ProjectResponse) GetData() *Project {
	if p == nil {
		return nil
	}
	return p.Data
}

// HasData checks if ProjectsResponse has any Data.
func (p *ProjectsResponse) HasData() bool {
	if p == nil || p.Data == nil {
		return false
	}
	if len(p.Data) == 0 {
		return false
	}
	return true
}

// GetAuthor returns the Author field.
func (p *ProjectStatus) GetAuthor() *User {
	if p == nil {
		return nil
	}
	return p.Author
}

// GetColor returns the Color field if it's non-nil, zero value otherwise.
func (p *ProjectStatus) GetColor() string {
	if p == nil || p.Color == nil {
		return ""
	}
	return *p.Color
}

// GetCreatedAt returns the CreatedAt field if it's non-nil, zero value otherwise.
func (p *ProjectStatus) GetCreatedAt() time.Time {
	if p == nil || p.CreatedAt == nil {
		return time.Time{}
	}
	return *p.CreatedAt
}

// GetCreatedBy returns the CreatedBy field.
func (p *ProjectStatus) GetCreatedBy() *User {
	if p == nil {
		return nil
	}
	return p.CreatedBy
}

// GetHtmlText returns the HtmlText field if it's non-nil, zero value otherwise.
func (p *ProjectStatus) GetHtmlText() string {
	if p == nil || p.HtmlText == nil {
		return ""
	}
	return *p.HtmlText
}

// GetModifiedAt returns the ModifiedAt field if it's non-nil, zero value otherwise.
func (p *ProjectStatus) GetModifiedAt() string {
	if p == nil || p.ModifiedAt == nil {
		return ""
	}
	return *p.ModifiedAt
}

// GetText returns the Text field if it's non-nil, zero value otherwise.
func (p *ProjectStatus) GetText() string {
	if p == nil || p.Text == nil {
		return ""
	}
	return *p.Text
}

// GetTitle returns the Title field if it's non-nil, zero value otherwise.
func (p *ProjectStatus) GetTitle() string {
	if p == nil || p.Title == nil {
		return ""
	}
	return *p.Title
}

// GetDescription returns the Description field if it's non-nil, zero value otherwise.
func (t *Team) GetDescription() string {
	if t == nil || t.Description == nil {
		return ""
	}
	return *t.Description
}

// GetHtmlDescription returns the HtmlDescription field if it's non-nil, zero value otherwise.
func (t *Team) GetHtmlDescription() string {
	if t == nil || t.HtmlDescription == nil {
		return ""
	}
	return *t.HtmlDescription
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (t *Team) GetName() string {
	if t == nil || t.Name == nil {
		return ""
	}
	return *t.Name
}

// GetOrganization returns the Organization field.
func (t *Team) GetOrganization() *Workspace {
	if t == nil {
		return nil
	}
	return t.Organization
}

// GetPermalinkURL returns the PermalinkURL field if it's non-nil, zero value otherwise.
func (t *Team) GetPermalinkURL() string {
	if t == nil || t.PermalinkURL == nil {
		return ""
	}
	return *t.PermalinkURL
}

// GetEmail returns the Email field if it's non-nil, zero value otherwise.
func (u *User) GetEmail() string {
	if u == nil || u.Email == nil {
		return ""
	}
	return *u.Email
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (u *User) GetName() string {
	if u == nil || u.Name == nil {
		return ""
	}
	return *u.Name
}

// GetPhoto returns the Photo field.
func (u *User) GetPhoto() *UserPhoto {
	if u == nil {
		return nil
	}
	return u.Photo
}

// HasWorkspaces checks if User has any Workspaces.
func (u *User) HasWorkspaces() bool {
	if u == nil || u.Workspaces == nil {
		return false
	}
	if len(u.Workspaces) == 0 {
		return false
	}
	return true
}

// GetImage128 returns the Image128 field if it's non-nil, zero value otherwise.
func (u *UserPhoto) GetImage128() string {
	if u == nil || u.Image128 == nil {
		return ""
	}
	return *u.Image128
}

// GetImage21 returns the Image21 field if it's non-nil, zero value otherwise.
func (u *UserPhoto) GetImage21() string {
	if u == nil || u.Image21 == nil {
		return ""
	}
	return *u.Image21
}

// GetImage27 returns the Image27 field if it's non-nil, zero value otherwise.
func (u *UserPhoto) GetImage27() string {
	if u == nil || u.Image27 == nil {
		return ""
	}
	return *u.Image27
}

// GetImage36 returns the Image36 field if it's non-nil, zero value otherwise.
func (u *UserPhoto) GetImage36() string {
	if u == nil || u.Image36 == nil {
		return ""
	}
	return *u.Image36
}

// GetImage60 returns the Image60 field if it's non-nil, zero value otherwise.
func (u *UserPhoto) GetImage60() string {
	if u == nil || u.Image60 == nil {
		return ""
	}
	return *u.Image60
}

// HasEmailDomains checks if Workspace has any EmailDomains.
func (w *Workspace) HasEmailDomains() bool {
	if w == nil || w.EmailDomains == nil {
		return false
	}
	if len(w.EmailDomains) == 0 {
		return false
	}
	return true
}

// GetIsOrganization returns the IsOrganization field if it's non-nil, zero value otherwise.
func (w *Workspace) GetIsOrganization() bool {
	if w == nil || w.IsOrganization == nil {
		return false
	}
	return *w.IsOrganization
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (w *Workspace) GetName() string {
	if w == nil || w.Name == nil {
		return ""
	}
	return *w.Name
}
