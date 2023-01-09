package dms

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/chnsz/golangsdk"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
	"github.com/jmespath/go-jmespath"
)

func ResourceDmsRocketMQConsumerGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDmsRocketMQConsumerGroupCreate,
		UpdateContext: resourceDmsRocketMQConsumerGroupUpdate,
		ReadContext:   resourceDmsRocketMQConsumerGroupRead,
		DeleteContext: resourceDmsRocketMQConsumerGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the ID of the rocketMQ instance.`,
			},
			"brokers": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the list of associated brokers of the consumer group.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the name of the consumer group.`,
				ValidateFunc: validation.All(
					validation.StringMatch(regexp.MustCompile(`^[A-Za-z|%-_0-9]*$`),
						"An instance name starts with a letter and can contain only letters, digits,"+
							"vertical lines(|), percent sign(%), underscores (_), and hyphens (-)"),
					validation.StringLenBetween(3, 64),
				),
			},
			"retry_max_times": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  `Specifies the maximum number of retry times.`,
				ValidateFunc: validation.IntBetween(1, 16),
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the consumer group is enabled or not. Default to true.`,
			},
			"broadcast": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: `Specifies whether to broadcast of the consumer group.`,
			},
		},
	}
}

func resourceDmsRocketMQConsumerGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	// createRocketmqConsumerGroup: create DMS rocketmq consumer group
	var (
		createRocketmqConsumerGroupHttpUrl = "v2/{project_id}/instances/{instance_id}/groups"
		createRocketmqConsumerGroupProduct = "dms"
	)
	createRocketmqConsumerGroupClient, err := config.NewServiceClient(createRocketmqConsumerGroupProduct, region)
	if err != nil {
		return diag.Errorf("error creating DmsRocketMQConsumerGroup Client: %s", err)
	}

	instanceID := d.Get("instance_id").(string)
	createRocketmqConsumerGroupPath := createRocketmqConsumerGroupClient.Endpoint + createRocketmqConsumerGroupHttpUrl
	createRocketmqConsumerGroupPath = strings.ReplaceAll(createRocketmqConsumerGroupPath, "{project_id}", createRocketmqConsumerGroupClient.ProjectID)
	createRocketmqConsumerGroupPath = strings.ReplaceAll(createRocketmqConsumerGroupPath, "{instance_id}", instanceID)

	createRocketmqConsumerGroupOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
	}
	createRocketmqConsumerGroupOpt.JSONBody = utils.RemoveNil(buildCreateRocketmqConsumerGroupBodyParams(d, config))
	createRocketmqConsumerGroupResp, err := createRocketmqConsumerGroupClient.Request("POST", createRocketmqConsumerGroupPath, &createRocketmqConsumerGroupOpt)
	if err != nil {
		return diag.Errorf("error creating DmsRocketMQConsumerGroup: %s", err)
	}

	createRocketmqConsumerGroupRespBody, err := utils.FlattenResponse(createRocketmqConsumerGroupResp)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := jmespath.Search("name", createRocketmqConsumerGroupRespBody)
	if err != nil {
		return diag.Errorf("error creating DmsRocketMQConsumerGroup: ID is not found in API response")
	}

	d.SetId(instanceID + "/" + id.(string))

	return resourceDmsRocketMQConsumerGroupRead(ctx, d, meta)
}

func buildCreateRocketmqConsumerGroupBodyParams(d *schema.ResourceData, config *config.Config) map[string]interface{} {
	var enabled interface{} = true
	if v, ok := d.GetOk("enabled"); ok {
		enabled = v
	}
	bodyParams := map[string]interface{}{
		"enabled":        enabled,
		"broadcast":      utils.ValueIngoreEmpty(d.Get("broadcast")),
		"brokers":        utils.ValueIngoreEmpty(d.Get("brokers")),
		"name":           utils.ValueIngoreEmpty(d.Get("name")),
		"retry_max_time": utils.ValueIngoreEmpty(d.Get("retry_max_times")),
	}
	return bodyParams
}

func resourceDmsRocketMQConsumerGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	updateRocketmqConsumerGrouphasChanges := []string{
		"enabled",
		"broadcast",
		"retry_max_times",
	}

	if d.HasChanges(updateRocketmqConsumerGrouphasChanges...) {
		// updateRocketmqConsumerGroup: update DMS rocketmq consumer group
		var (
			updateRocketmqConsumerGroupHttpUrl = "v2/{project_id}/instances/{instance_id}/groups/{group}"
			updateRocketmqConsumerGroupProduct = "dms"
		)
		updateRocketmqConsumerGroupClient, err := config.NewServiceClient(updateRocketmqConsumerGroupProduct, region)
		if err != nil {
			return diag.Errorf("error creating DmsRocketMQConsumerGroup Client: %s", err)
		}

		parts := strings.SplitN(d.Id(), "/", 2)
		if len(parts) != 2 {
			return diag.Errorf("invalid id format, must be <instance_id>/<consumerGroup>")
		}
		instanceID := parts[0]
		name := parts[1]
		updateRocketmqConsumerGroupPath := updateRocketmqConsumerGroupClient.Endpoint + updateRocketmqConsumerGroupHttpUrl
		updateRocketmqConsumerGroupPath = strings.ReplaceAll(updateRocketmqConsumerGroupPath, "{project_id}", updateRocketmqConsumerGroupClient.ProjectID)
		updateRocketmqConsumerGroupPath = strings.ReplaceAll(updateRocketmqConsumerGroupPath, "{instance_id}", instanceID)
		updateRocketmqConsumerGroupPath = strings.ReplaceAll(updateRocketmqConsumerGroupPath, "{group}", name)

		updateRocketmqConsumerGroupOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
			OkCodes: []int{
				204,
			},
		}
		updateRocketmqConsumerGroupOpt.JSONBody = utils.RemoveNil(buildUpdateRocketmqConsumerGroupBodyParams(d, config))
		_, err = updateRocketmqConsumerGroupClient.Request("PUT", updateRocketmqConsumerGroupPath, &updateRocketmqConsumerGroupOpt)
		if err != nil {
			return diag.Errorf("error updating DmsRocketMQConsumerGroup: %s", err)
		}
	}

	return resourceDmsRocketMQConsumerGroupRead(ctx, d, meta)
}

func buildUpdateRocketmqConsumerGroupBodyParams(d *schema.ResourceData, config *config.Config) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"broadcast":      utils.ValueIngoreEmpty(d.Get("broadcast")),
		"retry_max_time": utils.ValueIngoreEmpty(d.Get("retry_max_times")),
	}
	enabled := utils.ValueIngoreEmpty(d.Get("enabled"))
	if enabled != nil {
		bodyParams["enabled"] = enabled
	}
	return bodyParams
}

func resourceDmsRocketMQConsumerGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	var mErr *multierror.Error

	// getRocketmqConsumerGroup: query DMS rocketmq consumer group
	var (
		getRocketmqConsumerGroupHttpUrl = "v2/{project_id}/instances/{instance_id}/groups/{group}"
		getRocketmqConsumerGroupProduct = "dms"
	)
	getRocketmqConsumerGroupClient, err := config.NewServiceClient(getRocketmqConsumerGroupProduct, region)
	if err != nil {
		return diag.Errorf("error creating DmsRocketMQConsumerGroup Client: %s", err)
	}

	parts := strings.SplitN(d.Id(), "/", 2)
	if len(parts) != 2 {
		return diag.Errorf("invalid id format, must be <instance_id>/<consumerGroup>")
	}
	instanceID := parts[0]
	name := parts[1]
	getRocketmqConsumerGroupPath := getRocketmqConsumerGroupClient.Endpoint + getRocketmqConsumerGroupHttpUrl
	getRocketmqConsumerGroupPath = strings.ReplaceAll(getRocketmqConsumerGroupPath, "{project_id}", getRocketmqConsumerGroupClient.ProjectID)
	getRocketmqConsumerGroupPath = strings.ReplaceAll(getRocketmqConsumerGroupPath, "{instance_id}", instanceID)
	getRocketmqConsumerGroupPath = strings.ReplaceAll(getRocketmqConsumerGroupPath, "{group}", name)

	getRocketmqConsumerGroupOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
	}
	getRocketmqConsumerGroupResp, err := getRocketmqConsumerGroupClient.Request("GET", getRocketmqConsumerGroupPath, &getRocketmqConsumerGroupOpt)

	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving DmsRocketMQConsumerGroup")
	}

	getRocketmqConsumerGroupRespBody, err := utils.FlattenResponse(getRocketmqConsumerGroupResp)
	if err != nil {
		return diag.FromErr(err)
	}

	mErr = multierror.Append(
		mErr,
		d.Set("region", region),
		d.Set("instance_id", instanceID),
		d.Set("enabled", utils.PathSearch("enabled", getRocketmqConsumerGroupRespBody, nil)),
		d.Set("broadcast", utils.PathSearch("broadcast", getRocketmqConsumerGroupRespBody, nil)),
		d.Set("brokers", utils.PathSearch("brokers", getRocketmqConsumerGroupRespBody, nil)),
		d.Set("name", name),
		d.Set("retry_max_times", utils.PathSearch("retry_max_time", getRocketmqConsumerGroupRespBody, nil)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func resourceDmsRocketMQConsumerGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	// deleteRocketmqConsumerGroup: delete DMS rocketmq consumer group
	var (
		deleteRocketmqConsumerGroupHttpUrl = "v2/{project_id}/instances/{instance_id}/groups/{group}"
		deleteRocketmqConsumerGroupProduct = "dms"
	)
	deleteRocketmqConsumerGroupClient, err := config.NewServiceClient(deleteRocketmqConsumerGroupProduct, region)
	if err != nil {
		return diag.Errorf("error creating DmsRocketMQConsumerGroup Client: %s", err)
	}

	parts := strings.SplitN(d.Id(), "/", 2)
	if len(parts) != 2 {
		return diag.Errorf("invalid id format, must be <instance_id>/<consumerGroup>")
	}
	instanceID := parts[0]
	name := parts[1]
	deleteRocketmqConsumerGroupPath := deleteRocketmqConsumerGroupClient.Endpoint + deleteRocketmqConsumerGroupHttpUrl
	deleteRocketmqConsumerGroupPath = strings.ReplaceAll(deleteRocketmqConsumerGroupPath, "{project_id}", deleteRocketmqConsumerGroupClient.ProjectID)
	deleteRocketmqConsumerGroupPath = strings.ReplaceAll(deleteRocketmqConsumerGroupPath, "{instance_id}", instanceID)
	deleteRocketmqConsumerGroupPath = strings.ReplaceAll(deleteRocketmqConsumerGroupPath, "{group}", name)

	deleteRocketmqConsumerGroupOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			204,
		},
	}
	_, err = deleteRocketmqConsumerGroupClient.Request("DELETE", deleteRocketmqConsumerGroupPath, &deleteRocketmqConsumerGroupOpt)
	if err != nil {
		return diag.Errorf("error deleting DmsRocketMQConsumerGroup: %s", err)
	}

	d.SetId("")

	return nil
}
