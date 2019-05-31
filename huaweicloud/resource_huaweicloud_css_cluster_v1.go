// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://www.github.com/huaweicloud/magic-modules
//
// ----------------------------------------------------------------------------

package huaweicloud

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/huaweicloud/golangsdk"
)

func resourceCssClusterV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceCssClusterV1Create,
		Read:   resourceCssClusterV1Read,
		Update: resourceCssClusterV1Update,
		Delete: resourceCssClusterV1Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"engine_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"node_config": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"flavor": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"network_info": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"security_group_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"vpc_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
								},
							},
						},
						"volume": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"size": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
									"volume_type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
								},
							},
						},
						"availability_zone": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},

			"engine_type": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},

			"expect_node_num": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},

			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceCssClusterV1UserInputParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"terraform_resource_data": d,
		"engine_type":             d.Get("engine_type"),
		"engine_version":          d.Get("engine_version"),
		"expect_node_num":         d.Get("expect_node_num"),
		"name":                    d.Get("name"),
		"node_config":             d.Get("node_config"),
	}
}

func resourceCssClusterV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "css", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	opts := resourceCssClusterV1UserInputParams(d)

	arrayIndex := map[string]int{
		"node_config.network_info": 0,
		"node_config.volume":       0,
		"node_config":              0,
	}

	params, err := buildCssClusterV1CreateParameters(opts, arrayIndex)
	if err != nil {
		return fmt.Errorf("Error building the request body of api(create), err=%s", err)
	}
	r, err := sendCssClusterV1CreateRequest(d, params, client)
	if err != nil {
		return fmt.Errorf("Error creating CssClusterV1, err=%s", err)
	}

	timeout := d.Timeout(schema.TimeoutCreate)
	obj, err := asyncWaitCssClusterV1Create(d, config, r, client, timeout)
	if err != nil {
		return err
	}
	id, err := navigateValue(obj, []string{"id"}, nil)
	if err != nil {
		return fmt.Errorf("Error constructing id, err=%s", err)
	}
	d.SetId(id.(string))

	return resourceCssClusterV1Read(d, meta)
}

func resourceCssClusterV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "css", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	res := make(map[string]interface{})

	v, err := sendCssClusterV1ReadRequest(d, client)
	if err != nil {
		return err
	}
	res["read"] = fillCssClusterV1ReadRespBody(v)

	return setCssClusterV1Properties(d, res)
}

func resourceCssClusterV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "css", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	opts := resourceCssClusterV1UserInputParams(d)

	arrayIndex := map[string]int{
		"node_config.network_info": 0,
		"node_config.volume":       0,
		"node_config":              0,
	}
	timeout := d.Timeout(schema.TimeoutUpdate)

	params, err := buildCssClusterV1ExtendClusterParameters(opts, arrayIndex)
	if err != nil {
		return fmt.Errorf("Error building the request body of api(extend_cluster), err=%s", err)
	}
	if e, _ := isEmptyValue(reflect.ValueOf(params)); !e {
		r, err := sendCssClusterV1ExtendClusterRequest(d, params, client)
		if err != nil {
			return err
		}
		_, err = asyncWaitCssClusterV1ExtendCluster(d, config, r, client, timeout)
		if err != nil {
			return err
		}
	}

	return resourceCssClusterV1Read(d, meta)
}

func resourceCssClusterV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "css", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	url, err := replaceVars(d, "clusters/{id}", nil)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	log.Printf("[DEBUG] Deleting Cluster %q", d.Id())
	r := golangsdk.Result{}
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		OkCodes:      successHTTPCodes,
		JSONBody:     nil,
		JSONResponse: nil,
		MoreHeaders:  map[string]string{"Content-Type": "application/json"},
	})
	if r.Err != nil {
		return fmt.Errorf("Error deleting Cluster %q, err=%s", d.Id(), r.Err)
	}

	_, err = waitToFinish(
		[]string{"Done"}, []string{"Pending"},
		d.Timeout(schema.TimeoutCreate),
		1*time.Second,
		func() (interface{}, string, error) {
			_, err := client.Get(url, nil, &golangsdk.RequestOpts{
				MoreHeaders: map[string]string{"Content-Type": "application/json"}})
			if err != nil {
				if _, ok := err.(golangsdk.ErrDefault404); ok {
					return true, "Done", nil
				}
				return nil, "", nil
			}
			return true, "Pending", nil
		},
	)
	return err
}

func buildCssClusterV1CreateParameters(opts map[string]interface{}, arrayIndex map[string]int) (interface{}, error) {
	params := make(map[string]interface{})

	v, err := expandCssClusterV1CreateDatastore(opts, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["datastore"] = v
	}

	v, err = expandCssClusterV1CreateInstance(opts, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["instance"] = v
	}

	v, err = navigateValue(opts, []string{"expect_node_num"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["instanceNum"] = v
	}

	v, err = navigateValue(opts, []string{"name"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["name"] = v
	}

	if len(params) == 0 {
		return params, nil
	}

	params = map[string]interface{}{"cluster": params}

	return params, nil
}

func expandCssClusterV1CreateDatastore(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	req := make(map[string]interface{})

	v, err := navigateValue(d, []string{"engine_type"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		req["type"] = v
	} else {
		req["type"] = "elasticsearch"
	}

	v, err = navigateValue(d, []string{"engine_version"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		req["version"] = v
	}

	return req, nil
}

func expandCssClusterV1CreateInstance(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	req := make(map[string]interface{})

	v, err := navigateValue(d, []string{"node_config", "availability_zone"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		req["availability_zone"] = v
	}

	v, err = navigateValue(d, []string{"node_config", "flavor"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		req["flavorRef"] = v
	}

	v, err = expandCssClusterV1CreateInstanceNics(d, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		req["nics"] = v
	}

	v, err = expandCssClusterV1CreateInstanceVolume(d, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		req["volume"] = v
	}

	return req, nil
}

func expandCssClusterV1CreateInstanceNics(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	req := make(map[string]interface{})

	v, err := navigateValue(d, []string{"node_config", "network_info", "subnet_id"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		req["netId"] = v
	}

	v, err = navigateValue(d, []string{"node_config", "network_info", "security_group_id"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		req["securityGroupId"] = v
	}

	v, err = navigateValue(d, []string{"node_config", "network_info", "vpc_id"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		req["vpcId"] = v
	}

	return req, nil
}

func expandCssClusterV1CreateInstanceVolume(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	req := make(map[string]interface{})

	v, err := navigateValue(d, []string{"node_config", "volume", "size"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		req["size"] = v
	}

	v, err = navigateValue(d, []string{"node_config", "volume", "volume_type"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		req["volume_type"] = v
	}

	return req, nil
}

func sendCssClusterV1CreateRequest(d *schema.ResourceData, params interface{},
	client *golangsdk.ServiceClient) (interface{}, error) {
	url := client.ServiceURL("clusters")

	r := golangsdk.Result{}
	_, r.Err = client.Post(url, params, &r.Body, &golangsdk.RequestOpts{
		OkCodes: successHTTPCodes,
	})
	if r.Err != nil {
		return nil, fmt.Errorf("Error running api(create), err=%s", r.Err)
	}
	return r.Body, nil
}

func asyncWaitCssClusterV1Create(d *schema.ResourceData, config *Config, result interface{},
	client *golangsdk.ServiceClient, timeout time.Duration) (interface{}, error) {

	data := make(map[string]string)
	pathParameters := map[string][]string{
		"id": []string{"cluster", "id"},
	}
	for key, path := range pathParameters {
		value, err := navigateValue(result, path, nil)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving async operation path parameter, err=%s", err)
		}
		data[key] = value.(string)
	}

	url, err := replaceVars(d, "clusters/{id}", data)
	if err != nil {
		return nil, err
	}
	url = client.ServiceURL(url)

	return waitToFinish(
		[]string{"200"},
		[]string{"100"},
		timeout, 1*time.Second,
		func() (interface{}, string, error) {
			r := golangsdk.Result{}
			_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
				MoreHeaders: map[string]string{"Content-Type": "application/json"}})
			if r.Err != nil {
				return nil, "", nil
			}

			status, err := navigateValue(r.Body, []string{"status"}, nil)
			if err != nil {
				return nil, "", nil
			}
			return r.Body, status.(string), nil
		},
	)
}

func buildCssClusterV1ExtendClusterParameters(opts map[string]interface{}, arrayIndex map[string]int) (interface{}, error) {
	params := make(map[string]interface{})

	v, err := expandCssClusterV1ExtendClusterNodeNum(opts, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["modifySize"] = v
	}

	if len(params) == 0 {
		return params, nil
	}

	params = map[string]interface{}{"grow": params}

	return params, nil
}

func sendCssClusterV1ExtendClusterRequest(d *schema.ResourceData, params interface{},
	client *golangsdk.ServiceClient) (interface{}, error) {
	url, err := replaceVars(d, "clusters/{id}/extend", nil)
	if err != nil {
		return nil, err
	}
	url = client.ServiceURL(url)

	r := golangsdk.Result{}
	_, r.Err = client.Post(url, params, &r.Body, &golangsdk.RequestOpts{
		OkCodes: successHTTPCodes,
	})
	if r.Err != nil {
		return nil, fmt.Errorf("Error running api(extend_cluster), err=%s", r.Err)
	}
	return r.Body, nil
}

func asyncWaitCssClusterV1ExtendCluster(d *schema.ResourceData, config *Config, result interface{},
	client *golangsdk.ServiceClient, timeout time.Duration) (interface{}, error) {

	url, err := replaceVars(d, "clusters/{id}", nil)
	if err != nil {
		return nil, err
	}
	url = client.ServiceURL(url)

	return waitToFinish(
		[]string{"Done"}, []string{"Pending"}, timeout, 1*time.Second,
		func() (interface{}, string, error) {
			r := golangsdk.Result{}
			_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
				MoreHeaders: map[string]string{"Content-Type": "application/json"}})
			if r.Err != nil {
				return nil, "", nil
			}

			if checkCssClusterV1ExtendClusterFinished(r.Body) {
				return r.Body, "Done", nil
			}
			return r.Body, "Pending", nil
		},
	)
}

func sendCssClusterV1ReadRequest(d *schema.ResourceData, client *golangsdk.ServiceClient) (interface{}, error) {
	url, err := replaceVars(d, "clusters/{id}", nil)
	if err != nil {
		return nil, err
	}
	url = client.ServiceURL(url)

	r := golangsdk.Result{}
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"}})
	if r.Err != nil {
		return nil, fmt.Errorf("Error running api(read) for resource(CssClusterV1), err=%s", r.Err)
	}

	return r.Body, nil
}

func fillCssClusterV1ReadRespBody(body interface{}) interface{} {
	result := make(map[string]interface{})
	val, ok := body.(map[string]interface{})
	if !ok {
		val = make(map[string]interface{})
	}

	if v, ok := val["actionProgress"]; ok {
		result["actionProgress"] = v
	} else {
		result["actionProgress"] = nil
	}

	if v, ok := val["actions"]; ok {
		result["actions"] = v
	} else {
		result["actions"] = nil
	}

	if v, ok := val["created"]; ok {
		result["created"] = v
	} else {
		result["created"] = nil
	}

	if v, ok := val["datastore"]; ok {
		result["datastore"] = fillCssClusterV1ReadRespDatastore(v)
	} else {
		result["datastore"] = nil
	}

	if v, ok := val["endpoint"]; ok {
		result["endpoint"] = v
	} else {
		result["endpoint"] = nil
	}

	if v, ok := val["failed_reasons"]; ok {
		result["failed_reasons"] = fillCssClusterV1ReadRespFailedReasons(v)
	} else {
		result["failed_reasons"] = nil
	}

	if v, ok := val["id"]; ok {
		result["id"] = v
	} else {
		result["id"] = nil
	}

	if v, ok := val["instances"]; ok {
		result["instances"] = fillCssClusterV1ReadRespInstances(v)
	} else {
		result["instances"] = nil
	}

	if v, ok := val["name"]; ok {
		result["name"] = v
	} else {
		result["name"] = nil
	}

	if v, ok := val["securityGroupId"]; ok {
		result["securityGroupId"] = v
	} else {
		result["securityGroupId"] = nil
	}

	if v, ok := val["status"]; ok {
		result["status"] = v
	} else {
		result["status"] = nil
	}

	if v, ok := val["subnetId"]; ok {
		result["subnetId"] = v
	} else {
		result["subnetId"] = nil
	}

	if v, ok := val["updated"]; ok {
		result["updated"] = v
	} else {
		result["updated"] = nil
	}

	if v, ok := val["vpcId"]; ok {
		result["vpcId"] = v
	} else {
		result["vpcId"] = nil
	}

	return result
}

func fillCssClusterV1ReadRespDatastore(value interface{}) interface{} {
	if value == nil {
		return nil
	}

	value1, ok := value.(map[string]interface{})
	if !ok {
		value1 = make(map[string]interface{})
	}
	result := make(map[string]interface{})

	if v, ok := value1["type"]; ok {
		result["type"] = v
	} else {
		result["type"] = nil
	}

	if v, ok := value1["version"]; ok {
		result["version"] = v
	} else {
		result["version"] = nil
	}

	return result
}

func fillCssClusterV1ReadRespFailedReasons(value interface{}) interface{} {
	if value == nil {
		return nil
	}

	value1, ok := value.(map[string]interface{})
	if !ok {
		value1 = make(map[string]interface{})
	}
	result := make(map[string]interface{})

	if v, ok := value1["error_code"]; ok {
		result["error_code"] = v
	} else {
		result["error_code"] = nil
	}

	if v, ok := value1["error_msg"]; ok {
		result["error_msg"] = v
	} else {
		result["error_msg"] = nil
	}

	return result
}

func fillCssClusterV1ReadRespInstances(value interface{}) interface{} {
	if value == nil {
		return nil
	}

	value1, ok := value.([]interface{})
	if !ok || len(value1) == 0 {
		return nil
	}

	n := len(value1)
	result := make([]interface{}, n, n)
	for i := 0; i < n; i++ {
		val := make(map[string]interface{})
		item := value1[i].(map[string]interface{})

		if v, ok := item["id"]; ok {
			val["id"] = v
		} else {
			val["id"] = nil
		}

		if v, ok := item["name"]; ok {
			val["name"] = v
		} else {
			val["name"] = nil
		}

		if v, ok := item["status"]; ok {
			val["status"] = v
		} else {
			val["status"] = nil
		}

		if v, ok := item["type"]; ok {
			val["type"] = v
		} else {
			val["type"] = nil
		}

		result[i] = val
	}

	return result
}

func setCssClusterV1Properties(d *schema.ResourceData, response map[string]interface{}) error {
	opts := resourceCssClusterV1UserInputParams(d)

	v, err := navigateValue(response, []string{"read", "created"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:created, err: %s", err)
	}
	if err = d.Set("created", v); err != nil {
		return fmt.Errorf("Error setting Cluster:created, err: %s", err)
	}

	v, err = navigateValue(response, []string{"read", "endpoint"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:endpoint, err: %s", err)
	}
	if err = d.Set("endpoint", v); err != nil {
		return fmt.Errorf("Error setting Cluster:endpoint, err: %s", err)
	}

	v, err = navigateValue(response, []string{"read", "datastore", "type"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:engine_type, err: %s", err)
	}
	if err = d.Set("engine_type", v); err != nil {
		return fmt.Errorf("Error setting Cluster:engine_type, err: %s", err)
	}

	v, err = navigateValue(response, []string{"read", "datastore", "version"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:engine_version, err: %s", err)
	}
	if err = d.Set("engine_version", v); err != nil {
		return fmt.Errorf("Error setting Cluster:engine_version, err: %s", err)
	}

	v, err = navigateValue(response, []string{"read", "name"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:name, err: %s", err)
	}
	if err = d.Set("name", v); err != nil {
		return fmt.Errorf("Error setting Cluster:name, err: %s", err)
	}

	v, _ = opts["node_config"]
	v, err = flattenCssClusterV1NodeConfig(response, nil, v)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:node_config, err: %s", err)
	}
	if err = d.Set("node_config", v); err != nil {
		return fmt.Errorf("Error setting Cluster:node_config, err: %s", err)
	}

	v, _ = opts["nodes"]
	v, err = flattenCssClusterV1Nodes(response, nil, v)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:nodes, err: %s", err)
	}
	if err = d.Set("nodes", v); err != nil {
		return fmt.Errorf("Error setting Cluster:nodes, err: %s", err)
	}

	return nil
}

func flattenCssClusterV1NodeConfig(d interface{}, arrayIndex map[string]int, currentValue interface{}) (interface{}, error) {
	result, ok := currentValue.([]interface{})
	if !ok || len(result) == 0 {
		result = make([]interface{}, 1, 1)
	}
	hasInitValue := (result[0] != nil)
	if !hasInitValue {
		result[0] = make(map[string]interface{})
	}
	r := result[0].(map[string]interface{})

	if _, ok := r["availability_zone"]; !ok {
		r["availability_zone"] = nil
	}

	if _, ok := r["flavor"]; !ok {
		r["flavor"] = nil
	}

	v, _ := r["network_info"]
	v, err := flattenCssClusterV1NodeConfigNetworkInfo(d, arrayIndex, v)
	if err != nil {
		return nil, fmt.Errorf("Error reading Cluster:network_info, err: %s", err)
	}
	r["network_info"] = v

	if _, ok := r["volume"]; !ok {
		r["volume"] = nil
	}

	if hasInitValue {
		return result, nil
	}
	for _, v := range r {
		if v != nil {
			return result, nil
		}
	}
	return currentValue, nil
}

func flattenCssClusterV1NodeConfigNetworkInfo(d interface{}, arrayIndex map[string]int, currentValue interface{}) (interface{}, error) {
	result, ok := currentValue.([]interface{})
	if !ok || len(result) == 0 {
		result = make([]interface{}, 1, 1)
	}
	hasInitValue := (result[0] != nil)
	if !hasInitValue {
		result[0] = make(map[string]interface{})
	}
	r := result[0].(map[string]interface{})

	v, err := navigateValue(d, []string{"read", "securityGroupId"}, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error reading Cluster:security_group_id, err: %s", err)
	}
	r["security_group_id"] = v

	v, err = navigateValue(d, []string{"read", "subnetId"}, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error reading Cluster:subnet_id, err: %s", err)
	}
	r["subnet_id"] = v

	if _, ok := r["vpc_id"]; !ok {
		r["vpc_id"] = nil
	}

	if hasInitValue {
		return result, nil
	}
	for _, v := range r {
		if v != nil {
			return result, nil
		}
	}
	return currentValue, nil
}

func flattenCssClusterV1Nodes(d interface{}, arrayIndex map[string]int, currentValue interface{}) (interface{}, error) {
	n := 0
	hasInitValue := true
	result, ok := currentValue.([]interface{})
	if !ok || len(result) == 0 {
		v, err := navigateValue(d, []string{"read", "instances"}, arrayIndex)
		if err != nil {
			return nil, err
		}
		if v1, ok := v.([]interface{}); ok && len(v1) > 0 {
			n = len(v1)
		} else {
			return currentValue, nil
		}
		result = make([]interface{}, 0, n)
		hasInitValue = false
	} else {
		n = len(result)
	}

	newArrayIndex := make(map[string]int)
	if arrayIndex != nil {
		for k, v := range arrayIndex {
			newArrayIndex[k] = v
		}
	}

	for i := 0; i < n; i++ {
		newArrayIndex["read.instances"] = i

		var r map[string]interface{}
		if len(result) >= (i+1) && result[i] != nil {
			r = result[i].(map[string]interface{})
		} else {
			r = make(map[string]interface{})
		}

		v, err := navigateValue(d, []string{"read", "instances", "id"}, newArrayIndex)
		if err != nil {
			return nil, fmt.Errorf("Error reading Cluster:id, err: %s", err)
		}
		r["id"] = v

		v, err = navigateValue(d, []string{"read", "instances", "name"}, newArrayIndex)
		if err != nil {
			return nil, fmt.Errorf("Error reading Cluster:name, err: %s", err)
		}
		r["name"] = v

		v, err = navigateValue(d, []string{"read", "instances", "type"}, newArrayIndex)
		if err != nil {
			return nil, fmt.Errorf("Error reading Cluster:type, err: %s", err)
		}
		r["type"] = v

		if len(result) >= (i + 1) {
			if result[i] == nil {
				result[i] = r
			}
		} else {
			for _, v := range r {
				if v != nil {
					result = append(result, r)
					break
				}
			}
		}
	}

	if hasInitValue || len(result) > 0 {
		return result, nil
	}
	return currentValue, nil
}
