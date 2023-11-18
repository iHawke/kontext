package pkg

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

const KONTEXT_FILE = "KONTEXT_FILE"

type Cluster struct {
	CertificateAuthorityData string `yaml:"certificate-authority-data,omitempty"`
	CertificateAuthority     string `yaml:"certificate-authority,omitempty"`
	Server                   string `yaml:"server"`
}

type NamedCluster struct {
	*Cluster `yaml:"cluster,omitempty"`
	Name     string `yaml:"name"`
}

type Context struct {
	Cluster   string `yaml:"cluster"`
	User      string `yaml:"user"`
	Namespace string `yaml:"namespace,omitempty"`
}

type NamedContext struct {
	*Context `yaml:"context"`
	Name     string `yaml:"name"`
}

type Exec struct {
	ApiVersion string      `yaml:"apiVersion"`
	Args       []string    `yaml:"args"`
	Command    string      `yaml:"command"`
	Env        interface{} `yaml:"env"`
}

type User struct {
	Token                 string `yaml:"token,omitempty"`
	ClientCertificate     string `yaml:"client-certificate,omitempty"`
	ClientKey             string `yaml:"client-key,omitempty"`
	ClientCertificateData string `yaml:"client-certificate-data,omitempty"`
	ClientKeyData         string `yaml:"client-key-data,omitempty"`
	*Exec                 `yaml:"exec,omitempty"`
}

type NamedUser struct {
	Name  string `yaml:"name"`
	*User `yaml:"user,omitempty"`
}

type Preferences struct {
}

type Config struct {
	ApiVersion     string         `yaml:"apiVersion"`
	Kind           string         `yaml:"kind"`
	Preferences    Preferences    `yaml:"preferences"`
	Clusters       []NamedCluster `yaml:"clusters,omitempty"`
	Contexts       []NamedContext `yaml:"contexts"`
	CurrentContext string         `yaml:"current-context"`
	Users          []NamedUser    `yaml:"users"`
}

var DefaultConfig *Config
var DefaultConfigFileName string

func ReadDefaultConfig() {
	if CMD == NS || CMD == NoArgs {
		IS_DEBUG = false
	}

	defaultConfigFileName := os.Getenv(KONTEXT_FILE)

	if defaultConfigFileName != "" {
		Debug("config file [" + DefaultConfigFileName + "] found in KONTEXT_FILE env variable.")
	} else {
		Debug("KONTEXT_FILE env variable does not have 'config' file location.")
		Debug("Looking for [$HOME/.kube/config file]")
		homeDir, err := os.UserHomeDir()
		if err != nil {
			Fatal("Failed find User-Home-Dir!  Where is the config file? 🤔")
		}
		defaultConfigFileName = homeDir + "/.kube/config"
	}

	if _, err := os.Stat(defaultConfigFileName); os.IsNotExist(err) {
		Fatal("Could not read [" + defaultConfigFileName + "], permission issue? 🤔")
	}

	DefaultConfigFileName = defaultConfigFileName
	config, err := ReadConfig(defaultConfigFileName)
	if err != nil {
		Fatal(err.Error())
	}
	DefaultConfig = config
}

func ReadConfig(configFileName string) (config *Config, err error) {
	input, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return
	}

	err = yaml.Unmarshal([]byte(input), &config)
	if err != nil {
		return
	}
	return
}

func backup() {
	baseDir := filepath.Dir(DefaultConfigFileName)

	// Let's backup
	ext := time.Now().Format("2006-01-02T15:04:05")
	backupFile := baseDir + "/config." + ext

	source, err := os.Open(DefaultConfigFileName)
	defer source.Close()

	destination, err := os.Create(backupFile)
	if err != nil {
		Fatal("Could not create backup file. Check permission. Error: " + err.Error())
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		msg := fmt.Sprintf("Error creating backup file [%s]! No change was made! Error: [%s]", backupFile, err.Error())
		Fatal(msg)
	}
}

func Commit(doBackup bool) {
	if doBackup {
		backup()
	}

	content, err := yaml.Marshal(&DefaultConfig)
	if err != nil {
		Fatal("error: %v", err)
	}

	err = ioutil.WriteFile(DefaultConfigFileName, content, 0644)
	if err != nil {
		Fatal("Error creating backup file!  Err:", err)
	}
}
