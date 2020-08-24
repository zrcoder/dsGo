package ringbuffer

import "fmt"

func Example() {
	buf := NewWithSize(4)
	buf.Write(123)
	buf.Write(456)
	buf.Write(789)
	buf.Write(666)

	fmt.Println(buf.Read())
	fmt.Println(buf.Read())
	fmt.Println(buf.Read())

	buf.Write(333)
	buf.Write(555)

	fmt.Println(buf.Read())
	fmt.Println(buf.Read())
	fmt.Println(buf.Read())
	fmt.Println(buf.Read())

	// output:
	// 123
	// 456
	// 789
	// 666
	// 333
	// 555
	// <nil>
}
