package cce

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API CCE POST /v2/charts
// @API CCE GET /v2/charts/{chart_id}
// @API CCE PUT /v2/charts/{chart_id}
// @API CCE DELETE /v2/charts/{chart_id}

func ResourceChart() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceChartCreate,
		ReadContext:   resourceChartRead,
		UpdateContext: resourceChartUpdate,
		DeleteContext: resourceChartDelete,

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
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parameters": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"translate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instruction": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"chart_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceChartCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	file, err := os.Open(d.Get("content").(string))
	if err != nil {
		return diag.Errorf("error opening chart file: %s", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("content", file.Name())
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return diag.FromErr(err)
	}

	if v, ok := d.GetOk("parameters"); ok {
		err = writer.WriteField("parameters", v.(string))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	writer.Close()

	var (
		createChartHttpUrl = "v2/charts"
		createChartProduct = "cce"
	)
	createChartClient, err := cfg.NewServiceClient(createChartProduct, region)
	if err != nil {
		return diag.Errorf("error creating CCE client: %s", err)
	}

	createChartHttpPath := createChartClient.Endpoint + createChartHttpUrl

	createChartOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": writer.FormDataContentType(),
		},
		RawBody: body,
	}

	createChartResp, err := createChartClient.Request("POST", createChartHttpPath, &createChartOpt)
	if err != nil {
		return diag.Errorf("error uploading chart: %s", err)
	}

	createChartRespBody, err := utils.FlattenResponse(createChartResp)
	if err != nil {
		return diag.FromErr(err)
	}

	id := utils.PathSearch("id", createChartRespBody, nil)
	if id == nil {
		return diag.Errorf("error uploading chart: ID is not found in API response")
	}

	d.SetId(id.(string))

	return resourceChartRead(ctx, d, meta)
}

func resourceChartRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	var (
		getChartHttpUrl = "v2/charts/{chart_id}"
		getChartProduct = "cce"
	)
	getChartClient, err := cfg.NewServiceClient(getChartProduct, region)
	if err != nil {
		return diag.Errorf("error creating CCE client: %s", err)
	}

	getChartHttpPath := getChartClient.Endpoint + getChartHttpUrl
	getChartHttpPath = strings.ReplaceAll(getChartHttpPath, "{chart_id}", d.Id())

	getChartOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	getChartResp, err := getChartClient.Request("GET", getChartHttpPath, &getChartOpt)
	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving CCE chart")
	}

	getChartRespBody, err := utils.FlattenResponse(getChartResp)
	if err != nil {
		return diag.FromErr(err)
	}

	mErr := multierror.Append(nil,
		d.Set("region", cfg.GetRegion(d)),
		d.Set("name", utils.PathSearch("name", getChartRespBody, nil)),
		d.Set("value", utils.PathSearch("values", getChartRespBody, nil)),
		d.Set("translate", utils.PathSearch("translate", getChartRespBody, nil)),
		d.Set("instruction", utils.PathSearch("instruction", getChartRespBody, nil)),
		d.Set("version", utils.PathSearch("version", getChartRespBody, nil)),
		d.Set("description", utils.PathSearch("description", getChartRespBody, nil)),
		d.Set("source", utils.PathSearch("source", getChartRespBody, nil)),
		d.Set("public", utils.PathSearch("public", getChartRespBody, nil)),
		d.Set("chart_url", utils.PathSearch("chart_url", getChartRespBody, nil)),
		d.Set("created_at", utils.PathSearch("create_at", getChartRespBody, nil)),
		d.Set("updated_at", utils.PathSearch("update_at", getChartRespBody, nil)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func resourceChartUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	file, err := os.Open(d.Get("content").(string))
	if err != nil {
		return diag.Errorf("error opening chart file: %s", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("content", file.Name())
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return diag.FromErr(err)
	}

	if v, ok := d.GetOk("parameters"); ok {
		err = writer.WriteField("parameters", v.(string))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	writer.Close()

	var (
		updateChartHttpUrl = "v2/charts/{chart_id}"
		updateChartProduct = "cce"
	)
	updateChartClient, err := cfg.NewServiceClient(updateChartProduct, region)
	if err != nil {
		return diag.Errorf("error creating CCE client: %s", err)
	}

	updateChartHttpPath := updateChartClient.Endpoint + updateChartHttpUrl
	updateChartHttpPath = strings.ReplaceAll(updateChartHttpPath, "{chart_id}", d.Id())

	updateChartOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": writer.FormDataContentType(),
		},
		RawBody: body,
	}

	_, err = updateChartClient.Request("PUT", updateChartHttpPath, &updateChartOpt)
	if err != nil {
		return diag.Errorf("error updating CCE chart: %s", err)
	}

	return resourceChartRead(ctx, d, meta)
}

func resourceChartDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	var (
		deleteChartHttpUrl = "v2/charts/{chart_id}"
		deleteChartProduct = "cce"
	)
	deleteChartClient, err := cfg.NewServiceClient(deleteChartProduct, region)
	if err != nil {
		return diag.Errorf("error creating CCE client: %s", err)
	}

	deleteChartHttpPath := deleteChartClient.Endpoint + deleteChartHttpUrl
	deleteChartHttpPath = strings.ReplaceAll(deleteChartHttpPath, "{chart_id}", d.Id())

	deleteChartOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	_, err = deleteChartClient.Request("DELETE", deleteChartHttpPath, &deleteChartOpt)
	if err != nil {
		return common.CheckDeletedDiag(d, err, "error deleting CCE chart")
	}

	return nil
}
