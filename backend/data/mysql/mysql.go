package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GuiaBolso/darwin"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/sampamodas-system/backend/data/migrations"
	"github.com/diegoclair/sampamodas-system/backend/domain/contract"
	"github.com/diegoclair/sampamodas-system/backend/infra/config"
	"github.com/go-sql-driver/mysql"
)

// DBManager is the MySQL connection manager
type DBManager struct {
	db *sql.DB
}

//Instance returns an instance of a RepoManager
func Instance() (contract.RepoManager, error) {
	cfg := config.GetDBConfig()

	dataSourceName := fmt.Sprintf("%s:root@tcp(%s:%s)/%s?charset=utf8",
		cfg.Username, cfg.Host, cfg.Port, cfg.DBName,
	)

	log.Println("Connecting to database...")
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	log.Println("Database Ping...")
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Creating database...")
	if _, err = db.Exec("CREATE DATABASE IF NOT EXISTS sampamodas_db;"); err != nil {
		logger.Error("Create Database error: ", err)
		return nil, err
	}

	if _, err = db.Exec("USE sampamodas_db;"); err != nil {
		logger.Error("Default Database error: ", err)
		return nil, err
	}

	err = mysql.SetLogger(logger.GetLogger())
	if err != nil {
		return nil, err
	}
	logger.Info("Database successfully configured")

	logger.Info("Running the migrations")
	driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})

	d := darwin.New(driver, migrations.Migrations, nil)

	err = d.Migrate()
	if err != nil {
		logger.Error("Migrate Error: ", err)
		return nil, err
	}

	logger.Info("Migrations executed")

	instance := &DBManager{
		db: db,
	}

	return instance, nil
}

//Company returns the company set
func (c *DBManager) Company() contract.CompanyRepo {
	return newCompanyRepo(c.db)
}

//Lead returns the company set
func (c *DBManager) Lead() contract.LeadRepo {
	return newLeadRepo(c.db)
}
