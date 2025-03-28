package main

import "fmt"

func main() {
	args := ParseFlags()

	if args.Backup {
		TakeBackup(args)
	} else if args.Load {
		LoadBackup(args)
	} else {
		fmt.Println("User should have declared one of the backup or load arguments")
	}

}
