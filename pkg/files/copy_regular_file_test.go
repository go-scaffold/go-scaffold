package files_test

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pasdam/go-project-template/pkg/files"
	"github.com/stretchr/testify/assert"
)

func Test_CopyRegularFile_Success_ShouldCopyFile(t *testing.T) {
	src := "test/file_to_copy.txt"
	dst, err := randomFile()
	assert.Nil(t, err)
	err = files.CopyRegularFile(src, dst)
	assert.Nil(t, err)

	_, err = os.Stat(dst)
	assert.Nil(t, err)
	copiedContent, _ := ioutil.ReadFile(dst)
	assert.Equal(t, "file-to-copy-content\n", string(copiedContent))

	os.Remove(dst)
}

func Test_CopyRegularFile_Fail_ShouldReturnErrorIfFileDoesNotExist(t *testing.T) {
	src := "test/file_that_does_not_exist.txt"
	dst, err := randomFile()
	assert.Nil(t, err)
	err = files.CopyRegularFile(src, dst)
	assert.NotNil(t, err)

	_, err = os.Stat(dst)
	assert.NotNil(t, err)
}

func Test_CopyRegularFile_Fail_SourceAndDestinationAreSame(t *testing.T) {
	src := "test/file_to_copy.txt"
	dst := src
	err := files.CopyRegularFile(src, dst)
	assert.NotNil(t, err)

	_, err = os.Stat(dst)
	assert.Nil(t, err)
}

func Test_CopyRegularFile_Fail_SourceIsNotARegularFile(t *testing.T) {
	src := "test/test_folder"
	dst, err := randomFile()
	assert.Nil(t, err)
	err = files.CopyRegularFile(src, dst)
	assert.NotNil(t, err)
}

func Test_CopyRegularFile_Fail_CannotCreateDestination(t *testing.T) {
	src := "test/file_to_copy.txt"
	dst := "test/folder_that_does_not_exist/destination_file"
	err := files.CopyRegularFile(src, dst)
	assert.NotNil(t, err)

	_, err = os.Stat(dst)
	assert.NotNil(t, err)
}

func randomFile() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("/var/tmp/goscaffold_files_test-%X", b[0:]), nil
}
