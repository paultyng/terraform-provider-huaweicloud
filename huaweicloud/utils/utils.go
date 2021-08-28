package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/rts/v1/stacks"
	"gopkg.in/yaml.v2"
)

// ConvertStructToMap converts an instance of struct to a map object, and
// changes each key of fileds to the value of 'nameMap' if the key in it
// or to its corresponding lowercase.
func ConvertStructToMap(obj interface{}, nameMap map[string]string) (map[string]interface{}, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("Error converting struct to map, marshal failed:%v", err)
	}

	m, err := regexp.Compile(`"[a-z0-9A-Z_]+":`)
	if err != nil {
		return nil, fmt.Errorf("Error converting struct to map, compile regular express failed")
	}
	nb := m.ReplaceAllFunc(
		b,
		func(src []byte) []byte {
			k := fmt.Sprintf("%s", src[1:len(src)-2])
			v, ok := nameMap[k]
			if !ok {
				v = strings.ToLower(k)
			}
			return []byte(fmt.Sprintf("\"%s\":", v))
		},
	)
	log.Printf("[DEBUG]convertStructToMap:: before change b =%s", b)
	log.Printf("[DEBUG]convertStructToMap:: after change nb=%s", nb)

	p := make(map[string]interface{})
	err = json.Unmarshal(nb, &p)
	if err != nil {
		return nil, fmt.Errorf("Error converting struct to map, unmarshal failed:%v", err)
	}
	log.Printf("[DEBUG]convertStructToMap:: map= %#v\n", p)
	return p, nil
}

// ExpandToStringList takes the result for an array of strings and returns a []string
func ExpandToStringList(v []interface{}) []string {
	s := make([]string, 0, len(v))
	for _, val := range v {
		if strVal, ok := val.(string); ok && strVal != "" {
			s = append(s, strVal)
		}
	}

	return s
}

// Takes list of pointers to strings. Expand to an array
// of raw strings and returns a []interface{}
func flattenToStringList(list []*string) []interface{} {
	vs := make([]interface{}, 0, len(list))
	for _, v := range list {
		vs = append(vs, *v)
	}
	return vs
}

func pointersMapToStringList(pointers map[string]*string) map[string]interface{} {
	list := make(map[string]interface{}, len(pointers))
	for i, v := range pointers {
		list[i] = *v
	}
	return list
}

// Takes a value containing JSON string and passes it through
// the JSON parser to normalize it, returns either a parsing
// error or normalized JSON string.
func NormalizeJsonString(jsonString interface{}) (string, error) {
	var j interface{}

	if jsonString == nil || jsonString.(string) == "" {
		return "", nil
	}

	s := jsonString.(string)

	err := json.Unmarshal([]byte(s), &j)
	if err != nil {
		return s, err
	}

	// The error is intentionally ignored here to allow empty policies to passthrough validation.
	// This covers any interpolated values
	bytes, _ := json.Marshal(j)

	return string(bytes[:]), nil
}

// Takes a value containing YAML string and passes it through
// the YAML parser. Returns either a parsing
// error or original YAML string.
func checkYamlString(yamlString interface{}) (string, error) {
	var y interface{}

	if yamlString == nil || yamlString.(string) == "" {
		return "", nil
	}

	s := yamlString.(string)

	err := yaml.Unmarshal([]byte(s), &y)
	if err != nil {
		return s, err
	}

	return s, nil
}

func NormalizeStackTemplate(templateString interface{}) (string, error) {
	if looksLikeJsonString(templateString) {
		return NormalizeJsonString(templateString.(string))
	}

	return checkYamlString(templateString)
}

func FlattenStackOutputs(stackOutputs []*stacks.Output) map[string]string {
	outputs := make(map[string]string, len(stackOutputs))
	for _, o := range stackOutputs {
		outputs[*o.OutputKey] = *o.OutputValue
	}
	return outputs
}

// FlattenStackParameters is flattening list of
// stack Parameters and only returning existing
// parameters to avoid clash with default values
func FlattenStackParameters(stackParams map[string]string,
	originalParams map[string]interface{}) map[string]string {
	params := make(map[string]string, len(stackParams))
	for key, value := range stackParams {
		_, isConfigured := originalParams[key]
		if isConfigured {
			params[key] = value
		}
	}
	return params
}

// StrSliceContains checks if a given string is contained in a slice
// When anybody asks why Go needs generics, here you go.
func StrSliceContains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func JsonMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	enc := json.NewEncoder(buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(t)
	return buffer.Bytes(), err
}

// RemoveDuplicateElem removes duplicate elements from slice
func RemoveDuplicateElem(s []string) []string {
	result := []string{}
	tmpMap := map[string]byte{}
	for _, e := range s {
		l := len(tmpMap)
		tmpMap[e] = 0
		if len(tmpMap) != l {
			result = append(result, e)
		}
	}
	return result
}

func RemoveNil(data map[string]interface{}) map[string]interface{} {
	withoutNil := make(map[string]interface{})

	for k, v := range data {
		if v == nil {
			continue
		}

		switch v.(type) {
		case map[string]interface{}:
			withoutNil[k] = RemoveNil(v.(map[string]interface{}))
		default:
			withoutNil[k] = v
		}
	}

	return withoutNil
}

func IsResourceNotFound(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(golangsdk.ErrDefault404)
	return ok
}

// Method FormatTimeStampRFC3339 is used to unify the time format to RFC-3339 and return a time string.
func FormatTimeStampRFC3339(timestamp int64) string {
	createTime := time.Unix(timestamp, 0)
	return createTime.Format(time.RFC3339)
}

// Method EncodeBase64String is used to encode a string by base64.
func EncodeBase64String(str string) string {
	strByte := []byte(str)
	return base64.StdEncoding.EncodeToString(strByte)
}

// Method EncodeBase64IfNot is used to encode a string by base64 if it not a base64 string.
func EncodeBase64IfNot(str string) string {
	if _, err := base64.StdEncoding.DecodeString(str); err != nil {
		return base64.StdEncoding.EncodeToString([]byte(str))
	}
	return str
}
