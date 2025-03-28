package main

type Args struct {
	User             string
	DBType           string
	Backup           bool
	BackupDir        string
	Load             bool
	DumpFileLocation string
	DBName           string
}
