package command

import (
	"fmt"
	"os"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var cfg = conf.NewConfig()

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Root command for our app",
	Long:  `Root command for our app, the main purpose is to help setup subcommands`,
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate cmd is used for database migration",
	Long:  `migrate cmd is used for database migration: migrate < up | down >`,
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "migrate to v1 command",
	Long:  `Command to install version 1 of our application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running migrate up command")

		db := db_repository.InitDB(cfg)
		dbDriver, err := mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			fmt.Printf("instance error: %v \n", err)
		}

		fileSource, err := (&file.File{}).Open("file://cmd/migrate/migrations")
		if err != nil {
			fmt.Printf("opening file error: %v \n", err)
		}

		m, err := migrate.NewWithInstance("file", fileSource, "delivery_db", dbDriver)
		if err != nil {
			fmt.Printf("migrate error: %v \n", err)
		}

		if err = m.Up(); err != nil {
			fmt.Printf("migrate up error: %v \n", err)
		}

		fmt.Println("Migrate up done with success")

	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "migrate from v2 to v1",
	Long:  `Command to downgrade database from v2 to v1`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running migrate down command")

		db := db_repository.InitDB(cfg)
		dbDriver, err := mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			fmt.Printf("instance error: %v \n", err)
		}

		fileSource, err := (&file.File{}).Open("file://cmd/migrate/migrations")
		if err != nil {
			fmt.Printf("opening file error: %v \n", err)
		}

		m, err := migrate.NewWithInstance("file", fileSource, "delivery_db", dbDriver)
		if err != nil {
			fmt.Printf("migrate error: %v \n", err)
		}

		if err = m.Down(); err != nil {
			fmt.Printf("migrate down error: %v \n", err)
		}

		fmt.Println("Migrate down done with success")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
