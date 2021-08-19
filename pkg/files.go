package pkg

import "os"

func OneFile() {
	fileName := Args.File
	newConfig, err := ReadConfig(fileName)
	if err != nil {
		pwd, _ := os.Getwd()
		Fatal("Could not read [" + fileName + "], PWD [" + pwd + "].  Perhaps try the absolute path. 😨")
	}

	// Merge the configs
	MergeContexts(newConfig)

	Commit(true)
}
