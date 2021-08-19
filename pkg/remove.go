package pkg

func RemoveContext() {
	errMsg := ContextExist(Args.Remove)
	if errMsg != "" {
		Fatal(errMsg)
	}
	deleteContext(Args.Remove)
	Commit(true)
}
