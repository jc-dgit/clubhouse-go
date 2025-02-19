package clubhouse_test

import (
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
	"testing"

	// Frameworks
	ch "github.com/dgitsystems/clubhouse-go/v2"
)

// In order to run tests, you will need to put your
// API token in a file called .clubhouse in your
// home directory

func Test_Download_000(t *testing.T) {
	if client := ch.New(""); client == nil {
		t.Error("Invalid client == nil")
	} else {
		t.Log(client)
	}
}

func Test_Download_001(t *testing.T) {
	if token, err := getDownloadToken(); err != nil {
		t.Error(err)
	} else if client := ch.New(token).SetDebug(true); client == nil {
		t.Error("Invalid client == nil")
	} else if files, err := client.FileList(); err != nil {
		t.Error(err)
	} else {
		t.Log(files)
	}
}

func Test_Download_002(t *testing.T) {
	if token, err := getDownloadToken(); err != nil {
		t.Error(err)
	} else if client := ch.New(token).SetDebug(true); client == nil {
		t.Error("Invalid client == nil")
	} else if files, err := client.FileList(); err != nil {
		t.Error(err)
	} else if len(files) > 0 {
		if dir, err := ioutil.TempDir("", "clubhouse"); err != nil {
			t.Error(err)
		} else if filename, err := client.Download("GET", files[0].URL, dir); err != nil {
			t.Error(err)
		} else {
			t.Log("url", files[0].URL)
			t.Log("filename", filename)
		}
	}
}

/////////////////////////////////////////////////////

func getDownloadToken() (string, error) {
	if user, err := user.Current(); err != nil {
		return "", err
	} else if bytes, err := ioutil.ReadFile(filepath.Join(user.HomeDir, ".clubhouse")); err != nil {
		return "", err
	} else {
		return strings.TrimSpace(string(bytes)), nil
	}
}
