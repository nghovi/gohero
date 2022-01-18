package controllers
import (
	"database/sql"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"vietanh.com/gohero/models"
)
var dbmap = initDb()
func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:root@/cryptohero_local")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	return dbmap
}
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func GetTickerDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var ticker models.Ticker
	err := dbmap.SelectOne(&ticker, "SELECT * FROM tickers WHERE id=? LIMIT 1", id)
	if err == nil {
		c.JSON(200, &ticker)
	} else {
		c.JSON(200, gin.H{"error": err.Error()})
	}
}
