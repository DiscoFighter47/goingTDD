package pointers

import "fmt"

// Bitcoin represents bitcoin type
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
