package main

import (
	"flag"
	"log"
)

func ParseFlags() Args {
	var userFlag = flag.String("user", "username-Undefined", "Username for Database")
	var dbTypeFlag = flag.String("db-type", "dbtype-Undefined", "DB Type selection")
	var backupFlag = flag.Bool("backup", false, "Is it gonna back Up ?")
	var backupDirFlag = flag.String("backup-dir", "backupdir-Undefined", "Dump file Location")
	var loadFlag = flag.Bool("load", false, "Is it gonna Load ?")
	var dumpFileLocationFlag = flag.String("dump-file", "dumpfilelocation-Undefined", "Dump file location for load")
	var dbNameFlag = flag.String("db-name", "dbname-Undefined", "Db name")

	flag.Parse()

	return Args{User: *userFlag, DBType: *dbTypeFlag, Backup: *backupFlag, BackupDir: *backupDirFlag, Load: *loadFlag, DumpFileLocation: *dumpFileLocationFlag, DBName: *dbNameFlag}

}

func TakeBackup(argument Args) {
	var dbChoice int // 0:sqlite  1:mysql 2:postgre
	checkValues(argument)
	dbChoice = decideDB(argument.DBType)
	switch dbChoice {
	case 0:
		TakeBackupSqlite(argument)
	case 1:
		TakeBackupMySql(argument)
	case 2:

	case 3:
	}
}

func LoadBackup(argument Args) {
	var dbChoice int // 0:sqlite  1:mysql 2:postgre
	checkLoadValues(argument)
	dbChoice = decideDB(argument.DBType)
	switch dbChoice {
	case 0:
		LoadBackupSqlite(argument)
	case 1:
		LoadBackupMySql(argument)
	case 2:

	case 3:
	}
}

func checkLoadValues(argument Args) {

	if argument.DBType == "dbtype-Undefined" {

		log.Fatal("db-type not defined")
	}
	if argument.User == "username-Undefined" {

		log.Fatal("user not defined")
	}
	if argument.DumpFileLocation == "dumpfilelocation-Undefined" {

		log.Fatal("Dump file location is not defined")
	}
	if argument.DBName == "dbname-Undefined" {

		log.Fatal("database name not defined")
	}

}

func checkValues(argument Args) {
	if argument.DBType == "dbtype-Undefined" {
		log.Fatal("db-type not defined")
	}
	if argument.User == "username-Undefined" {
		log.Fatal("user not defined")
	}
	if argument.BackupDir == "backupdir-Undefined" {
		log.Fatal("backup directory not defined")
	}
	if argument.DBName == "dbname-Undefined" {
		log.Fatal("database name not defined")
	}
}

func decideDB(dbtype string) int {
	if dbtype == "sqlite" {
		return 0
	} else if dbtype == "mysql" {
		return 1
	} else if dbtype == "postgresql" {
		return 2
	} else {
		log.Fatal("Wrong dbtype input. It should one of sqlite,mysql,postgresql")
	}
	return 3
}
