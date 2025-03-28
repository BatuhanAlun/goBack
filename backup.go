package main

import (
	"fmt"
	"os"
	"os/exec"
)

func TakeBackupSqlite(argument Args) {
	cmd := exec.Command("sqlite3", argument.DBName+".db", ".dump")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = os.WriteFile("backup.sql", output, 0644) //join at bitti
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
