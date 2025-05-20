package main

import (
	"fmt"

	"github.com/rivo/tview"
)

const (
	EnvironmentDefaultLocation = "swedencentral"
)

// Function to create Environments submenu
func (m *MenuManager) createEnvironmentsMenu() *tview.List {

	environments := m.parent.getEnvironments()
	environmentsMenu := tview.NewList()
	idx := 1
	for _, environment := range environments {
		environmentsMenu.AddItem(environment.Name, fmt.Sprintf("Subscription ID: %s, Default Location: %s", environment.SubscriptionID, environment.DefaultLocation), rune('0'+idx), func() {
			m.showEditEnvironmentForm(environment)
		})
		idx += 1
	}

	environmentsMenu.
		AddItem("+ Add", "Add a new environment", '+', func() {
			m.showAddEnvironmentForm()
		}).
		AddItem("Back", "", 'b', func() {
			m.app.SetRoot(m.mainMenu, true)
		})

	environmentsMenu.SetBorder(true).SetTitle("Environments").SetTitleAlign(MenuDefaultAlign)
	return environmentsMenu
}

// Function to show Add Environment form
func (m *MenuManager) showAddEnvironmentForm() {
	environment := AzureEnvironment{}
	form := tview.NewForm().
		AddInputField("Name", "", 40, nil, func(text string) {
			environment.Name = text
		}).
		AddInputField("Subscription ID", "", 40, nil, func(text string) {
			environment.SubscriptionID = text
		}).
		AddInputField("Default Location", EnvironmentDefaultLocation, 20, nil, func(text string) {
			environment.DefaultLocation = text
		}).
		AddButton("Save", func() {
			if environment.Name == "" {
				m.app.SetRoot(tview.NewModal().SetText("Name is required").AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					m.app.SetRoot(m.environmentsMenu, true)
				}), true)
				return
			}
			if environment.SubscriptionID == "" {
				m.app.SetRoot(tview.NewModal().SetText("Subscription ID is required").AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					m.app.SetRoot(m.environmentsMenu, true)
				}), true)
				return
			}
			if environment.DefaultLocation == "" {
				environment.DefaultLocation = EnvironmentDefaultLocation
			}
			m.parent.addEnvironment(environment)
			m.showAddEnvironmentSaveConfirmation(environment)
		}).
		AddButton("Cancel", func() {
			m.app.SetRoot(m.environmentsMenu, true)
		})

	form.SetBorder(true).SetTitle("Add New Environment").SetTitleAlign(MenuDefaultAlign)
	m.app.SetRoot(form, true)
}

// Function to show a confirmation modal
func (m *MenuManager) showAddEnvironmentSaveConfirmation(environment AzureEnvironment) {
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Environment '%s' saved successfully!", environment.Name)).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			m.environmentsMenu = m.createEnvironmentsMenu()
			m.app.SetRoot(m.environmentsMenu, true)
		})

	m.app.SetRoot(modal, true)
}

// Function to show Edit Environment form
func (m *MenuManager) showEditEnvironmentForm(environment AzureEnvironment) {
	oldEnvironment := environment
	newEnvironment := environment
	form := tview.NewForm().
		AddInputField("Name", environment.Name, 40, nil, func(text string) {
			newEnvironment.Name = text
		}).
		AddInputField("Subscription ID", environment.SubscriptionID, 40, nil, func(text string) {
			newEnvironment.SubscriptionID = text
		}).
		AddInputField("Default Location", environment.DefaultLocation, 20, nil, func(text string) {
			newEnvironment.DefaultLocation = text
		}).
		AddButton("Save", func() {
			if newEnvironment.Name == "" {
				m.app.SetRoot(tview.NewModal().SetText("Name is required").AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					m.app.SetRoot(m.environmentsMenu, true)
				}), true)
				return
			}
			if newEnvironment.SubscriptionID == "" {
				m.app.SetRoot(tview.NewModal().SetText("Subscription ID is required").AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					m.app.SetRoot(m.environmentsMenu, true)
				}), true)
				return
			}
			m.parent.updateEnvironment(oldEnvironment, newEnvironment)
			m.showEditEnvironmentSaveConfirmation(oldEnvironment)
		}).
		AddButton("Delete", func() {
			m.parent.deleteEnvironment(oldEnvironment)
			m.showDeleteEnvironmentSaveConfirmation(oldEnvironment)
		}).
		AddButton("Cancel", func() {
			m.app.SetRoot(m.environmentsMenu, true)
		})

	form.SetBorder(true).SetTitle("Edit Environment").SetTitleAlign(MenuDefaultAlign)
	m.app.SetRoot(form, true)
}

// Function to show a confirmation modal for editing
func (m *MenuManager) showEditEnvironmentSaveConfirmation(environment AzureEnvironment) {
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Environment '%s' updated successfully!", environment.Name)).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			m.environmentsMenu = m.createEnvironmentsMenu()
			m.app.SetRoot(m.environmentsMenu, true)
		})

	m.app.SetRoot(modal, true)
}

func (m *MenuManager) showDeleteEnvironmentSaveConfirmation(environment AzureEnvironment) {
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Environment '%s' deleted successfully!", environment.Name)).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			m.environmentsMenu = m.createEnvironmentsMenu()
			m.app.SetRoot(m.environmentsMenu, true)
		})

	m.app.SetRoot(modal, true)
}
