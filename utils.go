package main

import (
	"flag"
	"fmt"
)

func ParseFlags() {
	var userFlag = flag.String("user", "Undefined", "Username for Database")
	var dbTypeFlag = flag.String("db-type", "Undefined", "DB Type selection")
	var backupFlag = flag.Bool("backup", false, "Is it gonna back Up ?")
	var backupDirFlag = flag.String("backup-dir", "/", "Dump file Location")
	var loadFlag = flag.Bool("load", false, "Is it gonna Load ?")
	var dumpFileLocationFlag = flag.String("dump-file", "Undefined", "Dump file location for load")
	var dbNameFlag = flag.String("db-name", "Undefined", "Db name")

	flag.Parse()
	type Args struct {
		User             string
		DBType           string
		Backup           bool
		BackupDir        string
		Load             bool
		DumpFileLocation string
		DBName           string
	}

	Arguments := Args{User: *userFlag, DBType: *dbTypeFlag, Backup: *backupFlag, BackupDir: *backupDirFlag, Load: *loadFlag, DumpFileLocation: *dumpFileLocationFlag, DBName: *dbNameFlag}
	fmt.Println(Arguments)

}
