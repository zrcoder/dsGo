package set

import (
	"fmt"

	. "github.com/zrcoder/dsGo"
)

func ExampleRange() {
	const tolal = 10
	set := NewWithCapacity(tolal)
	for i := 1; i <= tolal; i++ {
		set.Add(i)
	}

	sum := 0
	set.Range(func(item Any) bool {
		switch v := item.(type) {
		case int:
			sum += v
			return false
		default:
			return true
		}
	})

	fmt.Println(sum)

	// OutPut:
	// 55
}
