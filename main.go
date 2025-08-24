package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"task3/part2"
)

//第一部分main方法
//func main() {
//	// 配置MySQL连接（请替换为你的实际数据库信息）
//	dsn := "root:1234@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
//
//	// 连接数据库
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatalf("数据库连接失败: %v", err)
//	}
//
//	// 创建学生服务实例
//	//studentService := part1.NewStudentService(db)
//
//	// 初始化表结构（自动迁移）
//	//if err := studentService.InitTable(); err != nil {
//	//	log.Fatalf("初始化表结构失败: %v", err)
//	//}
//
//	// 1. 新增学生：张三，20岁，三年级
//	//zhangSan := &part1.Student{
//	//	Name:  "张三",
//	//	Age:   20,
//	//	Grade: "三年级",
//	//}
//	//if err := studentService.CreateStudent(zhangSan); err != nil {
//	//	log.Printf("新增学生失败: %v", err)
//	//} else {
//	//	log.Printf("成功新增学生，ID: %d", zhangSan.ID)
//	//}
//
//	// 2. 查询年龄大于18岁的学生
//	//adults, err := studentService.FindAdults()
//	//if err != nil {
//	//	log.Printf("查询成年学生失败: %v", err)
//	//} else {
//	//	fmt.Printf("年龄大于18岁的学生共有 %d 名:\n", len(adults))
//	//	for _, student := range adults {
//	//		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d, 年级: %s\n",
//	//			student.ID, student.Name, student.Age, student.Grade)
//	//	}
//	//}
//	//
//	//// 3. 更新张三的年级为四年级
//	//rowsAffected, err := studentService.UpdateGradeByName("张三", "四年级")
//	//if err != nil {
//	//	log.Printf("更新年级失败: %v", err)
//	//} else {
//	//	log.Printf("成功更新 %d 条记录", rowsAffected)
//	//}
//
//	//// 4. 删除年龄小于15岁的学生
//	//deletedCount, err := studentService.DeleteByAgeLessThan(15)
//	//if err != nil {
//	//	log.Printf("删除学生失败: %v", err)
//	//} else {
//	//	log.Printf("成功删除 %d 条记录", deletedCount)
//	//}
//
//	// 创建交易服务实例
//	transactionService := part1.NewTransactionService(db)
//
//	// 初始化表结构（自动迁移）
//	//if err := transactionService.InitTables(); err != nil {
//	//	log.Fatalf("初始化账户和交易表失败: %v", err)
//	//}
//
//	// 创建测试账户A和B
//	accountAID, err := transactionService.CreateAccount(500) // A账户初始500元
//	if err != nil {
//		log.Fatalf("创建账户A失败: %v", err)
//	}
//	accountBID, err := transactionService.CreateAccount(300) // B账户初始300元
//	if err != nil {
//		log.Fatalf("创建账户B失败: %v", err)
//	}
//
//	log.Printf("初始状态:")
//	balanceA, _ := transactionService.GetAccountBalance(accountAID)
//	balanceB, _ := transactionService.GetAccountBalance(accountBID)
//	log.Printf("账户A余额: %.2f 元", balanceA)
//	log.Printf("账户B余额: %.2f 元", balanceB)
//
//	// 执行转账操作：从A向B转账100元
//	log.Println("开始转账100元...")
//	err = transactionService.Transfer(accountAID, accountBID, 100)
//	if err != nil {
//		if err == errors.New("insufficient balance") {
//			log.Println("转账失败: 账户余额不足")
//		} else {
//			log.Printf("转账失败: %v", err)
//		}
//		return
//	}
//
//	// 查询转账后余额
//	log.Printf("转账后状态:")
//	balanceA, _ = transactionService.GetAccountBalance(accountAID)
//	balanceB, _ = transactionService.GetAccountBalance(accountBID)
//	log.Printf("账户A余额: %.2f 元", balanceA)
//	log.Printf("账户B余额: %.2f 元", balanceB)
//	log.Println("转账成功！")
//
//}

// 第二部门main方法
func main() {
	// 数据库配置
	dbConfig := mysql.Config{
		User:   "root",
		Passwd: "1234",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "gorm",
		Params: map[string]string{
			"charset":   "utf8mb4",
			"parseTime": "True",
			"loc":       "Local",
		},
	}
	dsn := dbConfig.FormatDSN()

	// 初始化part2（SQLx）相关功能
	sqlxDB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("SQLx数据库连接失败: %v", err)
	}
	defer sqlxDB.Close()
	//第二部分题目一----------------------------------------------------------------------------------------

	//// 初始化员工表（类似AutoMigrate的效果）
	//if err := InitTable(sqlxDB); err != nil {
	//	log.Fatalf("初始化员工表失败: %v", err)
	//}

	// 添加测试数据（类似GORM的Create）
	//_ = CreateEmployee(sqlxDB, Employee{Name: "张三", Department: "技术部", Salary: 9000})
	//_ = CreateEmployee(sqlxDB, Employee{Name: "李四", Department: "技术部", Salary: 12000})
	//_ = CreateEmployee(sqlxDB, Employee{Name: "王五", Department: "市场部", Salary: 8000})
	//_ = CreateEmployee(sqlxDB, Employee{Name: "赵六", Department: "技术部", Salary: 15000})

	// 1. 查询技术部所有员工
	//fmt.Println("----- 技术部员工列表 -----")
	//techEmployees, err := part2.FindTechDepartmentEmployees(sqlxDB)
	//if err != nil {
	//	log.Printf("查询技术部员工失败: %v", err)
	//} else {
	//	for _, emp := range techEmployees {
	//		fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n",
	//			emp.ID, emp.Name, emp.Department, emp.Salary)
	//	}
	//}
	//
	//// 2. 查询工资最高的员工
	//fmt.Println("\n----- 工资最高的员工 -----")
	//topEmp, err := part2.FindHighestSalaryEmployee(sqlxDB)
	//if err != nil {
	//	log.Printf("查询最高工资员工失败: %v", err)
	//} else {
	//	fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n",
	//		topEmp.ID, topEmp.Name, topEmp.Department, topEmp.Salary)
	//}

	//第二部分题目二---------------------------------------------------------------------------------------------

	// 初始化书籍表（首次运行时需要）
	if err := part2.InitBookTable(sqlxDB); err != nil {
		log.Fatalf("初始化书籍表失败: %v", err)
	}

	// 添加测试书籍数据
	_ = part2.CreateBook(sqlxDB, part2.Book{Title: "Go编程实战", Author: "张三", Price: 69.9})
	_ = part2.CreateBook(sqlxDB, part2.Book{Title: "SQL入门", Author: "李四", Price: 45.5})
	_ = part2.CreateBook(sqlxDB, part2.Book{Title: "算法导论", Author: "王五", Price: 89.0})
	_ = part2.CreateBook(sqlxDB, part2.Book{Title: "数据结构", Author: "赵六", Price: 59.9})

	// 题目2核心功能：查询价格大于50元的书籍
	fmt.Println("----- 价格大于50元的书籍列表 -----")
	expensiveBooks, err := part2.FindBooksByPriceGreaterThan(sqlxDB, 50)
	if err != nil {
		log.Printf("查询高价书籍失败: %v", err)
	} else {
		for _, book := range expensiveBooks {
			fmt.Printf("ID: %d, 标题: %s, 作者: %s, 价格: %.2f\n",
				book.ID, book.Title, book.Author, book.Price)
		}
	}
}
