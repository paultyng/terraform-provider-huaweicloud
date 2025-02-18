// ---------------------------------------------------------------
// *** AUTO GENERATED CODE ***
// @Product MRS
// ---------------------------------------------------------------

package mrs

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func DataSourceMrsFlavors() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceMrsFlavorsRead,
		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The version of cluster.`,
			},
			"availability_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The AZ name.`,
			},
			"node_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The node type supported by this flavor.`,
			},
			"flavors": {
				Type:        schema.TypeList,
				Elem:        mrsFlavorsFlavorSchema(),
				Computed:    true,
				Description: `List of available cluster flavors.`,
			},
		},
	}
}

func mrsFlavorsFlavorSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"flavor_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The flavor ID.`,
			},
			"version_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The version of cluster.`,
			},
			"availability_zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The availability zone.`,
			},
			"node_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The node type supported by this flavor.`,
			},
		},
	}
	return &sc
}

func resourceMrsFlavorsRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	var mErr *multierror.Error

	var (
		getFlavorsHttpUrl = "v2/{project_id}/metadata/version/{version_name}/available-flavor"
		getFlavorsProduct = "mrs"
	)
	getFlavorsClient, err := cfg.NewServiceClient(getFlavorsProduct, region)
	if err != nil {
		return diag.Errorf("error creating MRS client: %s", err)
	}

	getFlavorsPath := getFlavorsClient.Endpoint + getFlavorsHttpUrl
	getFlavorsPath = strings.ReplaceAll(getFlavorsPath, "{project_id}", getFlavorsClient.ProjectID)
	getFlavorsPath = strings.ReplaceAll(getFlavorsPath, "{version_name}", d.Get("version_name").(string))

	getFlavorsOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	}

	getFlavorsResp, err := getFlavorsClient.Request("GET", getFlavorsPath, &getFlavorsOpt)

	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving flavors")
	}

	getFlavorsRespBody, err := utils.FlattenResponse(getFlavorsResp)
	if err != nil {
		return diag.FromErr(err)
	}

	uuid, err := uuid.GenerateUUID()
	if err != nil {
		return diag.Errorf("unable to generate ID: %s", err)
	}
	d.SetId(uuid)

	mErr = multierror.Append(
		mErr,
		d.Set("region", region),
		d.Set("flavors", flattenListFlavorsBody(getFlavorsRespBody, d)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func flattenListFlavorsBody(resp interface{}, d *schema.ResourceData) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("available_flavors", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0)
	for _, v := range curArray {
		az := utils.PathSearch("az_code", v, nil)
		if param, ok := d.GetOk("availability_zone"); ok && fmt.Sprint(param) != fmt.Sprint(az) {
			continue
		}

		masterFlavors := utils.PathSearch("master[*].flavor_name", v, make([]interface{}, 0))
		coreFlavors := utils.PathSearch("core[*].flavor_name", v, make([]interface{}, 0))
		taskFlavors := utils.PathSearch("task[*].flavor_name", v, make([]interface{}, 0))

		rst = append(rst, flattenFlavors(masterFlavors.([]interface{}), az.(string), "master", d)...)
		rst = append(rst, flattenFlavors(coreFlavors.([]interface{}), az.(string), "core", d)...)
		rst = append(rst, flattenFlavors(taskFlavors.([]interface{}), az.(string), "task", d)...)
	}
	return rst
}

func flattenFlavors(all []interface{}, az, nodeType string, d *schema.ResourceData) []interface{} {
	rst := make([]interface{}, 0, len(all))
	if param, ok := d.GetOk("node_type"); ok && fmt.Sprint(param) != fmt.Sprint(nodeType) {
		return rst
	}
	versionName := d.Get("version_name")
	for _, v := range all {
		rst = append(rst, map[string]interface{}{
			"availability_zone": az,
			"node_type":         nodeType,
			"flavor_id":         v,
			"version_name":      versionName,
		})
	}
	return rst
}
