package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// create a folder PT.SE.RenderAndProps
// var configPath = "C:\\Users\\recs\\OneDrive - Premier Tech\\Bureau\\Golang\\userConfig.json"
// var configPath = "C:\\repos\\userConfig.json"
var configPath = "C:\\repos\\PT.SE\\PT.SE.RenderAndProps\\userConfig.json"

func trimWhiteSpace(_user string) string {
	return _user[:len(_user)-2]
}

func userPrompter(reader *bufio.Reader, print string) string {
	fmt.Println(print)
	entry, _ := reader.ReadString('\n')
	entry = trimWhiteSpace(entry)
	entry = strings.Trim(entry, `"`)
	entry = strings.ReplaceAll(entry, "\\", "\\\\")
	return entry
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	_user := userPrompter(reader, "Enter User ID:")
	_password := userPrompter(reader, "Enter password:")
	_group := userPrompter(reader, "Enter Group:")
	_role := userPrompter(reader, "Enter Role:")
	_server := userPrompter(reader, "Enter Server:")
	_folder := userPrompter(reader, "Enter the path of the folder where you want to image to be downloaded:")

	file, err := os.Create(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var configContent = fmt.Sprintf(`{"User": "%s","Password": "%s","Group": "%s","Role": "%s","Server": "%s","DownloadFolder": "%s"}`, _user, _password, _group, _role, _server, _folder)
	file.WriteString(configContent)
	defer file.Close()

	fmt.Println("[+] userConfig.json created successfully")
	fmt.Println(configPath)
}
