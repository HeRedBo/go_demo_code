package breaker

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/eapache/go-resiliency/breaker"
)

// 模拟一个可能失败的外部服务调用
func callExternalService() error {
	// 模拟 60% 失败率
	if time.Now().Unix()%10 < 6 {
		return errors.New("service unavailable")
	}
	fmt.Println("Call succeeded!")
	return nil
}

func BreakMain() {
	// 创建断路器：
	// - 统计最近 10 次调用
	// - 如果失败率 >= 60%，则熔断
	// - 熔断后 10 秒进入半开状态
	cb := breaker.New(10, 6, 10*time.Second)

	for i := 0; i < 30; i++ {
		err := cb.Run(func() error {
			return callExternalService()
		})

		if err != nil {
			if errors.Is(err, breaker.ErrBreakerOpen) {
				log.Println("Circuit breaker is OPEN, skipping call")
			} else {
				log.Printf("Call failed: %v", err)
			}
		}

		time.Sleep(500 * time.Millisecond)
	}
}
