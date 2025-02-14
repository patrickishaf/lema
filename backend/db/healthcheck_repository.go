package db

type HealthCheckRepository struct{}

func NewHealthCheckRepository() *HealthCheckRepository {
	return &HealthCheckRepository{}
}

func (r *HealthCheckRepository) IsDatabaseOnline() bool {
	sqlDB, err := getDB().DB()
	if err != nil {
		return false
	}
	err = sqlDB.Ping()
	return err == nil
}
