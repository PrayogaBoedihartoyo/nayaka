package service

import (
    "fmt"
    "os"
    "nayaka/internal/model"
    "nayaka/internal/repository"
    "nayaka/internal/utils"
)

type EmployeeService struct {
    employeeRepo *repository.EmployeeRepository
}

func NewEmployeeService(employeeRepo *repository.EmployeeRepository) *EmployeeService {
    return &EmployeeService{employeeRepo: employeeRepo}
}

func (s *EmployeeService) GetActiveSmithEmployees() ([]model.Employee, error) {
    employees, err := s.employeeRepo.GetActiveSmithEmployees()
    if err != nil {
        return nil, err
    }

    response := model.Response{
        Status: "success",
        Data:   employees,
    }

    err = utils.SaveJSONToFile(response, "contoh2.txt")
    if err != nil {
        return nil, fmt.Errorf("error saving to file: %v", err)
    }

    return employees, nil
}

func (s *EmployeeService) GetEmployeesWithoutReviews() ([]model.Employee, error) {
    employees, err := s.employeeRepo.GetEmployeesWithoutReviews()
    if err != nil {
        return nil, err
    }

    response := model.Response{
        Status: "success",
        Data:   employees,
    }

    err = utils.SaveJSONToFile(response, "contoh3.txt")
    if err != nil {
        return nil, fmt.Errorf("error saving to file: %v", err)
    }

    return employees, nil
}

func (s *EmployeeService) GetHireDateDifference() (int, error) {
    difference, err := s.employeeRepo.GetHireDateDifference()
    if err != nil {
        return 0, err
    }

    response := model.Response{
        Status: "success",
        Data:   difference,
    }

    err = utils.SaveJSONToFile(response, "contoh4.txt")
    if err != nil {
        return 0, fmt.Errorf("error saving to file: %v", err)
    }

    return difference, nil
}

func (s *EmployeeService) GetEmployeeSalaryProjections() ([]model.EmployeeSalaryProjection, error) {
    projections, err := s.employeeRepo.GetEmployeeSalaryProjections()
    if err != nil {
        return nil, err
    }

    response := model.Response{
        Status: "success",
        Data:   projections,
    }

    err = utils.SaveJSONToFile(response, "contoh5.txt")
    if err != nil {
        return nil, fmt.Errorf("error saving to file: %v", err)
    }

    return projections, nil
}

func (s *EmployeeService) ReadJSONFile(filename string) (map[string]interface{}, error) {
    data, err := utils.ReadJSONFromFile(filename)
    if err != nil {
        if os.IsNotExist(err) {
            return nil, os.ErrNotExist
        }
        return nil, fmt.Errorf("error reading JSON file: %v", err)
    }
    return data, nil
}