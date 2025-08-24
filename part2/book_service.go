package part2

import (
	"github.com/jmoiron/sqlx"
)

// Book 书籍结构体，与books表字段对应
type Book struct {
	ID     int     `db:"id"`     // 书籍ID，对应表中id字段
	Title  string  `db:"title"`  // 书籍标题，对应表中title字段
	Author string  `db:"author"` // 作者，对应表中author字段
	Price  float64 `db:"price"`  // 价格，对应表中price字段
}

// InitBookTable 初始化books表结构（类似AutoMigrate）
func InitBookTable(db *sqlx.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL COMMENT '书籍标题',
		author VARCHAR(255) NOT NULL COMMENT '作者',
		price DECIMAL(10, 2) NOT NULL COMMENT '价格'
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='书籍表';`

	_, err := db.Exec(createTableSQL)
	return err
}

// CreateBook 新增书籍（模拟GORM的Create方法）
func CreateBook(db *sqlx.DB, book Book) error {
	_, err := db.NamedExec(`
		INSERT INTO books (title, author, price)
		VALUES (:title, :author, :price)`, book)
	return err
}

// FindBooksByPriceGreaterThan 查询价格大于指定值的书籍
func FindBooksByPriceGreaterThan(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book
	// 使用SQLx的Select方法实现类型安全映射
	err := db.Select(&books, "SELECT * FROM books WHERE price > ?", minPrice)
	return books, err
}
