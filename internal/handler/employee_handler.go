package handler

import (
    "net/http"
    "os"
    "nayaka/internal/service"
    "nayaka/internal/model"
    "github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
    employeeService *service.EmployeeService
}

func NewEmployeeHandler(employeeService *service.EmployeeService) *EmployeeHandler {
    return &EmployeeHandler{employeeService: employeeService}
}

func (h *EmployeeHandler) GetActiveSmithEmployees(c *gin.Context) {
    employees, err := h.employeeService.GetActiveSmithEmployees()
    if err != nil {
        c.JSON(http.StatusInternalServerError, model.Response{
            Status: "error",
            Data:   err.Error(),
        })
        return
    }
    
    c.JSON(http.StatusOK, model.Response{
        Status: "success",
        Data:   employees,
    })
}

func (h *EmployeeHandler) GetEmployeesWithoutReviews(c *gin.Context) {
    employees, err := h.employeeService.GetEmployeesWithoutReviews()
    if err != nil {
        c.JSON(http.StatusInternalServerError, model.Response{
            Status: "error",
            Data:   err.Error(),
        })
        return
    }
    
    c.JSON(http.StatusOK, model.Response{
        Status: "success",
        Data:   employees,
    })
}

func (h *EmployeeHandler) GetHireDateDifference(c *gin.Context) {
    difference, err := h.employeeService.GetHireDateDifference()
    if err != nil {
        c.JSON(http.StatusInternalServerError, model.Response{
            Status: "error",
            Data:   err.Error(),
        })
        return
    }
    
    c.JSON(http.StatusOK, model.Response{
        Status: "success",
        Data: model.HireDateDifference{
            DaysDifference: difference,
        },
    })
}

func (h *EmployeeHandler) GetEmployeeSalaryProjections(c *gin.Context) {
    projections, err := h.employeeService.GetEmployeeSalaryProjections()
    if err != nil {
        c.JSON(http.StatusInternalServerError, model.Response{
            Status: "error",
            Data:   err.Error(),
        })
        return
    }
    
    c.JSON(http.StatusOK, model.Response{
        Status: "success",
        Data:   projections,
    })
}

func (h *EmployeeHandler) ReadJSONFile(c *gin.Context) {
    filename := c.Query("filename")
    if filename == "" {
        c.JSON(http.StatusBadRequest, model.Response{
            Status: "error",
            Data:   "Filename is required",
        })
        return
    }

    data, err := h.employeeService.ReadJSONFile(filename)
    if err != nil {
        if os.IsNotExist(err) {
            // File tidak ditemukan
            c.JSON(http.StatusNotFound, model.Response{
                Status: "error",
                Data:   "File not found: " + filename,
            })
        } else {
            // Error lainnya
            c.JSON(http.StatusInternalServerError, model.Response{
                Status: "error",
                Data:   "Error reading file: " + err.Error(),
            })
        }
        return
    }

    c.JSON(http.StatusOK, data)
}