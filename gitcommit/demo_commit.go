package gitcommit

import (
	"fmt"

	"github.com/gookit/goutil/dump"
)

func main() {

}

func DemoAction1() {
	fmt.Println("Demo Action 1")
	dump.P("this is add  DemoAction1 Content")
}

func DemoAction2() {
	fmt.Println("Demo Action 2")
	dump.P("this is add  DemoAction2 Content")
}

func DemoAction3() {
	fmt.Println("Demo Action 3")
}
