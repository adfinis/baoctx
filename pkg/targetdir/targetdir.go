package targetdir

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
)

// TargetHome returns the location of the target folder, usually $HOME/.target
func TargetHome() string {
	return fmt.Sprintf("%s/.target", xdg.Home)
}

// TargetHomeCreate checks for the target directory
// and profiles.json file and creates if they don't exist
func TargetHomeCreate() {
	var defaultConfig = "{\n\t\"openbao\": {}\n}"
	targetHome := TargetHome()
	if _, err := os.Stat(targetHome); os.IsNotExist(err) {
		os.Mkdir(targetHome, 0755) //nolint:errcheck
	}

	f := fmt.Sprintf("%s/profiles.json", targetHome)

	if _, err := os.Stat(f); os.IsNotExist(err) {
		// Create and write the default configuration to the file
		err := os.WriteFile(f, []byte(defaultConfig), 0644)
		if err != nil {
			fmt.Printf("Error creating and writing to profiles.json: %v\n", err)
		}
	}

	defaultsDir := targetHome + "/defaults"
	if _, err := os.Stat(defaultsDir); os.IsNotExist(err) {
		os.Mkdir(defaultsDir, 0755) //nolint:errcheck
	}

	tokensDir := targetHome + "/tokens"
	if _, err := os.Stat(tokensDir); os.IsNotExist(err) {
		os.Mkdir(tokensDir, 0700) //nolint:errcheck
	}
}
