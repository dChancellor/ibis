package db

import (
	"os"
	"fmt"
	"path/filepath"

	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"

)

var DB *sql.DB

// GetDBPath returns the path where the SQLite database should be stored
func GetDBPath(appName, dbName string) (string, error) {
    // Get the user-specific config directory (cross-platform)
    dir, err := os.UserConfigDir()
    if err != nil {
        return "", err
    }

    // Create a directory for your app
    appDir := filepath.Join(dir, appName)
    err = os.MkdirAll(appDir, os.ModePerm) // Make the directory if it doesn't exist
    if err != nil {
        return "", err
    }

    // Return the full path to the SQLite database
    return filepath.Join(appDir, dbName), nil
}

func InitializeDB(appName string) {
    var err error
	var dbName = appName + ".db"
	// Get the full path to the SQLite database file
	dbPath, err := GetDBPath(appName, dbName)
	if err != nil {
		log.Fatal("Error getting DB path:", err)
	}

	fmt.Println("Database path:", dbPath)

    DB, err = sql.Open("sqlite3", dbPath)
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println("DATABASE OPENED SUCCESSFULLY")
    // Create a table if it doesn't exist
    statement, err := DB.Prepare("CREATE TABLE IF NOT EXISTS skills (id INTEGER PRIMARY KEY, name TEXT, svg TEXT)")
    if err != nil {
        log.Fatal(err)
    }
	
	fmt.Println("DATABASE PREPARED SUCCESSFULLY")
    statement.Exec()
	fmt.Println("DATABASE COMPLETE")
}

func GetRow[T any](table_name string, row_name string) (T, error){
    var exists bool
    var row T
    err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM ? WHERE name = ?)", table_name,row_name).Scan(&exists)
    if(err != nil) {
        return row, err
    }
    return row, nil
}