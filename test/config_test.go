package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ihawke/kontext/pkg"
)

func TestCommit(t *testing.T) {
	path, _ := os.Getwd()
	inputConfig := path + "/samples/config1"

	config1, err := pkg.ReadConfig(inputConfig)
	if err != nil {
		t.Fatal(err)
	}

	outputConfig := path + "/samples/config.out"
	pkg.DefaultConfig = config1
	pkg.DefaultConfigFileName = outputConfig
	pkg.Commit(false)

	config2, err := pkg.ReadConfig(outputConfig)
	if err != nil {
		t.Fatal(err)
	}

	ConfigExist(config1, config2, t)
}

func ConfigExist(primaryConfig *pkg.Config, newConfig *pkg.Config, t *testing.T) {

	for _, c := range newConfig.Contexts {
		namedContext := pkg.GetContextByName(c.Name, primaryConfig)
		if namedContext == nil {
			t.Fatal("Context not found in primary config! Name: ", c.Name)
		}
	}

	for _, c := range newConfig.Clusters {
		namedCluster := pkg.GetClusterByName(c.Name, primaryConfig)
		if namedCluster == nil {
			t.Fatal("Cluster not found in primary config! Name: ", c.Name)
		}
	}

	for _, u := range newConfig.Users {
		namedUser := pkg.GetUserByName(u.Name, primaryConfig)
		if namedUser == nil {
			t.Fatal("User not found in primary config! Name: ", u.Name)
		}
	}
}

func TestReadSimpleConfig(t *testing.T) {
	path, _ := os.Getwd()
	configFileName := path + "/samples/config1"

	c, err := pkg.ReadConfig(configFileName)
	if err != nil {
		t.Fatal(err)
	}

	if c == nil {
		t.Fatal("Failed to read file: ", configFileName)
	}
}

func TestReadComplexConfig(t *testing.T) {
	path, _ := os.Getwd()
	configFileName := path + "/samples/config2"

	c, err := pkg.ReadConfig(configFileName)
	if err != nil {
		t.Fatal(err)
	}

	if c == nil {
		t.Fatal("Failed to read file: ", configFileName)
	}
}

func TestArray(t *testing.T) {
	scores := []int{1,2,3,4,5}
	slice := scores[2:4]
	fmt.Println("slice is:", slice)
	slice[0] = 999
	fmt.Println("score is", scores)
	if scores[2] != 999 {
		t.Fatal("Slice does not match underlining array!")
	}
}