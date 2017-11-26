package main

import (
	"database/sql"
	"errors"
)

type Employee struct {
	NIK    int    `json:"id"`
	Nama   string `json:"employee_name"`
	Salary string `json:"employee_salary"`
	Age    int    `json:"employee_age"`
}

func getEmployees(db *sql.DB, start, count int) ([]Employee, error) {
	return nil, errors.New("Not implemented")
}

func (e *Employee) getEmployee(db *sql.DB) error {
	return db.QueryRow("SELECT id, employee_name, employee_salary, employee_age FROM employee WHERE id=?", e.NIK).Scan(&e.NIK, &e.Nama, &e.Salary, &e.Age)
}

func (e *Employee) updateEmployee(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (e *Employee) deleteEmployee(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (e *Employee) createEmployee(db *sql.DB) error {
	return errors.New("Not implemented")
}
