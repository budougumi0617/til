package snippet

import "testing"

// 技術書典8用のサンプルコード
func Operation(x, y int, op string) int {
	var r int
	switch op {
	case "+":
		r = x + y
	case "-":
		r = x - y
	default:
		// 異常系は省略
	}
	return r
}

var ops = map[string]func(int, int) int{
	"+": func(x int, y int) int {
		return x + y
	},
	"-": func(x int, y int) int {
		return x - y
	},
}

func Operation2(x, y int, op string) int {
	if f, ok := ops[op]; ok {
		return f(x, y)
	}
	return 0 // 異常系は省略
}

func TestOperation(t *testing.T) {
	if Operation(4, 2, "+") != Operation2(4, 2, "+") {
		t.Errorf("want %d, but %d\n", Operation(4, 2, "+"), Operation2(4, 2, "+"))
	}
}

type Item struct {
	name  string
	price int
}

func (i *Item) Price() int { return i.price }

func (i *Item) SetPrice(p int) { i.price = p }

type ItemWithTax struct {
	Item
}

func (i *ItemWithTax) Price() int {
	return int(float64(i.Item.Price()) * 1.1)
}

func TestItem(t *testing.T) {
	iwt := ItemWithTax{}
	iwt.SetPrice(1000)
	if iwt.Price() != 1100 {
		t.Errorf("want 1100, but %d\n", iwt.Price())
	}
}

type Modem struct{}

func (Modem) Dial()   {}
func (Modem) Hangup() {}
func (Modem) Sender() {}
func (Modem) Recv()   {}

type Receiver interface {
	Recv()
}
