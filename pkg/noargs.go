package pkg

import "github.com/AlecAivazis/survey/v2"

func NoArgsRun() {
	contexts, currentContext := ListDefaultContexts()
	if len(contexts) == 0 {
		Fatal("No contexts found in [" + DefaultConfigFileName + "]")
	}

	// The questions to ask
	newContext := ""
	prompt := &survey.Select{
		Message: "Choose a context:",
		Options: contexts,
		Default: currentContext,
	}
	survey.AskOne(prompt, &newContext)

	if newContext != "" && DefaultConfig.CurrentContext != newContext {
		DefaultConfig.CurrentContext = newContext
		Commit(false)
	}
}
