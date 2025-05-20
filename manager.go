package main

import "github.com/rivo/tview"

func (m *AzureEnvironmentManager) self() *AzureEnvironmentManager {
	return m
}

func (m *AzureEnvironmentManager) getCompanyName() string {
	return m.settings.CompanyName
}

func (m *AzureEnvironmentManager) getTenantID() string {
	return m.settings.TenantID
}

func (m *AzureEnvironmentManager) setCompanyName(name string) {
	m.settings.CompanyName = name
}

func (m *AzureEnvironmentManager) setTenantID(id string) {
	m.settings.TenantID = id
}

func (m *AzureEnvironmentManager) getConfig() *Config {
	return &Config{
		Settings:     m.settings,
		Environments: m.environments,
		Deployments:  m.deployments,
	}
}

func (m *AzureEnvironmentManager) getEnvironments() []AzureEnvironment {
	return m.environments
}

func (m *AzureEnvironmentManager) listEnvironments() []string {
	environments := []string{}
	for _, environment := range m.environments {
		environments = append(environments, environment.Name)
	}
	return environments
}

func (m *AzureEnvironmentManager) addEnvironment(environment AzureEnvironment) {
	m.environments = append(m.environments, environment)
	saveConfig(m.getConfig(), "config.yaml")
}

func (m *AzureEnvironmentManager) updateEnvironment(oldEnvironment AzureEnvironment, newEnvironment AzureEnvironment) {
	for i, e := range m.environments {
		if e.Name == oldEnvironment.Name {
			m.environments[i] = newEnvironment
			break
		}
	}
	saveConfig(m.getConfig(), "config.yaml")
}

func (m *AzureEnvironmentManager) deleteEnvironment(environment AzureEnvironment) {
	for i, e := range m.environments {
		if e.Name == environment.Name {
			m.environments = append(m.environments[:i], m.environments[i+1:]...)
			break
		}
	}
	saveConfig(m.getConfig(), "config.yaml")
}

func (m *AzureEnvironmentManager) addDeployment(environmentName string, deployment AzureDeployment) {
	m.deployments[environmentName] = append(m.deployments[environmentName], deployment)
	saveConfig(m.getConfig(), "config.yaml")
}

func (m *AzureEnvironmentManager) init() {
	config, err := loadConfig("config.yaml")
	if err != nil {
		config = &Config{
			Settings:     AzureSettings{CompanyName: "Contoso Ltd", TenantID: "12345678-90ab-cdef-1234-567890abcdef"},
			Environments: []AzureEnvironment{},
			Deployments:  map[string][]AzureDeployment{},
		}
	}

	m.view = tview.NewApplication()
	m.environments = config.Environments
	m.deployments = config.Deployments
	m.settings = config.Settings
	m.menu = NewMenuManager(m.view, m.self())
}
