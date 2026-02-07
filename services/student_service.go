package services

import (
	"errors"

	"example.com/student-api/models"
	"example.com/student-api/repositories"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

func ValidateStudent(student models.Student, requireID bool) error {
	if requireID && student.Id == "" {
		return errors.New("id must not be empty")
	}
	if student.Name == "" {
		return errors.New("name must not be empty")
	}
	if student.GPA < 0.0 || student.GPA > 4.0 {
		return errors.New("gpa must be between 0.00 and 4.00")
	}
	return nil
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.Repo.GetByID(id)
}

func (s *StudentService) CreateStudent(student models.Student) error {
	if err := ValidateStudent(student, true); err != nil {
		return err
	}
	return s.Repo.Create(student)
}

func (s *StudentService) UpdateStudent(id string, student models.Student) (*models.Student, error) {
	if err := ValidateStudent(student, false); err != nil {
		return nil, err
	}
	if err := s.Repo.Update(id, student); err != nil {
		return nil, err
	}
	return s.Repo.GetByID(id)
}

func (s *StudentService) DeleteStudent(id string) error {
	return s.Repo.Delete(id)
}
