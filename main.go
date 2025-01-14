package main

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/checks"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/hcl"
	"os"
)

var (
	version string
)

func init() {
	if version == "" {
		version = "dev"
	}
}

func main() {
	exitCode := 0
	defer func() {
		if exitCode == 0 {
			fmt.Println("INFO: terraform-validator ran successfully")
		}
		os.Exit(exitCode)
	}()

	// Get the configuration
	globalConfig := config.GenerateGlobalConfig(version)

	// Get the terraform files informations
	folderParsedContent := hcl.GetFolderParsedContents(globalConfig.WorkDir)

	// Verify files normes and conventions
	for _, fileParsedContent := range folderParsedContent {
		var authorizedBlocks []string
		fileConfig, fileConfigExists := globalConfig.TerraformConfig.Files[fileParsedContent.Name]
		if fileConfigExists {
			authorizedBlocks = fileConfig.AuthorizedBlocks
		} else {
			authorizedBlocks = globalConfig.TerraformConfig.Files["default"].AuthorizedBlocks
		}
		ok := checks.VerifyFile(fileParsedContent,
			globalConfig.TerraformConfig.BlockPatternName,
			authorizedBlocks)

		if !ok {
			exitCode = 1
		}
	}

	if globalConfig.TerraformConfig.EnsureProvidersVersion {
		ok := checks.VerifyProvidersVersion(folderParsedContent)
		if !ok {
			exitCode = 1
		}
	}

	if globalConfig.TerraformConfig.EnsureTerraformVersion {
		ok := checks.VerifyTerraformVersion(folderParsedContent)
		if !ok {
			exitCode = 1
		}
	}
}
