package main

import (
	"fmt"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    float64
}

type EmployeeDB struct {
	employees map[int]Employee
}

func (db *EmployeeDB) addEmployee(emp Employee) {
	db.employees[emp.ID] = emp
	fmt.Printf("Nhân viên %s %s đã được thêm vào hệ thống.\n", emp.FirstName, emp.LastName)
}

func (db *EmployeeDB) updateEmployee(id int, emp Employee) {
	if _, ok := db.employees[id]; ok {
		db.employees[id] = emp
		fmt.Printf("Thông tin của nhân viên có ID %d đã được cập nhật.\n", id)
	} else {
		fmt.Printf("Không tìm thấy nhân viên có ID %d trong hệ thống.\n", id)
	}
}

func (db *EmployeeDB) deleteEmployee(id int) {
	if _, ok := db.employees[id]; ok {
		delete(db.employees, id)
		fmt.Printf("Nhân viên có ID %d đã được xóa khỏi hệ thống.\n", id)
	} else {
		fmt.Printf("Không tìm thấy nhân viên có ID %d trong hệ thống.\n", id)
	}
}

func (db EmployeeDB) showEmployees() {
	if len(db.employees) == 0 {
		fmt.Println("Không có nhân viên nào trong hệ thống.")
		return
	}
	fmt.Println("Danh sách nhân viên:")
	for id, emp := range db.employees {
		fmt.Printf("ID: %d, Họ và tên: %s %s, Chức vụ: %s, Lương: %.2f\n", id, emp.FirstName, emp.LastName, emp.Position, emp.Salary)
	}
}

func main() {
	var myEmployeeDB EmployeeDB
	myEmployeeDB.employees = make(map[int]Employee)

	myEmployeeDB.addEmployee(Employee{ID: 1, FirstName: "John", LastName: "Doe", Position: "Developer", Salary: 50000})
	myEmployeeDB.addEmployee(Employee{ID: 2, FirstName: "Jane", LastName: "Smith", Position: "Manager", Salary: 70000})
	myEmployeeDB.showEmployees()

	myEmployeeDB.updateEmployee(1, Employee{ID: 1, FirstName: "John", LastName: "Doe", Position: "Senior Developer", Salary: 60000})
	myEmployeeDB.showEmployees()

	myEmployeeDB.deleteEmployee(2)
	myEmployeeDB.showEmployees()
}
