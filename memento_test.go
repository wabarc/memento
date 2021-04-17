// Copyright 2021 Wayback Archiver. All rights reserved.
// Use of this source code is governed by the GNU GPL v3
// license that can be found in the LICENSE file.

package memento // import "github.com/wabarc/memento"

import (
	"testing"
)

func TestMemento(t *testing.T) {
	var got map[string]string

	tests := []struct {
		name string
		urls []string
		got  int
	}{
		{
			name: "Without URLs",
			urls: []string{},
			got:  0,
		},
		{
			name: "Has one invalid URL",
			urls: []string{"foo bar", "https://example.com/"},
			got:  1,
		},
		{
			name: "URLs full matches",
			urls: []string{"https://example.com/", "https://example.org/"},
			got:  2,
		},
	}

	mem := &Memento{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ = mem.Mementos(test.urls)
			if len(got) != test.got {
				t.Errorf("got = %d; want %d", len(got), test.got)
			}
			for orig, dest := range got {
				if testing.Verbose() {
					t.Log(orig, "=>", dest)
				}
			}
		})
	}
}
