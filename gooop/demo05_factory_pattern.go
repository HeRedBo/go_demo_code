package gooop

// 产品接口
type Database interface {
	Connect() string
}

// 具体产品
type MySQL struct{}

func (m MySQL) Connect() string { return "MySQL connected" }

type PostgreSQL struct{}

func (p PostgreSQL) Connect() string { return "PostgreSQL connected" }

// 工厂
func CreateDatabase(dbType string) Database {
	switch dbType {
	case "mysql":
		return MySQL{}
	case "postgres":
		return PostgreSQL{}
	default:
		return nil
	}
}
