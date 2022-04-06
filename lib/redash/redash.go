package redash

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	Token      string
	retMax     int
}

func NewRedash(uri, token string, httpclient *http.Client) (*Client, error) {
	var c Client
	u, err := url.Parse(uri)
	if err != nil {
		return &c, err
	}
	return &Client{
		URL:        u,
		Token:      token,
		HTTPClient: httpclient,
		retMax:     3,
	}, nil
}

func (c Client) newRequest(method, spath string, body io.Reader, params map[string]string) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("api_key", c.Token)

	for k, v := range params {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "Content-Type: application/json")
	req.Header.Set("User-Agent", "redash-client")
	return req, nil
}

func isErrorRetryable(resp *http.Response) bool {
	if resp.StatusCode >= http.StatusInternalServerError {
		return true
	}
	return false
}

func (c Client) DoQuery(method string, queryID int, params map[string]string) ([]byte, error) {
	t := "/api/queries/"
	path := fmt.Sprintf("%s/%d/%s", t, queryID, "results")
	req, err := c.newRequest(method, path, nil, params)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if isErrorRetryable(resp) {
		ret := 0
		for {
			ret++
			resp, err := c.HTTPClient.Do(req)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			if resp.StatusCode == 200 {
				break
			}

			if c.retMax < ret {
				return nil, errors.New("retry max exceeded")
			}
			time.Sleep(time.Duration(ret) * time.Second)
		}
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("%s: %d", "error not 200", resp.StatusCode))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
