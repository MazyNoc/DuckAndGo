package persistence

import (
	"sync"
	"annat.nu/data/sample"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type single struct {
	db *sql.DB
}

var instantiated *single
var once sync.Once

func New() *single {
	once.Do(func() {
		instantiated = &single{}
		db, _ := sql.Open("sqlite3", "./foo.db")
		db.Exec("create table nodeping (id INTEGER PRIMARY KEY AUTOINCREMENT, value int, timestamp date)")
		instantiated.db = db
		log.Println("singleton executed")
	})
	return instantiated
}

func GetNodePing() sample.Sample {
	var sample = sample.Sample{}
	sample.Key = "nodeping"
	sample.Typ = "status"

	rows, _ := New().db.Query("select value, timestamp from nodeping order by id desc limit 1")
	if (rows.Next()) {
		rows.Scan(&sample.Data.Status, &sample.Data.Timestamp)
	}
	rows.Close()
	return sample
}

func SetNodePing(nodeping sample.Sample) {
	stmt, err := New().db.Prepare("insert into nodeping (value, timestamp) values (?,?)")
	log.Println(err)
	stmt.Exec(nodeping.Data.Status, nodeping.Data.Timestamp)
}
