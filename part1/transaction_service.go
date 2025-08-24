package part1

import (
	"errors"
	"gorm.io/gorm"
)

// Account 账户表结构体
type Account struct {
	ID      uint    `gorm:"primaryKey"`
	Balance float64 `gorm:"not null"` // 账户余额
}

// Transaction 交易记录表结构体
type Transaction struct {
	ID            uint    `gorm:"primaryKey"`
	FromAccountID uint    `gorm:"not null"` // 转出账户ID
	ToAccountID   uint    `gorm:"not null"` // 转入账户ID
	Amount        float64 `gorm:"not null"` // 转账金额
}

// TransactionService 事务服务结构体
type TransactionService struct {
	db *gorm.DB
}

// NewTransactionService 创建事务服务实例
func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{db: db}
}

// InitTables 初始化账户表和交易记录表
func (s *TransactionService) InitTables() error {
	// 自动迁移创建表结构
	if err := s.db.AutoMigrate(&Account{}, &Transaction{}); err != nil {
		return err
	}
	return nil
}

// Transfer 实现从A账户向B账户转账的事务操作
func (s *TransactionService) Transfer(fromAccountID, toAccountID uint, amount float64) error {
	// 开始事务
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 1. 查询转出账户信息
		var fromAccount Account
		if err := tx.First(&fromAccount, fromAccountID).Error; err != nil {
			return err // 查询失败，回滚事务
		}

		// 2. 检查余额是否充足
		if fromAccount.Balance < amount {
			return errors.New("insufficient balance") // 余额不足，回滚事务（使用预定义错误类型）
		}

		// 3. 扣减转出账户余额
		if err := tx.Model(&Account{}).
			Where("id = ?", fromAccountID).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}

		// 4. 增加转入账户余额
		if err := tx.Model(&Account{}).
			Where("id = ?", toAccountID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		// 5. 记录交易信息
		transaction := Transaction{
			FromAccountID: fromAccountID,
			ToAccountID:   toAccountID,
			Amount:        amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		// 所有操作成功，事务将自动提交
		return nil
	})
}

// CreateAccount 辅助方法：创建账户
func (s *TransactionService) CreateAccount(balance float64) (uint, error) {
	account := Account{Balance: balance}
	if err := s.db.Create(&account).Error; err != nil {
		return 0, err
	}
	return account.ID, nil
}

// GetAccountBalance 辅助方法：查询账户余额
func (s *TransactionService) GetAccountBalance(accountID uint) (float64, error) {
	var account Account
	if err := s.db.First(&account, accountID).Error; err != nil {
		return 0, err
	}
	return account.Balance, nil
}
