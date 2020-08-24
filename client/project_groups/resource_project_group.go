package project_groups

import (
	"models"
	"project_groups"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceProjectGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectGroupCreate,
		//Read:   resourceProjectGroupRead,
		//Update: resourceProjectGroupUpdate,
		Delete: resourceProjectGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func buildProjectGroupResource(d *schema.ResourceData) *project_groups.ClientService {
	newCreateProjectGroupParams := project_groups.NewCreateProjectGroupParams()
	createProjectGroupCreated, err := (*project_groups.Client).CreateProjectGroup(newCreateProjectGroupParams, authInfo)

}

func resourceProjectGroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*project_groups.Client)
	name := d.Get("name").(string)

	newProjectGroup := buildProjectGroupResource(d)

	newCreateProjectGroupParams := project_groups.NewCreateProjectGroupParams()
	newCreateProjectGroupParams.WithProjectGroupResource(&models.ProjectGroupResource{Name: name})
	createProjectGroupCreated, err := client.ProjectGroups.CreateProjectGroup(newCreateProjectGroupParams, authInfo)

	d.SetId(createProjectGroupCreated)
	return createProjectGroupCreated.Payload
}
