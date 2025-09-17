// File: atlas.go - Main Atlas Entry Point
// Template: pkg/app/entry_point.go
// Purpose: Public interface for Atlas functionality providing main Start() function
// Reference: Follows lazygit's entry point pattern for application startup
package atlas

import (
	"github.com/jesseduffield/lazygit/pkg/app"
	integrationTypes "github.com/jesseduffield/lazygit/pkg/integration/types"

	atlasApp "github.com/kaka-ruto/atlas/pkg/atlas/app"
)

// Start is the main Atlas entry point that replaces lazygit's app.Start
// This function is called from main.go to start the Atlas application
func Start(buildInfo *app.BuildInfo, integrationTest integrationTypes.IntegrationTest) {
	atlasApp.Start(buildInfo, integrationTest)
}