package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/goutil/dump"
	"log"
)

func main() {
	// 1. 从文件加载模型和策略（适合初始化或简单场景）
	e, err := casbin.NewEnforcer("casbindemo/rbac_model.conf", "casbindemo/policy.csv")
	if err != nil {
		dump.Print("初始化Enforcer失败: %v", err)
		log.Fatalf("初始化Enforcer失败: %v", err)
	}

	// 2. 更常见的做法：使用数据库适配器（如MySQL）
	// adapter, _ := gormadapter.NewAdapter("mysql", "mysql_conn_string")
	// e, _ := casbin.NewEnforcer("path/to/rbac_model.conf", adapter)

	// 3. 权限检查
	sub, obj, act := "alice", "/api/user", "GET"
	obj = "/api/profile"
	// 为用户添加角色
	//_, err = e.AddGroupingPolicy("alice", "user")
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		dump.Println(err)
		log.Fatal(err)
	}
	if ok {
		dump.Println("允许访问")
		fmt.Println("允许访问")
	} else {
		dump.Println("拒绝访问")
		fmt.Println("拒绝访问")
	}

	// 4. 管理角色和策略（通常在管理后台调用）
	// 为用户添加角色
	//e.AddGroupingPolicy("charlie", "user")
	// 为角色添加权限
	//e.AddPolicy("user", "/api/order", "POST")
	// 保存策略回持久层（如果使用了适配器）
	// e.SavePolicy()
}
