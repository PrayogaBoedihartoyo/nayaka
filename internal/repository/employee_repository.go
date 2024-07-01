package repository

import (
    "database/sql"

    "nayaka/internal/model"
)

type EmployeeRepository struct {
    db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
    return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) GetActiveSmithEmployees() ([]model.Employee, error) {
    rows, err := r.db.Query(`
        SELECT first_name, last_name
        FROM employees
        WHERE is_active = true AND last_name LIKE 'Smith%'
        ORDER BY last_name, first_name
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var employees []model.Employee
    for rows.Next() {
        var e model.Employee
        if err := rows.Scan(&e.FirstName, &e.LastName); err != nil {
            return nil, err
        }
        employees = append(employees, e)
    }
    return employees, nil
}

func (r *EmployeeRepository) GetEmployeesWithoutReviews() ([]model.Employee, error) {
    query := `
        SELECT e.first_name, e.last_name, e.hire_date
        FROM auth.employees e
        LEFT JOIN reviews r ON e.id = r.employee_id
        WHERE r.id IS NULL
        ORDER BY e.hire_date
    `
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var employees []model.Employee
    for rows.Next() {
        var e model.Employee
        if err := rows.Scan(&e.FirstName, &e.LastName, &e.HireDate); err != nil {
            return nil, err
        }
        employees = append(employees, e)
    }
    return employees, nil
}

func (r *EmployeeRepository) GetHireDateDifference() (int, error) {
    var daysDifference int
    err := r.db.QueryRow(`
        SELECT (MAX(hire_date) - MIN(hire_date))::int AS days_difference
        FROM employees
        WHERE terminationdate IS NULL OR terminationdate > CURRENT_DATE
    `).Scan(&daysDifference)
    if err != nil {
        return 0, err
    }
    return daysDifference, nil
}

func (r *EmployeeRepository) GetEmployeeSalaryProjections() ([]model.EmployeeSalaryProjection, error) {
    query := `
    SELECT 
        e.first_name, 
        e.last_name, 
        e.salary AS initial_salary,
        e.salary * POWER(1.15, LEAST(2016 - EXTRACT(YEAR FROM e.hire_date), 7)) AS salary_2016,
        COUNT(r.id) AS total_reviews
    FROM 
        employees e
    LEFT JOIN 
        reviews r ON e.id = r.employee_id
    WHERE 
        e.hire_date <= '2016-12-31'
    GROUP BY 
        e.id, e.first_name, e.last_name, e.salary, e.hire_date
    ORDER BY 
        salary_2016 DESC, total_reviews ASC
    `

    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var projections []model.EmployeeSalaryProjection
    for rows.Next() {
        var p model.EmployeeSalaryProjection
        if err := rows.Scan(&p.FirstName, &p.LastName, &p.InitialSalary, &p.Salary2016, &p.TotalReviews); err != nil {
            return nil, err
        }
        projections = append(projections, p)
    }

    return projections, nil
}