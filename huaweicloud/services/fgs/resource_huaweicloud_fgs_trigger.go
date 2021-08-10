package fgs

import (
	"log"
	"regexp"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/huaweicloud/golangsdk/openstack/fgs/v2/trigger"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils/fmtp"
)

const (
	timingTrigger  = "TIMER"
	statusActive   = "ACTIVE"
	statusDisabled = "DISABLED"
)

func ResourceFunctionGraphTriggerV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceFunctionGraphTriggerV2Create,
		Read:   resourceFunctionGraphTriggerV2Read,
		Update: resourceFunctionGraphTriggerV2Update,
		Delete: resourceFunctionGraphTriggerV2Delete,

		Importer: &schema.ResourceImporter{
			State: resourceTriggerResourceImportState,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"function_urn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					timingTrigger,
				}, false),
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					statusActive, statusDisabled,
				}, false),
				Default: statusActive,
			},
			"timer": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem:     timerSchemaResource(),
			},
		},
	}
}

func timerSchemaResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^([A-Za-z][A-Za-z0-9-_]{0,63})$"),
					"The name can contains of 1 to 32 characters and start with a letter."+
						"Only letters, digits, hyphens (-) and underscores (_) are allowed."),
			},
			"schedule_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Cron", "Rate",
				}, false),
			},
			"schedule": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"additional_information": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func buildTimingTriggerParameters(d *schema.ResourceData) (map[string]interface{}, string) {
	event := make(map[string]interface{})

	event["name"] = d.Get("timer.0.name").(string)
	event["schedule"] = d.Get("timer.0.schedule").(string)
	event["schedule_type"] = d.Get("timer.0.schedule_type").(string)
	event["user_event"] = d.Get("timer.0.additional_information").(string)

	return event, "MessageCreated"
}

func buildFunctionGraphTriggerParameters(d *schema.ResourceData, config *config.Config) trigger.CreateOpts {
	triggerType := d.Get("type").(string)

	event, eventTypeCode := buildTimingTriggerParameters(d)
	return trigger.CreateOpts{
		TriggerTypeCode: triggerType,
		TriggerStatus:   d.Get("status").(string),
		EventTypeCode:   eventTypeCode,
		EventData:       event,
	}
}

func resourceFunctionGraphTriggerV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config.Config)
	client, err := config.FgsV2Client(config.GetRegion(d))
	if err != nil {
		return fmtp.Errorf("Error creating HuaweiCloud FunctionGraph client: %s", err)
	}

	opts := buildFunctionGraphTriggerParameters(d, config)
	log.Printf("[DEBUG] The create options is: %#v", opts)
	urn := d.Get("function_urn").(string)
	resp, err := trigger.Create(client, opts, urn).Extract()
	if err != nil {
		return fmtp.Errorf("Error creating HuaweiCloud FunctionGraph trigger for function (%s): %s", urn, err)
	}
	d.SetId(resp.TriggerId)

	return resourceFunctionGraphTriggerV2Read(d, meta)
}

func setTimerParameters(d *schema.ResourceData, triggerType string, eventData map[string]interface{}) error {
	if triggerType != timingTrigger {
		return nil
	}
	var info string
	if val, ok := eventData["additional_information"]; ok {
		info = val.(string)
	}
	timer := []map[string]interface{}{
		{
			"name":                   eventData["name"].(string),
			"schedule":               eventData["schedule"].(string),
			"schedule_type":          eventData["schedule_type"].(string),
			"additional_information": info,
		},
	}
	return d.Set("timer", timer)
}

func setTimingTriggerParamters(d *schema.ResourceData, resp *trigger.Trigger) error {
	triggerType := resp.TriggerTypeCode
	if triggerType != timingTrigger {
		return nil
	}
	mErr := multierror.Append(nil,
		d.Set("type", triggerType),
		d.Set("status", resp.Status),
		setTimerParameters(d, triggerType, resp.EventData),
	)
	if mErr.ErrorOrNil() != nil {
		return mErr
	}
	return nil
}

func resourceFunctionGraphTriggerV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config.Config)
	client, err := config.FgsV2Client(config.GetRegion(d))
	if err != nil {
		return fmtp.Errorf("Error creating HuaweiCloud FunctionGraph client: %s", err)
	}

	urn := d.Get("function_urn").(string)
	pages, err := trigger.List(client, urn).AllPages()
	if err != nil {
		return fmtp.Errorf("Error retrieving FunctionGraph trigger: %s", err)
	}
	triggerList, err := trigger.ExtractList(pages)
	if len(triggerList) > 0 {
		for _, v := range triggerList {
			if v.TriggerId == d.Id() {
				mErr := multierror.Append(nil,
					d.Set("region", config.GetRegion(d)),
					setTimingTriggerParamters(d, &v),
				)
				if mErr.ErrorOrNil() != nil {
					return mErr
				}
				return nil
			}
		}
	}

	return fmtp.Errorf("Unable to find the FunctionGraph trigger (%s) form function (%s): %s", d.Id(), urn, err)
}

func resourceFunctionGraphTriggerV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config.Config)
	client, err := config.FgsV2Client(config.GetRegion(d))
	if err != nil {
		return fmtp.Errorf("Error creating HuaweiCloud FunctionGraph client: %s", err)
	}

	urn := d.Get("function_urn").(string)
	triggerType := d.Get("type").(string)

	opts := trigger.UpdateOpts{
		TriggerStatus: d.Get("status").(string),
	}
	err = trigger.Update(client, opts, urn, triggerType, d.Id()).ExtractErr()
	if err != nil {
		return fmtp.Errorf("Updating HuaweiCloud FunctionGraph trigger failed: %s", err)
	}
	return resourceFunctionGraphTriggerV2Read(d, meta)
}

func resourceFunctionGraphTriggerV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config.Config)
	client, err := config.FgsV2Client(config.GetRegion(d))
	if err != nil {
		return fmtp.Errorf("Error creating HuaweiCloud FunctionGraph client: %s", err)
	}

	urn := d.Get("function_urn").(string)
	triggerType := d.Get("type").(string)
	err = trigger.Delete(client, urn, triggerType, d.Id()).ExtractErr()
	if err != nil {
		return fmtp.Errorf("Error deleting HuaweiCloud FunctionGraph trigger (%s) from the function (%s): %s",
			d.Id(), urn, err)
	}
	d.SetId("")
	return nil
}

func resourceTriggerResourceImportState(d *schema.ResourceData,
	meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.SplitN(d.Id(), "/", 2)
	if len(parts) != 2 {
		return nil, fmtp.Errorf("Invalid format specified for import id, must be <function_urn>/<id>")
	}
	d.SetId(parts[1])
	d.Set("function_urn", parts[0])
	return []*schema.ResourceData{d}, nil
}
