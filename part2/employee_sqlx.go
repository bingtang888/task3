package part2

import (
	"github.com/jmoiron/sqlx"
)

// Employee 员工结构体，与数据库表字段对应
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// InitTable 初始化员工表结构
func InitTable(db *sqlx.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS employees (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		department VARCHAR(255) NOT NULL,
		salary DECIMAL(10, 2) NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`

	_, err := db.Exec(createTableSQL)
	return err
}

// CreateEmployee 新增员工（模拟GORM的Create方法）
func CreateEmployee(db *sqlx.DB, emp Employee) error {
	_, err := db.NamedExec(`
		INSERT INTO employees (name, department, salary)
		VALUES (:name, :department, :salary)`, emp)
	return err
}

// FindTechDepartmentEmployees 查询技术部所有员工
func FindTechDepartmentEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	return employees, err
}

// FindHighestSalaryEmployee 查询工资最高的员工
func FindHighestSalaryEmployee(db *sqlx.DB) (*Employee, error) {
	var employee Employee
	err := db.Get(&employee, "SELECT * FROM employees WHERE salary = (SELECT MAX(salary) FROM employees)")
	if err != nil {
		return nil, err
	}
	return &employee, nil
}
