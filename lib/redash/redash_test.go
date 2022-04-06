package redash

import (
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNewRedash(t *testing.T) {
	type args struct {
		uri        string
		token      string
		httpclient *http.Client
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRedash(tt.args.uri, tt.args.token, tt.args.httpclient)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRedash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_newRequest(t *testing.T) {
	type fields struct {
		URL        *url.URL
		HTTPClient *http.Client
		Token      string
		retMax     int
	}
	type args struct {
		method string
		spath  string
		body   io.Reader
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				URL:        tt.fields.URL,
				HTTPClient: tt.fields.HTTPClient,
				Token:      tt.fields.Token,
				retMax:     tt.fields.retMax,
			}
			got, err := c.newRequest(tt.args.method, tt.args.spath, tt.args.body, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.newRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.newRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isErrorRetryable(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isErrorRetryable(tt.args.resp); got != tt.want {
				t.Errorf("isErrorRetryable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DoQuery(t *testing.T) {
	type fields struct {
		URL        *url.URL
		HTTPClient *http.Client
		Token      string
		retMax     int
	}
	type args struct {
		method  string
		queryID int
		params  map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				URL:        tt.fields.URL,
				HTTPClient: tt.fields.HTTPClient,
				Token:      tt.fields.Token,
				retMax:     tt.fields.retMax,
			}
			got, err := c.DoQuery(tt.args.method, tt.args.queryID, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.DoQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.DoQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
