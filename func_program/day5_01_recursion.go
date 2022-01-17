package func_program

import (
	"fmt"
	"time"
)





/*1+2+...+n的和——循环实现*/
func GetSum(n int) (sum int)  {
	for i:=1; i<=n; i++ {
		sum += i
	}
	return sum
}

/*1+2+...+n的和——递归实现*/
func ReF(n int) (sum int) {
	//终止条件：由递转归
	if n == 1{
		return 1
	}
	//自己调用自己
	return n + ReF(n-1)
}

// 测试执行函数
func TimeIt(f func(int) int, arg int) {
	// 记录开始时间
	startTime := time.Now()
	// 调用 f 函数 ，丢入器参数 arg
	f(arg)
	endTime := time.Now()
	//求出endTime和startTime的时间差——函数f的执行时间
	fmt.Println("执行耗时=",endTime.Sub(startTime))
}
