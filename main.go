package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"task3/part3"
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
//func main() {
//	// 数据库配置
//	dbConfig := mysql.Config{
//		User:   "root",
//		Passwd: "1234",
//		Net:    "tcp",
//		Addr:   "127.0.0.1:3306",
//		DBName: "gorm",
//		Params: map[string]string{
//			"charset":   "utf8mb4",
//			"parseTime": "True",
//			"loc":       "Local",
//		},
//	}
//	dsn := dbConfig.FormatDSN()
//
//	// 初始化part2（SQLx）相关功能
//	sqlxDB, err := sqlx.Connect("mysql", dsn)
//	if err != nil {
//		log.Fatalf("SQLx数据库连接失败: %v", err)
//	}
//	defer sqlxDB.Close()
//	//第二部分题目一----------------------------------------------------------------------------------------
//
//	//// 初始化员工表（类似AutoMigrate的效果）
//	//if err := InitTable(sqlxDB); err != nil {
//	//	log.Fatalf("初始化员工表失败: %v", err)
//	//}
//
//	// 添加测试数据（类似GORM的Create）
//	//_ = CreateEmployee(sqlxDB, Employee{Name: "张三", Department: "技术部", Salary: 9000})
//	//_ = CreateEmployee(sqlxDB, Employee{Name: "李四", Department: "技术部", Salary: 12000})
//	//_ = CreateEmployee(sqlxDB, Employee{Name: "王五", Department: "市场部", Salary: 8000})
//	//_ = CreateEmployee(sqlxDB, Employee{Name: "赵六", Department: "技术部", Salary: 15000})
//
//	// 1. 查询技术部所有员工
//	//fmt.Println("----- 技术部员工列表 -----")
//	//techEmployees, err := part2.FindTechDepartmentEmployees(sqlxDB)
//	//if err != nil {
//	//	log.Printf("查询技术部员工失败: %v", err)
//	//} else {
//	//	for _, emp := range techEmployees {
//	//		fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n",
//	//			emp.ID, emp.Name, emp.Department, emp.Salary)
//	//	}
//	//}
//	//
//	//// 2. 查询工资最高的员工
//	//fmt.Println("\n----- 工资最高的员工 -----")
//	//topEmp, err := part2.FindHighestSalaryEmployee(sqlxDB)
//	//if err != nil {
//	//	log.Printf("查询最高工资员工失败: %v", err)
//	//} else {
//	//	fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n",
//	//		topEmp.ID, topEmp.Name, topEmp.Department, topEmp.Salary)
//	//}
//
//	//第二部分题目二---------------------------------------------------------------------------------------------
//
//	// 初始化书籍表（首次运行时需要）
//	if err := part2.InitBookTable(sqlxDB); err != nil {
//		log.Fatalf("初始化书籍表失败: %v", err)
//	}
//
//	// 添加测试书籍数据
//	_ = part2.CreateBook(sqlxDB, part2.Book{Title: "Go编程实战", Author: "张三", Price: 69.9})
//	_ = part2.CreateBook(sqlxDB, part2.Book{Title: "SQL入门", Author: "李四", Price: 45.5})
//	_ = part2.CreateBook(sqlxDB, part2.Book{Title: "算法导论", Author: "王五", Price: 89.0})
//	_ = part2.CreateBook(sqlxDB, part2.Book{Title: "数据结构", Author: "赵六", Price: 59.9})
//
//	// 题目2核心功能：查询价格大于50元的书籍
//	fmt.Println("----- 价格大于50元的书籍列表 -----")
//	expensiveBooks, err := part2.FindBooksByPriceGreaterThan(sqlxDB, 50)
//	if err != nil {
//		log.Printf("查询高价书籍失败: %v", err)
//	} else {
//		for _, book := range expensiveBooks {
//			fmt.Printf("ID: %d, 标题: %s, 作者: %s, 价格: %.2f\n",
//				book.ID, book.Title, book.Author, book.Price)
//		}
//	}
//}

// 第三部分main方法
func main() {
	// 1. 数据库连接配置
	//dsn := "root:1234@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatalf("数据库连接失败: %v", err)
	//}

	// 2. 初始化表结构+添加测试数据
	//if err := part3.InitBlogTables(db); err != nil {
	//	log.Fatalf("初始化博客表失败: %v", err)
	//}
	//if err := part3.CreateTestData(db); err != nil {
	//	log.Fatalf("添加测试数据失败: %v", err)
	//}
	//log.Println("✅ 表初始化及测试数据添加完成")

	// 3. 需求1：查询指定用户（ID=1）的所有文章及评论
	//user, err := part3.GetUserWithPostsAndComments(db, 3)
	//if err != nil {
	//	log.Printf("❌ 查询用户文章评论失败: %v", err)
	//} else {
	//	fmt.Printf("\n📌 用户信息：ID=%d, 用户名=%s\n", user.ID, user.Username)
	//	fmt.Printf("📝 该用户共发布 %d 篇文章：\n", len(user.Posts))
	//	for _, post := range user.Posts {
	//		fmt.Printf("  - 文章ID=%d, 标题=%s\n", post.ID, post.Title)
	//		fmt.Printf("    文章内容：%s\n", post.Content)
	//		fmt.Printf("    评论数：%d 条\n", len(post.Comments))
	//		for i, comment := range post.Comments {
	//			fmt.Printf("      %d. 评论：%s\n", i+1, comment.Content)
	//		}
	//	}
	//}
	//
	// 4. 需求2：查询评论数量最多的文章
	//topPost, err := part3.GetPostWithMostComments(db)
	//if err != nil {
	//	log.Printf("❌ 查询评论最多文章失败: %v", err)
	//} else {
	//	fmt.Printf("\n🔥 评论最多的文章：\n")
	//	fmt.Printf("  文章ID：%d\n", topPost.ID)
	//	fmt.Printf("  标题：%s\n", topPost.Title)
	//	fmt.Printf("  作者：%s\n", topPost.User.Username)
	//	fmt.Printf("  评论数：%d 条\n", len(topPost.Comments))
	//	fmt.Printf("  热门评论：%s\n", topPost.Comments[0].Content)
	//}

	//第三部分作业二----------------------------------------------------------------------------------------------------
	// 数据库连接配置（使用新表，避免与旧表冲突）
	// 1. 数据库连接（使用新库 gorm_hooks，避免旧表冲突）
	dsn := "root:1234@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 2. 创建新表（明确使用 User_hook、Post_hook、Comment，GORM 会按模型的 table 标签生成表名）
	//if err := db.AutoMigrate(
	//	&part3.User_hook{},     // 生成表：user_hooks
	//	&part3.Post_hook{},     // 生成表：post_hooks
	//	&part3.Comment_hooks{}, // 生成表：comment_hooks（模型中已指定）
	//); err != nil {
	//	log.Fatalf("创建新表失败: %v", err)
	//}
	//log.Println("✅ 新表创建成功（user_hooks、post_hooks、comment_hooks）")

	// 3. 测试钩子函数（使用新模型）
	//testPostCreateHook(db) // 测试文章创建钩子
	testCommentDeleteHook(db) // 测试评论删除钩子
}

// 测试 Post_hook 的 BeforeCreate 钩子（创建文章时更新 User_hook 的 post_count）
func testPostCreateHook(db *gorm.DB) {
	fmt.Println("\n----- 测试文章创建钩子函数 -----")

	// 步骤1：创建测试用户（User_hook 模型）
	testUser := part3.User_hook{Username: "hook_test_user_001"} // 用户名加后缀，避免重复
	if err := db.Create(&testUser).Error; err != nil {
		log.Printf("创建用户失败: %v", err)
		return
	}
	fmt.Printf("初始用户信息: ID=%d, 用户名=%s, 文章数量=%d\n",
		testUser.ID, testUser.Username, testUser.PostCount) // 预期：文章数量=0

	// 步骤2：创建文章（Post_hook 模型，触发 BeforeCreate 钩子）
	testPost := part3.Post_hook{
		Title:   "测试钩子：文章创建",
		Content: "创建这篇文章会自动增加用户的文章数量",
		UserID:  testUser.ID, // 关联上面创建的用户
	}
	if err := db.Create(&testPost).Error; err != nil {
		log.Printf("创建文章失败: %v", err)
		return
	}

	// 步骤3：重新查询用户，验证 post_count 是否+1
	var updatedUser part3.User_hook
	if err := db.First(&updatedUser, testUser.ID).Error; err != nil {
		log.Printf("查询更新后用户失败: %v", err)
		return
	}
	fmt.Printf("创建文章后用户信息: ID=%d, 用户名=%s, 文章数量=%d\n",
		updatedUser.ID, updatedUser.Username, updatedUser.PostCount) // 预期：文章数量=1
}

// 测试 Comment 的 AfterDelete 钩子（删除最后一条评论时更新 Post_hook 的 comment_status）
func testCommentDeleteHook(db *gorm.DB) {
	fmt.Println("\n----- 测试评论删除钩子函数 -----")

	// 步骤1：先获取一个测试用户（如果没有则创建）
	var testUser part3.User_hook
	if err := db.Where("username = ?", "hook_test_user_001").First(&testUser).Error; err != nil {
		// 若之前的用户不存在，新建一个
		testUser = part3.User_hook{Username: "hook_test_user_002"}
		db.Create(&testUser)
	}

	// 步骤2：创建测试文章（Post_hook 模型，初始评论状态为"有评论"）
	testPost := part3.Post_hook{
		Title:   "测试钩子：评论删除",
		Content: "删除这篇文章的最后一条评论会更新状态为'无评论'",
		UserID:  testUser.ID,
	}
	if err := db.Create(&testPost).Error; err != nil {
		log.Printf("创建测试文章失败: %v", err)
		return
	}
	fmt.Printf("初始文章状态: ID=%d, 标题=%s, 评论状态=%s\n",
		testPost.ID, testPost.Title, testPost.CommentStatus) // 预期：评论状态="有评论"

	// 步骤3：给文章添加一条评论（Comment 模型）
	testComment := part3.Comment_hooks{
		Content: "这是测试评论，删除后会触发钩子",
		PostID:  testPost.ID, // 关联上面创建的文章
	}
	if err := db.Create(&testComment).Error; err != nil {
		log.Printf("创建测试评论失败: %v", err)
		return
	}

	// 步骤4：删除这条评论（触发 AfterDelete 钩子）
	if err := db.Delete(&testComment).Error; err != nil {
		log.Printf("删除评论失败: %v", err)
		return
	}

	// 步骤5：重新查询文章，验证 comment_status 是否更新为"无评论"
	var updatedPost part3.Post_hook
	if err := db.First(&updatedPost, testPost.ID).Error; err != nil {
		log.Printf("查询更新后文章失败: %v", err)
		return
	}
	fmt.Printf("删除最后一条评论后文章状态: ID=%d, 标题=%s, 评论状态=%s\n",
		updatedPost.ID, updatedPost.Title, updatedPost.CommentStatus) // 预期：评论状态="无评论"
}
