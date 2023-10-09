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

func MyAppend(src []int) []int {
	return append(src, 10)
}

func TestAppend(t *testing.T) {
	t.Run("noarmal", func(t *testing.T) {
		org := []int{1, 2, 3}
		got := MyAppend(org[:2])
		if want := []int{1, 2, 10}; !reflect.DeepEqual(org, want) {
			t.Errorf("want %v, but got %v", want, org)
		}
		if want := []int{1, 2, 10}; !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, but got %v", want, got)
		}
	})
	t.Run("完全スライス式", func(t *testing.T) {
		org := []int{1, 2, 3}
		got := MyAppend(org[:2:2])
		// 元のスライスの3要素が変更されない
		if want := []int{1, 2, 3}; !reflect.DeepEqual(org, want) {
			t.Errorf("want %v, but got %v", want, org)
		}
		// 新しいバッファを使っている
		if want := []int{1, 2, 10}; !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, but got %v", want, got)
		}
	})
}
