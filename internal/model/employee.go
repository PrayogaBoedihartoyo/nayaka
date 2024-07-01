package model

import "time"

type Employee struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    HireDate  time.Time `json:"hire_date"`
}

type Response struct {
    Status string      `json:"status"`
    Data   interface{} `json:"data"`
}

type HireDateDifference struct {
    DaysDifference int `json:"days_difference"`
}

type EmployeeSalaryProjection struct {
    FirstName      string  `json:"first_name"`
    LastName       string  `json:"last_name"`
    InitialSalary  float64 `json:"initial_salary"`
    Salary2016     float64 `json:"salary_2016"`
    TotalReviews   int     `json:"total_reviews"`
}