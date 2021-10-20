package testcases

import (
	"fmt"
	"testing"

	"github.com/cyverse/go-irodsclient/fs"
	"github.com/stretchr/testify/assert"
)

func TestFS(t *testing.T) {
	setup()

	t.Run("test PrepareSamples", testPrepareSamples)

	t.Run("test ListEntries", testListEntries)
	t.Run("test ListACLs", testListACLs)
	t.Run("test ReadWrite", testReadWrite)

	shutdown()
}

func testListEntries(t *testing.T) {
	account := GetTestAccount()

	account.ClientServerNegotiation = false

	fsConfig := fs.NewFileSystemConfigWithDefault("go-irodsclient-test")

	fs, err := fs.NewFileSystem(account, fsConfig)
	assert.NoError(t, err)
	defer fs.Release()

	homedir := fmt.Sprintf("/%s/home/%s", account.ClientZone, account.ClientUser)

	collections, err := fs.List(homedir)
	assert.NoError(t, err)

	collectionPaths := []string{}

	for _, collection := range collections {
		collectionPaths = append(collectionPaths, collection.Path)
		assert.NotEmpty(t, collection.ID)
	}

	expected := []string{}
	expected = append(expected, GetTestDirs()...)
	expected = append(expected, GetTestFiles()...)

	assert.Equal(t, len(collections), len(expected))
	assert.ElementsMatch(t, collectionPaths, expected)
}

/*
func testListEntriesByMeta(t *testing.T) {
	logger := log.WithFields(log.Fields{
		"package":  "fs",
		"function": "TestListEntriesByMeta",
	})

	setup()

	entries, err := fs.SearchByMeta("ipc_UUID", "3241af9a-c199-11e5-bd90-3c4a92e4a804")
	if err != nil {
		t.Errorf("err - %v", err)
		panic(err)
	}

	if len(entries) == 0 {
		logger.Debug("There is no entries")
	} else {
		for _, entry := range entries {
			logger.Debugf("Entry : %v", entry)
		}
	}

	shutdown()
}
*/

func testListACLs(t *testing.T) {
	account := GetTestAccount()

	account.ClientServerNegotiation = false

	fsConfig := fs.NewFileSystemConfigWithDefault("go-irodsclient-test")

	fs, err := fs.NewFileSystem(account, fsConfig)
	assert.NoError(t, err)
	defer fs.Release()

	objectPath := GetTestFiles()[0]

	acls, err := fs.ListACLsWithGroupUsers(objectPath)
	assert.NoError(t, err)

	assert.GreaterOrEqual(t, len(acls), 1)
	for _, acl := range acls {
		assert.Equal(t, objectPath, acl.Path)
	}
}

func testReadWrite(t *testing.T) {
	account := GetTestAccount()

	account.ClientServerNegotiation = false

	fsConfig := fs.NewFileSystemConfigWithDefault("go-irodsclient-test")

	fs, err := fs.NewFileSystem(account, fsConfig)
	assert.NoError(t, err)
	defer fs.Release()

	homedir := fmt.Sprintf("/%s/home/%s", account.ClientZone, account.ClientUser)

	newDataObjectFilename := "testobj123"
	newDataObjectPath := homedir + "/" + newDataObjectFilename

	text := "HELLO WORLD!<?!'\">"

	// create
	handle, err := fs.CreateFile(newDataObjectPath, "")
	assert.NoError(t, err)

	err = handle.Write([]byte(text))
	assert.NoError(t, err)

	err = handle.Close()
	assert.NoError(t, err)

	assert.True(t, fs.Exists(newDataObjectPath))

	// read
	newHandle, err := fs.OpenFile(newDataObjectPath, "", "r")
	assert.NoError(t, err)

	readData, err := newHandle.Read(1024)
	assert.NoError(t, err)

	err = newHandle.Close()
	assert.NoError(t, err)

	assert.Equal(t, text, string(readData))

	// delete
	err = fs.RemoveFile(newDataObjectPath, true)
	assert.NoError(t, err)

	assert.False(t, fs.Exists(newDataObjectPath))
}
