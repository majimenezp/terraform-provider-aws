package aws

import (
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAwsMediaConvertJobTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsMediaConvertJobTemplateCreate,
		Read:   resourceAwsMediaConvertJobTemplateRead,
		Update: resourceAwsMediaConvertJobTemplateUpdate,
		Delete: resourceAwsMediaConvertJobTemplateDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"queue": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_update_interval": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"acceleration_settings": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								mediaconvert.AccelerationModeDisabled,
								mediaconvert.AccelerationModeEnabled,
								mediaconvert.AccelerationModePreferred,
							}, false),
						},
					},
				},
			},
			"hop_destinations": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"queue": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"wait_minutes ": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntAtLeast(1),
						},
					},
				},
			},
			"settings": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ad_avail_offset": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"avail_blanking": {
							Type:     schema.TypeList,
							MinItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"avail_blanking_image": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"tags": tagsSchema(),
		},
	}
}

func resourceAwsMediaConvertJobTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceAwsMediaConvertJobTemplateRead(d, meta)
}

func resourceAwsMediaConvertJobTemplateRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsMediaConvertJobTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceAwsMediaConvertJobTemplateRead(d, meta)
}

func resourceAwsMediaConvertJobTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
