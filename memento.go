// Copyright 2021 Wayback Archiver. All rights reserved.
// Use of this source code is governed by the GNU GPL v3
// license that can be found in the LICENSE file.

package memento

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Memento represents a Memento client.
type Memento struct {
	Client *http.Client
}

func (m *Memento) Mementos(_ context.Context, input *url.URL) (string, error) {
	if m.Client == nil {
		m.Client = &http.Client{Timeout: 60 * time.Second}
	}

	dst, err := m.timegate(input.String())
	if err != nil {
		return "", err
	}

	return dst, nil
}

func (m *Memento) timegate(uri string) (string, error) {
	if m.Client == nil {
		return "", fmt.Errorf("Client must not nil")
	}
	m.Client.CheckRedirect = noRedirect

	var endpoint = "http://timetravel.mementoweb.org/timegate/"
	var url = endpoint + uri

	resp, err := m.Client.Head(url)
	if err != nil {
		return "", fmt.Errorf("The requested URI: %s is not available in an archive.", url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusFound {
		return "", fmt.Errorf(resp.Status)
	}

	location := resp.Header.Get("Location")
	if location == "" {
		location = "No Found"
	}

	return location, nil
}

func noRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}
