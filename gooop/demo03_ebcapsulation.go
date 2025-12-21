package gooop

/**
* 封装（Encapsulation）
 */

// 私有字段（小写开头）
type account struct {
	owner   string
	balance float64
}

// 公开构造函数
func NewAccount(owner string) *account {
	return &account{
		owner:   owner,
		balance: 0,
	}
}

// 公开方法（大写开头）
func (a *account) Deposit(amount float64) {
	if amount > 0 {
		a.balance += amount
	}
}

func (a *account) GetBalance() float64 {
	return a.balance
}
