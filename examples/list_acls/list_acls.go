package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cyverse/go-irodsclient/fs"
	"github.com/cyverse/go-irodsclient/irods/types"

	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.WithFields(log.Fields{
		"package":  "main",
		"function": "main",
	})

	// Parse cli parameters
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Give an iRODS path!\n")
		os.Exit(1)
	}

	inputPath := args[0]

	// Read account configuration from YAML file
	yaml, err := ioutil.ReadFile("account.yml")
	if err != nil {
		logger.Errorf("err - %v", err)
		panic(err)
	}

	account, err := types.CreateIRODSAccountFromYAML(yaml)
	if err != nil {
		logger.Errorf("err - %v", err)
		panic(err)
	}

	logger.Debugf("Account : %v", account.MaskSensitiveData())

	// Create a file system
	appName := "list_acls"
	filesystem, err := fs.NewFileSystemWithDefault(account, appName)
	if err != nil {
		logger.Errorf("err - %v", err)
		panic(err)
	}

	defer filesystem.Release()

	accesses, err := filesystem.ListACLs(inputPath)
	if err != nil {
		logger.Errorf("err - %v", err)
		panic(err)
	}

	if len(accesses) == 0 {
		fmt.Printf("Found no acls - %s\n", inputPath)
	} else {
		fmt.Printf("%s\n", inputPath)
		for _, access := range accesses {
			fmt.Printf("> User: %s (%s) = %s\n", access.UserName, access.UserType, access.AccessLevel)
		}
	}
}
