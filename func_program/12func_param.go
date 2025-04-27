package func_program

import (
	"fmt"

	"github.com/gookit/goutil/dump"
)

func SayLove2Go() {

	dump.P("人生苦短，快学GO")
	dump.P("此时不浪，更待何时")
	dump.P("师姐在干嘛~")

	fmt.Println("人生苦短，必须性感")
	fmt.Println("扬天大笑出门去，我辈岂是逢蒿人")

}

func SayLoveToSb(name string) {
	fmt.Println("人生苦短，必须性感")
	fmt.Printf("我爱%s\n", name)
	word := fmt.Sprintf("我爱%s", name)
	dump.P(word)

}
