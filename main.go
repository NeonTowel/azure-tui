package main

import (
	"fmt"
)

// Main function to initialize the application
func main() {
	manager := &AzureEnvironmentManager{}
	manager.init()
	if err := manager.view.SetRoot(manager.menu.mainMenu, true).Run(); err != nil {
		fmt.Println("Error running application:", err)
	}
}
