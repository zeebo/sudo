package sudo_test

import (
	"fmt"
	"reflect"

	"github.com/zeebo/sudo"
)

func ExampleSudo() {
	var s struct{ x int }
	x := reflect.ValueOf(&s).Elem().FieldByName("x")

	// Because x went through an unexported field, we can't set it.
	fmt.Println(x.CanSet(), s.x)

	// But if we Sudo the reflect.Value
	x = sudo.Sudo(x)

	// then our wildest dreams will come true.
	x.SetInt(10)
	fmt.Println(x.CanSet(), s.x)

	// output:
	// false 0
	// true 10
}
