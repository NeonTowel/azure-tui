package main

import "github.com/rivo/tview"

type AzureEnvironment struct {
	Name            string
	SubscriptionID  string
	DefaultLocation string
}

type AzureDeployment struct {
	Name        string
	Version     string
	Environment string
}

type AzureSettings struct {
	CompanyName string
	TenantID    string
}

type AzureEnvironmentManager struct {
	view         *tview.Application
	menu         *MenuManager
	environments []AzureEnvironment
	deployments  map[string][]AzureDeployment
	settings     AzureSettings
}
