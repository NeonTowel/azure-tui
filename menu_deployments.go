package main

import "github.com/rivo/tview"

// Function to create Deployments submenu
func (m *MenuManager) createDeploymentsMenu() *tview.List {
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
	return deploymentsMenu
}

// Function to show Add Deployment form
func (m *MenuManager) showAddDeploymentForm() {
	deployment := AzureDeployment{}
	form := tview.NewForm().
		AddInputField("Deployment Name", "", 20, nil, func(text string) {
			deployment.Name = text
		}).
		AddInputField("Version", "", 20, nil, func(text string) {
			deployment.Version = text
		}).
		AddDropDown("Environment", m.parent.listEnvironments(), 0, func(text string, index int) {
			deployment.Environment = text
		}).
		AddButton("Save", func() {
			m.parent.addDeployment(deployment.Environment, deployment)
			m.showAddDeploymentSaveConfirmation()
		}).
		AddButton("Cancel", func() {
			m.app.SetRoot(m.deploymentsMenu, true)
		})

	form.SetBorder(true).SetTitle("Add New Deployment").SetTitleAlign(MenuDefaultAlign)
	m.app.SetRoot(form, true)
}

// Function to show a confirmation modal
func (m *MenuManager) showAddDeploymentSaveConfirmation() {
	modal := tview.NewModal().
		SetText("Deployment saved successfully!").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			m.app.SetRoot(m.mainMenu, true)
		})

	m.app.SetRoot(modal, true)
}
