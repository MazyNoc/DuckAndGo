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

func Get() *single {
	once.Do(func() {
		instantiated = &single{}
		db, err := sql.Open("sqlite3", "./foo.db")
		if err != nil {
			panic("Can't open the database, exiting")
		}

		db.Exec("create table nodeping (id INTEGER PRIMARY KEY AUTOINCREMENT, value TEXT, timestamp date)")
		instantiated.db = db
		log.Println("singleton created")
	})
	return instantiated
}

func GetNodePing() sample.Sample {
	sample := sample.Sample{Key:"nodeping", Typ:"status"}
	rows, _ := Get().db.Query("select value, timestamp from nodeping order by id desc limit 1")
	defer rows.Close()

	if (rows.Next()) {
		rows.Scan(&sample.Data.Status, &sample.Data.Timestamp)
	}
	return sample
}

func SetNodePing(nodePing sample.Sample) {
	stmt, _ := Get().db.Prepare("insert into nodeping (value, timestamp) values (?,?)")
	data := nodePing.Data;
	stmt.Exec(data.Status, data.Timestamp)
}
