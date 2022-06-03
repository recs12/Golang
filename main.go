package main

//todo: prompt user for configuration parameters.
//todo: split the code in small functions.

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// join home directory and user directory
// create a folder PT.SE.RenderAndProps
var configPath = "C:\\Users\\recs\\OneDrive - Premier Tech\\Bureau\\Golang\\userConfig.json"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter User ID:")
	_user, _ := reader.ReadString('\n')
	fmt.Println(_user)

	fmt.Println("Enter password:")
	_password, _ := reader.ReadString('\n')
	fmt.Println(_password)

	fmt.Println("Enter Group:")
	_group, _ := reader.ReadString('\n')
	fmt.Println(_group)

	fmt.Println("Enter Role:")
	_role, _ := reader.ReadString('\n')
	fmt.Println(_role)

	fmt.Println("Enter Server:")
	_server, _ := reader.ReadString('\n')
	fmt.Println(_server)

	fmt.Println("Enter Folder:")
	_folder, _ := reader.ReadString('\n')
	fmt.Println(_folder)

	file, err := os.Create(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// loop trough the list of string
	configContent := []string{
		"{\n",
		"\t\"User\": \"recs\",\n",
		"\t\"Password\": \"recs\",\n",
		"\t\"Group\": \"Engineering\",\n",
		"\t\"Role\": \"Designer\",\n",
		"\t\"Server\": \"UAT_TC\",\n",
		"\t\"DownloadFolder\": \"C:\\\\Repos\"\n",
		"}\n",
	}

	for _, line := range configContent {
		file.WriteString(line)
	}
	defer file.Close()

	fmt.Println("[+] userConfig.json created successfully")

}
