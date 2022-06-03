package main

//todo: prompt user for configuration parameters.
//todo: split the code in small functions.

import (
	"fmt"
	"log"
	"os"
)

// join home directory and user directory
// create a folder PT.SE.RenderAndProps
var configPath = "C:\\Users\\recs\\OneDrive - Premier Tech\\Bureau\\Golang\\userConfig.json"

func main() {
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

	fmt.Println("[+] userConfig.json created successfully")
	defer file.Close()

}
