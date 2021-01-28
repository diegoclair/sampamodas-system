package entity

// InitialConfig entity
type InitialConfig struct {
	Username         string
	Password         string
	Host             string
	Port             string
	DBName           string
	MaxLifeInMinutes int
	MaxIdleConns     int
	MaxOpenConns     int
}
