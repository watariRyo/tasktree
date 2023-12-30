package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/watariRyo/tasktree/server/domain/repository"
)

var defaultConnectionInfo = &connectionInfo{
	username:  "root",
	password:  "root",
	host:      "db",
	port:      "3306",
	schema:    "balance",
	debugMode: false,
}

type connectionInfo struct {
	username  string
	password  string
	host      string
	port      string
	schema    string
	debugMode bool
}

func (c *connectionInfo) connectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charaset=utf8&parseTime=True&loc=Asia%%2FTokyo", c.username, c.password, c.host, c.port, c.schema)
}

type OptionFunc func(c *connectionInfo)

type ConnectionManager struct {
	connInfo *connectionInfo
}

func NewConnectionManager(options ...OptionFunc) *ConnectionManager {
	c := defaultConnectionInfo
	for _, option := range options {
		option(c)
	}
	return &ConnectionManager{connInfo: c}
}

func (cm *ConnectionManager) DSN(c *connectionInfo) string {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	mc := mysql.Config{
		DBName:    c.schema,
		User:      c.username,
		Passwd:    c.password,
		Addr:      c.host,
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	return mc.FormatDSN()
}

func Username(username string) OptionFunc {
	return func(c *connectionInfo) {
		c.username = username
	}
}

func Password(password string) OptionFunc {
	return func(c *connectionInfo) {
		c.password = password
	}
}

func Host(host string) OptionFunc {
	return func(c *connectionInfo) {
		c.host = host
	}
}

func Port(port string) OptionFunc {
	return func(c *connectionInfo) {
		c.port = port
	}
}

func Schema(schema string) OptionFunc {
	return func(c *connectionInfo) {
		c.schema = schema
	}
}

func DebugMode(debugMode bool) OptionFunc {
	return func(c *connectionInfo) {
		c.debugMode = debugMode
	}
}

func (cm *ConnectionManager) Open() (repository.DBConnection, error) {
	ci := cm.connInfo
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", ci.username, ci.password, ci.host, ci.schema))
	if err != nil {
		return nil, err
	}

	boil.SetDB(db)
	boil.DebugMode = ci.debugMode

	return db, nil
}

func Transaction(ctx context.Context, conn repository.DBConnection, txFunc func(ctx context.Context, tx repository.DBConnection) error, options *sql.TxOptions) (err error) {
	db, ok := conn.(*sql.DB)
	if !ok {
		return errors.New("invalid connection")
	}

	tx, err := db.BeginTx(ctx, options)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
		}
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	err = txFunc(ctx, tx)
	return err
}
