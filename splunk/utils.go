package splunk

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

type splunkClient struct {
	HTTPClient *http.Client
	URL        string
	Username   string
	Password   string
	AuthToken  string
}

func (c *splunkClient) BuildSplunkURL(path string, params interface{}) url.URL {

	var q url.Values
	var err error

	if params != nil {
		q, err = query.Values(params)
		if err != nil {
			// Return me?
			panic(err)
		}
	} else {
		q = url.Values{}
	}

	// Default to JSON output
	if q.Get("output_mode") == "" {
		q.Set("output_mode", "json")
	}

	// Default to returning all rows (10,000,000)
	if q.Get("count") == "" {
		q.Set("count", "0")
	}

	u := url.URL{
		Scheme:   "https",
		Host:     c.URL,
		Path:     path,
		RawQuery: q.Encode(),
	}

	return u

}

func (c *splunkClient) Get(getURL url.URL) ([]byte, error) {
	return c.DoRequest("GET", getURL, nil)
}

func (c *splunkClient) DoRequest(method string, requestURL url.URL, reqBody interface{}) ([]byte, error) {
	var buffer *bytes.Buffer
	if contentBytes, ok := reqBody.([]byte); ok {
		buffer = bytes.NewBuffer(contentBytes)
	} else {
		if content, err := c.EncodeRequestBody(reqBody); err == nil {
			buffer = bytes.NewBuffer(content)
		} else {
			return nil, err
		}
	}

	// Build the request
	request, err := c.NewRequest(method, requestURL.String(), buffer)
	if err != nil {
		return nil, err
	}

	// Run the request
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	var body []byte

	// Read the body (if available)
	if response.Body != nil {
		defer response.Body.Close()
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		//httpErr.Body = string(body)
	}

	if response != nil && (response.StatusCode < 200 || response.StatusCode >= 400) {
		httpErr := &HTTPError{
			Status:  response.StatusCode,
			Message: response.Status,
		}
		if response.Body != nil {
			httpErr.Body = string(body)
		}
		return nil, httpErr
	}

	return body, nil
}

func (c *splunkClient) NewRequest(httpMethod, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}
	if c.AuthToken != "" {
		request.Header.Add("Authorization", "Bearer "+c.AuthToken)
	} else {
		request.SetBasicAuth(c.Username, c.Password)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "Steampipe")
	return request, nil
}

// EncodeRequestBody takes a json string or object and serializes it to be used in request body
func (c *splunkClient) EncodeRequestBody(content interface{}) ([]byte, error) {
	if content != nil {
		switch value := reflect.ValueOf(content); value.Kind() {
		case reflect.String:
			return []byte(value.String()), nil
		case reflect.Map:
			return c.EncodeObject(value.Interface())
		case reflect.Struct:
			return c.EncodeObject(value.Interface())
		default:
			return nil, &HTTPError{Status: 400, Message: "Bad Request"}
		}
	}
	return nil, nil
}

// EncodeObject encodes an object into url-encoded string
func (c *splunkClient) EncodeObject(content interface{}) ([]byte, error) {
	URLValues := url.Values{}
	marshalContent, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}
	var valueMap map[string]interface{}
	if err := json.Unmarshal(marshalContent, &valueMap); err != nil {
		return nil, err
	}
	for k, v := range valueMap {
		//k = strings.ToLower(k)
		switch val := v.(type) {
		case []interface{}:
			for _, ele := range val {
				if encoded, err := encodeValue(ele); err == nil && len(encoded) > 0 {
					URLValues.Add(k, encoded)
				}
			}
		case map[string]interface{}:
			for innerK, innerV := range val {
				if encoded, err := encodeValue(innerV); err == nil && len(encoded) > 0 {
					URLValues.Set(innerK, encoded)
				}
			}
		default:
			if encoded, err := encodeValue(val); err == nil && len(encoded) > 0 {
				URLValues.Set(k, encoded)
			}
		}
	}
	return []byte(URLValues.Encode()), nil
}

func encodeValue(v interface{}) (string, error) {
	switch val := v.(type) {
	case string:
		return val, nil
	case bool:
		return strconv.FormatBool(val), nil
	case int:
		return strconv.FormatInt(int64(val), 10), nil
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(float64(val), 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("could not encode type %T", v)
	}
}

// HTTPError is raised when status code is not 2xx
type HTTPError struct {
	Status  int
	Message string
	Body    string
}

// This allows HTTPError to satisfy the error interface
func (he *HTTPError) Error() string {
	return fmt.Sprintf("Http Error: [%v] %v %v", he.Status, he.Message, he.Body)
}

func connect(_ context.Context, d *plugin.QueryData) (*splunkClient, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "splunk"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*splunkClient), nil
	}

	// Default to the env var settings
	url := os.Getenv("SPLUNK_URL")
	username := os.Getenv("SPLUNK_USERNAME")
	password := os.Getenv("SPLUNK_PASSWORD")
	authToken := os.Getenv("SPLUNK_AUTH_TOKEN")
	insecureSkipVerify := os.Getenv("SPLUNK_INSECURE_SKIP_VERIFY") != "false"

	// Prefer config settings
	splunkConfig := GetConfig(d.Connection)
	if splunkConfig.URL != nil {
		url = *splunkConfig.URL
	}
	if splunkConfig.Username != nil {
		username = *splunkConfig.Username
	}
	if splunkConfig.Password != nil {
		password = *splunkConfig.Password
	}
	if splunkConfig.AuthToken != nil {
		authToken = *splunkConfig.AuthToken
	}
	if splunkConfig.InsecureSkipVerify != nil {
		insecureSkipVerify = *splunkConfig.InsecureSkipVerify
	}

	// Defaults
	timeout := 30
	if url == "" {
		url = "localhost:8089"
	}

	// Error if the minimum config is not set
	if (username == "" || password == "") && authToken == "" {
		return nil, errors.New("username and password, or an auth_token, must be configured")
	}

	httpClient, err := newHTTPClient(time.Second*time.Duration(timeout), insecureSkipVerify)
	if err != nil {
		return nil, err
	}

	conn := &splunkClient{
		HTTPClient: httpClient,
		URL:        url,
		Username:   username,
		Password:   password,
		AuthToken:  authToken,
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func newHTTPClient(timeout time.Duration, insecureSkipVerify bool) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Timeout: timeout,
		Jar:     jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
		},
	}
	return client, nil
}
