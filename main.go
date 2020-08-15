package main

import (
	"database/sql"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/test_database")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("connected")
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	layout := "2006-01-02"

	value := "2020-08-01"
	t, err := time.Parse(layout, value)

	if err != nil {
		log.Fatal("parse time: ", err)
	}

	var wg sync.WaitGroup

	for {
		if t.After(time.Now()) {
			break
		}
		wg.Add(1)
		go func(db *sql.DB, t time.Time) {
			genVal(db, t)
			wg.Done()
		}(db, t)
		t = t.Add(1 * time.Hour)
	}
	wg.Wait()
	log.Println("done")
}

func execSQL(db *sql.DB, sqlIns string, vals []interface{}) {
	stmtIns, err := db.Prepare(sqlIns)
	if err != nil {
		log.Fatal("insert ", err)
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(vals...)
	if err != nil {
		log.Fatal("insert exec ", err)
	}
}

func genVal(db *sql.DB, t time.Time) {
	start := 111111111
	for i := start; i < start+999; i++ {
		vs := []interface{}{}
		// id1
		if i%10 != 0 {
			vs = append(vs, i%10)
		} else {
			continue
		}
		// id2
		if i/10%10 != 0 {
			vs = append(vs, i/10%10)
		} else {
			continue
		}
		// id3
		if i/100%10 != 0 {
			vs = append(vs, i/100%10)
		} else {
			continue
		}
		// id4
		if i/1000%10 != 0 {
			vs = append(vs, i/1000%10)
		} else {
			continue
		}
		// id5
		if i/10000%10 != 0 {
			vs = append(vs, i/10000%10)
		} else {
			continue
		}
		// id6
		if i/100000%10 != 0 {
			vs = append(vs, i/100000%10)
		} else {
			continue
		}
		// id7
		if i/1000000%10 != 0 {
			vs = append(vs, i/1000000%10)
		} else {
			continue
		}
		// id8
		if i/10000000%10 != 0 {
			vs = append(vs, i/10000000%10)
		} else {
			continue
		}
		// id9
		if i/100000000%10 != 0 {
			vs = append(vs, i/100000000%10)
		} else {
			continue
		}
		// log_time
		vs = append(vs, t)
		// request
		vs = append(vs, 100000)
		// imp
		vs = append(vs, 10000)
		// click
		vs = append(vs, 100)
		// revenue
		vs = append(vs, 1)
		sqlIns := "INSERT INTO logs(id1, id2, id3, id4, id5, id6, id7, id8, id9, log_time, request, imp, click, revenue) VALUES " + "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE updated_at = CURRENT_TIMESTAMP"
		execSQL(db, sqlIns, vs)
		//log.Println(vss)
	}
}
