package cmd

import (
	"github.com/shea11012/go_blog/global"
	"github.com/shea11012/go_blog/internal/models"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "数据库迁移",
	Long:  "数据库迁移",
	Run: func(cmd *cobra.Command, args []string) {
		err := global.DBEngine.AutoMigrate(
			&models.Tag{},
			&models.Article{},
			&models.ArticleTag{},
			&models.Auth{},
		)
		global.CheckError(err)
	},
}
