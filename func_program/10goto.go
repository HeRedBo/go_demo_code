package func_program

import "fmt"

func GoTo()  {
	fmt.Println("日照香炉生紫烟")
	fmt.Println("遥看破布挂前川")
	fmt.Println("飞流直下三千尺")
	goto GAMEOVER

LASTWORD:
	fmt.Println("疑是银河落九天")
	return

GAMEOVER:
	fmt.Println("GAME OVER!")
	goto LASTWORD
}