package main

import (
    "log"
    "os"

    "nayaka/internal/config"
    "nayaka/internal/handler"
    "nayaka/internal/repository"
    "nayaka/internal/service"
    "nayaka/pkg/database"

    "github.com/gin-gonic/gin"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    db, err := database.NewPostgresConnection(cfg.DatabaseURL)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    employeeRepo := repository.NewEmployeeRepository(db)
    employeeService := service.NewEmployeeService(employeeRepo)
    employeeHandler := handler.NewEmployeeHandler(employeeService)

    router := gin.Default()

    router.GET("/active-smith-employees", employeeHandler.GetActiveSmithEmployees)
    router.GET("/employees-without-reviews", employeeHandler.GetEmployeesWithoutReviews)
    router.GET("/hire-date-difference", employeeHandler.GetHireDateDifference)
    router.GET("/employee-salary-projections", employeeHandler.GetEmployeeSalaryProjections)
    router.GET("/read-json-file", employeeHandler.ReadJSONFile)

    port := os.Getenv("PORT")
    if port == "" {
        port = "9000"
    }

    log.Printf("Server is running on port %s\n", port)
    log.Fatal(router.Run(":" + port))
}