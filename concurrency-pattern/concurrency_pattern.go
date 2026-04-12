package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// region 生产者 - 消费者模式（最基础）
// 核心
//一个 / 多个协程生产数据，一个 / 多个协程消费数据，Channel 做通信。
//适用场景
//日志收集、数据处理、任务生产

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i // 生产数据
	}
	close(ch) // 生产完毕关闭通道
}

func consumer(ch chan int) {
	for v := range ch { // 消费数据
		fmt.Println("消费:", v)
	}
}

// endregion

// region worker 池模式（控制并发数）
/*
核心
	固定数量的协程处理任务，防止协程爆炸（你写的调度就是它的升级版）。
适用场景
	爬虫、接口请求、批量任务
*/
// 固定3个worker
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker%d 处理任务%d\n", id, j)
		time.Sleep(100)
		results <- j * 2
	}
}

// endregion

// region  select 多路复用模式
/*
核心
	同时监听多个 Channel，哪个就绪执行哪个，非阻塞调度。
适用场景
	任务调度、超时控制、多事件监听
*/

func mutChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "任务1完成"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "任务2完成"
	}()

	// 同时监听两个通道
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

// endregion

// region  上下文（Context）模式（取消 / 超时）
func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // 收到取消信号
			fmt.Println("任务停止")
			return
		default:
			fmt.Println("任务执行中...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// endregion

// region 5 信号量模式（控制并发度）

// 核心
// 用带缓冲 Channel做令牌，控制同时运行的协程数量。
// 适用场景
// 限流、接口并发控制
func signChannel() {
	// 最多同时运行2个任务（令牌数=2）
	sem := make(chan struct{}, 2)

	for i := 1; i <= 5; i++ {
		sem <- struct{}{} // 获取令牌
		go func(n int) {
			defer func() { <-sem }() // 释放令牌
			fmt.Printf("任务%d 执行\n", n)
			time.Sleep(500 * time.Millisecond)
		}(i)
	}

	// 等待所有任务完成
	time.Sleep(3 * time.Second)
}

// endregion

// region 6 等待组（WaitGroup）模式（等待协程完成）
func Worker(n int, wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成，计数器-1
	fmt.Printf("任务%d 执行\n", n)
	time.Sleep(200 * time.Millisecond)

}

func waitGroup() {
	var wg sync.WaitGroup

	wg.Add(3) // 设置等待3个协程
	for i := 1; i <= 3; i++ {
		go Worker(i, &wg)
	}

	wg.Wait() // 阻塞直到计数器为0
	fmt.Println("所有任务完成")
}

// endregion

// region   7 单例模式（sync.Once）

var once sync.Once
var conn string

func initDB() {
	conn = "数据库连接成功"
	fmt.Println("初始化数据库（只执行一次）")
}

func getConn() string {
	once.Do(initDB) // 保证只执行一次
	return conn
}

func syncOnce() {
	var wg sync.WaitGroup
	wg.Add(3)

	// 3个协程同时调用，只初始化一次
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(getConn())
		}()
	}
	wg.Wait()
}

//endregion

// region 管道（Pipeline）模式（数据流式处理）
/*
核心
	数据一步一步流式处理，协程串联，前一个输出是后一个输入。
适用场景
	数据清洗、日志解析、流式计算
*/

// 生成数据
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// 平方处理
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func pipelineDeme() {
	// 管道：生成 -> 平方 -> 输出
	c := gen(1, 2, 3)
	out := square(c)

	for v := range out {
		fmt.Println(v) // 1,4,9
	}
}

// endregion

// region 9. 扇入 / 扇出模式（并行聚合）
/*
核心
	扇出：一个任务分发给多个协程并行执行
	扇入：多个协程结果聚合到一个通道
适用场景
	并发请求多接口、并行计算、结果汇总
*/

// 扇出：多个worker并行执行
//
//	worker：从输入读，处理后写入自己的输出通道
func chanWorker(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		fmt.Printf("worker %d 处理: %d\n", id, n)
		out <- n * 2
	}
}

// 扇入 merge：把多个 chan 汇聚成一个（官方标准写法）
func merge(chs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	// 每个channel启动一个goroutine转发
	for _, ch := range chs {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	// 全部转发完 → 关闭out
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func inOutDemo() {
	const workerCount = 2 // 扇出：2个worker

	// 1. 任务通道
	jobs := make(chan int, 5)

	// 2. 每个worker一个独立输出channel（扇出核心）
	outChs := make([]chan int, workerCount)
	for i := 0; i < workerCount; i++ {
		outChs[i] = make(chan int)
	}

	// 3. 启动worker（扇出）
	var wg sync.WaitGroup
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go chanWorker(i, jobs, outChs[i], &wg)
	}

	// 4. 发送任务
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
		close(jobs) // 任务发完关闭
	}()

	// 5. 等待所有worker结束 → 关闭它们的输出channel
	go func() {
		wg.Wait()
		for _, ch := range outChs {
			close(ch)
		}
	}()

	// 6. 扇入：合并所有结果
	finalResult := merge(outChs...)

	// 7. 输出最终结果
	fmt.Println("\n==== 最终结果 ====")
	for val := range finalResult {
		fmt.Println("结果:", val)
	}
}

//endregion

func main() {
	// 信号量模式（控制并发度）
	//signChannel()

	//waitGroup()
	//syncOnce()
	//pipelineDeme()
	inOutDemo()

	// select 多路复用模式
	//mutChannel()

	// 创建1秒超时的context
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	//
	//go work(ctx)
	//
	//// 等待超时
	//<-ctx.Done()
	//fmt.Println("主程序：任务超时退出")

	return

	jobs := make(chan int, 5)
	results := make(chan int, 5)
	// 启动3个worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 发送5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
