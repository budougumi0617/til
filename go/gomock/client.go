package gomock

// Client is interface for check gomock
type Client interface {
	Do(in string) (string, error)
}

// Foo is interface for check gomock
type Foo interface {
	Method(s string) int
}
