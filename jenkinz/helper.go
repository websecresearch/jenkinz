package jenkinz

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

var (
	Version = "-> jenkinz v1.0\n-> by Corben Leo (@hacker_)\n\n"
	home    = HomeDir()
)

func Usage() {
	help := `Usage: jenkinz -d https://<jenkins-instance> [options]

  -d string
		url of jenkins instance		

  -creds string
		Credentials for Jenkins instance (format = username:apikey)
	
  -c int
		Number of concurrent fetchers (default 20)

  -timeout int
		Timeout for the tool in seconds (default 30)`
	fmt.Printf(help)
}

func Exists(loc string) bool {
	_, err := os.Stat(loc)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true

}
func HomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return usr.HomeDir
}
func CreateHost(dir string) {
	CreateDir("output")
	dir = fmt.Sprintf("output/%s", dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}
func CreateDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}
