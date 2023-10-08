package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	// 環境変数を取得する
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	// データベースに接続する
	_, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, database))
	if err != nil {
		panic(err)
	}

	// 後でマイグレーションのモジュールを作成する

}

func main() {

	// 後でAPIエンドポイントのモジュールを作成する

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
