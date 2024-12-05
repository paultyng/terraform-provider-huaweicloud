// ---------------------------------------------------------------
// *** AUTO GENERATED CODE ***
// @Product CPH
// ---------------------------------------------------------------

package cph

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API CPH PUT /v1/{project_id}/cloud-phone/servers/{server_id}
// @API CPH GET /v1/{project_id}/cloud-phone/servers/{server_id}
// @API CPH PUT /v1/{project_id}/cloud-phone/servers/open-access
// @API CPH POST /v2/{project_id}/cloud-phone/servers
// @API CPH POST /v2/{project_id}/cloud-phone/servers/{server_id}/change
// @API BSS POST /v2/orders/subscriptions/resources/unsubscribe
func ResourceCphServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCphServerCreate,
		UpdateContext: resourceCphServerUpdate,
		ReadContext:   resourceCphServerRead,
		DeleteContext: resourceCphServerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(90 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					return newValue+"-1" == oldValue
				},
				Description: `Server name.`,
			},
			"server_flavor": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The CPH server flavor.`,
			},
			"phone_flavor": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The cloud phone flavor.`,
			},
			"image_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The cloud phone image ID.`,
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The ID of VPC which the cloud server belongs to`,
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The ID of subnet which the cloud server belongs to`,
			},
			"eip_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"eip_type", "bandwidth"},
				Description:   `The ID of an **existing** EIP assigned to the server.`,
			},
			"eip_type": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"eip_id"},
				RequiredWith:  []string{"bandwidth"},
				Description:   `The type of an EIP that will be automatically assigned to the cloud server.`,
			},
			"bandwidth": {
				Type:          schema.TypeList,
				MaxItems:      1,
				ConflictsWith: []string{"eip_id"},
				RequiredWith:  []string{"eip_type"},
				Elem:          cphServerBandWidthSchema(),
				Optional:      true,
				Description:   `The bandwidth used by the cloud phone.`,
			},
			"period_unit": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The charging period unit.`,
				ValidateFunc: validation.StringInSlice([]string{
					"month", "year",
				}, false),
			},
			"period": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: `The charging period.`,
			},
			"auto_renew": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Whether auto renew is enabled. Valid values are "true" and "false".`,
			},
			"availability_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: `The name of the AZ where the cloud server is located.`,
			},
			"keypair_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The key pair name, which is used for logging in to the cloud phone through ADB.`,
			},
			"enterprise_project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `the enterprise project ID.`,
			},
			"ports": {
				Type:        schema.TypeList,
				Elem:        cphServerApplicationPortSchema(),
				Optional:    true,
				Computed:    true,
				Description: `The application port enabled by the cloud phone.`,
			},
			"phone_data_volume": {
				Type:        schema.TypeList,
				Elem:        cphPhoneDataVolumeSchema(),
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: `The phone data volume.`,
			},
			"server_share_data_volume": {
				Type:        schema.TypeList,
				Elem:        cphShareDataVolumeSchema(),
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: `The server share data volume.`,
			},
			"order_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The order ID.`,
			},
			"addresses": {
				Type:        schema.TypeList,
				Elem:        cphServerAddressSchema(),
				Computed:    true,
				Description: `The IP addresses of the CPH server.`,
			},
			"security_groups": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: `The list of the security groups bound to the extension NIC of the CPH server.`,
			},
			"status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The CPH server status.`,
			},
		},
	}
}

func cphServerBandWidthSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"share_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The bandwidth type.`,
				ValidateFunc: validation.StringInSlice([]string{
					"0", "1",
				}, false),
			},
			"id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				Description:   `The bandwidth ID.`,
				ConflictsWith: []string{"bandwidth.0.size", "bandwidth.0.charge_mode"},
			},
			"size": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				Description:  `The bandwidth (Mbit/s).`,
				RequiredWith: []string{"bandwidth.0.charge_mode"},
			},
			"charge_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Which the bandwidth used by the CPH server is billed.`,
				ValidateFunc: validation.StringInSlice([]string{
					"0", "1",
				}, false),
				RequiredWith: []string{"bandwidth.0.size"},
			},
		},
	}
	return &sc
}

func cphServerApplicationPortSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The application port name, which can contain a maximum of 16 bytes.`,
			},
			"listen_port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `The port number, which ranges from 10000 to 50000.`,
			},
			"internet_accessible": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Whether public network access is mapped.`,
			},
		},
	}
	return &sc
}

func cphServerAddressSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"server_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The internal IP address of the CPH server.`,
			},
			"public_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The public IP address of the CPH server.`,
			},
		},
	}
	return &sc
}

func cphPhoneDataVolumeSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"volume_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Specifies the volume type.`,
			},
			"volume_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: `Specifies the volume size, the unit is GB.`,
			},
			"volume_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The volume name.`,
			},
			"volume_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The volume ID.`,
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The creation time.`,
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The update time.`,
			},
		},
	}
	return &sc
}

func cphShareDataVolumeSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"volume_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Specifies the share volume type.`,
			},
			"size": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: `Specifies the share volume size, the unit is GB.`,
			},
			"version": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The share volume type.`,
			},
		},
	}
	return &sc
}

func resourceCphServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// createCphServer: create a CPH.
	var (
		createCphServerHttpUrl = "v2/{project_id}/cloud-phone/servers"
		createCphServerProduct = "cph"
	)
	createCphServerClient, err := cfg.NewServiceClient(createCphServerProduct, region)
	if err != nil {
		return diag.Errorf("error creating CPH Client: %s", err)
	}

	createCphServerPath := createCphServerClient.Endpoint + createCphServerHttpUrl
	createCphServerPath = strings.ReplaceAll(createCphServerPath, "{project_id}", createCphServerClient.ProjectID)

	createCphServerOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
	}
	createCphServerOpt.JSONBody = utils.RemoveNil(buildCreateCphServerBodyParams(d, cfg))
	createCphServerResp, err := createCphServerClient.Request("POST", createCphServerPath, &createCphServerOpt)
	if err != nil {
		return diag.Errorf("error creating CPH Server: %s", err)
	}

	createCphServerRespBody, err := utils.FlattenResponse(createCphServerResp)
	if err != nil {
		return diag.FromErr(err)
	}

	serverId := utils.PathSearch("server_ids[0]", createCphServerRespBody, "").(string)
	if serverId == "" {
		return diag.Errorf("unable to find the CPH Server ID from the API response")
	}

	d.SetId(serverId)

	err = createCphServerWaitingForStateCompleted(ctx, d, meta, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.Errorf("error waiting for the Create of CPH server (%s) to complete: %s", d.Id(), err)
	}

	orderId := utils.PathSearch("order_id", createCphServerRespBody, "").(string)
	if orderId == "" {
		return diag.Errorf("unable to find the order ID of the CPH Server from the API response")
	}

	bssClient, err := cfg.BssV2Client(region)
	if err != nil {
		return diag.Errorf("error creating BSS v2 client: %s", err)
	}
	err = common.WaitOrderComplete(ctx, bssClient, orderId, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.Errorf("error waiting for the Create of CPH server (%s) to complete: %s", d.Id(), err)
	}

	return resourceCphServerRead(ctx, d, meta)
}

func buildCreateCphServerBodyParams(d *schema.ResourceData, cfg *config.Config) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"server_name":       utils.ValueIgnoreEmpty(d.Get("name")),
		"server_model_name": utils.ValueIgnoreEmpty(d.Get("server_flavor")),
		"phone_model_name":  utils.ValueIgnoreEmpty(d.Get("phone_flavor")),
		"image_id":          utils.ValueIgnoreEmpty(d.Get("image_id")),
		"count":             1,
		"tenant_vpc_id":     utils.ValueIgnoreEmpty(d.Get("vpc_id")),
		"nics": []map[string]interface{}{
			{
				"subnet_id": utils.ValueIgnoreEmpty(d.Get("subnet_id")),
			},
		},
		"public_ip":                buildCreateCphServerRequestBodyPublicIp(d),
		"band_width":               buildCreateCphServerRequestBodyBandWidth(d.Get("bandwidth")),
		"keypair_name":             utils.ValueIgnoreEmpty(d.Get("keypair_name")),
		"availability_zone":        utils.ValueIgnoreEmpty(d.Get("availability_zone")),
		"ports":                    buildCreateCphServerRequestBodyApplicationPort(d.Get("ports")),
		"phone_data_volume":        buildCreateCphServerRequestBodyPhoneDataVolume(d.Get("phone_data_volume")),
		"server_share_data_volume": buildCreateCphServerRequestBodyShareDataVolume(d.Get("server_share_data_volume")),
	}

	extendParam := map[string]interface{}{
		"charging_mode":         0,
		"is_auto_pay":           1,
		"period_num":            utils.ValueIgnoreEmpty(d.Get("period")),
		"enterprise_project_id": utils.ValueIgnoreEmpty(cfg.GetEnterpriseProjectID(d)),
	}

	periodUnit := d.Get("period_unit").(string)
	if periodUnit == "month" {
		extendParam["period_type"] = 2
	} else {
		extendParam["period_type"] = 3
	}

	autoRenew := d.Get("auto_renew").(string)
	if autoRenew == "true" {
		extendParam["is_auto_renew"] = 1
	} else {
		extendParam["is_auto_renew"] = 0
	}

	bodyParams["extend_param"] = extendParam

	return bodyParams
}

func buildCreateCphServerRequestBodyBandWidth(rawParams interface{}) map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}
		raw := rawArray[0].(map[string]interface{})
		params := map[string]interface{}{
			"band_width_id":   utils.ValueIgnoreEmpty(raw["id"]),
			"band_width_size": utils.ValueIgnoreEmpty(raw["size"]),
		}

		shareType, _ := strconv.Atoi(utils.ValueIgnoreEmpty(raw["share_type"]).(string))
		params["band_width_share_type"] = shareType

		chargeMode := utils.ValueIgnoreEmpty(raw["charge_mode"]).(string)
		if len(chargeMode) > 0 {
			chargeModeInteger, _ := strconv.Atoi(chargeMode)
			params["band_width_charge_mode"] = chargeModeInteger
		}

		return params
	}
	return nil
}

func buildCreateCphServerRequestBodyPublicIp(d *schema.ResourceData) map[string]interface{} {
	params := make(map[string]interface{})
	if value, ok := d.GetOk("eip_type"); ok {
		params["eip"] = map[string]interface{}{
			"type": utils.ValueIgnoreEmpty(value),
		}
	}
	if value, ok := d.GetOk("eip_id"); ok {
		params["ids"] = []string{value.(string)}
	}
	return params
}

func buildCreateCphServerRequestBodyApplicationPort(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"name":                utils.ValueIgnoreEmpty(raw["name"]),
				"listen_port":         utils.ValueIgnoreEmpty(raw["listen_port"]),
				"internet_accessible": utils.ValueIgnoreEmpty(raw["internet_accessible"]),
			}
		}
		return rst
	}
	return nil
}

func buildCreateCphServerRequestBodyPhoneDataVolume(rawParams interface{}) map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		raw := rawArray[0].(map[string]interface{})
		params := map[string]interface{}{
			"volume_type": utils.ValueIgnoreEmpty(raw["volume_type"]),
			"size":        utils.ValueIgnoreEmpty(raw["volume_size"]),
		}
		return params
	}

	return nil
}

func buildCreateCphServerRequestBodyShareDataVolume(rawParams interface{}) map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		raw := rawArray[0].(map[string]interface{})
		params := map[string]interface{}{
			"volume_type": utils.ValueIgnoreEmpty(raw["volume_type"]),
			"size":        utils.ValueIgnoreEmpty(raw["size"]),
		}
		return params
	}

	return nil
}

func createCphServerWaitingForStateCompleted(ctx context.Context, d *schema.ResourceData, meta interface{}, t time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"PENDING"},
		Target:  []string{"COMPLETED"},
		Refresh: func() (interface{}, string, error) {
			cfg := meta.(*config.Config)
			region := cfg.GetRegion(d)
			var (
				createCphServerHttpUrl = "v1/{project_id}/cloud-phone/servers/{server_id}"
				createCphServerProduct = "cph"
			)
			createCphServerClient, err := cfg.NewServiceClient(createCphServerProduct, region)
			if err != nil {
				return nil, "ERROR", fmt.Errorf("error creating CPH Client: %s", err)
			}

			createCphServerPath := createCphServerClient.Endpoint + createCphServerHttpUrl
			createCphServerPath = strings.ReplaceAll(createCphServerPath, "{project_id}", createCphServerClient.ProjectID)
			createCphServerPath = strings.ReplaceAll(createCphServerPath, "{server_id}", d.Id())

			createCphServerOpt := golangsdk.RequestOpts{
				KeepResponseBody: true,
				OkCodes: []int{
					200,
				},
			}
			createCphServerResp, err := createCphServerClient.Request("GET", createCphServerPath, &createCphServerOpt)
			if err != nil {
				return nil, "ERROR", err
			}

			createCphServerRespBody, err := utils.FlattenResponse(createCphServerResp)
			if err != nil {
				return nil, "ERROR", err
			}
			statusRaw := utils.PathSearch(`status`, createCphServerRespBody, nil)
			if statusRaw == nil {
				return nil, "ERROR", fmt.Errorf("unable to find the status from the API response")
			}
			status := fmt.Sprint(statusRaw)

			targetStatus := []string{
				"5", "8", "10",
			}
			if utils.StrSliceContains(targetStatus, status) {
				return createCphServerRespBody, "COMPLETED", nil
			}

			unexpectedStatus := []string{
				"2",
			}
			if utils.StrSliceContains(unexpectedStatus, status) {
				return createCphServerRespBody, status, nil
			}

			return createCphServerRespBody, "PENDING", nil
		},
		Timeout:      t,
		Delay:        120 * time.Second,
		PollInterval: 10 * time.Second,
	}
	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

func resourceCphServerRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	var mErr *multierror.Error

	// getCphServer: Query the CPH instance
	var (
		getCphServerHttpUrl = "v1/{project_id}/cloud-phone/servers/{server_id}"
		getCphServerProduct = "cph"
	)
	getCphServerClient, err := cfg.NewServiceClient(getCphServerProduct, region)
	if err != nil {
		return diag.Errorf("error creating CPH client: %s", err)
	}

	getCphServerPath := getCphServerClient.Endpoint + getCphServerHttpUrl
	getCphServerPath = strings.ReplaceAll(getCphServerPath, "{project_id}", getCphServerClient.ProjectID)
	getCphServerPath = strings.ReplaceAll(getCphServerPath, "{server_id}", d.Id())

	getCphServerOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}
	getCphServerResp, err := getCphServerClient.Request("GET", getCphServerPath, &getCphServerOpt)

	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving CPH server")
	}

	getCphServerRespBody, err := utils.FlattenResponse(getCphServerResp)
	if err != nil {
		return diag.FromErr(err)
	}
	statusRaw := utils.PathSearch("status", getCphServerRespBody, nil)

	if fmt.Sprint(statusRaw) == "6" {
		return common.CheckDeletedDiag(d, golangsdk.ErrDefault404{}, "error retrieving CPH server")
	}

	mErr = multierror.Append(
		mErr,
		d.Set("region", region),
		d.Set("name", utils.PathSearch("server_name", getCphServerRespBody, nil)),
		d.Set("server_flavor", utils.PathSearch("server_model_name", getCphServerRespBody, nil)),
		d.Set("phone_flavor", utils.PathSearch("phone_model_name", getCphServerRespBody, nil)),
		d.Set("keypair_name", utils.PathSearch("keypair_name", getCphServerRespBody, nil)),
		d.Set("vpc_id", utils.PathSearch("vpc_id", getCphServerRespBody, nil)),
		d.Set("subnet_id", utils.PathSearch("subnet_id", getCphServerRespBody, nil)),
		d.Set("order_id", utils.PathSearch("metadata.order_id", getCphServerRespBody, nil)),
		d.Set("addresses", flattenGetCphServerResponseBodyAddress(getCphServerRespBody)),
		d.Set("bandwidth", flattenGetCphServerResponseBodyBandWidth(getCphServerRespBody)),
		d.Set("availability_zone", utils.PathSearch("availability_zone", getCphServerRespBody, nil)),
		d.Set("enterprise_project_id", utils.PathSearch("enterprise_project_id", getCphServerRespBody, nil)),
		d.Set("status", utils.PathSearch("status", getCphServerRespBody, nil)),
		d.Set("security_groups", utils.PathSearch("security_groups", getCphServerRespBody, nil)),
		d.Set("phone_data_volume", flattenPhoneDataVolume(getCphServerRespBody)),
		d.Set("server_share_data_volume", flattenShareDataVolume(getCphServerRespBody)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func flattenPhoneDataVolume(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("volumes", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"volume_type": utils.PathSearch("volume_type", v, nil),
			"volume_size": utils.PathSearch("volume_size", v, nil),
			"volume_name": utils.PathSearch("volume_name", v, nil),
			"volume_id":   utils.PathSearch("volume_id", v, nil),
			"created_at":  utils.PathSearch("create_time", v, nil),
			"updated_at":  utils.PathSearch("update_time", v, nil),
		})
	}
	return rst
}

func flattenShareDataVolume(resp interface{}) []map[string]interface{} {
	shareVolume := utils.PathSearch("share_volume_info", resp, nil)
	if shareVolume == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"volume_type": utils.PathSearch("volume_type", shareVolume, nil),
			"size":        utils.PathSearch("size", shareVolume, nil),
			"version":     utils.PathSearch("version", shareVolume, nil),
		},
	}
}

func flattenGetCphServerResponseBodyAddress(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("addresses", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"server_ip": utils.PathSearch("server_ip", v, nil),
			"public_ip": utils.PathSearch("public_ip", v, nil),
		})
	}
	return rst
}

func flattenGetCphServerResponseBodyBandWidth(resp interface{}) []interface{} {
	var rst []interface{}
	curJson := utils.PathSearch("band_widths[0]", resp, make(map[string]interface{})).(map[string]interface{})
	if len(curJson) < 1 {
		return rst
	}

	rst = []interface{}{
		map[string]interface{}{
			"share_type":  fmt.Sprintf("%v", utils.PathSearch("band_width_share_type", curJson, nil)),
			"id":          utils.PathSearch("band_width_id", curJson, nil),
			"size":        utils.PathSearch("band_width_size", curJson, nil),
			"charge_mode": fmt.Sprintf("%v", utils.PathSearch("band_width_charge_mode", curJson, nil)),
		},
	}
	return rst
}

func resourceCphServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	client, err := cfg.NewServiceClient("cph", region)
	if err != nil {
		return diag.Errorf("error creating CPH client: %s", err)
	}

	if d.HasChange("name") {
		err := updateServerName(client, d)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("keypair_name") {
		err := updateKeypair(ctx, client, d)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	updateServerSwitchChanges := []string{
		"phone_flavor",
		"keypair_name",
		"ports",
		"tenant_vpc_id",
		"nics",
		"public_ip",
		"band_width",
		"image_id",
		"extend_param",
		"phone_data_volume",
		"server_share_data_volume",
	}
	if d.HasChanges(updateServerSwitchChanges...) {
		err := updateServerChange(ctx, client, d, cfg)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceCphServerRead(ctx, d, meta)
}

func resourceCphServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	if err := common.UnsubscribePrePaidResource(d, cfg, []string{d.Id()}); err != nil {
		return diag.Errorf("error unsubscribing CPH server: %s", err)
	}

	err := deleteCphServerWaitingForStateCompleted(ctx, d, meta, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return diag.Errorf("error waiting for the delete of CPH server (%s) to complete: %s", d.Id(), err)
	}
	return nil
}

func deleteCphServerWaitingForStateCompleted(ctx context.Context, d *schema.ResourceData, meta interface{}, t time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"PENDING"},
		Target:  []string{"COMPLETED"},
		Refresh: func() (interface{}, string, error) {
			cfg := meta.(*config.Config)
			region := cfg.GetRegion(d)
			var (
				getCphServerHttpUrl = "v1/{project_id}/cloud-phone/servers/{server_id}"
				getCphServerProduct = "cph"
			)
			getCphServerClient, err := cfg.NewServiceClient(getCphServerProduct, region)
			if err != nil {
				return nil, "ERROR", fmt.Errorf("error creating CPH client: %s", err)
			}

			getCphServerPath := getCphServerClient.Endpoint + getCphServerHttpUrl
			getCphServerPath = strings.ReplaceAll(getCphServerPath, "{project_id}", getCphServerClient.ProjectID)
			getCphServerPath = strings.ReplaceAll(getCphServerPath, "{server_id}", d.Id())

			getCphServerOpt := golangsdk.RequestOpts{
				KeepResponseBody: true,
			}
			getCphServerResp, err := getCphServerClient.Request("GET", getCphServerPath, &getCphServerOpt)

			if err != nil {
				if _, ok := err.(golangsdk.ErrDefault404); ok {
					return getCphServerResp, "COMPLETED", nil
				}

				return nil, "ERROR", err
			}

			getCphServerRespBody, err := utils.FlattenResponse(getCphServerResp)
			if err != nil {
				return nil, "ERROR", err
			}

			statusRaw := utils.PathSearch(`status`, getCphServerRespBody, nil)
			if statusRaw == nil {
				return nil, "ERROR", fmt.Errorf("unable to find the status from the API response")
			}
			status := fmt.Sprint(statusRaw)

			targetStatus := []string{
				"6",
			}
			if utils.StrSliceContains(targetStatus, status) {
				return getCphServerRespBody, "COMPLETED", nil
			}

			return getCphServerResp, "PENDING", nil
		},
		Timeout:      t,
		Delay:        10 * time.Second,
		PollInterval: 10 * time.Second,
	}
	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

func updateKeypair(ctx context.Context, client *golangsdk.ServiceClient, d *schema.ResourceData) error {
	// updateKeypair: update CPH server keypair
	updateKeypairHttpUrl := "v1/{project_id}/cloud-phone/servers/open-access"

	updateKeypairPath := client.Endpoint + updateKeypairHttpUrl
	updateKeypairPath = strings.ReplaceAll(updateKeypairPath, "{project_id}", client.ProjectID)

	updateKeypairOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}
	updateKeypairOpt.JSONBody = map[string]interface{}{
		"servers": []map[string]interface{}{
			{
				"server_id":    d.Id(),
				"keypair_name": d.Get("keypair_name").(string),
			},
		},
	}
	updateServerKeypairResp, err := client.Request("PUT", updateKeypairPath, &updateKeypairOpt)
	if err != nil {
		return fmt.Errorf("error updating CPH server keypair: %s", err)
	}

	updateServerKeypairRespBody, err := utils.FlattenResponse(updateServerKeypairResp)
	if err != nil {
		return fmt.Errorf("error flattening CPH server keypair: %s", err)
	}

	jobId := utils.PathSearch("jobs|[0].job_id", updateServerKeypairRespBody, "").(string)
	if jobId == "" {
		return fmt.Errorf("unable to find the job ID from the API response")
	}

	err = checkCphJobStatus(ctx, client, jobId, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf("error waiting for updating CPH server keypair to be completed: %s", err)
	}

	return nil
}

func updateServerName(client *golangsdk.ServiceClient, d *schema.ResourceData) error {
	// updateCphServerName: update CPH server name
	updateCphServerNameHttpUrl := "v1/{project_id}/cloud-phone/servers/{server_id}"

	updateCphServerNamePath := client.Endpoint + updateCphServerNameHttpUrl
	updateCphServerNamePath = strings.ReplaceAll(updateCphServerNamePath, "{project_id}", client.ProjectID)
	updateCphServerNamePath = strings.ReplaceAll(updateCphServerNamePath, "{server_id}", d.Id())

	updateCphServerNameOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}
	updateCphServerNameOpt.JSONBody = utils.RemoveNil(map[string]interface{}{
		"server_name": utils.ValueIgnoreEmpty(d.Get("name")),
	})

	_, err := client.Request("PUT", updateCphServerNamePath, &updateCphServerNameOpt)
	if err != nil {
		return fmt.Errorf("error updating CPH server: %s", err)
	}

	return nil
}

func checkCphJobStatus(ctx context.Context, client *golangsdk.ServiceClient, id string, timeout time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      jobStatusRefreshFunc(client, id),
		Timeout:      timeout,
		PollInterval: 10 * timeout,
		Delay:        10 * time.Second,
	}
	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func updateServerChange(ctx context.Context, client *golangsdk.ServiceClient, d *schema.ResourceData, cfg *config.Config) error {
	// updateCphServerName: update CPH server change
	updateServerChangeHttpUrl := "v2/{project_id}/cloud-phone/servers/{server_id}/change"
	updateServerChangePath := client.Endpoint + updateServerChangeHttpUrl
	updateServerChangePath = strings.ReplaceAll(updateServerChangePath, "{project_id}", client.ProjectID)
	updateServerChangePath = strings.ReplaceAll(updateServerChangePath, "{server_id}", d.Id())

	updateCphServerChangeOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}
	updateCphServerChangeOpt.JSONBody = utils.RemoveNil(buildUpdateCphServerChangeBodyParams(d, cfg))
	_, err := client.Request("POST", updateServerChangePath, &updateCphServerChangeOpt)
	if err != nil {
		return fmt.Errorf("error updating CPH server change: %s", err)
	}

	err = checkServerSwitchStatus(ctx, client, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	return nil
}

func buildUpdateCphServerChangeBodyParams(d *schema.ResourceData, cfg *config.Config) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"phone_model_name": utils.ValueIgnoreEmpty(d.Get("phone_flavor")),
		"keypair_name":     utils.ValueIgnoreEmpty(d.Get("keypair_name")),
		"ports":            buildCreateCphServerRequestBodyApplicationPort(d.Get("ports")),
		"tenant_vpc_id":    utils.ValueIgnoreEmpty(d.Get("vpc_id")),
		"nics": []map[string]interface{}{
			{
				"subnet_id": utils.ValueIgnoreEmpty(d.Get("subnet_id")),
			},
		},
		"public_ip":  buildCreateCphServerRequestBodyPublicIp(d),
		"band_width": buildUpdateServerRequestBodyBandWidth(d.Get("bandwidth")),
		"image_id":   utils.ValueIgnoreEmpty(d.Get("image_id")),
		"extend_param": map[string]interface{}{
			"enterprise_project_id": utils.ValueIgnoreEmpty(cfg.GetEnterpriseProjectID(d)),
		},
		"phone_data_volume":        buildCreateCphServerRequestBodyPhoneDataVolume(d.Get("phone_data_volume")),
		"server_share_data_volume": buildCreateCphServerRequestBodyShareDataVolume(d.Get("server_share_data_volume")),
		"image_type":               "gold",
	}

	return bodyParams
}

func buildUpdateServerRequestBodyBandWidth(rawParams interface{}) map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}
		raw := rawArray[0].(map[string]interface{})

		shareType, _ := strconv.Atoi(utils.ValueIgnoreEmpty(raw["share_type"]).(string))
		chargeModeInteger, _ := strconv.Atoi(utils.ValueIgnoreEmpty(raw["charge_mode"]).(string))

		params := map[string]interface{}{
			"band_width_charge_mode": chargeModeInteger,
			"band_width_share_type":  shareType,
		}

		if chargeModeInteger == 0 {
			params["band_width_id"] = utils.ValueIgnoreEmpty(raw["id"])
		}

		if chargeModeInteger == 1 {
			params["band_width_size"] = utils.ValueIgnoreEmpty(raw["size"])
		}

		return params
	}

	return nil
}

func checkServerSwitchStatus(ctx context.Context, client *golangsdk.ServiceClient, id string, timeout time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      serverStateRefreshFunc(client, id),
		Timeout:      timeout,
		PollInterval: 10 * timeout,
		Delay:        10 * time.Second,
	}
	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return fmt.Errorf("error waiting for CPH server switch to be completed: %s", err)
	}
	return nil
}
