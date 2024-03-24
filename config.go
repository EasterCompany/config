/*
Go Module: config

Allows a developer to contain a standardised configuration for applications
ran by a specific user on a specific system.

Part of the Go Extended Library by Easter Company
designed for use within the Overlord fullstack framework.

Version: 0.0.0
Author: Owen Cameron Easter
*/
package config

import (
	"os"
)

var configFilePath string

type templateConfig struct {
	DEBUG             bool
	LOCAL_BRANCH      string
	STAGING_BRANCH    string
	PRODUCTION_BRANCH string
	TIMEZONE          string
	LANGUAGE_CODE     string
	SECRETKEY         string
	PUBLICKEY         string
}

func SetFilePath(filePath string) string {
	configFilePath = filePath
	return configFilePath
}

func VerifyFileExists() bool {
	if _, err := os.Stat(configFilePath); err == nil {
		return true
	}
	return false
}

func GenerateFileTemplate() templateConfig {
	return templateConfig{true, "Lab", "Dev", "Prd", "utc", "en-gb", "...", "..."}
}

func SaveFile() {
}

func LoadFile() {
}
