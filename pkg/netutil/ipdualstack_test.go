// Copyright (C) 2022  mieru authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package netutil_test

import (
	"testing"

	"github.com/enfein/mieru/pkg/netutil"
)

func TestMaybeDecorateIPv6(t *testing.T) {
	testcases := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"google.com", "google.com"},
		{"0.0.0.0", "0.0.0.0"},
		{"::", "[::]"},
		{"2001:db8::", "[2001:db8::]"},
	}

	for _, tc := range testcases {
		if out := netutil.MaybeDecorateIPv6(tc.input); out != tc.want {
			t.Errorf("MaybeDecorateIPv6(%q) = %q, want %q", tc.input, out, tc.want)
		}
	}
}
