// Copyright 2021 Wayback Archiver. All rights reserved.
// Use of this source code is governed by the GNU GPL v3
// license that can be found in the LICENSE file.

package memento

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/wabarc/helper"
)

// Memento represents a Memento client.
type Memento struct {
	Client *http.Client
}

func (m *Memento) Mementos(links []string) (map[string]string, error) {
	collects, results := make(map[string]string), make(map[string]string)
	for _, link := range links {
		if helper.IsURL(link) {
			collects[link] = link
		}
	}
	if len(collects) == 0 {
		return results, fmt.Errorf("No found URL")
	}

	if m.Client == nil {
		m.Client = &http.Client{Timeout: 60 * time.Second}
	}

	ch := make(chan string, len(collects))
	defer close(ch)

	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, link := range collects {
		wg.Add(1)
		go func(link string) {
			mu.Lock()
			m.timegate(link, ch)
			results[link] = <-ch
			mu.Unlock()
			wg.Done()
		}(link)
		time.Sleep(time.Second)
	}
	wg.Wait()

	return results, nil
}

func (m *Memento) timegate(uri string, ch chan<- string) {
	if m.Client == nil {
		ch <- "Client must not nil"
		return
	}
	m.Client.CheckRedirect = noRedirect

	var endpoint = "http://timetravel.mementoweb.org/timegate/"
	var url = endpoint + uri

	resp, err := m.Client.Head(url)
	if err != nil {
		ch <- fmt.Sprintf("The requested URI: %s is not available in an archive.", url)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusFound {
		ch <- resp.Status
		return
	}

	location := resp.Header.Get("Location")
	if location == "" {
		location = "No Found"
	}

	ch <- location
}

func noRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}
