package main

import (
	"embed"
	"os"
	"fmt"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed all:frontend/dist
var assets embed.FS

var db *sql.DB

var appName = "Ibis"

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

// InitializeDB opens a connection to the SQLite database
func InitializeDB() {
    var err error
	var dbName = appName + ".db"
	// Get the full path to the SQLite database file
	dbPath, err := GetDBPath(appName, dbName)
	if err != nil {
		log.Fatal("Error getting DB path:", err)
	}

	fmt.Println("Database path:", dbPath)

    db, err = sql.Open("sqlite3", dbPath)
    if err != nil {
        log.Fatal(err)
    }

    // Create a table if it doesn't exist
    statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS skills (id INTEGER PRIMARY KEY, name TEXT)")
    if err != nil {
        log.Fatal(err)
    }
    statement.Exec()
}

func main() {
	// Create an instance of the app structure
	app := NewApp()
	InitializeDB()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Ibis",
		Width:  1024,
		Height: 1024,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
