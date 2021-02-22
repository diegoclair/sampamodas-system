package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/GuiaBolso/darwin"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/sampamodas-system/backend/contract"
	"github.com/diegoclair/sampamodas-system/backend/infra/data/migrations"
	"github.com/diegoclair/sampamodas-system/backend/util/config"
	mysqlDriver "github.com/go-sql-driver/mysql"
)

var (
	conn    *mysqlConn
	onceDB  sync.Once
	connErr error
)

// mysqlConn is the database connection manager
type mysqlConn struct {
	db *sql.DB
}

//Instance returns an instance of a MySQLRepo
func Instance() (contract.MySQLRepo, error) {
	onceDB.Do(func() {
		cfg := config.GetDBConfig()

		dataSourceName := fmt.Sprintf("%s:root@tcp(%s:%s)/%s?charset=utf8",
			cfg.Username, cfg.Host, cfg.Port, cfg.DBName,
		)

		log.Println("Connecting to database...")
		db, connErr := sql.Open("mysql", dataSourceName)
		if connErr != nil {
			return
		}

		log.Println("Database Ping...")
		connErr = db.Ping()
		if connErr != nil {
			return
		}

		log.Println("Creating database...")
		if _, connErr = db.Exec("CREATE DATABASE IF NOT EXISTS sampamodas_db;"); connErr != nil {
			logger.Error("Create Database error: ", connErr)
			return
		}

		if _, connErr = db.Exec("USE sampamodas_db;"); connErr != nil {
			logger.Error("Default Database error: ", connErr)
			return
		}

		connErr = mysqlDriver.SetLogger(logger.GetLogger())
		if connErr != nil {
			return
		}
		logger.Info("Database successfully configured")

		logger.Info("Running the migrations")
		driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})

		d := darwin.New(driver, migrations.Migrations, nil)

		connErr = d.Migrate()
		if connErr != nil {
			logger.Error("Migrate Error: ", connErr)
			return
		}

		logger.Info("Migrations executed")

		conn = &mysqlConn{
			db: db,
		}
	})

	return conn, nil
}

// Begin starts a transaction
func (c *mysqlConn) Begin() (contract.MysqlTransaction, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}

	return newTransaction(tx), nil
}

// Close closes the db connection
func (c *mysqlConn) Close() (err error) {
	return c.db.Close()
}

//Business returns the business set
func (c *mysqlConn) Business() contract.BusinessRepo {
	return newBusinessRepo(c.db)
}

//Company returns the company set
func (c *mysqlConn) Company() contract.CompanyRepo {
	return newCompanyRepo(c.db)
}

//Lead returns the lead set
func (c *mysqlConn) Lead() contract.LeadRepo {
	return newLeadRepo(c.db)
}

//Product returns the product set
func (c *mysqlConn) Product() contract.ProductRepo {
	return newProductRepo(c.db)
}

//Sale returns the sale set
func (c *mysqlConn) Sale() contract.SaleRepo {
	return newSaleRepo(c.db)
}
