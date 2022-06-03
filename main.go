package main

//todo: prompt user for configuration parameters.
//todo: split the code in small functions.

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

// join home directory and user directory
var configPath = "C:\\Users\\recs\\OneDrive - Premier Tech\\Bureau\\Golang\\userConfig.json"

func main() {
	file, err := os.Create(configPath)
	if err != nil {
		log.Fatal(err)
	}

	x := []string{"{\n",
		"\t\"User\": \"recs\",\n",
		"\t\"Password\": \"recs\",\n",
		"\t\"Group\": \"Engineering\",\n",
	}

	file.WriteString("{\n")
	file.WriteString("\t\"User\": \"recs\",\n")
	file.WriteString("\t\"Password\": \"recs\",\n")
	file.WriteString("\t\"Group\": \"Engineering\",\n")
	file.WriteString("\t\"Role\": \"Designer\",\n")
	file.WriteString("\t\"Server\": \"UAT_TC\",\n")
	file.WriteString("\t\"DownloadFolder\": \"C:\\Repos\",\n")
	file.WriteString("{\n")

	fmt.Println("[+] userConfig.json created successfully")
	defer file.Close()

	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := user.Username
	homedir := user.HomeDir
	name := user.Name

	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Homedir: %s\n", homedir)
	fmt.Printf("Name: %s\n", name)
}
