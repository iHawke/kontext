package pkg

import (
	"github.com/atotto/clipboard"
	"gopkg.in/yaml.v3"
)

func PasteFromClipboard() {
	Info("Reading clipboard...")
	text, err := clipboard.ReadAll()
	if err != nil {
		Fatal("Could not read from clipboard.  Perhaps add the context from a file...")
	}

	var newConfig Config
	err = yaml.Unmarshal([]byte(text), &newConfig)
	if err != nil {
		Fatal("Content from clipboard could not be unmarshalled.  Error: " + err.Error())
	}

	MergeContexts(&newConfig)
	Commit(true)
}
