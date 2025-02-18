package httphelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/tidwall/gjson"

	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/pagination"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/filters"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

type HttpHelper struct {
	client      *golangsdk.ServiceClient
	requestOpts *golangsdk.RequestOpts
	url         string
	method      string
	body        map[string]any
	query       map[string]any
	queryExt    map[string]any
	filters     []*filters.JsonFilter

	offsetStart int
	pager       func(r pagination.PageResult) pagination.Page

	responseBody []byte
	result       golangsdk.Result
}

func New(client *golangsdk.ServiceClient) *HttpHelper {
	return &HttpHelper{
		client:   client,
		queryExt: make(map[string]any),
		filters:  make([]*filters.JsonFilter, 0),

		requestOpts: &golangsdk.RequestOpts{
			MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
		},
	}
}

//nolint:revive
func (c *HttpHelper) URI(url string) *HttpHelper {
	c.url = url
	return c
}

func (c *HttpHelper) Method(method string) *HttpHelper {
	c.method = method
	return c
}

func (c *HttpHelper) Body(body map[string]any) *HttpHelper {
	c.body = body
	return c
}

func (c *HttpHelper) Query(query map[string]any) *HttpHelper {
	c.query = query
	return c
}

func (c *HttpHelper) Headers(headers map[string]string) *HttpHelper {
	for key, val := range headers {
		c.requestOpts.MoreHeaders[key] = val
	}
	return c
}

func (c *HttpHelper) OkCode(okCodes ...int) *HttpHelper {
	c.requestOpts.OkCodes = okCodes
	return c
}

func (c *HttpHelper) MarkerPager(dataPath, nextExp, markerKey string) *HttpHelper {
	timestamp, _ := uuid.GenerateUUID()
	c.pager = func(r pagination.PageResult) pagination.Page {
		p := MarkerPager{
			MarkerPageBase: pagination.MarkerPageBase{PageResult: r},
			uuid:           timestamp,
			DataPath:       dataPath,
			NextExp:        nextExp,
			MarkerKey:      markerKey,
		}
		p.Owner = p
		return p
	}

	return c
}

func (c *HttpHelper) CustomPager(dataPath string, nextUrlFunc URLFunc) *HttpHelper {
	timestamp, _ := uuid.GenerateUUID()
	c.pager = func(r pagination.PageResult) pagination.Page {
		return CustomPager{
			PageResult:  r,
			uuid:        timestamp,
			DataPath:    dataPath,
			NextURLFunc: nextUrlFunc,
		}
	}
	return c
}

func (c *HttpHelper) PageSizePager(dataPath, pageNumKey, perPageKey string, perPage int) *HttpHelper {
	if perPage > 0 {
		c.queryExt[perPageKey] = perPage
	}

	timestamp, _ := uuid.GenerateUUID()
	c.pager = func(r pagination.PageResult) pagination.Page {
		return PageSizePager{
			OffsetPageBase: pagination.OffsetPageBase{PageResult: r},
			uuid:           timestamp,
			DataPath:       dataPath,
			PageNumKey:     pageNumKey,
			PerPageKey:     perPageKey,
		}
	}

	return c
}

func (c *HttpHelper) LinkPager(dataPath, linkExp string) *HttpHelper {
	timestamp, _ := uuid.GenerateUUID()
	c.pager = func(r pagination.PageResult) pagination.Page {
		return LinkPager{
			LinkedPageBase: pagination.LinkedPageBase{PageResult: r},
			uuid:           timestamp,
			DataPath:       dataPath,
			LinkExp:        linkExp,
		}
	}

	return c
}

// OffsetStart set the default value of offset.
// Note: must be executed before setting pager.
func (c *HttpHelper) OffsetStart(offsetStart int) *HttpHelper {
	c.offsetStart = offsetStart
	return c
}

func (c *HttpHelper) OffsetPager(dataPath, offsetKey, limitKey string, defaultLimit int) *HttpHelper {
	if defaultLimit > 0 {
		c.queryExt[limitKey] = defaultLimit
		c.queryExt[offsetKey] = c.offsetStart
	}
	timestamp, _ := uuid.GenerateUUID()

	c.pager = func(r pagination.PageResult) pagination.Page {
		return OffsetPager{
			OffsetPageBase: pagination.OffsetPageBase{PageResult: r},
			uuid:           timestamp,
			DataPath:       dataPath,
			DefaultLimit:   defaultLimit,
			OffsetKey:      offsetKey,
			LimitKey:       limitKey,
		}
	}

	return c
}

func (c *HttpHelper) Filter(filter *filters.JsonFilter) *HttpHelper {
	c.filters = append(c.filters, filter)
	return c
}

func (c *HttpHelper) Request() *HttpHelper {
	if c.result.Err != nil {
		return c
	}
	if c.method == "" {
		c.result.Err = fmt.Errorf("`method` is empty, please specify the client through Client(method string)")
		return c
	}

	c.buildURL()
	c.appendQueryParams()

	if c.pager != nil {
		c.requestWithPage()
		c.doFilter()
		return c
	}
	c.requestNoPage()
	c.doFilter()
	return c
}

func (c *HttpHelper) Send() (*gjson.Result, error) {
	if c.result.Err != nil {
		return nil, c.result.Err
	}
	if c.method == "" {
		return nil, fmt.Errorf("`method` is empty, please specify the client through Client(method string)")
	}

	c.buildURL()
	c.appendQueryParams()
	c.requestOpts.JSONBody = c.body

	_, err := c.client.Request(c.method, c.url, c.requestOpts)

	return nil, err
}

func (c *HttpHelper) buildURL() *HttpHelper {
	endpoint := strings.TrimRight(c.client.Endpoint, "/")
	c.url = fmt.Sprintf("%s/%s", endpoint, strings.TrimLeft(c.url, "/"))
	c.url = strings.ReplaceAll(c.url, "{project_id}", c.client.ProjectID)
	c.url = strings.ReplaceAll(c.url, "{domain_id}", c.client.DomainID)
	return c
}

func (c *HttpHelper) appendQueryParams() {
	query := make(map[string]any)
	for k, v := range c.query {
		query[k] = v
	}
	for k, v := range c.queryExt {
		if _, ok := query[k]; ok {
			continue
		}
		query[k] = v
	}
	if len(query) == 0 {
		return
	}

	params := marshalQueryParams(query)
	if strings.Contains(c.url, "?") {
		c.url = c.url + "&" + strings.TrimLeft(params, "?")
	} else {
		c.url += params
	}
}

func (c *HttpHelper) ErrorOrNil() error {
	return c.result.Err
}

func (c *HttpHelper) Result() (*gjson.Result, error) {
	if c.result.Err != nil {
		return nil, c.result.Err
	}

	jsonData := gjson.ParseBytes(c.responseBody)
	if !jsonData.Exists() {
		return nil, golangsdk.ErrDefault404{}
	}

	return &jsonData, nil
}

func (c *HttpHelper) Data() (map[string]any, error) {
	if c.result.Err != nil {
		return nil, c.result.Err
	}

	data := make(map[string]any)
	err := c.ExtractInto(&data)
	return data, err
}

func (c *HttpHelper) ExtractInto(to any) error {
	if c.result.Err != nil {
		return c.result.Err
	}

	return json.Unmarshal(c.responseBody, to)
}

func (c *HttpHelper) ToFile(dataPath, filePath string) error {
	err := utils.WriteToFile(filePath, "", false)
	if err != nil {
		return err
	}

	var fileErr error
	count := 0
	err = c.EachPage(func(jsonData *gjson.Result, err error) bool {
		if err != nil {
			fileErr = err
			return true
		}

		arr := jsonData.Get(dataPath).Array()
		for _, v := range arr {
			if err := utils.WriteToFile(filePath, fmt.Sprintf("%s\n", v.Raw), true); err != nil {
				log.Printf("[ERROR] unable to write file: %s, error: %s", filePath, err)
				count++
				fileErr = fmt.Errorf("%v items that failed to be written to the file %s", count, err)
			}
		}
		return true
	})

	mErr := multierror.Append(nil, err, fileErr)
	return mErr.ErrorOrNil()
}

func (c *HttpHelper) EachPage(handler func(*gjson.Result, error) bool) error {
	if c.method == "" {
		c.result.Err = fmt.Errorf("`method` is empty, please specify the client through Client(method string)")
	}
	if c.result.Err != nil {
		return c.result.Err
	}

	if c.pager == nil {
		return fmt.Errorf("`EachPage` only supports paginated query data")
	}

	c.buildURL()
	c.appendQueryParams()

	pager := pagination.NewPager(c.client, c.url, c.pager)
	pager.Headers = c.requestOpts.MoreHeaders

	return pager.EachPage(func(page pagination.Page) (bool, error) {
		jsonData, err := bodyToGJson(page.GetBody())
		b := handler(jsonData, err)
		return b, nil
	})
}

func (c *HttpHelper) requestWithPage() {
	body := make(map[string]any)
	pager := pagination.NewPager(c.client, c.url, c.pager)
	pager.Headers = c.requestOpts.MoreHeaders

	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		b := page.GetBody().(map[string]interface{})
		mergeMaps(body, b)
		return true, nil
	})

	if err != nil {
		c.result.Err = err
		return
	}

	c.result.Body = body
	c.parseRspBody()
}

func (c *HttpHelper) requestNoPage() {
	var err error

	switch c.method {
	case "HEAD":
		_, err = c.client.Head(c.url, c.requestOpts)
	case "GET":
		_, err = c.client.Get(c.url, &c.result.Body, c.requestOpts)
	case "POST":
		_, err = c.client.Post(c.url, c.body, &c.result.Body, c.requestOpts)
	case "PUT":
		_, err = c.client.Put(c.url, c.body, &c.result.Body, c.requestOpts)
	case "PATCH":
		_, err = c.client.Patch(c.url, c.body, &c.result.Body, c.requestOpts)
	case "DELETE":
		_, err = c.client.DeleteWithBodyResp(c.url, c.body, &c.result.Body, c.requestOpts)
	}

	c.result.Err = err
	c.parseRspBody()
}

func (c *HttpHelper) parseRspBody() {
	if c.result.Err != nil {
		return
	}

	b, err := bodyToBytes(c.result.Body)
	c.responseBody = b
	c.result.Err = err
}

func (c *HttpHelper) doFilter() {
	if len(c.filters) == 0 || c.result.Err != nil {
		return
	}

	var data any
	if err := json.Unmarshal(c.responseBody, &data); err != nil {
		c.result.Err = err
		return
	}

	for _, filter := range c.filters {
		query := filters.New().
			Data(data).
			From(filter.GetFrom()).
			Filter(filter.GetFilter())

		for _, q := range filter.GetQueries() {
			query = query.Where(q.Key, q.Operator, q.Value)
		}

		r, err := query.Get()
		if err != nil {
			c.result.Err = err
			return
		}
		data = r
	}

	b, err := json.Marshal(data)
	if err != nil {
		c.result.Err = err
		return
	}
	c.responseBody = b
}

func marshalQueryParams(params map[string]any) string {
	query := url.Values{}

	for key, val := range params {
		v := reflect.ValueOf(val)
		if !v.IsValid() {
			continue
		}

		switch v.Kind() {
		case reflect.String:
			if !v.IsZero() {
				query.Add(key, v.String())
			}
		case reflect.Bool:
			query.Add(key, strconv.FormatBool(v.Bool()))
		case reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				if v.Index(i).Type().Kind() == reflect.String && v.Index(i).IsZero() {
					continue
				}
				query.Add(key, fmt.Sprintf("%v", v.Index(i).Interface()))
			}
		case reflect.Map:
			if v.Type().Key().Kind() == reflect.String && v.Type().Elem().Kind() == reflect.String {
				var s []string
				for _, k := range v.MapKeys() {
					value := v.MapIndex(k).String()
					s = append(s, fmt.Sprintf("'%s':'%s'", k.String(), value))
				}
				query.Add(key, fmt.Sprintf("{%s}", strings.Join(s, ", ")))
			}
		default:
			query.Add(key, fmt.Sprintf("%v", v.Interface()))
		}
	}

	u := &url.URL{RawQuery: query.Encode()}
	return u.String()
}

func bodyToGJson(body any) (*gjson.Result, error) {
	b, err := bodyToBytes(body)
	if err != nil {
		return nil, err
	}
	result := gjson.ParseBytes(b)
	return &result, nil
}

func bodyToBytes(body any) ([]byte, error) {
	if reader, ok := body.(io.Reader); ok {
		if readCloser, ok := reader.(io.Closer); ok {
			defer readCloser.Close()
		}

		return io.ReadAll(reader)
	}

	buffer := &bytes.Buffer{}
	enc := json.NewEncoder(buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(body)
	return buffer.Bytes(), err
}

func mergeMaps(target, source map[string]any) map[string]any {
	for key, sv := range source {
		tagVal, ok := target[key]
		if !ok {
			target[key] = sv
			continue
		}

		switch tv := tagVal.(type) {
		case map[string]any:
			if v, ok := sv.(map[string]any); ok {
				target[key] = mergeMaps(tv, v)
			}
		case []any:
			if v, ok := sv.([]any); ok {
				target[key] = append(tv, v...)
			}
		default:
			target[key] = sv
		}
	}
	return target
}
