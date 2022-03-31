package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"seckill/api"
	"sync"
)

var db *sql.DB
var err error

type user struct {
	id   int
	name string
	age  int
}

func Query(tel string) (api.Login, error) {
	sqlStr := "select tel,pass from test where tel=? "
	var u api.Login
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, tel).Scan(&u.Tel, &u.Passwd)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return u, err
	}
	return u, err
}

func PrepareQuery(tel string) (api.Login, error) {
	//WG.Add(1)
	//defer WG.Done()
	sqlstr := "select tel,pass from test where tel=? "
	stmt, err := db.Prepare(sqlstr)
	var inf api.Login
	if err != nil {
		fmt.Println("cant prepare,error:", err)
		return inf, err
	}
	rowobj, err := stmt.Query(tel)
	if err != nil {
		fmt.Println("cantread error:", err)
		return inf, err
	}
	defer rowobj.Close()
	for rowobj.Next() {
		err := rowobj.Scan(&inf.Tel, &inf.Passwd)
		if err != nil {
			fmt.Println("cant read ,error", err)
			//return
		}
		fmt.Printf("tel:%v,passwd:%v\n", inf.Tel, inf.Passwd)
	}
	return inf, nil

}
func prepareDelete(id int) {
	sqlstr := "delete from user where id =?"
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Println("cant prepare error:", err)
		return
	}

	result, err := stmt.Exec(id)
	if err != nil {
		fmt.Println("delee fauid,error", err)
		return
	}
	i, _ := result.LastInsertId()
	j, _ := result.RowsAffected()
	fmt.Println(i, j)
}

func InitDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("error1", err)
		return err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("error2", err)
		return err
	}
	fmt.Println("数据库连接成功")
	db.SetMaxOpenConns(1)
	return
}
func Insert(name string, id string, email string, tel string, pass string) {
	sqlinsert := "insert into test(name,id,email,tel,pass) values(?,?,?,?,?)"
	ret, err := db.Exec(sqlinsert, name, id, email, tel, pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	line, err := ret.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(line)
}
func delete(id int) {
	sqlstr := "delete from user where id=?"
	ret, err := db.Exec(sqlstr, id)
	if err != nil {
		fmt.Println("a", err)
		return
	}
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("b", err)
		return
	}
	fmt.Println(num)
}
func updateRow(age int, name string) {
	sqlstr := "update user set age=? where name =?"
	ret, err := db.Exec(sqlstr, age, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println(id)

}

var WG sync.WaitGroup
var num = make(chan int, 10)
