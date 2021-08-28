package deprecated

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var successHTTPCodes = []int{200, 201, 202, 203, 204, 205, 206, 207, 208, 226}

func isEmptyValue(v reflect.Value) (bool, error) {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0, nil
	case reflect.Bool:
		return !v.Bool(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0, nil
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0, nil
	case reflect.Interface, reflect.Ptr:
		return v.IsNil(), nil
	case reflect.Invalid:
		return true, nil
	}
	return false, fmt.Errorf("isEmptyValue:: unknown type")
}

func addQueryParams(rawurl string, params map[string]string) (string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func replaceVars(d *schema.ResourceData, linkTmpl string, kv map[string]interface{}) (string, error) {
	re := regexp.MustCompile("{([[:word:]]+)}")

	replaceFunc := func(s string) string {
		m := re.FindStringSubmatch(s)[1]
		if kv != nil {
			if v, ok := kv[m]; ok {
				return convertToStr(v)
			}
		}
		if m == "project" {
			return "replace_holder"
		}
		if d != nil {
			if m == "id" {
				return d.Id()
			}
			v, ok := d.GetOk(m)
			if ok {
				return convertToStr(v)
			}
		}
		return ""
	}

	s := re.ReplaceAllStringFunc(linkTmpl, replaceFunc)
	return strings.Replace(s, "replace_holder/", "", 1), nil
}

func navigateMap(d interface{}, index []string) (interface{}, error) {
	for _, i := range index {
		d1, ok := d.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("navigateMap:: Can not convert to map")
		}
		d, ok = d1[i]
		if !ok {
			return nil, fmt.Errorf("navigateMap:: '%s' may not exist", i)
		}
	}
	return d, nil
}

func navigateValue(d interface{}, index []string, arrayIndex map[string]int) (interface{}, error) {
	for n, i := range index {
		if d == nil {
			return nil, nil
		}
		if d1, ok := d.(map[string]interface{}); ok {
			d, ok = d1[i]
			if !ok {
				msg := fmt.Sprintf("navigate value with index(%s)", strings.Join(index, "."))
				return nil, fmt.Errorf("%s: '%s' may not exist", msg, i)
			}
		} else {
			msg := fmt.Sprintf("navigate value with index(%s)", strings.Join(index, "."))
			return nil, fmt.Errorf("%s: Can not convert (%s) to map", msg, reflect.TypeOf(d))
		}

		if arrayIndex != nil {
			if j, ok := arrayIndex[strings.Join(index[:n+1], ".")]; ok {
				if d == nil {
					return nil, nil
				}
				if d2, ok := d.([]interface{}); ok {
					if len(d2) == 0 {
						return nil, nil
					}
					if j >= len(d2) {
						msg := fmt.Sprintf("navigate value with index(%s)", strings.Join(index, "."))
						return nil, fmt.Errorf("%s: The index is out of array", msg)
					}

					d = d2[j]
				} else {
					msg := fmt.Sprintf("navigate value with index(%s)", strings.Join(index, "."))
					return nil, fmt.Errorf("%s: Can not convert (%s) to array, index=%s.%v", msg, reflect.TypeOf(d), i, j)
				}
			}
		}
	}

	return d, nil
}

func isUserInput(d *schema.ResourceData, index []string, arrayIndex map[string]int) bool {
	var r = make([]string, 0, len(index)*2)
	for n, i := range index {
		r = append(r, i)

		if arrayIndex != nil {
			if j, ok := arrayIndex[strings.Join(index[:n+1], ".")]; ok {
				r = append(r, strconv.Itoa(j))
			}
		}
	}
	_, e := d.GetOkExists(strings.Join(r[:len(r)], "."))
	return e
}

func convertToInt(v interface{}) (int64, error) {
	s := fmt.Sprintf("%v", v)
	r, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return r, err
	}

	if i, ok := v.(int); ok {
		return int64(i), nil

	} else if i, ok := v.(float64); ok {
		return int64(i), nil

	} else if i, ok := v.(float32); ok {
		return int64(i), nil
	}

	return 0, fmt.Errorf("can not convert to integer")
}

func convertToStr(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func convertSeconds2Str(v int64) string {
	t := time.Unix(v, 0)
	return t.Format("2006-01-02 15:04:05")
}

func waitToFinish(target, pending []string, timeout, interval time.Duration, f resource.StateRefreshFunc) (interface{}, error) {
	stateConf := &resource.StateChangeConf{
		Target:     target,
		Pending:    pending,
		Refresh:    f,
		Timeout:    timeout,
		Delay:      5 * time.Second,
		MinTimeout: interval,
	}

	return stateConf.WaitForState()
}

func checkCsClusterV1DeleteFinished(data interface{}) bool {
	c, err := navigateValue(data, []string{"error_id"}, nil)
	if err != nil {
		return false
	}
	return "CS.20005" == convertToStr(c)
}

func checkCsPeeringConnectV1DeleteFinished(data interface{}) bool {
	c, err := navigateValue(data, []string{"error_id"}, nil)
	if err != nil {
		return false
	}
	return "CS.22103" == convertToStr(c)
}
