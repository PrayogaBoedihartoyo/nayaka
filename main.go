package main

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "log"
    "os"
    _ "github.com/lib/pq"
)

type Employee struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}

var db *sql.DB

func main() {
    var err error
    db, err = sql.Open("postgres", "user=postgres.jmxllisawahmpkdkpzvo password=Lordrazor890# host=aws-0-us-east-1.pooler.supabase.com port=6543 dbname=postgres")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    router := gin.Default()

    router.GET("/active-smith-employees", getActiveSmithEmployeesHandler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "9000"
    }

    log.Printf("Server is running on port %s\n", port)
    log.Fatal(router.Run(":" + port))
}

func getActiveSmithEmployeesHandler(c *gin.Context) {
    employees, err := getActiveSmithEmployees(db)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, employees)
}

func getActiveSmithEmployees(db *sql.DB) ([]Employee, error) {
    rows, err := db.Query(`
        SELECT first_name, last_name
        FROM employees
        WHERE is_active = true AND last_name LIKE 'Smith%'
        ORDER BY last_name, first_name
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var employees []Employee
    for rows.Next() {
        var e Employee
        if err := rows.Scan(&e.FirstName, &e.LastName); err != nil {
            return nil, err
        }
        employees = append(employees, e)
    }
    return employees, nil
}