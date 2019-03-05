package gomock

// Client is interface for check gomock
type Client interface {
	Do(req string) (string, error)
}

// Foo is interface for check gomock
type Foo interface {
	Method(s string) int
}
