package asana

import (
	"context"
	"fmt"
	"github.com/davidji99/terraform-provider-asana/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"log"
	"regexp"
	"strings"
)

func resourceAsanaProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAsanaProjectCreate,
		ReadContext:   resourceAsanaProjectRead,
		UpdateContext: resourceAsanaProjectUpdate,
		DeleteContext: resourceAsanaProjectDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"default_view": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"list", "board", "calendar", "timeline"}, false),
			},

			"team_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},

			"archived": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"color": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"custom_fields": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},

			"due_on": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateDates,
			},

			"followers": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"html_notes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"is_template": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"notes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"owner": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"is_public": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"start_on": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateDates,
				RequiredWith: []string{"due_on"},
			},

			//"current_status": {},
		},
	}
}

func validateDates(v interface{}, k string) (ws []string, errors []error) {
	date := v.(string)
	if !regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`).MatchString(date) {
		errors = append(errors, fmt.Errorf("%s attribute value "+
			"must follow this format: YYYY-MM-DD", k))
	}
	return
}

func resourceAsanaProjectImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return nil, nil
}

func constructProjectOpts(d *schema.ResourceData) *api.ProjectRequestOpts {
	opts := &api.ProjectRequestOpts{}

	if v, ok := d.GetOk("custom_fields"); ok {
		vs := v.(map[string]interface{})
		log.Printf("[DEBUG] project custom_fields is : %v", vs)
		opts.CustomFields = vs
	}

	if v, ok := d.GetOk("default_view"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] project default_view is : %v", vs)
		opts.DefaultView = vs
	}

	if v, ok := d.GetOk("html_notes"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] project html_notes is : %v", vs)
		opts.HTMLNotes = vs
	}

	if v, ok := d.GetOk("notes"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] project notes is : %v", vs)
		opts.Notes = vs
	}

	if v, ok := d.GetOk("name"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] project name is : %v", vs)
		opts.Name = vs
	}

	if v, ok := d.GetOk("workspace_id"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] project workspace_id is : %v", vs)
		opts.Workspace = vs
	}

	opts.Public = d.Get("is_public").(bool)
	log.Printf("[DEBUG] project is_public is : %v", opts.Public)

	opts.Archived = d.Get("archived").(bool)
	log.Printf("[DEBUG] project archived is : %v", opts.Archived)

	// The extra use of GetOk is needed here as project templates are a paid plan feature.
	if _, ok := d.GetOk("is_template"); ok {
		opts.IsTemplate = Bool(d.Get("is_template").(bool))
		log.Printf("[DEBUG] project is_template is : %v", opts.IsTemplate)
	}

	return opts
}

func resourceAsanaProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Config).API
	opts := constructProjectOpts(d)
	var diags diag.Diagnostics

	if v, ok := d.GetOk("color"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] new project color is : %v", vs)
		opts.Color = &vs
	}

	if v, ok := d.GetOk("due_on"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] new project due_on is : %v", vs)
		opts.DueOn = &vs
	}

	if v, ok := d.GetOk("owner"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] new project owner is : %v", vs)
		opts.Owner = &vs
	}

	if v, ok := d.GetOk("start_on"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] new project start_on is : %v", vs)
		opts.StartOn = &vs
	}

	if v, ok := d.GetOk("followers"); ok {
		vs := v.(*schema.Set).List()
		fl := make([]string, 0)
		for _, f := range vs {
			fl = append(fl, f.(string))
		}

		opts.Followers = strings.Join(fl, ",")
		log.Printf("[DEBUG] project followers is : %v", opts.Followers)
	}

	if v, ok := d.GetOk("team_id"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] project team_id is : %v", vs)
		opts.Team = vs
	}

	log.Printf("[DEBUG] Creating new project with the following options: %v", opts)

	p, _, createErr := client.Projects.Create(opts, &api.InputOutputOpts{
		Fields: []string{"html_notes", "is_template"},
	})
	if createErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create new project",
			Detail:   createErr.Error(),
		})
		return diags
	}

	log.Printf("[DEBUG] Created new project")

	d.SetId(p.GetData().GetGID())

	return resourceAsanaProjectRead(ctx, d, meta)
}

func resourceAsanaProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Config).API

	var diags diag.Diagnostics

	p, _, readErr := client.Projects.Get(d.Id())
	if readErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Unable to read/fetch info about project %s", d.Id()),
			Detail:   readErr.Error(),
		})
		return diags
	}

	d.Set("name", p.GetData().GetName())
	d.Set("workspace_id", p.GetData().GetWorkspace().GetGID())
	d.Set("archived", p.GetData().GetArchived())
	d.Set("color", p.GetData().GetColor())
	d.Set("default_view", p.GetData().GetDefaultView())
	d.Set("due_on", p.GetData().GetDueOn())
	d.Set("html_notes", p.GetData().GetHtmlNotes())
	d.Set("is_template", p.GetData().GetIsTemplate())
	d.Set("notes", p.GetData().GetNotes())
	d.Set("is_public", p.GetData().GetPublic())
	d.Set("start_on", p.GetData().GetStartOn())

	if p.GetData().GetTeam() != nil {
		d.Set("team_id", p.GetData().GetTeam().GetGID())

	} else {
		d.Set("team_id", "")
	}

	// Construct a schema friendly version of custom_fields
	cf := make(map[string]interface{}, 0)
	for _, f := range p.GetData().CustomFields {
		cf[f.GetGID()] = f.GetTextValue()
	}
	d.Set("custom_fields", cf)

	// Construct a schema friendly version of followers
	followers := make([]string, 0)
	for _, f := range p.GetData().Followers {
		followers = append(followers, f.GetGID())
	}
	d.Set("followers", followers)

	return diags
}

func resourceAsanaProjectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Config).API
	opts := constructProjectOpts(d)
	var diags diag.Diagnostics

	if d.HasChange("color") {
		_, n := d.GetChange("color")

		if n == "null" {
			opts.Color = nil
		} else {
			opts.Color = String(n.(string))
		}
		log.Printf("[DEBUG] updated project color is : %v", opts.Color)
	}

	if d.HasChange("due_on") {
		_, n := d.GetChange("due_on")

		if n == "null" {
			opts.DueOn = nil
		} else {
			opts.DueOn = String(n.(string))
		}
		log.Printf("[DEBUG] updated project due_on is : %v", opts.DueOn)
	}

	if d.HasChange("owner") {
		_, n := d.GetChange("owner")

		if n == "null" {
			opts.Owner = nil
		} else {
			opts.Owner = String(n.(string))
		}
		log.Printf("[DEBUG] updated project owner is : %v", opts.Owner)
	}

	if d.HasChange("start_on") {
		_, n := d.GetChange("start_on")

		if n == "null" {
			opts.StartOn = nil
		} else {
			opts.StartOn = String(n.(string))
		}
		log.Printf("[DEBUG] updated project start_on is : %v", opts.StartOn)
	}

	log.Printf("[DEBUG] Updating project with the following options: %v", opts)

	_, _, updateErr := client.Projects.Update(d.Id(), opts, &api.InputOutputOpts{
		Fields: []string{"html_notes", "is_template"},
	})
	if updateErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Unable to update existing project %s", d.Id()),
			Detail:   updateErr.Error(),
		})
		return diags
	}

	log.Printf("[DEBUG] Updated project")

	return resourceAsanaProjectRead(ctx, d, meta)
}

func resourceAsanaProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Config).API

	var diags diag.Diagnostics

	_, deleteErr := client.Projects.Delete(d.Id())
	if deleteErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Unable to delete project %s", d.Id()),
			Detail:   deleteErr.Error(),
		})

		return diags
	}

	d.SetId("")

	return diags
}
