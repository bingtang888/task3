package main

import (
	"errors"
	"log"
	"task3/part1"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 配置MySQL连接（请替换为你的实际数据库信息）
	dsn := "root:1234@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 创建学生服务实例
	//studentService := part1.NewStudentService(db)

	// 初始化表结构（自动迁移）
	//if err := studentService.InitTable(); err != nil {
	//	log.Fatalf("初始化表结构失败: %v", err)
	//}

	// 1. 新增学生：张三，20岁，三年级
	//zhangSan := &part1.Student{
	//	Name:  "张三",
	//	Age:   20,
	//	Grade: "三年级",
	//}
	//if err := studentService.CreateStudent(zhangSan); err != nil {
	//	log.Printf("新增学生失败: %v", err)
	//} else {
	//	log.Printf("成功新增学生，ID: %d", zhangSan.ID)
	//}

	// 2. 查询年龄大于18岁的学生
	//adults, err := studentService.FindAdults()
	//if err != nil {
	//	log.Printf("查询成年学生失败: %v", err)
	//} else {
	//	fmt.Printf("年龄大于18岁的学生共有 %d 名:\n", len(adults))
	//	for _, student := range adults {
	//		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d, 年级: %s\n",
	//			student.ID, student.Name, student.Age, student.Grade)
	//	}
	//}
	//
	//// 3. 更新张三的年级为四年级
	//rowsAffected, err := studentService.UpdateGradeByName("张三", "四年级")
	//if err != nil {
	//	log.Printf("更新年级失败: %v", err)
	//} else {
	//	log.Printf("成功更新 %d 条记录", rowsAffected)
	//}

	//// 4. 删除年龄小于15岁的学生
	//deletedCount, err := studentService.DeleteByAgeLessThan(15)
	//if err != nil {
	//	log.Printf("删除学生失败: %v", err)
	//} else {
	//	log.Printf("成功删除 %d 条记录", deletedCount)
	//}

	// 创建交易服务实例
	transactionService := part1.NewTransactionService(db)

	// 初始化表结构（自动迁移）
	//if err := transactionService.InitTables(); err != nil {
	//	log.Fatalf("初始化账户和交易表失败: %v", err)
	//}

	// 创建测试账户A和B
	accountAID, err := transactionService.CreateAccount(500) // A账户初始500元
	if err != nil {
		log.Fatalf("创建账户A失败: %v", err)
	}
	accountBID, err := transactionService.CreateAccount(300) // B账户初始300元
	if err != nil {
		log.Fatalf("创建账户B失败: %v", err)
	}

	log.Printf("初始状态:")
	balanceA, _ := transactionService.GetAccountBalance(accountAID)
	balanceB, _ := transactionService.GetAccountBalance(accountBID)
	log.Printf("账户A余额: %.2f 元", balanceA)
	log.Printf("账户B余额: %.2f 元", balanceB)

	// 执行转账操作：从A向B转账100元
	log.Println("开始转账100元...")
	err = transactionService.Transfer(accountAID, accountBID, 100)
	if err != nil {
		if err == errors.New("insufficient balance") {
			log.Println("转账失败: 账户余额不足")
		} else {
			log.Printf("转账失败: %v", err)
		}
		return
	}

	// 查询转账后余额
	log.Printf("转账后状态:")
	balanceA, _ = transactionService.GetAccountBalance(accountAID)
	balanceB, _ = transactionService.GetAccountBalance(accountBID)
	log.Printf("账户A余额: %.2f 元", balanceA)
	log.Printf("账户B余额: %.2f 元", balanceB)
	log.Println("转账成功！")

}
