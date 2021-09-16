package uuidrange

import (
	"reflect"
	"testing"
)

func TestUUIDRange(t *testing.T) {
	cases := []struct {
		name string
		n    uint8
		want Ranges
	}{
		{
			name: "split on 2",
			n:    2,
			want: Ranges{
				{From: "00000000-0000-0000-0000-000000000000", To: "7fffffff-ffff-ffff-ffff-ffffffffffff"},
				{From: "80000000-0000-0000-0000-000000000000", To: "ffffffff-ffff-ffff-ffff-ffffffffffff"},
			},
		},
		{
			name: "split on 3",
			n:    3,
			want: Ranges{
				{From: "00000000-0000-0000-0000-000000000000", To: "55555555-5555-5555-5555-555555555554"},
				{From: "55555555-5555-5555-5555-555555555555", To: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa9"},
				{From: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", To: "ffffffff-ffff-ffff-ffff-fffffffffffe"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := New(c.n)
			if !reflect.DeepEqual(c.want, got) {
				t.Errorf("unexpected split:\n- want: %v\n-  got: %v", c.want, got)
			}
		})
	}
}
