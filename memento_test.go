// Copyright 2021 Wayback Archiver. All rights reserved.
// Use of this source code is governed by the GNU GPL v3
// license that can be found in the LICENSE file.

package memento // import "github.com/wabarc/memento"

import (
	"context"
	"net/url"
	"testing"
)

func TestMemento(t *testing.T) {
	uri := "https://example.com"

	mem := &Memento{}
	in, err := url.Parse(uri)
	if err != nil {
		t.Fatal(err)
	}

	_, err = mem.Mementos(context.TODO(), in)
	if err != nil {
		t.Fatal(err)
	}
}
