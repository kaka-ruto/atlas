// File: atlas_app.go - Atlas Application Entry Point
// Template: pkg/app/app.go + pkg/app/entry_point.go
// Purpose: Atlas application lifecycle management and startup logic
// Reference: Combines lazygit's App struct pattern with entry_point.Start() flow
package app

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-errors/errors"

	"github.com/jesseduffield/lazygit/pkg/app"
	appTypes "github.com/jesseduffield/lazygit/pkg/app/types"
	"github.com/jesseduffield/lazygit/pkg/commands/git_commands"
	"github.com/jesseduffield/lazygit/pkg/commands/oscommands"
	"github.com/jesseduffield/lazygit/pkg/common"
	"github.com/jesseduffield/lazygit/pkg/config"
	"github.com/jesseduffield/lazygit/pkg/constants"
	"github.com/jesseduffield/lazygit/pkg/env"
	"github.com/jesseduffield/lazygit/pkg/gui"
	integrationTypes "github.com/jesseduffield/lazygit/pkg/integration/types"
	"github.com/jesseduffield/lazygit/pkg/updates"

	atlasGui "github.com/kaka-ruto/atlas/pkg/atlas/gui"
)

// AtlasApp is the struct that manages Atlas application lifecycle
// It follows the same pattern as lazygit's App but uses AtlasGUI
type AtlasApp struct {
	*common.Common
	closers   []io.Closer
	Config    config.AppConfigurer
	OSCommand *oscommands.OSCommand
	AtlasGUI  *atlasGui.AtlasGUI
}

// NewAtlasApp creates a new Atlas application instance
// This follows the same pattern as app.NewApp but creates AtlasGUI instead of gui.Gui
func NewAtlasApp(config config.AppConfigurer, test integrationTypes.IntegrationTest, common *common.Common) (*AtlasApp, error) {
	app := &AtlasApp{
		closers: []io.Closer{},
		Config:  config,
		Common:  common,
	}

	app.OSCommand = oscommands.NewOSCommand(common, config, oscommands.GetPlatform(), oscommands.NewNullGuiIO(app.Log))

	updater, err := updates.NewUpdater(common, config, app.OSCommand)
	if err != nil {
		return app, err
	}

	dirName, err := os.Getwd()
	if err != nil {
		return app, err
	}

	gitVersion, err := app.validateGitVersion()
	if err != nil {
		return app, err
	}

	showRecentRepos := len(config.GetAppState().RecentRepos) > 0 && app.shouldShowRecentRepos(dirName)

	// used for testing purposes
	if os.Getenv("SHOW_RECENT_REPOS") == "true" {
		showRecentRepos = true
	}

	// Create regular GUI first
	regularGui, err := gui.NewGui(common, config, gitVersion, updater, showRecentRepos, dirName, test)
	if err != nil {
		return app, err
	}

	// Wrap it in AtlasGUI
	app.AtlasGUI, err = atlasGui.NewAtlasGUI(regularGui)
	if err != nil {
		return app, err
	}

	return app, nil
}

// Run runs the Atlas application
func (app *AtlasApp) Run(startArgs appTypes.StartArgs) error {
	err := app.AtlasGUI.RunAndHandleError(startArgs)
	return err
}

// GetAtlasGUI returns the AtlasGUI instance
func (app *AtlasApp) GetAtlasGUI() *atlasGui.AtlasGUI {
	return app.AtlasGUI
}

// Close closes any resources
func (app *AtlasApp) Close() error {
	for _, closer := range app.closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// validateGitVersion validates that git version meets minimum requirements
func (app *AtlasApp) validateGitVersion() (*git_commands.GitVersion, error) {
	gitVersion, err := git_commands.GetGitVersion(app.OSCommand)
	if err != nil {
		return nil, err
	}

	return gitVersion, nil
}

// shouldShowRecentRepos determines if recent repos should be shown
func (app *AtlasApp) shouldShowRecentRepos(dirName string) bool {
	if env.GetGitDirEnv() != "" {
		return false
	}

	// Check if current directory is a git repository
	if isRepo, _ := isDirectoryAGitRepository(dirName); isRepo {
		return false
	}

	return true
}

// isDirectoryAGitRepository checks if directory contains git repository
func isDirectoryAGitRepository(dir string) (bool, error) {
	info, err := os.Stat(fmt.Sprintf("%s/.git", dir))
	return info != nil, err
}

// Start is the Atlas entry point that replaces app.Start
// It follows the exact same pattern as app.Start but uses AtlasApp
func Start(buildInfo *app.BuildInfo, integrationTest integrationTypes.IntegrationTest) {
	// For now, just use the regular app.Start pattern but insert AtlasApp
	// This avoids replicating all the CLI parsing and setup logic

	// Use a minimal approach: set up config and common, then run AtlasApp
	tempDir, err := os.MkdirTemp("", "atlas_temp")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer os.RemoveAll(tempDir)

	appConfig, err := config.NewAppConfig("atlas", buildInfo.Version, buildInfo.Commit, buildInfo.Date, buildInfo.BuildSource, false, tempDir)
	if err != nil {
		log.Fatal(err.Error())
	}

	common, err := app.NewCommon(appConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Use AtlasApp instead of regular app
	startArgs := appTypes.NewStartArgs("", appTypes.GitArgNone, "normal", integrationTest)

	atlasApp, err := NewAtlasApp(appConfig, integrationTest, common)
	if err == nil {
		err = atlasApp.Run(startArgs)
	}

	if err != nil {
		// Use the same error handling as regular app
		newErr := errors.Wrap(err, 0)
		stackTrace := newErr.ErrorStack()
		common.Log.Error(stackTrace)

		log.Fatalf("%s: %s\n\n%s", common.Tr.ErrorOccurred, constants.Links.Issues, stackTrace)
	}
}