package dms

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
	"github.com/jmespath/go-jmespath"
)

func ResourceDmsRocketMQUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDmsRocketMQUserCreate,
		UpdateContext: resourceDmsRocketMQUserUpdate,
		ReadContext:   resourceDmsRocketMQUserRead,
		DeleteContext: resourceDmsRocketMQUserDelete,
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
			"access_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the access key of the user.`,
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the secret key of the user.`,
			},
			"white_remote_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the IP address whitelist.`,
			},
			"admin": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: `Specifies whether the user is an administrator.`,
			},
			"default_topic_perm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: `Specifies the default topic permissions.
Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.`,
			},
			"default_group_perm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: `Specifies the default consumer group permissions.
Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.`,
			},
			"topic_perms": {
				Type:        schema.TypeList,
				Elem:        DmsRocketMQUserPermsRefSchema(),
				Optional:    true,
				Computed:    true,
				Description: `Specifies the special topic permissions.`,
			},
			"group_perms": {
				Type:        schema.TypeList,
				Elem:        DmsRocketMQUserPermsRefSchema(),
				Optional:    true,
				Computed:    true,
				Description: `Specifies the special consumer group permissions.`,
			},
		},
	}
}

func DmsRocketMQUserPermsRefSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Indicates the name of a topic or consumer group.`,
			},
			"perm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: `Indicates the permissions of the topic or consumer group.
Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.`,
			},
		},
	}
	return &sc
}

func resourceDmsRocketMQUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	// createRocketmqUser: create DMS rocketmq user
	var (
		createRocketmqUserHttpUrl = "v2/{project_id}/instances/{instance_id}/users"
		createRocketmqUserProduct = "dms"
	)
	createRocketmqUserClient, err := config.NewServiceClient(createRocketmqUserProduct, region)
	if err != nil {
		return diag.Errorf("error creating DmsRocketMQUser Client: %s", err)
	}

	instanceID := d.Get("instance_id").(string)
	createRocketmqUserPath := createRocketmqUserClient.Endpoint + createRocketmqUserHttpUrl
	createRocketmqUserPath = strings.ReplaceAll(createRocketmqUserPath, "{project_id}",
		createRocketmqUserClient.ProjectID)
	createRocketmqUserPath = strings.ReplaceAll(createRocketmqUserPath, "{instance_id}",
		fmt.Sprintf("%v", d.Get("instance_id")))

	createRocketmqUserOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
	}
	createRocketmqUserOpt.JSONBody = utils.RemoveNil(buildCreateRocketmqUserBodyParams(d, config))
	createRocketmqUserResp, err := createRocketmqUserClient.Request("POST", createRocketmqUserPath,
		&createRocketmqUserOpt)
	if err != nil {
		return diag.Errorf("error creating DmsRocketMQUser: %s", err)
	}

	createRocketmqUserRespBody, err := utils.FlattenResponse(createRocketmqUserResp)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := jmespath.Search("access_key", createRocketmqUserRespBody)
	if err != nil {
		return diag.Errorf("error creating DmsRocketMQUser: ID is not found in API response")
	}
	d.SetId(instanceID + "/" + id.(string))

	return resourceDmsRocketMQUserRead(ctx, d, meta)
}

func buildCreateRocketmqUserBodyParams(d *schema.ResourceData, config *config.Config) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"access_key":           utils.ValueIngoreEmpty(d.Get("access_key")),
		"secret_key":           utils.ValueIngoreEmpty(d.Get("secret_key")),
		"white_remote_address": utils.ValueIngoreEmpty(d.Get("white_remote_address")),
		"admin":                utils.ValueIngoreEmpty(d.Get("admin")),
		"default_topic_perm":   utils.ValueIngoreEmpty(d.Get("default_topic_perm")),
		"default_group_perm":   utils.ValueIngoreEmpty(d.Get("default_group_perm")),
		"topic_perms":          buildRocketmqUserPermsChildBody(d, "topic_perms"),
		"group_perms":          buildRocketmqUserPermsChildBody(d, "group_perms"),
	}
	return bodyParams
}

func resourceDmsRocketMQUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	updateRocketmqUserhasChanges := []string{
		"white_remote_address",
		"admin",
		"default_topic_perm",
		"default_group_perm",
		"topic_perms",
		"group_perms",
	}

	if d.HasChanges(updateRocketmqUserhasChanges...) {
		// updateRocketmqUser: update DMS rocketmq user
		var (
			updateRocketmqUserHttpUrl = "v2/{project_id}/instances/{instance_id}/users/{user_name}"
			updateRocketmqUserProduct = "dms"
		)
		updateRocketmqUserClient, err := config.NewServiceClient(updateRocketmqUserProduct, region)
		if err != nil {
			return diag.Errorf("error creating DmsRocketMQUser Client: %s", err)
		}

		parts := strings.SplitN(d.Id(), "/", 2)
		if len(parts) != 2 {
			return diag.Errorf("invalid id format, must be <instance_id>/<user>")
		}
		instanceID := parts[0]
		user := parts[1]
		updateRocketmqUserPath := updateRocketmqUserClient.Endpoint + updateRocketmqUserHttpUrl
		updateRocketmqUserPath = strings.ReplaceAll(updateRocketmqUserPath, "{project_id}",
			updateRocketmqUserClient.ProjectID)
		updateRocketmqUserPath = strings.ReplaceAll(updateRocketmqUserPath, "{instance_id}", instanceID)
		updateRocketmqUserPath = strings.ReplaceAll(updateRocketmqUserPath, "{user_name}", user)

		updateRocketmqUserOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
			OkCodes: []int{
				200,
			},
		}
		updateRocketmqUserOpt.JSONBody = utils.RemoveNil(buildUpdateRocketmqUserBodyParams(d, config))
		_, err = updateRocketmqUserClient.Request("PUT", updateRocketmqUserPath, &updateRocketmqUserOpt)
		if err != nil {
			return diag.Errorf("error updating DmsRocketMQUser: %s", err)
		}
	}
	return resourceDmsRocketMQUserRead(ctx, d, meta)
}

func buildUpdateRocketmqUserBodyParams(d *schema.ResourceData, config *config.Config) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"access_key":           fmt.Sprintf("%v", d.Get("access_key")),
		"secret_key":           fmt.Sprintf("%v", d.Get("secret_key")),
		"white_remote_address": utils.ValueIngoreEmpty(d.Get("white_remote_address")),
		"admin":                utils.ValueIngoreEmpty(d.Get("admin")),
		"default_topic_perm":   utils.ValueIngoreEmpty(d.Get("default_topic_perm")),
		"default_group_perm":   utils.ValueIngoreEmpty(d.Get("default_group_perm")),
		"topic_perms":          buildRocketmqUserPermsChildBody(d, "topic_perms"),
		"group_perms":          buildRocketmqUserPermsChildBody(d, "group_perms"),
	}
	return bodyParams
}

func buildRocketmqUserPermsChildBody(d *schema.ResourceData, key string) []map[string]interface{} {
	rawParams := d.Get(key).([]interface{})
	if len(rawParams) == 0 {
		return nil
	}
	params := make([]map[string]interface{}, 0)
	for _, param := range rawParams {
		perm := make(map[string]interface{})
		perm["name"] = utils.PathSearch("name", param, nil)
		perm["perm"] = utils.PathSearch("perm", param, nil)
		params = append(params, perm)
	}
	return params
}

func resourceDmsRocketMQUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	var mErr *multierror.Error

	// getRocketmqUser: query DMS rocketmq user
	var (
		getRocketmqUserHttpUrl = "v2/{project_id}/instances/{instance_id}/users/{user_name}"
		getRocketmqUserProduct = "dms"
	)
	getRocketmqUserClient, err := config.NewServiceClient(getRocketmqUserProduct, region)
	if err != nil {
		return diag.Errorf("error creating DmsRocketMQUser Client: %s", err)
	}

	parts := strings.SplitN(d.Id(), "/", 2)
	if len(parts) != 2 {
		return diag.Errorf("invalid id format, must be <instance_id>/<user>")
	}
	instanceID := parts[0]
	user := parts[1]
	getRocketmqUserPath := getRocketmqUserClient.Endpoint + getRocketmqUserHttpUrl
	getRocketmqUserPath = strings.ReplaceAll(getRocketmqUserPath, "{project_id}", getRocketmqUserClient.ProjectID)
	getRocketmqUserPath = strings.ReplaceAll(getRocketmqUserPath, "{instance_id}", instanceID)
	getRocketmqUserPath = strings.ReplaceAll(getRocketmqUserPath, "{user_name}", user)

	getRocketmqUserOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
	}
	getRocketmqUserResp, err := getRocketmqUserClient.Request("GET", getRocketmqUserPath, &getRocketmqUserOpt)

	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving DmsRocketMQUser")
	}

	getRocketmqUserRespBody, err := utils.FlattenResponse(getRocketmqUserResp)
	if err != nil {
		return diag.FromErr(err)
	}

	mErr = multierror.Append(
		mErr,
		d.Set("region", region),
		d.Set("instance_id", instanceID),
		d.Set("access_key", utils.PathSearch("access_key", getRocketmqUserRespBody, nil)),
		d.Set("secret_key", utils.PathSearch("secret_key", getRocketmqUserRespBody, nil)),
		d.Set("white_remote_address", utils.PathSearch("white_remote_address",
			getRocketmqUserRespBody, nil)),
		d.Set("admin", utils.PathSearch("admin", getRocketmqUserRespBody, nil)),
		d.Set("default_topic_perm", utils.PathSearch("default_topic_perm",
			getRocketmqUserRespBody, nil)),
		d.Set("default_group_perm", utils.PathSearch("default_group_perm",
			getRocketmqUserRespBody, nil)),
		d.Set("topic_perms", flattenGetRocketmqUserResponseBodyPermsRef(getRocketmqUserRespBody,
			"topic_perms")),
		d.Set("group_perms", flattenGetRocketmqUserResponseBodyPermsRef(getRocketmqUserRespBody,
			"group_perms")),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func flattenGetRocketmqUserResponseBodyPermsRef(resp interface{}, expression string) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch(expression, resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"name": utils.PathSearch("name", v, nil),
			"perm": utils.PathSearch("perm", v, nil),
		})
	}
	return rst
}

func resourceDmsRocketMQUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	// deleteRocketmqUser: delete DMS rocketmq user
	var (
		deleteRocketmqUserHttpUrl = "v2/{project_id}/instances/{instance_id}/users/{user_name}"
		deleteRocketmqUserProduct = "dms"
	)
	deleteRocketmqUserClient, err := config.NewServiceClient(deleteRocketmqUserProduct, region)
	if err != nil {
		return diag.Errorf("error creating DmsRocketMQUser Client: %s", err)
	}

	parts := strings.SplitN(d.Id(), "/", 2)
	if len(parts) != 2 {
		return diag.Errorf("invalid id format, must be <instance_id>/<user>")
	}
	instanceID := parts[0]
	user := parts[1]
	deleteRocketmqUserPath := deleteRocketmqUserClient.Endpoint + deleteRocketmqUserHttpUrl
	deleteRocketmqUserPath = strings.ReplaceAll(deleteRocketmqUserPath, "{project_id}",
		deleteRocketmqUserClient.ProjectID)
	deleteRocketmqUserPath = strings.ReplaceAll(deleteRocketmqUserPath, "{instance_id}", instanceID)
	deleteRocketmqUserPath = strings.ReplaceAll(deleteRocketmqUserPath, "{user_name}", user)

	deleteRocketmqUserOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
	}
	_, err = deleteRocketmqUserClient.Request("DELETE", deleteRocketmqUserPath, &deleteRocketmqUserOpt)
	if err != nil {
		return diag.Errorf("error deleting DmsRocketMQUser: %s", err)
	}

	return nil
}
