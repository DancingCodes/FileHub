package main

import (
	"backend/internal/db"
	"backend/internal/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("未找到 .env 文件，将使用系统环境变量")
	}

	uploadPath := os.Getenv("UPLOAD_DIR")
	err := os.MkdirAll(uploadPath, 0755)
	if err != nil {
		log.Fatalf("无法创建上传目录 [%s]: %v", uploadPath, err)
	}

	db.InitDB()
	r := router.NewRouter()

	port := os.Getenv("PORT")

	if err := r.Run(port); err != nil {
		log.Fatalf("❌ 服务启动崩溃: %v", err)
	}
}
