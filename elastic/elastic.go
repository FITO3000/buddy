package elastic

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-8.2.0-windows-x86_64.zip
// https://artifacts.elastic.co/downloads/kibana/kibana-8.2.0-windows-x86_64.zip

// https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-8.2.0-linux-x86_64.tar.gz
// https://artifacts.elastic.co/downloads/kibana/kibana-8.2.0-linux-x86_64.tar.gz

const (
	VERSION       = "8.2.0"
	WIN_ZIP_URL   = "https://artifacts.elastic.co/downloads/%[1]s/%[1]s-%[2]s-windows-x86_64.zip"
	WIN_FILE_NAME = "%s-%s-windows-x86_64.zip"
)

type Installer struct {
	Dir string
	Url string
}

func NewInstaller(dir string) *Installer {
	return &Installer{
		Dir: dir,
	}
}

func (e *Installer) download(name string) error {
	destination := e.downloadDestination(name)
	if fileExists(destination) {
		fmt.Printf("%s already exists\n", destination)
		return nil
	} else {
		err := os.MkdirAll(path.Dir(destination), 0777)
		if err != nil {
			return err
		}
		out, err := os.Create(destination)
		if err != nil {
			return err
		}
		defer out.Close()
		resp, err := http.Get(fmt.Sprintf(WIN_ZIP_URL, name, VERSION))
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		_, err = io.Copy(out, resp.Body)
		return err
	}
}

func (e *Installer) downloadDestination(app string) string {
	return path.Join(e.Dir, "download", fmt.Sprintf(WIN_FILE_NAME, app, VERSION))
}

func (e *Installer) InstallElasticsearch() error {
	return e.download("elasticsearch")
}

func (e *Installer) InstallKibana() error {
	return e.download("kibana")
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		if err == nil {
			return true
		} else {
			panic(err)
		}
	}

}
