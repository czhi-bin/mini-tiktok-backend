package main

import (
	"github.com/czhi-bin/mini-tiktok-backend/biz/dal/db"

	"fmt"
)

func main() {
	db.Init()
	fmt.Println("DB init success")
}