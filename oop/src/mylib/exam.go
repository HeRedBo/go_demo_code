package mylib

import (
	"fmt"
	"math/rand"
	"time"
)

//全班学生名单
var Names []string

/*包初始化方法*/
func init() {
	fmt.Println("mylib包初始化")
	Names = make([]string,0)
}

/*显示成绩做降序排名*/
func ShowRanking(scoreMap map[string]int) {
	/*形成名单，锁定姓名顺序才能才开始排序*/
	names := make([]string, 0)
	for name,_ := range scoreMap{
		names = append(names, name)
	}

	/*对names按分数的降序排列*/
	for i:=0;i<len(names)-1 ;i++  {
		//将剩余的同志中，最高分数的人调整到当前位置
		for j:=i+1;j<len(names);j++{
			//锁定第i位:将i后面的每个人（j）的分数与第i个人比较
			//如果第j个人的分数比第i个人高，就互换位置
			if scoreMap[names[j]] > scoreMap[names[i]]{
				names[i],names[j] = names[j],names[i]
			}
		}
	}

	/*按names的顺序输出姓名和分数*/
	for i,name := range names{
		fmt.Printf("第%3d名:%s \t %3d\n",i+1,name,scoreMap[name])
	}

}

/*给examers考试，为每个人生成随机成绩*/
func TakeExam(examers ...string) (scoreMap map[string]int) {
	scoreMap = make(map[string]int)
	for _,name := range examers{
		score := GetRandomInt(100)
		scoreMap[name] = score
	}
	return
}

/*生成[0,n]之间的随机数*/
func GetRandomInt(n int) int {
	time.Sleep(time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n + 1)
}

func AddMember(newcomers ...string)  {
	//fmt.Printf("newcomers的类型是%T\n",newcomers)//[]string
	Names = append(Names, newcomers...)
}
