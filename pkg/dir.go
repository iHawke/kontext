package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

func Directory() {
	// Read dir
	dirName := Args.Dir
	newConfigs := ReadDir(dirName)

	// Merge the configs
	for _, newConfig := range newConfigs {
		MergeContexts(newConfig)
	}

	Commit(true)
}

func ReadDir(dirName string) (configs []*Config) {
	files := dir(dirName)
	for _, f := range files {
		Info("Processing [" + f + "]")
		c, e := ReadConfig(f)

		if e != nil {
			msg := fmt.Sprintf("Failed to read: [%s]...", f)
			Warn(msg)
			continue
		}

		configs = append(configs, c)
	}
	return
}

func dir(dirName string) (ymlFiles []string) {
	Info("Reading directory [" + dirName + "}")
	ymlFiles, err := filepath.Glob(dirName + "/*.yml")
	if err != nil {
		pwd, _ := os.Getwd()
		Error("Could not read dir [" + dirName + "], PWD [" + pwd + "].  Perhaps try the absolute path. 😨")
		Fatal("Error: " + err.Error())
	}
	yamlFiles, err := filepath.Glob(dirName + "/*.yaml")
	if err != nil {
		pwd, _ := os.Getwd()
		Error("Could not read dir [" + dirName + "], PWD [" + pwd + "].  Perhaps try the absolute path. 😨")
		Fatal("Error: " + err.Error())
	}
	ymlFiles = append(ymlFiles, yamlFiles...)

	msg := fmt.Sprintf("Found [%d] files", len(ymlFiles))
	Info(msg)
	return
}
