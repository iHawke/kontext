package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

func Directory() {
	// Read dir
	dirName := Args.Dir
	newConfigs, err := ReadDir(dirName)
	if err != nil {
		pwd, _ := os.Getwd()
		Fatal("Could not read dir [" + dirName + "], PWD [" + pwd + "].  Perhaps try the absolute path. 😨")
	}

	// Merge the configs
	for _, newConfig := range newConfigs {
		MergeContexts(newConfig)
	}

	Commit(true)
}

func ReadDir(dirName string) (configs []*Config, err error) {
	files := dir(dirName)
	// configs := make([]*Config, len(files))
	j := 0
	for _, f := range files {
		c, e := ReadConfig(f)

		if e != nil {
			msg := fmt.Sprintf("Failed to read: [%s]...", f)
			Warn(msg)
			continue
		}

		configs[j] = c
		j++
	}
	return
}

func dir(dirName string) (ymlFiles []string) {
	Info("Reading directory [" + dirName + "}")
	ymlFiles, err := filepath.Glob(dirName + "/*.yml")
	if err != nil {
		Fatal(err)
	}
	yamlFiles, err := filepath.Glob(dirName + "/*.yaml")
	if err != nil {
		Fatal(err)
	}
	ymlFiles = append(ymlFiles, yamlFiles...)

	msg := fmt.Sprintf("Found [%d] files", len(ymlFiles))
	Info(msg)
	return
}
