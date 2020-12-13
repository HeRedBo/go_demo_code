package func_program

import "fmt"

func Compare() {
	fmt.Println((1+1)==2)
	fmt.Println((1+1)==3)

	fmt.Println((1+1)!=2)
	fmt.Println((1+1)!=3)

	//true,false
	fmt.Println((2+3)>4)
	fmt.Println((2+3)<3)

	//true,true
	fmt.Println((2+3)>=5)
	fmt.Println((2+3)<=5)
}