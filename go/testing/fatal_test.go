package testing

import "testing"

// サブテストの中でfatalしても影響はない。
func TestFatal(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		t.Fatal("first")
	})
	t.Run("second", func(t *testing.T) {
		t.Fatal("")
	})
	t.Fatal("in parent")
	t.Run("third", func(t *testing.T) {
		t.Fatal("third")
	})
}

// shuffleしてもサブテストの実行順序は同じ
func TestTable(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "first"},
		{name: "second"},
		{name: "third"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Fatal(test.name)
		})
	}
}

// ランダムに実行される
func TestMapTable(t *testing.T) {
	tests := map[string]struct {
		name string
	}{
		"first":  {name: "first"},
		"second": {name: "second"},
		"third":  {name: "third"},
		"forth":  {name: "forth"},
		"5":      {name: "5"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Fatal(test.name)
		})
	}
}

func TestTableParallel(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "first"},
		{name: "second"},
		{name: "third"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			t.Fatal(test.name, "in parallel")
		})
	}
}
