package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gookit/goutil/dump"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// 全局 etcd 客户端（微服务中可封装为单例）
var etcdClient *clientv3.Client

// 初始化 etcd 客户端
func initEtcdClient() error {
	// 配置 etcd 连接（本地 Docker 地址）
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, // 你的 etcd 地址
		DialTimeout: 5 * time.Second,            // 连接超时
		// 若开启认证，需添加以下配置：
		// Username: "root",
		// Password: "123456",
	}

	// 创建客户端
	client, err := clientv3.New(config)
	if err != nil {
		return fmt.Errorf("创建etcd客户端失败: %v", err)
	}

	etcdClient = client
	log.Println("etcd 客户端初始化成功")
	return nil
}

// 场景1：基础KV操作（微服务配置读写）
func kvOperationDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 1. 写入配置（微服务中可存服务地址、数据库配置等）
	putResp, err := etcdClient.Put(ctx, "/micro/service/user/addr", "127.0.0.1:8080")
	dump.P(putResp, err)
	if err != nil {
		log.Printf("写入配置失败: %v", err)
		return
	}
	log.Printf("写入配置成功，版本号: %d", putResp.Header.Revision)

	// 2. 读取单个配置
	getResp, err := etcdClient.Get(ctx, "/micro/service/user/addr")
	if err != nil {
		log.Printf("读取配置失败: %v", err)
		return
	}
	for _, kv := range getResp.Kvs {
		log.Printf("读取配置：%s = %s", string(kv.Key), string(kv.Value))
	}

	// 3. 读取前缀匹配的配置（微服务中批量获取某类配置）
	prefixResp, err := etcdClient.Get(ctx, "/micro/service/", clientv3.WithPrefix())
	if err != nil {
		log.Printf("读取前缀配置失败: %v", err)
		return
	}
	log.Println("=== 前缀匹配的所有配置 ===")
	for _, kv := range prefixResp.Kvs {
		log.Printf("%s = %s", string(kv.Key), string(kv.Value))
	}

	// 4. 删除配置
	//delResp, err := etcdClient.Delete(ctx, "/micro/service/user/addr")
	//if err != nil {
	//	log.Printf("删除配置失败: %v", err)
	//	return
	//}
	//log.Printf("删除配置成功，删除数量: %d", delResp.Deleted)
}

// 场景2：监听配置变化（微服务配置热更新核心）
func watchConfigDemo() {
	// 监听 /micro/service/ 前缀的所有配置变化
	watchChan := etcdClient.Watch(context.Background(), "/micro/service/", clientv3.WithPrefix())

	log.Println("=== 开始监听配置变化（按Ctrl+C停止）===")
	// 阻塞监听
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			log.Printf(
				"配置变化：类型=%s, Key=%s, Value=%s",
				event.Type, // PUT/DELETE
				string(event.Kv.Key),
				string(event.Kv.Value),
			)
		}
	}
}

// 场景3：租约（Lease）- 微服务服务注册（临时节点自动下线）
func leaseDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 1. 创建 10 秒租约
	leaseResp, err := etcdClient.Grant(ctx, 10)
	if err != nil {
		log.Printf("创建租约失败: %v", err)
		return
	}
	leaseID := leaseResp.ID
	log.Printf("创建10秒租约成功，ID: %d", leaseID)

	// 2. 绑定租约写入临时节点（服务注册）
	_, err = etcdClient.Put(ctx, "/micro/service/order/addr", "127.0.0.1:8081", clientv3.WithLease(leaseID))
	if err != nil {
		log.Printf("写入临时节点失败: %v", err)
		return
	}
	log.Println("写入临时节点成功（10秒后自动删除，除非续租）")

	// 3. 续租（后台持续续租，模拟服务运行中）
	keepAliveChan, err := etcdClient.KeepAlive(context.Background(), leaseID)
	if err != nil {
		log.Printf("续租失败: %v", err)
		return
	}

	// 打印续租结果
	go func() {
		for keepAliveResp := range keepAliveChan {
			log.Printf("续租成功，剩余时间: %d秒", keepAliveResp.TTL)
		}
	}()

	// 模拟服务运行5秒后停止续租
	time.Sleep(5 * time.Second)
	// 关闭续租
	etcdClient.Revoke(context.Background(), leaseID)
	log.Println("停止续租，临时节点将在10秒内删除")
	// 等待验证
	time.Sleep(10 * time.Second)
}

// 场景4：分布式锁（微服务并发控制）

func main() {
	// 初始化etcd客户端
	if err := initEtcdClient(); err != nil {
		log.Fatal(err)
	}
	defer etcdClient.Close() // 程序退出时关闭客户端

	// 依次运行不同场景（注释/取消注释切换）
	// 场景1：基础KV操作
	kvOperationDemo()

	// 场景2：监听配置变化（运行后可在etcdctl中修改配置，观察控制台输出）
	// watchConfigDemo()

	// 场景3：租约（临时节点）
	// leaseDemo()

	// 场景4：分布式锁
	//lockDemo()
}
