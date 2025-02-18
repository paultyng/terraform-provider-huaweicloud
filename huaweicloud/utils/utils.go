package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jmespath/go-jmespath"

	"github.com/chnsz/golangsdk"
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
			k := string(src[1 : len(src)-2])
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

// ExpandToStringMap takes the result for a map of string and returns a map[string]string
func ExpandToStringMap(v map[string]interface{}) map[string]string {
	s := make(map[string]string)
	for key, val := range v {
		if strVal, ok := val.(string); ok && strVal != "" {
			s[key] = strVal
		}
	}

	return s
}

// ExpandToStringListPointer takes the result for an array of strings and returns a pointer of the array
func ExpandToStringListPointer(v []interface{}) *[]string {
	s := ExpandToStringList(v)

	return &s
}

// ExpandToIntList takes the result for an array of intgers and returns a []int
func ExpandToIntList(v []interface{}) []int {
	s := make([]int, 0, len(v))
	for _, val := range v {
		if intVal, ok := val.(int); ok {
			s = append(s, intVal)
		}
	}
	return s
}

// ExpandToInt32List takes the result for an array of intgers and returns a []int32
func ExpandToInt32List(v []interface{}) []int32 {
	s := make([]int32, 0, len(v))
	for _, val := range v {
		if intVal, ok := val.(int); ok {
			s = append(s, int32(intVal))
		}
	}
	return s
}

// ExpandToInt32ListPointer takes the result for an array of in32 and returns a pointer of the array
func ExpandToInt32ListPointer(v []interface{}) *[]int32 {
	s := ExpandToInt32List(v)

	return &s
}

// ExpandToStringListBySet takes the result for a set of strings and returns a []string
func ExpandToStringListBySet(v *schema.Set) []string {
	s := make([]string, 0, v.Len())
	for _, val := range v.List() {
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

// SliceContains checks if a target object is present in a slice (the type of the elemetes which same as the target
// object).
func SliceContains(slice []interface{}, target interface{}) bool {
	for _, v := range slice {
		if reflect.DeepEqual(v, target) {
			return true
		}
	}
	return false
}

// StrSliceContains checks if a given string is contained in a slice
// When anybody asks why Go needs generics, here you go.
func StrSliceContains(haystack []string, needle string) bool {
	return IsStrContainsSliceElement(needle, haystack, false, true)
}

// StrSliceContainsAnother checks whether a string slice (b) contains another string slice (s).
func StrSliceContainsAnother(b []string, s []string) bool {
	// The empty set is the subset of any set.
	if len(s) < 1 {
		return true
	}
	for _, v := range s {
		if !StrSliceContains(b, v) {
			return false
		}
	}
	return true
}

// IsStrContainsSliceElement returns true if the string exists in given slice or contains in one of slice elements when
// open exact flag. Also you can ignore case for this check.
func IsStrContainsSliceElement(str string, sl []string, ignoreCase, isExcat bool) bool {
	if ignoreCase {
		str = strings.ToLower(str)
	}
	for _, s := range sl {
		if ignoreCase {
			s = strings.ToLower(s)
		}
		if isExcat && s == str {
			return true
		}
		if !isExcat && strings.Contains(str, s) {
			return true
		}
	}
	return false
}

// IsSliceContainsAnyAnotherSliceElement is a method that used to determine whether a list contains any element of
// another list (including its fragments belonging to the current string), returns true if it contains.
// sl: The slice body used to determine the inclusion relationship.
// another: The included slice object used to determine the inclusion relationship.
// ignoreCase: Whether to ignore case.
// isExcat: Whether the inclusion relationship of string objects applies exact matching rules.
func IsSliceContainsAnyAnotherSliceElement(sl, another []string, ignoreCase, isExcat bool) bool {
	for _, elem := range sl {
		if IsStrContainsSliceElement(elem, another, ignoreCase, isExcat) {
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

		switch v := v.(type) {
		case map[string]interface{}:
			if len(v) > 0 {
				withoutNil[k] = RemoveNil(v)
			}
		case []map[string]interface{}:
			rv := make([]map[string]interface{}, 0, len(v))
			for _, vv := range v {
				rst := RemoveNil(vv)
				if len(rst) > 0 {
					rv = append(rv, rst)
				}
			}
			if len(rv) > 0 {
				withoutNil[k] = rv
			}
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

// IsIPv4Address is used to check whether the addr string is IPv4 format
func IsIPv4Address(addr string) bool {
	pattern := "^((25[0-5]|2[0-4]\\d|(1\\d{2}|[1-9]?\\d))\\.){3}(25[0-5]|2[0-4]\\d|(1\\d{2}|[1-9]?\\d))$"
	matched, _ := regexp.MatchString(pattern, addr)
	return matched
}

// This function compares whether there is a containment relationship between two maps, that is,
// whether map A (rawMap) contains map B (filterMap).
//
//	Map A is {'foo': 'bar'} and filter map B is {'foo': 'bar'} or {'foo': 'bar,dor'} will return true.
//	Map A is {'foo': 'bar'} and filter map B is {'foo': 'dor'} or {'foo1': 'bar'} will return false.
//	Map A is {'foo': 'bar'} and filter map B is {'foo': ''} will return true.
//	Map A is {'foo': 'bar'} and filter map B is {'': 'bar'} or {'': ''} will return false.
//
// The value of filter map 'bar,for' means that the object value can be either 'bar' or 'dor'.
// Note: There is no spaces before and after the delimiter (,).
func HasMapContains(rawMap map[string]string, filterMap map[string]interface{}) bool {
	if len(filterMap) < 1 {
		return true
	}

	hasContain := true
	for key, value := range filterMap {
		hasContain = hasContain && hasMapContain(rawMap, key, value.(string))
	}

	return hasContain
}

func hasMapContain(rawMap map[string]string, filterKey, filterValue string) bool {
	if rawTag, ok := rawMap[filterKey]; ok {
		if filterValue != "" {
			filterTagValues := strings.Split(filterValue, ",")
			return StrSliceContains(filterTagValues, rawTag)
		} else {
			return true
		}
	} else {
		return false
	}
}

// WriteToPemFile is used to write the keypair to Pem file.
func WriteToPemFile(path, privateKey string) (err error) {
	// If the private key exists, give it write permission for editing (-rw-------) for root user.
	if _, err = os.ReadFile(path); err == nil {
		err = os.Chmod(path, 0600)
		if err != nil {
			return
		}

		defer func() {
			// read-only permission (-r--------).
			mErr := multierror.Append(err, os.Chmod(path, 0400))
			err = mErr.ErrorOrNil()
		}()
	}
	if err = os.WriteFile(path, []byte(privateKey), 0600); err != nil {
		return err
	}
	return nil
}

/*
MarshalValue is used to marshal the value of struct in huaweicloud-sdk-go-v3, like this:
type Xxxx struct { value string }
*/
func MarshalValue(i interface{}) string {
	if i == nil {
		return ""
	}

	jsonRaw, err := json.Marshal(i)
	if err != nil {
		log.Printf("[WARN] failed to marshal %#v: %s", i, err)
		return ""
	}

	return strings.Trim(string(jsonRaw), `"`)
}

// RandomString returns a random string with a fixed length. You can also define a custom random character set.
// Note: make sure the number is not a negative integer or a big integer.
func RandomString(n int, allowedChars ...[]rune) (result string) {
	var letters []rune

	if len(allowedChars) == 0 {
		// Using default seed.
		letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	} else {
		letters = allowedChars[0]
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("[ERROR] The number (input n) cannot be a negative integer or a large integer: %#v", r)
		}
	}()
	b := make([]rune, n)
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[n.Int64()]
	}

	result = string(b)
	return
}

// IsDebugOrHigher returns a bool type parameter, which specifies whether to print log
var validLevels = []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR"}

func IsDebugOrHigher() bool {
	logLevel := os.Getenv("TF_LOG_PROVIDER")
	if logLevel == "" {
		logLevel = os.Getenv("TF_LOG")
	}

	if logLevel != "" {
		if isValidLogLevel(logLevel) {
			logLevel = strings.ToUpper(logLevel)
			return logLevel == "DEBUG" || logLevel == "TRACE"
		} else {
			log.Printf("[WARN] Invalid log level: %q. Valid levels are: %+v", logLevel, validLevels)
		}
	}
	return false
}

func isValidLogLevel(level string) bool {
	for _, l := range validLevels {
		if strings.ToUpper(level) == string(l) {
			return true
		}
	}

	return false
}

// PathSearch evaluates a JMESPath expression against input data and returns the result.
func PathSearch(expression string, obj interface{}, defaultValue interface{}) interface{} {
	v, err := jmespath.Search(expression, obj)
	if err != nil || v == nil {
		return defaultValue
	}
	return v
}

// FlattenResponse returns the api response body if it's not empty
func FlattenResponse(resp *http.Response) (interface{}, error) {
	var respBody interface{}
	defer resp.Body.Close()
	// Don't decode JSON when there is no content
	if resp.StatusCode == http.StatusNoContent {
		_, err := io.Copy(io.Discard, resp.Body)
		return resp, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}
	return respBody, nil
}

// Reverse is a function that used to reverse the order of the characters in the given string.
func Reverse(s string) string {
	bs := []byte(s)
	for left, right := 0, len(s)-1; left < right; left++ {
		bs[left], bs[right] = bs[right], bs[left]
		right--
	}
	return string(bs)
}

func jsonBytesEqual(b1, b2 []byte) bool {
	var o1 interface{}
	if err := json.Unmarshal(b1, &o1); err != nil {
		return false
	}

	var o2 interface{}
	if err := json.Unmarshal(b2, &o2); err != nil {
		return false
	}

	return reflect.DeepEqual(o1, o2)
}

// JSONStringsEqual is the function for comparing the contents of two json strings regardless of their formatting.
// Tabs (\r \n \t) and the order of elements are not included in the comparison.
// These json strings are same:
// + "{\n\"key1\":\"value1\",\n\"key2\":\"value2\"\n}"
// + "{\"key1\":\"value1\",\"key2\":\"value2\"}"
// + "{\"key2\":\"value2\",\"key1\":\"value1\"}"
func JSONStringsEqual(s1, s2 string) bool {
	b1 := bytes.NewBufferString("")
	if err := json.Compact(b1, []byte(s1)); err != nil {
		return false
	}

	b2 := bytes.NewBufferString("")
	if err := json.Compact(b2, []byte(s2)); err != nil {
		return false
	}

	return jsonBytesEqual(b1.Bytes(), b2.Bytes())
}

type SchemaDescInput struct {
	Internal   bool     `json:"Internal,omitempty"`
	Deprecated bool     `json:"Deprecated,omitempty"`
	Required   bool     `json:"Required,omitempty"`
	Computed   bool     `json:"Computed,omitempty"`
	ForceNew   bool     `json:"ForceNew,omitempty"`
	Unscope    []string `json:"Unscope,omitempty"`
	UsedBy     []string `json:"UsedBy,omitempty"`
}

func SchemaDesc(description string, schemaDescInput SchemaDescInput) string {
	if os.Getenv("HW_SCHEMA") == "" {
		return description
	}

	b, err := json.Marshal(schemaDescInput)
	if err == nil && string(b) != "" {
		return "schema:" + string(b) + ";" + description
	}

	return description
}

// ConvertMemoryUnit is a method that used to convert the memory unit.
// Parameters:
// + memory: The memory size of the current unit, only supports int value or string corresponding to int value.
// + diffLevel: Difference level between units before and after conversion.
// diffLevel greater than 0 means that the unit is converted to a higher level, and vice versa to a lower level, e.g.
// the unit of memory input is MB, -2 means it is converted from MB to B, and 2 means it is converted from MB to TB.
func ConvertMemoryUnit(memory interface{}, diffLevel int) int {
	var memoryInt int
	switch memory := memory.(type) {
	case int:
		memoryInt = memory
	case string:
		var err error
		memoryInt, err = strconv.Atoi(memory)
		if err != nil {
			log.Printf("convert string value (%v) to int fail: %s", memory, err)
			return -1
		}
	default:
		log.Printf("unsupported memory unit type, want 'int' or 'string', but got '%T'", memory)
		return -1
	}

	if diffLevel >= 0 {
		return memoryInt / Power(1024, diffLevel)
	}
	return memoryInt * Power(1024, -diffLevel)
}

// IsUUID is a method used to determine whether a string is in UUID format.
func IsUUID(uuid string) bool {
	// Using regular expressions to match UUID formats, with or without underscores.
	pattern := "[0-9a-fA-F]{8}(-?[0-9a-fA-F]{4}){3}-?[0-9a-fA-F]{12}"
	match, _ := regexp.MatchString(pattern, uuid)
	return match
}

func mkdirParentDir(filePath string) error {
	filePath = strings.ReplaceAll(filePath, "\\", "/")
	pos := strings.LastIndexByte(filePath, '/')
	if pos == -1 {
		return nil
	}
	directory := filePath[:pos]

	if err := os.MkdirAll(directory, os.ModePerm); err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create directory: %s, error: %s", directory, err)
	}
	return nil
}

func WriteToFile(path, content string, append bool) error {
	if err := mkdirParentDir(path); err != nil {
		return err
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("failed to create file: %s, error: %s", path, err)
		}
		defer file.Close()

		_, err = file.WriteString(content)
		if err != nil {
			return fmt.Errorf("failed to write data to %s, error: %s", path, err)
		}
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to stat file: %s, error: %s", path, err)
	}

	var file *os.File
	if append {
		file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open file for appending: %s, error: %s", path, err)
		}
	} else {
		file, err = os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open file for overwriting: %s, error: %s", path, err)
		}
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write data to %s, error: %s", path, err)
	}
	return nil
}
