package fgs

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API FunctionGraph POST /v2/{project_id}/fgs/functions
// @API FunctionGraph PUT /v2/{project_id}/fgs/functions/{function_urn}/config
// @API FunctionGraph PUT /v2/{project_id}/fgs/functions/{function_urn}/config-max-instance
// @API FunctionGraph GET /v2/{project_id}/fgs/functions/{function_urn}/config
// @API FunctionGraph GET /v2/{project_id}/fgs/functions/{function_urn}/versions
// @API FunctionGraph GET /v2/{project_id}/fgs/functions/reservedinstances
// @API FunctionGraph POST /v2/{project_id}/fgs/functions/{function_urn}/tags/create
// @API FunctionGraph DELETE /v2/{project_id}/fgs/functions/{function_urn}/tags/delete
// @API FunctionGraph PUT /v2/{project_id}/fgs/functions/{function_urn}/code
// @API FunctionGraph POST /v2/{project_id}/fgs/functions/{function_urn}/versions
// @API FunctionGraph GET /v2/{project_id}/fgs/functions/{function_urn}/aliases
// @API FunctionGraph POST /v2/{project_id}/fgs/functions/{function_urn}/aliases
// @API FunctionGraph DELETE /v2/{project_id}/fgs/functions/{function_urn}/aliases/{alias_name}
// @API FunctionGraph DELETE /v2/{project_id}/fgs/functions/{function_urn}
// @API FunctionGraph PUT /v2/{project_id}/fgs/functions/{function_urn}/reservedinstances
// @API FunctionGraph GET /v2/{project_id}/fgs/functions/reservedinstanceconfigs
func ResourceFgsFunctionV2() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFunctionCreate,
		ReadContext:   resourceFunctionRead,
		UpdateContext: resourceFunctionUpdate,
		DeleteContext: resourceFunctionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `The region where the function is located.`,
			},

			// Required parameters.
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the function.`,
			},
			"memory_size": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `The memory size allocated to the function, in MByte (MB).`,
			},
			"runtime": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The environment for executing the function.`,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `The timeout interval of the function, in seconds.`,
			},

			// Optional parameters but required in documentation.
			"app": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"package"},
				Description: utils.SchemaDesc(
					`The group to which the function belongs.`,
					utils.SchemaDescInput{
						Required: true,
					},
				),
			},
			"code_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: utils.SchemaDesc(
					`The code type of the function.`,
					utils.SchemaDescInput{
						Required: true,
					},
				),
			},
			"handler": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: utils.SchemaDesc(
					`The entry point of the function.`,
					utils.SchemaDescInput{
						Required: true,
					},
				),
			},

			// Optional parameters.
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The description of the function.`,
			},
			"functiongraph_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"v1", "v2",
				}, false), // The current default value is v1, which may be adjusted in the future.
				Description: `The description of the function.`,
			},
			"func_code": {
				Type:        schema.TypeString,
				Optional:    true,
				StateFunc:   utils.DecodeHashAndHexEncode,
				Description: `The function code.`,
			},
			"code_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The URL where the function code is stored in OBS.`,
			},
			"code_filename": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The name of the function file.`,
			},
			"depend_list": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: `The ID list of the dependencies.`,
			},
			"user_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The key/value information defined for the function.`,
			},
			"encrypted_user_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: `The key/value information defined to be encrypted for the function.`,
			},
			"agency": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"xrole"},
				Description:   `The agency configuration of the function.`,
			},
			"app_agency": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The execution agency enables you to obtain a token or an AK/SK for accessing other cloud services.`,
			},
			"initializer_handler": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The initializer of the function.`,
			},
			"initializer_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: `The maximum duration the function can be initialized.`,
			},
			"enterprise_project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: `The ID of the enterprise project to which the function belongs.`,
			},
			"vpc_id": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"network_id"},
				Description:  `The ID of the VPC to which the function belongs.`,
			},
			"network_id": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"vpc_id"},
				Description:  `The network ID of subnet.`,
			},
			"dns_list": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				RequiredWith: []string{"vpc_id"},
				Description:  `The private DNS configuration of the function network.`,
			},
			"mount_user_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: `The mount user ID.`,
			},
			"mount_user_group_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: `The mount user group ID.`,
			},
			"func_mounts": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mount_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The mount type.`,
						},
						"mount_resource": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The ID of the mounted resource (corresponding cloud service).`,
						},
						"mount_share_path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The remote mount path.`,
						},
						"local_mount_path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The function access path.`,
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The mount status.`,
						},
					},
				},
				Description: `The list of function mount configuration.`,
			},
			"custom_image": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The URL of SWR image.`,
						},
						"command": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The startup commands of the SWR image.`,
						},
						"args": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The command line arguments used to start the SWR image.`,
						},
						"working_dir": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: `The working directory of the SWR image.`,
						},
						"user_id": {
							Type:     schema.TypeString,
							Optional: true,
							Description: utils.SchemaDesc(
								`The user ID that used to run SWR image.`,
								utils.SchemaDescInput{
									Internal: true,
								},
							),
						},
						"user_group_id": {
							Type:     schema.TypeString,
							Optional: true,
							Description: utils.SchemaDesc(
								`The user group ID that used to run SWR image.`,
								utils.SchemaDescInput{
									Internal: true,
								},
							),
						},
					},
				},
				ConflictsWith: []string{
					"code_type",
				},
				Description: `The custom image configuration of the function.`,
			},
			"max_instance_num": {
				// The original type of this parameter is int, but its zero value is meaningful.
				// So, the following types of parameter passing are realized through the logic of terraform's implicit
				// conversion of int:
				//   + -1: the number of instances is unlimited.
				//   + 0: this function is disabled.
				//   + (0, +1000]: Specific value (2023.06.26).
				//   + empty: keep the default (latest updated) value.
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The maximum number of instances of the function.`,
			},
			"versions": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The version name.`,
						},
						"aliases": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The name of the version alias.`,
									},
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The description of the version alias.`,
									},
								},
							},
							Description: `The aliases management for specified version.`,
						},
					},
				},
				Description: `The versions management of the function.`,
			},
			"tags": common.TagsSchema(
				`The key/value pairs to associate with the function.`,
			),
			"log_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				RequiredWith: []string{
					"log_stream_id",
					"log_group_name",
					"log_stream_name",
				},
				Description: `The LTS group ID for collecting logs.`,
			},
			"log_group_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				RequiredWith: []string{"log_group_id"},
				Description:  `The LTS group name for collecting logs.`,
			},
			"log_stream_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				RequiredWith: []string{"log_group_id"},
				Description:  `The LTS stream ID for collecting logs.`,
			},
			"log_stream_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				RequiredWith: []string{"log_group_id"},
				Description:  `The LTS stream name for collecting logs.`,
			},
			"reserved_instances": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"qualifier_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The qualifier type of reserved instance.`,
						},
						"qualifier_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The version name or alias name.`,
						},
						"count": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: `The number of reserved instance.`,
						},
						"idle_mode": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Whether to enable the idle mode.`,
						},
						"tactics_config": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem:        tracticsConfigsSchema(),
							Description: `The auto scaling policies for reserved instance.`,
						},
					},
				},
				Description: `The reserved instance policies of the function.`,
			},
			// The value in the api document is -1 to 1000, After confirmation, when the parameter set to -1 or 0,
			// the actual number of concurrent requests is 1, so the value range is set to 1 to 1000, and the document
			// will be modified later (2024.02.29).
			"concurrency_num": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: `The number of concurrent requests of the function.`,
			},
			"gpu_type": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"gpu_memory"},
				Description:  `The GPU type of the function.`,
			},
			"gpu_memory": {
				Type:         schema.TypeInt,
				Optional:     true,
				RequiredWith: []string{"gpu_type"},
				Description:  `The GPU memory size allocated to the function, in MByte (MB).`,
			},
			// Currently, the "pre_stop_timeout" and "pre_stop_timeout" are not visible on the page,
			// so they are temporarily used as internal parameters.
			"pre_stop_handler": {
				Type:     schema.TypeString,
				Optional: true,
				Description: utils.SchemaDesc(
					`The pre-stop handler of a function.`,
					utils.SchemaDescInput{
						Internal: true,
					},
				),
			},
			"pre_stop_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: utils.SchemaDesc(
					`The maximum duration that the function can be initialized.`,
					utils.SchemaDescInput{
						Internal: true,
					},
				),
			},

			// Deprecated parameters.
			"package": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"app"},
				Deprecated:    `use app instead`,
			},
			"xrole": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"agency"},
				Deprecated:    `use agency instead`,
			},

			// Attributes.
			"urn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The URN (Uniform Resource Name) of the function.`,
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The version of the function.`,
			},
		},
	}
}

func tracticsConfigsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cron_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The name of scheduled policy configuration.`,
						},
						"cron": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The cron expression.`,
						},
						"count": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: `The number of reserved instance to which the policy belongs.`,
						},
						"start_time": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: `The effective timestamp of policy.`,
						},
						"expired_time": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: `The expiration timestamp of the policy.`,
						},
					},
				},
				Description: `The list of scheduled policy configurations.`,
			},
			"metric_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The name of metric policy.`,
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The type of metric policy.`,
						},
						"threshold": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: `The metric policy threshold.`,
						},
						"min": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: `The minimun of traffic.`,
						},
					},
				},
				Description: `The list of metric policy configurations.`,
			},
		},
	}
}

func buildFunctionCustomImage(imageConfigs []interface{}) map[string]interface{} {
	if len(imageConfigs) < 1 {
		return nil
	}

	imageConfig := imageConfigs[0]
	return map[string]interface{}{
		"enabled":     true,
		"image":       utils.ValueIgnoreEmpty(utils.PathSearch("url", imageConfig, nil)),
		"command":     utils.ValueIgnoreEmpty(utils.PathSearch("command", imageConfig, nil)),
		"args":        utils.ValueIgnoreEmpty(utils.PathSearch("args", imageConfig, nil)),
		"working_dir": utils.ValueIgnoreEmpty(utils.PathSearch("working_dir", imageConfig, nil)),
		"uid":         utils.ValueIgnoreEmpty(utils.PathSearch("user_id", imageConfig, nil)),
		"gid":         utils.ValueIgnoreEmpty(utils.PathSearch("user_group_id", imageConfig, nil)),
	}
}

func buildFunctionCodeConfig(funcCode string) map[string]interface{} {
	if funcCode == "" {
		return nil
	}

	return map[string]interface{}{
		"file": utils.TryBase64EncodeString(funcCode),
	}
}

func buildFunctionLogConfig(d *schema.ResourceData) map[string]interface{} {
	// If the LTS log parameters is not configured (in the creation phase), the service will automatically create an
	// LTS stream (group will also be created) and associate it with the function.
	// So, the tfstate records will always have the log group ID and the log stream ID.
	groupName, ok := d.GetOk("log_group_name") // Only log group name and the log stream name are specified by users.
	if !ok {
		return nil
	}

	return map[string]interface{}{
		"group_id":    d.Get("log_group_id"),
		"group_name":  groupName,
		"stream_id":   d.Get("log_stream_id"),
		"stream_name": utils.ValueIgnoreEmpty(d.Get("log_stream_name")),
	}
}

func buildCreateFunctionBodyParams(cfg *config.Config, d *schema.ResourceData) map[string]interface{} {
	// Parameter app is recommended to replace parameter package.
	pkg, ok := d.GetOk("app")
	if !ok {
		pkg = d.Get("package")
	}

	// Parameter agency is recommended to replace parameter xrole.
	agency, ok := d.GetOk("agency")
	if !ok {
		agency = d.Get("xrole")
	}

	return map[string]interface{}{
		// Required parameters.
		"func_name":   d.Get("name"),
		"runtime":     d.Get("runtime"),
		"timeout":     d.Get("timeout"),
		"memory_size": d.Get("memory_size"),
		// Optional parameters but required in documentation.
		"package":   utils.ValueIgnoreEmpty(pkg),
		"handler":   utils.ValueIgnoreEmpty(d.Get("handler")),
		"code_type": utils.ValueIgnoreEmpty(d.Get("code_type")),
		// Optional parameters.
		"description":           utils.ValueIgnoreEmpty(d.Get("description")),
		"type":                  utils.ValueIgnoreEmpty(d.Get("functiongraph_version")),
		"code_url":              utils.ValueIgnoreEmpty(d.Get("code_url")),
		"code_filename":         utils.ValueIgnoreEmpty(d.Get("code_filename")),
		"user_data":             utils.ValueIgnoreEmpty(d.Get("user_data")),
		"encrypted_user_data":   utils.ValueIgnoreEmpty(d.Get("encrypted_user_data")),
		"xrole":                 utils.ValueIgnoreEmpty(agency),
		"enterprise_project_id": cfg.GetEnterpriseProjectID(d),
		"custom_image":          buildFunctionCustomImage(d.Get("custom_image").([]interface{})),
		"gpu_memory":            utils.ValueIgnoreEmpty(d.Get("gpu_memory")),
		"gpu_type":              utils.ValueIgnoreEmpty(d.Get("gpu_type")),
		"pre_stop_handler":      utils.ValueIgnoreEmpty(d.Get("pre_stop_handler")),
		"pre_stop_timeout":      utils.ValueIgnoreEmpty(d.Get("pre_stop_timeout")),
		"func_code":             buildFunctionCodeConfig(d.Get("func_code").(string)),
		"log_config":            buildFunctionLogConfig(d),
	}
}

func createFunction(cfg *config.Config, client *golangsdk.ServiceClient, d *schema.ResourceData) (string, error) {
	httpUrl := "v2/{project_id}/fgs/functions"

	createPath := client.Endpoint + httpUrl
	createPath = strings.ReplaceAll(createPath, "{project_id}", client.ProjectID)
	createOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		JSONBody: utils.RemoveNil(buildCreateFunctionBodyParams(cfg, d)),
	}

	requestResp, err := client.Request("POST", createPath, &createOpt)
	if err != nil {
		return "", err
	}

	respBody, err := utils.FlattenResponse(requestResp)
	if err != nil {
		return "", err
	}

	return utils.PathSearch("func_urn", respBody, "").(string), nil
}

func buildFunctionVpcConfig(d *schema.ResourceData) map[string]interface{} {
	vpcId, ok := d.GetOk("vpc_id")
	if !ok {
		return nil
	}

	return map[string]interface{}{
		"vpc_id":    vpcId,
		"subnet_id": d.Get("network_id"),
	}
}

func parsedFunctionMountId(mountId int) int {
	if mountId < 1 {
		return -1
	}
	return mountId
}

func buildFunctionMountConfig(mounts []interface{}, mountUserId, mountGroupId int) map[string]interface{} {
	if len(mounts) < 1 {
		return nil
	}

	parsedMounts := make([]interface{}, 0, len(mounts))
	for _, mount := range mounts {
		parsedMounts = append(parsedMounts, map[string]interface{}{
			"mount_type":       utils.PathSearch("mount_type", mount, nil),
			"mount_resource":   utils.PathSearch("mount_resource", mount, nil),
			"mount_share_path": utils.PathSearch("mount_share_path", mount, nil),
			"local_mount_path": utils.PathSearch("local_mount_path", mount, nil),
		})
	}

	return map[string]interface{}{
		"mount_config": parsedMounts,
		"mount_user": map[string]interface{}{
			"mount_user_id":       parsedFunctionMountId(mountUserId),
			"mount_user_group_id": parsedFunctionMountId(mountGroupId),
		},
	}
}

func buildFunctionStrategyConfig(concurrencyNum int) map[string]interface{} {
	if concurrencyNum < 1 {
		return nil
	}

	return map[string]interface{}{
		"concurrent_num": concurrencyNum,
	}
}

func buildUpdateFunctionMetadataBodyParams(d *schema.ResourceData) map[string]interface{} {
	// Parameter app is recommended to replace parameter package.
	pkg, ok := d.GetOk("app")
	if !ok {
		pkg = d.Get("package")
	}

	// Parameter agency is recommended to replace parameter xrole.
	agency, ok := d.GetOk("agency")
	if !ok {
		pkg = d.Get("xrole")
	}

	return map[string]interface{}{
		// Required parameters.
		"runtime":     d.Get("runtime"),
		"timeout":     d.Get("timeout"),
		"memory_size": d.Get("memory_size"),
		// Optional parameters but required in documentation.
		"package": utils.ValueIgnoreEmpty(pkg),
		"handler": utils.ValueIgnoreEmpty(d.Get("handler")),
		// Optional parameters.
		"description":         d.Get("description"),
		"user_data":           utils.ValueIgnoreEmpty(d.Get("user_data")),
		"encrypted_user_data": utils.ValueIgnoreEmpty(d.Get("encrypted_user_data")),
		"xrole":               utils.ValueIgnoreEmpty(agency),
		"app_xrole":           utils.ValueIgnoreEmpty(d.Get("app_agency")),
		"custom_image":        buildFunctionCustomImage(d.Get("custom_image").([]interface{})),
		"gpu_memory":          utils.ValueIgnoreEmpty(d.Get("gpu_memory")),
		"gpu_type":            utils.ValueIgnoreEmpty(d.Get("gpu_type")),
		"pre_stop_handler":    utils.ValueIgnoreEmpty(d.Get("pre_stop_handler")),
		"pre_stop_timeout":    utils.ValueIgnoreEmpty(d.Get("pre_stop_timeout")),
		"log_config":          buildFunctionLogConfig(d),
		"domain_names":        utils.ValueIgnoreEmpty(d.Get("dns_list")),
		"func_vpc":            buildFunctionVpcConfig(d),
		"func_mounts": buildFunctionMountConfig(d.Get("func_mounts").([]interface{}),
			d.Get("mount_user_id").(int), d.Get("mount_user_group_id").(int)),
		"strategy_config": buildFunctionStrategyConfig(d.Get("concurrency_num").(int)),
	}
}

func updateFunctionMetadata(client *golangsdk.ServiceClient, d *schema.ResourceData, functionUrn string) error {
	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}/config"

	updatePath := client.Endpoint + httpUrl
	updatePath = strings.ReplaceAll(updatePath, "{project_id}", client.ProjectID)
	updatePath = strings.ReplaceAll(updatePath, "{function_urn}", functionUrn)
	updateOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		JSONBody: utils.RemoveNil(buildUpdateFunctionMetadataBodyParams(d)),
	}

	_, err := client.Request("PUT", updatePath, &updateOpt)
	if err != nil {
		return fmt.Errorf("failed to update the function metadata: %s", err)
	}
	return nil
}

func buildUpdateFunctionCodeBodyParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"code_type":           d.Get("code_type"),
		"code_url":            d.Get("code_url"),
		"code_filename":       d.Get("code_filename"),
		"depend_version_list": d.Get("depend_list").(*schema.Set).List(),
		"func_code":           buildFunctionCodeConfig(d.Get("func_code").(string)),
	}
}

func updateFunctionCode(client *golangsdk.ServiceClient, d *schema.ResourceData, functionUrn string) error {
	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}/code"

	updatePath := client.Endpoint + httpUrl
	updatePath = strings.ReplaceAll(updatePath, "{project_id}", client.ProjectID)
	updatePath = strings.ReplaceAll(updatePath, "{function_urn}", functionUrn)
	updateOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		JSONBody: utils.RemoveNil(buildUpdateFunctionCodeBodyParams(d)),
	}

	_, err := client.Request("PUT", updatePath, &updateOpt)
	if err != nil {
		return fmt.Errorf("failed to update the function code: %s", err)
	}
	return nil
}

func updateFunctionMaxInstanceNum(client *golangsdk.ServiceClient, functionUrn string, maxInstanceNum int) error {
	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}/config-max-instance"

	updatePath := client.Endpoint + httpUrl
	updatePath = strings.ReplaceAll(updatePath, "{project_id}", client.ProjectID)
	updatePath = strings.ReplaceAll(updatePath, "{function_urn}", functionUrn)
	updateOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		JSONBody: map[string]interface{}{
			"max_instance_num": maxInstanceNum,
		},
	}

	_, err := client.Request("PUT", updatePath, &updateOpt)
	if err != nil {
		return fmt.Errorf("failed to update the function max instance number: %s", err)
	}
	return nil
}

func buildFunctionTagsBodyParams(tags map[string]interface{}) map[string]interface{} {
	tagsList := make([]interface{}, 0, len(tags))

	for k, v := range tags {
		tagsList = append(tagsList, map[string]interface{}{
			"key":   k,
			"value": v,
		})
	}

	return map[string]interface{}{
		"tags": tagsList,
	}
}

func createFunctionTags(client *golangsdk.ServiceClient, functionUrn string, tags map[string]interface{}) error {
	if len(tags) < 1 {
		return nil
	}

	httpUrl := "v2/{project_id}/functions/{function_urn}/tags/create"

	createPath := client.Endpoint + httpUrl
	createPath = strings.ReplaceAll(createPath, "{project_id}", client.ProjectID)
	createPath = strings.ReplaceAll(createPath, "{function_urn}", functionUrn)
	createOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		OkCodes:  []int{204},
		JSONBody: buildFunctionTagsBodyParams(tags),
	}
	_, err := client.Request("POST", createPath, &createOpt)
	if err != nil {
		return fmt.Errorf("failed to create the function tags: %s", err)
	}
	return nil
}

func createFunctionVersion(client *golangsdk.ServiceClient, functionUrn, versionName string) error {
	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}/versions"

	createPath := client.Endpoint + httpUrl
	createPath = strings.ReplaceAll(createPath, "{project_id}", client.ProjectID)
	createPath = strings.ReplaceAll(createPath, "{function_urn}", functionUrn)
	createOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		JSONBody: map[string]interface{}{
			"version": versionName,
		},
	}

	_, err := client.Request("POST", createPath, &createOpt)
	if err != nil {
		return fmt.Errorf("failed to create the function version: %s", err)
	}
	return nil
}

func createFunctionVersionAlias(client *golangsdk.ServiceClient, functionUrn, versionName string, aliasCfg interface{}) error {
	if aliasCfg == nil {
		return nil
	}

	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}/aliases"

	createPath := client.Endpoint + httpUrl
	createPath = strings.ReplaceAll(createPath, "{project_id}", client.ProjectID)
	createPath = strings.ReplaceAll(createPath, "{function_urn}", functionUrn)
	createOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		JSONBody: map[string]interface{}{
			"version":     versionName,
			"name":        utils.ValueIgnoreEmpty(utils.PathSearch("name", aliasCfg, nil)),
			"description": utils.ValueIgnoreEmpty(utils.PathSearch("description", aliasCfg, nil)),
		},
	}

	_, err := client.Request("POST", createPath, &createOpt)
	if err != nil {
		return fmt.Errorf("failed to create the function version alias: %s", err)
	}
	return nil
}

func createFunctionVersions(client *golangsdk.ServiceClient, functionUrn string, versions []interface{}) error {
	for _, version := range versions {
		versionName := utils.PathSearch("name", version, "").(string)
		if versionName != "latest" {
			err := createFunctionVersion(client, functionUrn, versionName)
			if err != nil {
				return err
			}
		}
		// In the future, the function will support manage multiple versions, and will add the corresponding logic to
		// create versions based on the related API (Create) in this place.
		aliases := utils.PathSearch("aliases", version, make([]interface{}, 0)).([]interface{})
		for _, alias := range aliases {
			err := createFunctionVersionAlias(client, functionUrn, versionName, alias)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getFunctionVersionUrn(client *golangsdk.ServiceClient, functionUrn string, qualifierName string) (string, error) {
	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}/versions"

	// The query parameter 'marker' and 'maxitems' are not available.
	listPath := client.Endpoint + httpUrl
	listPath = strings.ReplaceAll(listPath, "{project_id}", client.ProjectID)
	listPath = strings.ReplaceAll(listPath, "{function_urn}", functionUrn)
	listOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}
	requestResp, err := client.Request("GET", listPath, &listOpt)
	if err != nil {
		return "", fmt.Errorf("failed to query the function versions: %s", err)
	}
	respBody, err := utils.FlattenResponse(requestResp)
	if err != nil {
		return "", err
	}

	return utils.PathSearch(fmt.Sprintf("versions[?version=='%s']|[0].func_urn", qualifierName), respBody, "").(string), nil
}

func getFunctionAliasUrn(client *golangsdk.ServiceClient, functionUrn string, qualifierName string) (string, error) {
	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}/aliases"

	listPath := client.Endpoint + httpUrl
	listPath = strings.ReplaceAll(listPath, "{project_id}", client.ProjectID)
	listPath = strings.ReplaceAll(listPath, "{function_urn}", functionUrn)
	listOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}
	requestResp, err := client.Request("GET", listPath, &listOpt)
	if err != nil {
		return "", fmt.Errorf("failed to query the function aliases: %s", err)
	}
	respBody, err := utils.FlattenResponse(requestResp)
	if err != nil {
		return "", err
	}

	return utils.PathSearch(fmt.Sprintf("[?version=='%s']|[0].func_urn", qualifierName), respBody, "").(string), nil
}

func getReservedInstanceUrn(client *golangsdk.ServiceClient, functionUrn, qualifierType, qualifierName string) (string, error) {
	if qualifierType == "version" {
		return getFunctionVersionUrn(client, functionUrn, qualifierName)
	}

	return getFunctionAliasUrn(client, functionUrn, qualifierName)
}

func deleteFunctionReservedInstances(client *golangsdk.ServiceClient, functionUrn string, policies []interface{}) error {
	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}/reservedinstances"

	basePath := client.Endpoint + httpUrl
	basePath = strings.ReplaceAll(basePath, "{project_id}", client.ProjectID)

	for _, policy := range policies {
		qualifierType := utils.PathSearch("qualifier_type", policy, "").(string)
		qualifierName := utils.PathSearch("qualifier_name", policy, "").(string)

		urn, err := getReservedInstanceUrn(client, functionUrn, qualifierType, qualifierName)
		if err != nil {
			return err
		}
		// Deleting the alias will also delete the corresponding reserved instance.
		if urn == "" {
			return nil
		}

		deletePath := basePath
		deletePath = strings.ReplaceAll(deletePath, "{function_urn}", urn)
		deleteOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
			MoreHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			JSONBody: map[string]interface{}{
				"count":     0,
				"idle_mode": false,
			},
		}

		_, err = client.Request("PUT", deletePath, &deleteOpt)
		if err != nil {
			return fmt.Errorf("failed to remove the function reversed instance: %s", err)
		}
	}

	return nil
}

func buildTracticsConfigCronConfigs(cronConfigs []interface{}) []map[string]interface{} {
	if len(cronConfigs) < 1 {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(cronConfigs))
	for _, cronConfig := range cronConfigs {
		result = append(result, map[string]interface{}{
			"name":         utils.ValueIgnoreEmpty(utils.PathSearch("name", cronConfig, nil)),
			"cron":         utils.ValueIgnoreEmpty(utils.PathSearch("cron", cronConfig, nil)),
			"count":        utils.ValueIgnoreEmpty(utils.PathSearch("count", cronConfig, nil)),
			"start_time":   utils.ValueIgnoreEmpty(utils.PathSearch("start_time", cronConfig, nil)),
			"expired_time": utils.ValueIgnoreEmpty(utils.PathSearch("expired_time", cronConfig, nil)),
		})
	}

	return result
}

func buildTracticsConfigMetricConfigs(metricConfigs []interface{}) []map[string]interface{} {
	if len(metricConfigs) < 1 {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(metricConfigs))
	for _, metricConfig := range metricConfigs {
		result = append(result, map[string]interface{}{
			"name":      utils.ValueIgnoreEmpty(utils.PathSearch("name", metricConfig, nil)),
			"type":      utils.ValueIgnoreEmpty(utils.PathSearch("type", metricConfig, nil)),
			"threshold": utils.ValueIgnoreEmpty(utils.PathSearch("threshold", metricConfig, nil)),
			"min":       utils.ValueIgnoreEmpty(utils.PathSearch("min", metricConfig, nil)),
		})
	}

	return result
}

func buildReservedInstanceTracticsConfigs(tacticsConfigs []interface{}) map[string]interface{} {
	if len(tacticsConfigs) < 1 {
		return nil
	}

	return map[string]interface{}{
		"cron_configs": buildTracticsConfigCronConfigs(utils.PathSearch("cron_configs",
			tacticsConfigs[0], make([]interface{}, 0)).([]interface{})),
		"metric_configs": buildTracticsConfigMetricConfigs(utils.PathSearch("metric_configs",
			tacticsConfigs[0], make([]interface{}, 0)).([]interface{})),
	}
}

func createFunctionReservedInstances(client *golangsdk.ServiceClient, functionUrn string, policies []interface{}) error {
	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}/reservedinstances"

	basePath := client.Endpoint + httpUrl
	basePath = strings.ReplaceAll(basePath, "{project_id}", client.ProjectID)

	for _, policy := range policies {
		qualifierType := utils.PathSearch("qualifier_type", policy, "").(string)
		qualifierName := utils.PathSearch("qualifier_name", policy, "").(string)

		urn, err := getReservedInstanceUrn(client, functionUrn, qualifierType, qualifierName)
		if err != nil {
			return err
		}

		createPath := basePath
		createPath = strings.ReplaceAll(createPath, "{function_urn}", urn)
		createOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
			MoreHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			JSONBody: map[string]interface{}{
				"count":     utils.PathSearch("count", policy, nil),
				"idle_mode": utils.PathSearch("idle_mode", policy, nil),
				"tactics_config": buildReservedInstanceTracticsConfigs(utils.PathSearch("tactics_config",
					policy, make([]interface{}, 0)).([]interface{})),
			},
		}

		_, err = client.Request("PUT", createPath, &createOpt)
		if err != nil {
			return fmt.Errorf("failed to configure the function reversed instance: %s", err)
		}
	}

	return nil
}

func updateFunctionReservedInstances(client *golangsdk.ServiceClient, d *schema.ResourceData, functionUrn string) error {
	oldVal, newVal := d.GetChange("reserved_instances")

	oldRaws := oldVal.(*schema.Set).Difference(newVal.(*schema.Set))
	newRaws := newVal.(*schema.Set).Difference(oldVal.(*schema.Set))

	if oldRaws.Len() > 0 {
		if err := deleteFunctionReservedInstances(client, functionUrn, oldRaws.List()); err != nil {
			return err
		}
	}

	if newRaws.Len() > 0 {
		if err := createFunctionReservedInstances(client, functionUrn, newRaws.List()); err != nil {
			return err
		}
	}

	return nil
}

func resourceFunctionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg    = meta.(*config.Config)
		region = cfg.GetRegion(d)
	)

	client, err := cfg.NewServiceClient("fgs", region)
	if err != nil {
		return diag.Errorf("error creating FunctionGraph client: %s", err)
	}

	funcUrn, err := createFunction(cfg, client, d)
	if err != nil {
		return diag.Errorf("error creating function: %s", err)
	}
	if funcUrn == "" {
		return diag.Errorf("unable to find the function URN from the API response")
	}
	d.SetId(funcUrn)
	funcUrnWithoutVersion := parseFunctionUrnWithoutVersion(funcUrn)

	// lintignore:R019
	if d.HasChanges("vpc_id", "func_mounts", "app_agency", "initializer_handler", "initializer_timeout", "concurrency_num") {
		err = updateFunctionMetadata(client, d, funcUrnWithoutVersion)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("depend_list") {
		err := updateFunctionCode(client, d, funcUrnWithoutVersion)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if strNum, ok := d.GetOk("max_instance_num"); ok {
		// If the maximum number of instances is omitted (after type conversion, the value is zero), means this feature
		// is disabled.
		maxInstanceNum, _ := strconv.Atoi(strNum.(string))
		err = updateFunctionMaxInstanceNum(client, funcUrnWithoutVersion, maxInstanceNum)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if tags, ok := d.GetOk("tags"); ok {
		if err := createFunctionTags(client, funcUrn, tags.(map[string]interface{})); err != nil {
			return diag.FromErr(err)
		}
	}

	if err = createFunctionVersions(client, funcUrnWithoutVersion, d.Get("versions").(*schema.Set).List()); err != nil {
		return diag.FromErr(err)
	}

	if d.HasChanges("reserved_instances") {
		if err = updateFunctionReservedInstances(client, d, funcUrnWithoutVersion); err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceFunctionRead(ctx, d, meta)
}

func GetFunctionMetadata(client *golangsdk.ServiceClient, functionUrn string) (interface{}, error) {
	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}/config"

	getPath := client.Endpoint + httpUrl
	getPath = strings.ReplaceAll(getPath, "{project_id}", client.ProjectID)
	getPath = strings.ReplaceAll(getPath, "{function_urn}", functionUrn)
	getOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}

	requestResp, err := client.Request("GET", getPath, &getOpt)
	if err != nil {
		return nil, fmt.Errorf("error querying function (%s) metadata: %s", functionUrn, err)
	}

	return utils.FlattenResponse(requestResp)
}

func flattenFgsCustomImage(imageConfig map[string]interface{}) []map[string]interface{} {
	if len(imageConfig) < 1 {
		return nil
	}

	return []map[string]interface{}{
		{
			"url":           utils.PathSearch("image", imageConfig, nil),
			"command":       utils.PathSearch("command", imageConfig, nil),
			"args":          utils.PathSearch("args", imageConfig, nil),
			"working_dir":   utils.PathSearch("working_dir", imageConfig, nil),
			"user_id":       utils.PathSearch("uid", imageConfig, nil),
			"user_group_id": utils.PathSearch("gid", imageConfig, nil),
		},
	}
}

func flattenFuncionMounts(mounts []interface{}) []map[string]interface{} {
	if len(mounts) < 1 {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(mounts))
	for _, mount := range mounts {
		result = append(result, map[string]interface{}{
			"mount_type":       utils.PathSearch("mount_type", mount, nil),
			"mount_resource":   utils.PathSearch("mount_resource", mount, nil),
			"mount_share_path": utils.PathSearch("mount_share_path", mount, nil),
			"local_mount_path": utils.PathSearch("local_mount_path", mount, nil),
			"status":           utils.PathSearch("status", mount, nil),
		})
	}

	return result
}

func getFunctionVersions(client *golangsdk.ServiceClient, functionUrn string) ([]interface{}, error) {
	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}/versions"

	listPath := client.Endpoint + httpUrl
	listPath = strings.ReplaceAll(listPath, "{project_id}", client.ProjectID)
	listPath = strings.ReplaceAll(listPath, "{function_urn}", functionUrn)
	listOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}

	requestResp, err := client.Request("GET", listPath, &listOpt)
	if err != nil {
		return nil, err
	}
	respBody, err := utils.FlattenResponse(requestResp)
	if err != nil {
		return nil, err
	}

	return utils.PathSearch("versions", respBody, make([]interface{}, 0)).([]interface{}), nil
}

func getFunctionAliases(client *golangsdk.ServiceClient, functionUrn string) ([]interface{}, error) {
	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}/aliases"

	listPath := client.Endpoint + httpUrl
	listPath = strings.ReplaceAll(listPath, "{project_id}", client.ProjectID)
	listPath = strings.ReplaceAll(listPath, "{function_urn}", functionUrn)
	listOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}

	requestResp, err := client.Request("GET", listPath, &listOpt)
	if err != nil {
		return nil, err
	}
	respBody, err := utils.FlattenResponse(requestResp)
	if err != nil {
		return nil, err
	}

	return utils.PathSearch("[]", respBody, make([]interface{}, 0)).([]interface{}), nil
}

func flattenFunctionVersionAliases(aliases []interface{}) []map[string]interface{} {
	if len(aliases) < 1 {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(aliases))
	for _, alias := range aliases {
		result = append(result, map[string]interface{}{
			"name":        utils.PathSearch("name", alias, nil),
			"description": utils.PathSearch("description", alias, nil),
		})
	}

	return result
}

func flattenFunctionVersions(client *golangsdk.ServiceClient, functionUrn string) ([]map[string]interface{}, error) {
	versionList, err := getFunctionVersions(client, functionUrn)
	if err != nil {
		return nil, err
	}
	aliasesConfig, err := getFunctionAliases(client, functionUrn)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0, len(versionList))
	for _, version := range versionList {
		versionName := utils.PathSearch("version", version, "").(string)
		if versionName == "" {
			log.Printf("[DEBUG] The version name is not found from the API response: %v", version)
			continue
		}
		aliases := utils.PathSearch(fmt.Sprintf("[?version=='%s']", versionName),
			aliasesConfig, make([]interface{}, 0)).([]interface{})
		if versionName == "latest" && len(aliases) < 1 {
			continue
		}
		result = append(result, utils.RemoveNil(map[string]interface{}{
			"name":    versionName,
			"aliases": flattenFunctionVersionAliases(aliases),
		}))
	}

	return result, nil
}

func getFunctionReservedInstances(client *golangsdk.ServiceClient, functionUrn string) ([]interface{}, error) {
	var (
		httpUrl = "v2/{project_id}/fgs/functions/reservedinstanceconfigs?function_urn={function_urn}&limit=100"
		marker  = 0
	)

	listPath := client.Endpoint + httpUrl
	listPath = strings.ReplaceAll(listPath, "{project_id}", client.ProjectID)
	listPath = strings.ReplaceAll(listPath, "{function_urn}", functionUrn)
	listOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}

	result := make([]interface{}, 0)
	for {
		listPathWithMarker := fmt.Sprintf("%s&marker=%d", listPath, marker)
		requestResp, err := client.Request("GET", listPathWithMarker, &listOpt)
		if err != nil {
			return nil, fmt.Errorf("error querying function reserved instances: %s", err)
		}
		respBody, err := utils.FlattenResponse(requestResp)
		if err != nil {
			return nil, err
		}
		reservedInstances := utils.PathSearch("reserved_instances", respBody, make([]interface{}, 0)).([]interface{})
		if len(reservedInstances) < 1 {
			break
		}
		result = append(result, reservedInstances...)
		marker += len(reservedInstances)
	}

	return result, nil
}

func flattenReservedInstanceCronConfig(cronConfigs []interface{}) []map[string]interface{} {
	if len(cronConfigs) < 1 {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(cronConfigs))
	for _, cronConfig := range cronConfigs {
		result = append(result, map[string]interface{}{
			"name":         utils.PathSearch("name", cronConfig, nil),
			"cron":         utils.PathSearch("cron", cronConfig, nil),
			"count":        utils.PathSearch("count", cronConfig, nil),
			"start_time":   utils.PathSearch("start_time", cronConfig, nil),
			"expired_time": utils.PathSearch("expired_time", cronConfig, nil),
		})
	}

	return result
}

func flattenReservedInstanceMetricConfig(metricConfigs []interface{}) []map[string]interface{} {
	if len(metricConfigs) < 1 {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(metricConfigs))
	for _, metricConfig := range metricConfigs {
		result = append(result, map[string]interface{}{
			"name":      utils.PathSearch("name", metricConfig, nil),
			"type":      utils.PathSearch("type", metricConfig, nil),
			"threshold": utils.PathSearch("threshold", metricConfig, nil),
			"min":       utils.PathSearch("min", metricConfig, nil),
		})
	}

	return result
}

func flattenReservedInstanceTracticsConfig(tracticsConfig map[string]interface{}) []map[string]interface{} {
	if len(tracticsConfig) < 1 {
		return nil
	}

	return []map[string]interface{}{
		{
			"cron_configs": flattenReservedInstanceCronConfig(utils.PathSearch("cron_configs",
				tracticsConfig, make([]interface{}, 0)).([]interface{})),
			"metric_configs": flattenReservedInstanceMetricConfig(utils.PathSearch("metric_configs",
				tracticsConfig, make([]interface{}, 0)).([]interface{})),
		},
	}
}

func flattenReservedInstances(reservedInstances []interface{}) []map[string]interface{} {
	if len(reservedInstances) < 1 {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(reservedInstances))
	for _, reservedInstance := range reservedInstances {
		result = append(result, map[string]interface{}{
			"count":          utils.PathSearch("min_count", reservedInstance, nil),
			"idle_mode":      utils.PathSearch("idle_mode", reservedInstance, nil),
			"qualifier_name": utils.PathSearch("qualifier_name", reservedInstance, nil),
			"qualifier_type": utils.PathSearch("qualifier_type", reservedInstance, nil),
			"tactics_config": flattenReservedInstanceTracticsConfig(utils.PathSearch("tactics_config",
				reservedInstance, make(map[string]interface{})).(map[string]interface{})),
		})
	}
	return result
}

func resourceFunctionRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg                   = meta.(*config.Config)
		region                = cfg.GetRegion(d)
		funcUrn               = d.Id()
		funcUrnWithoutVersion = parseFunctionUrnWithoutVersion(funcUrn)
	)
	client, err := cfg.NewServiceClient("fgs", region)
	if err != nil {
		return diag.Errorf("error creating FunctionGraph client: %s", err)
	}

	function, err := GetFunctionMetadata(client, funcUrnWithoutVersion)
	if err != nil {
		return common.CheckDeletedDiag(d, err, "FunctionGraph function")
	}

	log.Printf("[DEBUG] Retrieved Function %s: %+v", funcUrnWithoutVersion, function)
	mErr := multierror.Append(
		// Required parameters.
		d.Set("name", utils.PathSearch("func_name", function, nil)),
		d.Set("runtime", utils.PathSearch("runtime", function, nil)),
		d.Set("timeout", utils.PathSearch("timeout", function, nil)),
		d.Set("memory_size", utils.PathSearch("memory_size", function, nil)),
		// Optional parameters but required in documentation.
		d.Set("app", utils.PathSearch("package", function, nil)),
		d.Set("handler", utils.PathSearch("handler", function, nil)),
		d.Set("code_type", utils.PathSearch("code_type", function, nil)),
		// Optional parameters.
		d.Set("description", utils.PathSearch("description", function, nil)),
		d.Set("functiongraph_version", utils.PathSearch("type", function, nil)),
		d.Set("code_url", utils.PathSearch("code_url", function, nil)),
		d.Set("code_filename", utils.PathSearch("code_filename", function, nil)),
		d.Set("user_data", utils.PathSearch("user_data", function, nil)),
		d.Set("encrypted_user_data", utils.PathSearch("encrypted_user_data", function, nil)),
		d.Set("agency", utils.PathSearch("xrole", function, nil)),
		d.Set("enterprise_project_id", utils.PathSearch("enterprise_project_id", function, nil)),
		d.Set("custom_image", flattenFgsCustomImage(utils.PathSearch("custom_image",
			function, make(map[string]interface{})).(map[string]interface{}))),
		d.Set("gpu_type", utils.PathSearch("gpu_type", function, nil)),
		d.Set("gpu_memory", utils.PathSearch("gpu_memory", function, nil)),
		d.Set("pre_stop_handler", utils.PathSearch("pre_stop_handler", function, nil)),
		d.Set("pre_stop_timeout", utils.PathSearch("pre_stop_timeout", function, nil)),
		d.Set("log_group_id", utils.PathSearch("log_group_id", function, nil)),
		d.Set("log_stream_id", utils.PathSearch("log_stream_id", function, nil)),
		d.Set("app_agency", utils.PathSearch("app_xrole", function, nil)),
		d.Set("depend_list", utils.PathSearch("depend_version_list", function, nil)),
		d.Set("initializer_handler", utils.PathSearch("initializer_handler", function, nil)),
		d.Set("initializer_timeout", utils.PathSearch("initializer_timeout", function, nil)),
		d.Set("max_instance_num", strconv.Itoa(int(utils.PathSearch("strategy_config.concurrency",
			function, float64(0)).(float64)))),
		d.Set("concurrency_num", int(utils.PathSearch("strategy_config.concurrent_num",
			function, float64(0)).(float64))),
		d.Set("dns_list", utils.PathSearch("domain_names", function, nil)),
		d.Set("vpc_id", utils.PathSearch("func_vpc.vpc_id", function, nil)),
		d.Set("network_id", utils.PathSearch("func_vpc.subnet_id", function, nil)),
		d.Set("mount_user_id", utils.PathSearch("mount_config.mount_user.user_id", function, nil)),
		d.Set("mount_user_group_id", utils.PathSearch("mount_config.mount_user.user_group_id", function, nil)),
		d.Set("func_mounts", flattenFuncionMounts(utils.PathSearch("mount_config.func_mounts",
			function, make([]interface{}, 0)).([]interface{}))),
		// Attributes.
		d.Set("urn", utils.PathSearch("func_urn", function, nil)),
		d.Set("version", utils.PathSearch("version", function, nil)),
	)

	versionConfig, err := flattenFunctionVersions(client, funcUrnWithoutVersion)
	if err != nil {
		// Not all regions support the version related API calls.
		log.Printf("[ERROR] Unable to parsing the function versions: %s", err)
	}
	mErr = multierror.Append(mErr, d.Set("versions", versionConfig))

	reservedInstances, err := getFunctionReservedInstances(client, funcUrn)
	if err != nil {
		return diag.Errorf("error retrieving function reserved instance: %s", err)
	}
	mErr = multierror.Append(mErr, d.Set("reserved_instances", flattenReservedInstances(reservedInstances)))

	if err := mErr.ErrorOrNil(); err != nil {
		return diag.Errorf("error setting function fields: %s", err)
	}

	return nil
}

func resourceFunctionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg                   = meta.(*config.Config)
		region                = cfg.GetRegion(d)
		funcUrnWithoutVersion = parseFunctionUrnWithoutVersion(d.Id())
	)

	client, err := cfg.NewServiceClient("fgs", region)
	if err != nil {
		return diag.Errorf("error creating FunctionGraph client: %s", err)
	}

	// lintignore:R019
	if d.HasChanges("code_type", "code_url", "code_filename", "depend_list", "func_code") {
		err := updateFunctionCode(client, d, funcUrnWithoutVersion)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	// lintignore:R019
	if d.HasChanges("app", "handler", "memory_size", "timeout", "encrypted_user_data",
		"user_data", "agency", "app_agency", "description", "initializer_handler", "initializer_timeout",
		"vpc_id", "network_id", "dns_list", "mount_user_id", "mount_user_group_id", "func_mounts", "custom_image",
		"log_group_id", "log_stream_id", "log_group_name", "log_stream_name", "concurrency_num", "gpu_memory", "gpu_type") {
		err := updateFunctionMetadata(client, d, funcUrnWithoutVersion)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("max_instance_num") {
		// If the maximum number of instances is omitted (after type conversion, the value is zero), means this feature
		// is disabled.
		maxInstanceNum, _ := strconv.Atoi(d.Get("max_instance_num").(string))
		err = updateFunctionMaxInstanceNum(client, funcUrnWithoutVersion, maxInstanceNum)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("tags") {
		if err = updateFunctionTags(client, d, funcUrnWithoutVersion); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("versions") {
		if err = updateFunctionVersions(client, d); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChanges("reserved_instances") {
		if err = updateFunctionReservedInstances(client, d, funcUrnWithoutVersion); err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceFunctionRead(ctx, d, meta)
}

func deleteFunctionTags(client *golangsdk.ServiceClient, functionUrn string, tags map[string]interface{}) error {
	if len(tags) < 1 {
		return nil
	}

	httpUrl := "v2/{project_id}/functions/{function_urn}/tags/delete"

	deletePath := client.Endpoint + httpUrl
	deletePath = strings.ReplaceAll(deletePath, "{project_id}", client.ProjectID)
	deletePath = strings.ReplaceAll(deletePath, "{function_urn}", functionUrn)
	deleteOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		OkCodes:  []int{204},
		JSONBody: buildFunctionTagsBodyParams(tags),
	}
	_, err := client.Request("DELETE", deletePath, &deleteOpt)
	if err != nil {
		return fmt.Errorf("failed to delete the function tags: %s", err)
	}
	return nil
}

func updateFunctionTags(client *golangsdk.ServiceClient, d *schema.ResourceData, functionUrn string) error {
	oldVal, newVal := d.GetChange("tags")
	oldTags := oldVal.(map[string]interface{})
	newTags := newVal.(map[string]interface{})

	if len(oldTags) > 0 {
		if err := deleteFunctionTags(client, functionUrn, oldTags); err != nil {
			return err
		}
	}

	if len(newTags) > 0 {
		if err := createFunctionTags(client, functionUrn, newTags); err != nil {
			return err
		}
	}
	return nil
}

// Does not allow to delete the latest version of the function (URN with the latest version).
// To delete the entire function (including all versions), the URN should be without any version number/alias,
// such as 'urn:fss:cn-north-4:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx:function:default:xxxxxx'
func deleteFunctionOrVersion(client *golangsdk.ServiceClient, functionUrn string) error {
	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}"

	deletePath := client.Endpoint + httpUrl
	deletePath = strings.ReplaceAll(deletePath, "{project_id}", client.ProjectID)
	deletePath = strings.ReplaceAll(deletePath, "{function_urn}", functionUrn)
	deleteOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}

	_, err := client.Request("DELETE", deletePath, &deleteOpt)
	if err != nil {
		return fmt.Errorf("error deleting function (version): %s", err)
	}
	return nil
}

func deleteFunctionVersionAlias(client *golangsdk.ServiceClient, functionUrn, aliasName string) error {
	httpUrl := "v2/{project_id}/fgs/functions/{function_urn}/aliases/{alias_name}"

	deletePath := client.Endpoint + httpUrl
	deletePath = strings.ReplaceAll(deletePath, "{project_id}", client.ProjectID)
	deletePath = strings.ReplaceAll(deletePath, "{function_urn}", functionUrn)
	deletePath = strings.ReplaceAll(deletePath, "{alias_name}", aliasName)
	deleteOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}

	_, err := client.Request("DELETE", deletePath, &deleteOpt)
	if err != nil {
		return fmt.Errorf("error deleting function version alias (%s): %s", aliasName, err)
	}
	return nil
}

func deleteFunctionVersions(client *golangsdk.ServiceClient, functionUrn string, versions []interface{}) error {
	// In the future, the function will support manage multiple versions.
	for _, version := range versions {
		versionNum := utils.PathSearch("name", version, "").(string) // The version name, also name as the version number.
		if versionNum != "latest" {
			// Deletes a function version, also deleting all aliases beneath it.
			err := deleteFunctionOrVersion(client, fmt.Sprintf("%s:%s", functionUrn, versionNum))
			if err != nil {
				return err
			}
			continue
		}
		// Since the latest version cannot be deleted, only the version alias under it can be deleted.
		aliases := utils.PathSearch("aliases", version, make([]interface{}, 0)).([]interface{})
		for _, alias := range aliases {
			err := deleteFunctionVersionAlias(client, functionUrn, utils.PathSearch("name", alias, "").(string))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func updateFunctionVersions(client *golangsdk.ServiceClient, d *schema.ResourceData) error {
	var (
		functionUrn = parseFunctionUrnWithoutVersion(d.Id())

		oldSet, newSet = d.GetChange("versions")
		decrease       = oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
		increase       = newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	)

	err := deleteFunctionVersions(client, functionUrn, decrease.List())
	if err != nil {
		return fmt.Errorf("error deleting function versions: %s", err)
	}

	err = createFunctionVersions(client, functionUrn, increase.List())
	if err != nil {
		return fmt.Errorf("error creating function versions: %s", err)
	}

	return nil
}

func resourceFunctionDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg                   = meta.(*config.Config)
		region                = cfg.GetRegion(d)
		funcUrnWithoutVersion = parseFunctionUrnWithoutVersion(d.Id())
	)

	client, err := cfg.NewServiceClient("fgs", region)
	if err != nil {
		return diag.Errorf("error creating FunctionGraph client: %s", err)
	}

	// Using function URN delete function without the version means delete the function and all its versions (aliases).
	err = deleteFunctionOrVersion(client, funcUrnWithoutVersion)
	if err != nil {
		return diag.Errorf("error deleting function (%s): %s", funcUrnWithoutVersion, err)
	}
	return nil
}

// The function ID consists of the function URN and the current version.
// Some requests require the URN information without the version, so this function is used to extract it.
func parseFunctionUrnWithoutVersion(urn string) string {
	// The format of the function URN is: 'urn:fss:cn-north-4:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx:function:default:xxxxxx'.
	index := strings.LastIndex(urn, ":")
	if index != -1 {
		urn = urn[0:index]
	}
	return urn
}
