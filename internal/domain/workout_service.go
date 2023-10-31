package domain

import "gorm.io/gorm"

type WorkoutService struct {
	db *gorm.DB
}

func (s *WorkoutService) FetchLogs() ([]Log, error) {

	var logs []Log

	result := s.db.Find(&logs)

	if result.Error != nil {
		return logs, result.Error
	}

	return logs, nil
}

func (s *WorkoutService) FetchLogById(uid string) (Log, error) {

	var log Log

	result := s.db.Where("uid = ?", uid).First(&log)

	if result.Error != nil {
		return log, result.Error
	}

	return log, nil

}

func (s *WorkoutService) CreateLog(input NewLogInput) {

}

func (s *WorkoutService) UpdateLog() {

}

func (s *WorkoutService) DeleteLog() {

}
