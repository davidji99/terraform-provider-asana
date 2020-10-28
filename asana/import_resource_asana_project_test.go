package asana

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"
)

func TestAccAsanaProject_importBasic(t *testing.T) {
	var providers []*schema.Provider
	workspaceID := testAccConfig.GetWorkspaceIDorSkip(t)
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))
	defaultView := "board"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories(&providers),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAsanaProject_basic(workspaceID, name, defaultView),
			},
			{
				ResourceName:      "asana_project.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
