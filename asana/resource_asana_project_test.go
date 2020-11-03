package asana

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccAsanaProject_Basic(t *testing.T) {
	workspaceID := testAccConfig.GetWorkspaceIDorSkip(t)
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))
	defaultView := "board"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAsanaProject_basic(workspaceID, name, defaultView),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "workspace_id", workspaceID),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "name", name),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "default_view", defaultView),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "team_id", ""),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "archived", "false"),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "color", ""),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "custom_fields.%", "0"),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "due_on", ""),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "followers.#", "1"),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "html_notes", ""),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "is_template", "false"),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "notes", ""),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "is_public", "false"),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "start_on", ""),
				),
			},
			{
				Config: testAccCheckAsanaProject_Updatedbasic(workspaceID, fmt.Sprintf("%s_updated", name), defaultView),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "name", fmt.Sprintf("%s_updated", name)),
					resource.TestCheckResourceAttr(
						"asana_project.foobar", "is_public", "true"),
				),
			},
		},
	})
}

func testAccCheckAsanaProject_basic(workspaceID, name, defaultView string) string {
	return fmt.Sprintf(`
resource "asana_project" "foobar" {
	workspace_id = "%s"
	name = "%s"
	default_view = "%s"
	is_public = false
}
`, workspaceID, name, defaultView)
}

func testAccCheckAsanaProject_Updatedbasic(workspaceID, name, defaultView string) string {
	return fmt.Sprintf(`
resource "asana_project" "foobar" {
	workspace_id = "%s"
	name = "%s"
	default_view = "%s"
	is_public = true
}
`, workspaceID, name, defaultView)
}
