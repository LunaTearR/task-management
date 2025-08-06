package main

import (
	"fmt"
	"os"
	"task-management-task-service/configs"
	"task-management-task-service/core/models"
	"task-management-task-service/pkg"
	"time"

	"task-management-task-service/server"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "แอปพลิเคชันสำหรับ Task Management(task-service)",
		Long:  `แอปพลิเคชันนี้เป็น backend สำหรับ Task Management(task-service) ที่ใช้ในการจัดการข้อมูลและการเรียกใช้ API`,
	}

	var serveCmd = &cobra.Command{
		Use:   "start",
		Short: "เริ่มการทำงานของเซิร์ฟเวอร์",
		Run: func(cmd *cobra.Command, args []string) {
			// โหลดการตั้งค่า
			cfg := configs.NewConfig()

			// เชื่อมต่อกับฐานข้อมูล
			for {
				db, err := pkg.NewPostgresDB(cfg)
				if err != nil {
					fmt.Printf("ไม่สามารถเชื่อมต่อกับฐานข้อมูลได้: %v\n", err)
					fmt.Println("กำลังพยายามเชื่อมต่อกับฐานข้อมูลอีกครั้งใน 3 วินาที...")
					// Wait before retrying
					time.Sleep(3 * time.Second)
					continue
				}
				db.Migrate(&models.Project{}, &models.ProjectMember{}, &models.Task{}, &models.TaskUser{})
				// สร้างและเริ่มเซิร์ฟเวอร์
				server := server.NewServer(cfg, db)
				server.Start()
				break
			}
		},
	}
	rootCmd.AddCommand(serveCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
