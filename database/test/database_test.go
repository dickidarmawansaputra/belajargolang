package test

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestCreateDBConnection(t *testing.T) {
	db, err := sql.Open("mysql", "dickids:rahasia@tcp(localhost:3306)/belajar-golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func GetConnection() *sql.DB {
	// ini tidak di close, karna akan di pake di func test lain (hanya untuk belajar)
	// parseTime=true
	/*
		● Secara default, Driver MySQL untuk Golang akan melakukan query tipe data DATE, DATETIME,
		TIMESTAMP menjadi []byte / []uint8. Dimana ini bisa dikonversi menjadi String, lalu di parsing
		menjadi time.Time
		● Namun hal ini merepotkan jika dilakukan manual, kita bisa meminta Driver MySQL untuk Golang
		secara otomatis melakukan parsing dengan menambahkan parameter parseTime=true
	*/
	db, err := sql.Open("mysql", "dickids:rahasia@tcp(localhost:3306)/belajar-golang?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('dicki', 'Dicki')"

	// execContext digunakan selain query data yg tanpa mengembalikan result seperti insert, delete, etc
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("berhasil insert ke DB")
}

func TestQueryContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	// best practice: rows perlu ditutup jika tidak digunakan lagi
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id)
		fmt.Println("name:", name)
	}

}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// saran mention satu satu kolomnya, biar tahu posisi kolom jika ada alter column
	// urutan kolom harus sama ketika rows.Scan()
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	// best practice: rows perlu ditutup jika tidak digunakan lagi
	defer rows.Close()

	for rows.Next() {
		var id, name string
		// untuk tipe data nullable pake seperti ini
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		// hati - hati jika kolom nullable maka akan error ketika scan
		/*
			● Konversi secara otomatis NULL tidak didukung oleh Driver MySQL Golang
			● Oleh karena itu, khusus tipe kolom yang bisa NULL, kita perlu menggunakan tipe data yang ada
			dalam package sql
		*/
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		// akan error untuk tipe data time.Time
		/*
			● Secara default, Driver MySQL untuk Golang akan melakukan query tipe data DATE, DATETIME,
			TIMESTAMP menjadi []byte / []uint8. Dimana ini bisa dikonversi menjadi String, lalu di parsing
			menjadi time.Time
			● Namun hal ini merepotkan jika dilakukan manual, kita bisa meminta Driver MySQL untuk Golang
			secara otomatis melakukan parsing dengan menambahkan parameter parseTime=true
		*/

		fmt.Println("id:", id)
		fmt.Println("name:", name)
		if email.Valid {
			fmt.Println("email:", email.String)
		}
		fmt.Println("balance:", balance)
		fmt.Println("rating:", rating)
		fmt.Println("birthDate:", birthDate)
		fmt.Println("married:", married)
		fmt.Println("createdAt:", createdAt)
	}

}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// simulasi sql injection
	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	fmt.Println(script) // SELECT username FROM user WHERE username = 'admin'; #' AND password = 'salah' LIMIT 1

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("sukses login:", username)
	} else {
		fmt.Println("gagal login")
	}
}

// solusi SQL Injection
func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// simulasi sql injection
	username := "admin'; #"
	password := "salah"

	// gunakan tanda "?"
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(script) // SELECT username FROM user WHERE username = 'admin'; #' AND password = 'salah' LIMIT 1

	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("sukses login:", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestExecSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "dicki"
	password := "dicki"

	script := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("berhasil insert ke DB")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "dicki@mail.com"
	comment := "komen pertama"

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("insert ID:", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "dicki" + strconv.Itoa(i) + "@mail.com"
		comment := "Komen ke " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("comment ID:", id)
	}
}

func TestDatabaseTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	// do transaction
	for i := 0; i < 10; i++ {
		email := "transaction" + strconv.Itoa(i) + "@mail.com"
		comment := "Komen ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("comment ID:", id)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
