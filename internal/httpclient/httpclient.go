package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/leyle/go-crud-starter/configandcontext"
	"io"
	"net/http"
	"net/url"
	"time"
)

const HttpClientErrCode = 0

var defaultTimeout time.Duration = configandcontext.DefaultHTTPRequestTimeout * time.Second

const (
	defaultContentType = "application/json"
)

type ClientRequest struct {
	Method      string
	Url         string
	query       url.Values
	headers     http.Header
	timeout     time.Duration
	body        []byte
	debug       bool // if true, log response body
	extractBody bool // default is true
}

type ClientResponse struct {
	Code int   // http status code and default err code(0)
	Err  error // when program err occurred
	Body []byte
	Data []byte         // response.body.data
	Raw  *http.Response // be careful, response.Body can be read exactly once
}

func NewClientRequest(method, url string) *ClientRequest {
	c := &ClientRequest{
		Method:      method,
		Url:         url,
		timeout:     defaultTimeout,
		debug:       true,
		extractBody: false,
	}
	defaultHeaders := map[string][]string{
		"Content-Type": {defaultContentType},
	}
	c.headers = defaultHeaders
	return c
}

func (r *ClientRequest) SetHeader(key string, val []string) {
	if r.headers == nil {
		headers := make(map[string][]string)
		r.headers = headers
	}
	r.headers[key] = val
}

func (r *ClientRequest) AssignHeaders(headers http.Header) {
	r.headers = headers
}

func (r *ClientRequest) SetQueryVal(name, val string) {
	if r.query == nil {
		query := make(map[string][]string)
		r.query = query
	}
	r.query[name] = []string{val}
}

func (r *ClientRequest) SetQueryVals(name string, vals []string) {
	if r.query == nil {
		query := make(map[string][]string)
		r.query = query
	}
	r.query[name] = vals
}

func (r *ClientRequest) AssignQuery(query url.Values) {
	r.query = query
}

func (r *ClientRequest) SetTimeout(t time.Duration) {
	r.timeout = t
}

func (r *ClientRequest) SetBody(body []byte) {
	r.body = body
}

func (r *ClientRequest) CloseDebug() {
	r.debug = false
}

func (r *ClientRequest) ExtractBody(yes bool) {
	r.extractBody = yes
}

func (r *ClientRequest) Do(ctx *configandcontext.APIContext) *ClientResponse {
	if r.Method == "" {
		return &ClientResponse{
			Code: HttpClientErrCode,
			Err:  errors.New("no http method in ClientRequest object"),
		}
	}

	resp := httpRequest(ctx, r)
	return resp
}

func httpRequest(ctx *configandcontext.APIContext, req *ClientRequest) *ClientResponse {
	var err error
	startT := time.Now()
	resp := &ClientResponse{
		Code: HttpClientErrCode,
		Err:  err,
	}

	logger := ctx.Logger.With().Str("method", req.Method).Str("url", req.Url).Logger()
	logger.Info().Msg("start http request...")

	// generate req
	var newReq *http.Request
	if req.body != nil {
		newReq, err = http.NewRequest(req.Method, req.Url, bytes.NewBuffer(req.body))
	} else {
		newReq, err = http.NewRequest(req.Method, req.Url, nil)
	}
	if err != nil {
		logger.Error().Err(err).Send()
		resp.Err = err
		return resp
	}

	// process url query string
	if req.query != nil {
		urlV := url.Values{}
		for k, vs := range req.query {
			if len(vs) == 1 {
				urlV.Set(k, vs[0])
			} else if len(vs) > 1 {
				for _, v := range vs {
					urlV.Add(k, v)
				}
			}
		}
		newReq.URL.RawQuery = urlV.Encode()
		logger.Debug().Str("fullUrl", newReq.URL.String()).Send()
	}

	// process headers
	for k, v := range req.headers {
		for _, vv := range v {
			newReq.Header.Add(k, vv)

		}
	}

	// timeout
	timeout := defaultTimeout
	if req.timeout > 0 {
		timeout = req.timeout
	}
	client := &http.Client{
		Timeout: timeout,
	}

	doResp, err := client.Do(newReq)
	if err != nil {
		logger.Error().Err(err).Send()
		resp.Err = err
		return resp
	}
	resp.Raw = doResp
	resp.Code = doResp.StatusCode

	// print debug info
	logger.Info().Int("statusCode", doResp.StatusCode).Str("elapsed", time.Since(startT).String()).Msg("http response")

	// if it has response body, then read it
	cl := doResp.ContentLength
	if cl > 0 {
		respBody, err := io.ReadAll(doResp.Body)
		if err != nil {
			logger.Error().Err(err).Send()
			resp.Err = err
			return resp
		}
		defer doResp.Body.Close()
		resp.Body = respBody
	}

	if cl > 0 && req.extractBody {
		var apiData APIResponse
		err = json.Unmarshal(resp.Body, &apiData)
		if err != nil {
			logger.Warn().Err(err).Msg("try to extract response to general api response failed")
		} else {
			resp.Data = apiData.Data
		}
	}

	return resp
}
