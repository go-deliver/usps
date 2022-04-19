package api

import (
	"bytes"
	"encoding/xml"
	"errors"
	"html"
	"io"
	"net/http"

	"github.com/p-lau/usps/types"
)

func (c *Client) Call(request request, response response) error {
	// USERID is required
	if c.username == "" {
		return errors.New("username is required")
	}
	request.SetUser(c.username)

	// Uses default http client if not set
	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}

	// Marshals request to XML, in bytes
	req, err := xml.Marshal(request)
	if err != nil {
		return err
	}

	// Appends &XML to the body
	body := bytes.NewBuffer(append([]byte("&XML="), req...))

	// Creates the request and sends it
	resp, err := c.httpClient.Post(URL+request.API(), "text/xml", body)
	if err != nil {
		return err
	}

	// Closes the response body
	defer resp.Body.Close()

	// Checks for network errors
	if resp.StatusCode != http.StatusOK {
		return errors.New("HTTP Error: " + resp.Status)
	}

	// Reads the response body
	bodyBytes, err := io.ReadAll(resp.Body)

	// Check for XML/API errors
	USPSError := new(types.Error)
	if err = xml.Unmarshal(bodyBytes, USPSError); err != nil {
		return err
	}
	if USPSError.Description != "" {
		return errors.New(html.UnescapeString(USPSError.Description))
	}

	// Unmarshals the response body
	return xml.Unmarshal(bodyBytes, response)
}
