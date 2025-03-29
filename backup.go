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

func TakeBackupMySql(argument Args) {
	scriptContent := `#!/bin/bash
sudo mysqldump -u ` + argument.User + " " + argument.DBName + ` > ` + argument.BackupDir + `/mysqlbackup.sql`

	tmpFile, err := os.CreateTemp("", "backup_*.sh")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(scriptContent)
	if err != nil {
		fmt.Println("Error writing to temporary file:", err)
		return
	}

	err = tmpFile.Close()
	if err != nil {
		fmt.Println("Error closing temporary file:", err)
		return
	}

	err = os.Chmod(tmpFile.Name(), 0755)
	if err != nil {
		fmt.Println("Error setting file permissions:", err)
		return
	}

	cmd := exec.Command(tmpFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing script:", err)
	}

}

func LoadBackupMySql(argument Args) {
	scriptContent := `#!/bin/bash
	sudo mysql -u ` + argument.User + " " + argument.DBName + ` < ` + argument.DumpFileLocation

	tmpFile, err := os.CreateTemp("", "backup_*.sh")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(scriptContent)
	if err != nil {
		fmt.Println("Error writing to temporary file:", err)
		return
	}

	err = tmpFile.Close()
	if err != nil {
		fmt.Println("Error closing temporary file:", err)
		return
	}

	err = os.Chmod(tmpFile.Name(), 0755)
	if err != nil {
		fmt.Println("Error setting file permissions:", err)
		return
	}

	cmd := exec.Command(tmpFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing script:", err)
	}
}
