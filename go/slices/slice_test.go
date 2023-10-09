package slices

import (
	"reflect"
	"testing"
)

func MyCopy[T any](src []T) []T {
	return append([]T(nil), src...)
}

func TestCopy(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in   []any
		want []any
	}{
		"string": {
			in:   []any{"a", "b", "c"},
			want: []any{"a", "b", "c"},
		},
		"int": {
			in:   []any{1, 2, 3},
			want: []any{1, 2, 3},
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := MyCopy(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}
