package main

import (
	"fmt"
	"os"
	"task-management-backend/configs"
	"task-management-backend/pkg"
	"task-management-backend/servers"

	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "แอปพลิเคชันสำหรับ Task Management",
		Long:  `แอปพลิเคชันนี้ช่วยให้คุณจัดการงานต่างๆ ได้อย่างมีประสิทธิภาพ`,
	}

	var serveCmd = &cobra.Command{
		Use:   "start",
		Short: "เริ่มต้นเซิร์ฟเวอร์",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := configs.NewConfig()
			
			logger := pkg.NewAppLogZap()
			
			server := servers.NewServer(cfg, logger)
			server.Start()
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "แสดงเวอร์ชันของแอปพลิเคชัน",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	var checkCmd = &cobra.Command{
		Use:   "check",
		Short: "ตรวจสอบสถานะของแอปพลิเคชัน",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(checkCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
