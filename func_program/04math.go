package func_program

import (
	"fmt"
	"math"
)

// 基本的数学运算
func BaseMath() {

	//+
	fmt.Println(10 + 3)

	// -
	fmt.Println(10  - 3)

	// *
	fmt.Println(10  * 3)
	// /
	fmt.Println(10  / 3)

	// /
	fmt.Println(10.0 / 3.0)

	//1求余/求模
	fmt.Println(10 % 3)


	//四则运算，使用括号代表优先计算
	//2.25
	fmt.Println(((1+2)*3/4.0))


}

/*四舍五入、绝对值、乘方、开方、三角函数*/
func HighMath() {
	// 四舍五入
	fmt.Println("四舍五入数学运算：")
	fmt.Println(math.Round(3.4))
	fmt.Println(math.Round(3.5))

	//先对绝对值四舍五入再加负号
	fmt.Println(math.Round(-3.4))
	fmt.Println(math.Round(-3.5))

	/*纯舍，纯入*/
	fmt.Println("纯舍去地板",math.Floor(3.99))
	fmt.Println("纯入去天花板",math.Ceil(3.01))

	//绝对值
	fmt.Println(math.Abs(-3.4))

	//2^3=8 乘方
	fmt.Println(math.Pow(2,3))

	//3 = 9的平方根square root
	fmt.Println(math.Sqrt(9))
	//对8开立方，即求8的三分之一次方
	fmt.Println(math.Pow(8,(1.0/3)))

	//正弦余弦正切
	fmt.Println(math.Sin((30.0/180)*math.Pi))
	fmt.Println(math.Cos((30.0/180)*math.Pi))
	fmt.Println(math.Tan((30.0/180)*math.Pi))

	//0.52，反正弦，求正弦为0.5的弧度,0.52，即math.Pi/6，即30度
	fmt.Println(math.Asin(0.5))
}

func Int22Float() {
	//把整数化成浮点数
	var a = 123
	b := float64(a)
	fmt.Printf("b的类型是%T,值是%v\n",b,b)

	//将浮点数化为整数
	var c = 123.45
	d := int32(c)
	fmt.Printf("d的类型是%T,值是%v\n",d,d)
}


/*自加运算*/
func SelfAutoMath() {
	var a int = 123
	//a = a + 2
	//a += 2
	//a -= 2
	//a *= 2
	//a /= 2
	//a %= 2
	//
	//a++
	a--
	fmt.Println(a)

}

