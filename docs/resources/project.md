---
layout: "asana"
page_title: "Asana: asana_project"
sidebar_current: "docs-asana-resource-project"
description: |-
  Provides a resource to manage an Asana Project.
---

# asana\_project

This resource manages an Asana project. A project represents a prioritized list of tasks in Asana or a board with
columns of tasks represented as cards. It exists in a single workspace or organization and is accessible to a subset
of users in that workspace or organization, depending on its permissions.

## Example Usage

```hcl-terraform
resource "asana_project" "foobar" {
	workspace_id = "12345"
	name = "my_foobar_project"
	default_view = "board"
	is_public = false
}
```

## Argument Reference

The following arguments are supported:

* `workspace_id` - (Required) `<string>` The workspace you wish to create the project in.

* `name` - (Required) `<string>` Name of the project. This is generally a short sentence fragment that fits on a line
in the UI for maximum readability. However, it can be longer.

* `default_view` - (Required) `<string>` The default view of a project.
Acceptable values: `list`, `board`, `calendar`, `timeline`. Case-sensitive.

* `team_id` - `<string>` The team that this project is shared with. This field only exists for projects in organizations.
This attribute is only available on initial project creation. Updates will result in resource destruction/recreation.

* `archived` - `<boolean>` 	True if the project is archived, false if not.
Archived projects do not show in the UI by default and may be treated differently for queries. Default: `false`.

* `color` - `<string>` Color of the project. Set value to `"null"` to remove the color.

* `custom_fields` - `<map(string)>` Define custom field values.

* `due_on` - `<string>` The day on which this project is due. This takes a date with format YYYY-MM-DD.
Set value to `"null"` to remove the date.

* `followers` - `<list(string)>` List of users. Followers are a subset of members who receive all
notifications for a project, the default notification setting when adding members to a project in-product.
This attribute is only available on initial project creation. Updates will result in resource destruction/recreation.

* `html_notes` - `<string>` The notes of the project with formatting as HTML.

* `is_template` - `<boolean>` Determines if the project is a template. This attribute is only available for accounts
on a paid plan.

* `notes` - `<string>` More detailed, free-form textual information associated with the project.

* `owner` - `<string>` The current owner of the project. Set value to `"null"` to remove the owner.

* `is_public` - `<string>` 	True if the project is public to the organization. If false, do not share this project
with other users in this organization without explicitly checking to see if they have access.
Defaults to `false`.

* `start_on` - `<string>` The day on which work for this project begins, or `"null"` if the project has no start date.
This takes a date with YYYY-MM-DD format. Note: the `due_on` attribute must be defined in your configuration when setting
or unsetting the this attribute. Additionally, `start_on` and `due_on` cannot be the same date.

## Attributes Reference

The following attributes are exported:

N/A

## Import

An existing project can be imported using the project ID.

For example:

```shell script
$ terraform import asana_project.foobar <ID>
```