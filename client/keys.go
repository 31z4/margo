package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// GetKeysPath computes a request path to the get action of keys.
func GetKeysPath(key string) string {
	return fmt.Sprintf("/keys/%v", key)
}

// Get the value of a key.
func (c *Client) GetKeys(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetKeysRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetKeysRequest create the request corresponding to the get action endpoint of the keys resource.
func (c *Client) NewGetKeysRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListKeysPath computes a request path to the list action of keys.
func ListKeysPath() string {
	return fmt.Sprintf("/keys")
}

// Retrieve all keys.
func (c *Client) ListKeys(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListKeysRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListKeysRequest create the request corresponding to the list action endpoint of the keys resource.
func (c *Client) NewListKeysRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// RemoveKeysPath computes a request path to the remove action of keys.
func RemoveKeysPath(key string) string {
	return fmt.Sprintf("/keys/%v", key)
}

// Delete a key.
func (c *Client) RemoveKeys(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRemoveKeysRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRemoveKeysRequest create the request corresponding to the remove action endpoint of the keys resource.
func (c *Client) NewRemoveKeysRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// SetKeysPayload is the keys set action payload.
type SetKeysPayload interface{}

// SetKeysPath computes a request path to the set action of keys.
func SetKeysPath(key string) string {
	return fmt.Sprintf("/keys/%v", key)
}

// Set the value of a key.
func (c *Client) SetKeys(ctx context.Context, path string, payload SetKeysPayload, contentType string) (*http.Response, error) {
	req, err := c.NewSetKeysRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSetKeysRequest create the request corresponding to the set action endpoint of the keys resource.
func (c *Client) NewSetKeysRequest(ctx context.Context, path string, payload SetKeysPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// UpdateKeysPayload is the keys update action payload.
type UpdateKeysPayload interface{}

// UpdateKeysPath computes a request path to the update action of keys.
func UpdateKeysPath(key string) string {
	return fmt.Sprintf("/keys/%v", key)
}

// Update the value of a key.
func (c *Client) UpdateKeys(ctx context.Context, path string, payload UpdateKeysPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateKeysRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateKeysRequest create the request corresponding to the update action endpoint of the keys resource.
func (c *Client) NewUpdateKeysRequest(ctx context.Context, path string, payload UpdateKeysPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
