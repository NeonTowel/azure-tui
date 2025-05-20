package main

import "github.com/rivo/tview"

// Function to create Azure settings submenu
func (m *MenuManager) createAzureSettingsMenu() *tview.Form {
	azureSettingsForm := tview.NewForm().
		AddInputField("Company Name", m.parent.getCompanyName(), 50, nil, func(text string) {
			m.parent.setCompanyName(text)
		}).
		AddInputField("Tenant ID", m.parent.getTenantID(), 50, nil, func(text string) {
			m.parent.setTenantID(text)
		}).
		AddButton("Save", func() {
			saveConfig(m.parent.getConfig(), "config.yaml")
			m.app.SetRoot(m.mainMenu, true)
		}).
		AddButton("Back", func() {
			m.app.SetRoot(m.mainMenu, true)
		})

	azureSettingsForm.SetBorder(true).SetTitle("Azure Settings").SetTitleAlign(MenuDefaultAlign)
	return azureSettingsForm
}
