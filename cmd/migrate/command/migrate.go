package command

import (
	"database/sql"
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
	Long:  `migrate cmd is used for database migration: migrate < create | up | down >`,
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create db",
	Long:  `Command to create db for further migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running migrate create command")

		db, err := sql.Open(
			"mysql",
			fmt.Sprintf(
				"root:%s@tcp(%s:%s)/?multiStatements=true",
				cfg.ConfigDB.DbRootPassword,
				cfg.ConfigDB.Host,
				cfg.ConfigDB.Port,
			),
		)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + cfg.ConfigDB.DBName)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(fmt.Sprintf("GRANT ALL PRIVILEGES ON %s.* TO %s@%s;", cfg.ConfigDB.DBName, cfg.ConfigDB.Username, cfg.ConfigDB.Host))
		if err != nil {
			panic(err)
		}

		db.Close()

		fmt.Println("Creation done with success")
	},
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

		m, err := migrate.NewWithInstance("file", fileSource, cfg.ConfigDB.DBName, dbDriver)
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

		m, err := migrate.NewWithInstance("file", fileSource, cfg.ConfigDB.DBName, dbDriver)
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
	migrateCmd.AddCommand(migrateCreateCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
