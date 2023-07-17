package http

import (
	"card-service/internal/util"
	"card-service/internal/util/log"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"moul.io/http2curl"
)

// HTTPClient ...
type HTTPClient struct {
	logger *log.Logger
}

// HTTPSetting ...
type HTTPSetting struct {
	RequestType  string
	URL          string
	Proxy        string
	IsEnableCert bool
	CertPath     string
	Timeout      int64
	Headers      map[string]string
	Parameters   map[string]string
	Body         io.Reader
}

// HTTPResponse ...
type HTTPResponse struct {
	StatusCode     int
	Header         map[string][]string
	Body           []byte
	RequestString  string
	ResponseString string
}

// SendRequest ...
func (h HTTPClient) SendRequest(setting *HTTPSetting) (*HTTPResponse, error) {
	if setting.Timeout == 0 {
		setting.Timeout = 30
	}

	transport, err := setProxy(setting)

	if err != nil {
		return nil, err
	}

	if err = setTLS(transport, setting); err != nil {
		return nil, err
	}

	client := http.Client{
		Transport: transport,
		Timeout:   time.Duration(setting.Timeout) * time.Second,
	}

	request, err := http.NewRequest(setting.RequestType, setting.URL, setting.Body)

	if err != nil {
		return nil, err
	}

	setHeaders(request, setting.Headers)
	setParameters(request, setting.Parameters)

	requestString, err := util.ConvertRequestToString(request)

	if err != nil {
		return nil, err
	}

	curl, err := http2curl.GetCurlCommand(request)

	if err != nil {
		return nil, err
	}

	h.logger.WithFields(map[string]interface{}{
		"Request": struct {
			Curl   string
			Detail string
		}{
			fmt.Sprintf("%v", curl),
			requestString,
		},
	}).Info()

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseString, err := util.ConvertResponseToString(response)

	if err != nil {
		return nil, err
	}

	h.logger.WithFields(map[string]interface{}{
		"Response": responseString,
	}).Info()

	result, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return &HTTPResponse{
		StatusCode:     response.StatusCode,
		Header:         response.Header,
		Body:           result,
		RequestString:  requestString,
		ResponseString: responseString,
	}, nil
}

func setProxy(setting *HTTPSetting) (*http.Transport, error) {
	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	if setting.Proxy != "" {
		path, err := url.Parse(setting.Proxy)

		if err != nil {
			return nil, err
		}

		transport.Proxy = http.ProxyURL(path)
	}

	return transport, nil
}

func setTLS(transport *http.Transport, setting *HTTPSetting) error {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	if setting.IsEnableCert {
		cert, err := ioutil.ReadFile(setting.CertPath)

		if err != nil {
			return err
		}

		certPool := x509.NewCertPool()

		certPool.AppendCertsFromPEM(cert)

		tlsConfig.RootCAs = certPool
	}

	transport.TLSClientConfig = tlsConfig

	return nil
}

func setHeaders(request *http.Request, headers map[string]string) {
	if len(headers) > 0 {
		for key, value := range headers {
			request.Header.Set(key, value)
		}
	}
}

func setParameters(request *http.Request, parameters map[string]string) {
	if len(parameters) > 0 {
		query := request.URL.Query()

		for key, value := range parameters {
			query.Add(key, value)
		}

		request.URL.RawQuery = query.Encode()
	}
}

// NewHTTPClient ...
func NewHTTPClient(l *log.Logger) *HTTPClient {
	return &HTTPClient{
		logger: l,
	}
}
