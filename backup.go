package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func TakeBackupSqlite(argument Args) {
	cmd := exec.Command("sqlite3", argument.DBName+".db", ".dump")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fp := filepath.Join(argument.BackupDir, "backup.sql")
	err = os.WriteFile(fp, output, 0644) //join at bitti
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func LoadBackupSqlite(argument Args) {
	fp := filepath.Join(argument.DumpFileLocation, "backup.sql")
	loadCommand := "sqlite3 " + argument.DBName + ".db" + " < " + fp
	cmd := exec.Command("sh", "-c", loadCommand)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

}
