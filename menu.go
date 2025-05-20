package main

import "github.com/rivo/tview"

// Constants for menu items
const (
	AzureSettingsTitle = "Azure settings"
	EnvironmentsTitle  = "Environments"
	DeploymentsTitle   = "Deployments"
	QuitTitle          = "Quit"
	MenuDefaultAlign   = tview.AlignLeft
)

// MenuManager manages the application's menus
type MenuManager struct {
	app               *tview.Application
	mainMenu          *tview.List
	azureSettingsMenu *tview.Form
	environmentsMenu  *tview.List
	deploymentsMenu   *tview.List
	parent            *AzureEnvironmentManager
}

// NewMenuManager creates a new MenuManager
func NewMenuManager(app *tview.Application, parent *AzureEnvironmentManager) *MenuManager {
	manager := &MenuManager{app: app, parent: parent}
	manager.initMenus()
	return manager
}

// Initialize menus
func (m *MenuManager) initMenus() {
	m.mainMenu = m.createMainMenu()
	m.azureSettingsMenu = m.createAzureSettingsMenu()
	m.environmentsMenu = m.createEnvironmentsMenu()
	m.deploymentsMenu = m.createDeploymentsMenu()
}

// Function to create the main menu
func (m *MenuManager) createMainMenu() *tview.List {
	list := tview.NewList().
		AddItem(AzureSettingsTitle, "Modify Azure settings for the tenant", 'a', func() {
			m.app.SetRoot(m.azureSettingsMenu, true)
		}).
		AddItem(EnvironmentsTitle, "Manage environments for the tenant", 'e', func() {
			m.app.SetRoot(m.environmentsMenu, true)
		}).
		AddItem(DeploymentsTitle, "Manage deployments for the tenant", 'd', func() {
			m.app.SetRoot(m.deploymentsMenu, true)
		}).
		AddItem(QuitTitle, "", 'q', func() {
			m.app.Stop()
		})

	list.SetBorder(true).SetTitle("Azure Environments Manager").SetTitleAlign(MenuDefaultAlign)
	list.SetTitle("Azure Environments Manager")
	list.SetTitleAlign(MenuDefaultAlign)

	return list
}

// Function to show Azure settings submenu
func (m *MenuManager) showAzureSettingsMenu() {
	azureSettingsMenu := tview.NewList().
		AddItem("Setting 1", "", '1', nil).
		AddItem("Setting 2", "", '2', nil).
		AddItem("Back", "", 'b', func() {
			m.app.SetRoot(m.mainMenu, true)
		})

	azureSettingsMenu.SetBorder(true).SetTitle("Azure Settings").SetTitleAlign(MenuDefaultAlign)
	m.app.SetRoot(azureSettingsMenu, true)
}

// Function to show Environments submenu
func (m *MenuManager) showEnvironmentsMenu() {
	environmentsMenu := tview.NewList().
		AddItem("Environment 1", "", '1', nil).
		AddItem("Environment 2", "", '2', nil).
		AddItem("+ Add", "Add a new environment", '+', func() {
			m.showAddEnvironmentForm()
		}).
		AddItem("Back", "", 'b', func() {
			m.app.SetRoot(m.mainMenu, true)
		})

	environmentsMenu.SetBorder(true).SetTitle("Environments").SetTitleAlign(MenuDefaultAlign)
	m.app.SetRoot(environmentsMenu, true)
}

// Function to show Deployments submenu
func (m *MenuManager) showDeploymentsMenu() {
	deploymentsMenu := tview.NewList().
		AddItem("Deployment 1", "", '1', nil).
		AddItem("Deployment 2", "", '2', nil).
		AddItem("+ Add", "Add a new deployment", '+', func() {
			m.showAddDeploymentForm()
		}).
		AddItem("Back", "", 'b', func() {
			m.app.SetRoot(m.mainMenu, true)
		})

	deploymentsMenu.SetBorder(true).SetTitle("Deployments").SetTitleAlign(MenuDefaultAlign)
	m.app.SetRoot(deploymentsMenu, true)
}
