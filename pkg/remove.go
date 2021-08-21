package pkg

func RemoveContext() {
	removed := DeleteContext(Args.Remove)

	if removed {
		Commit(true)
	}
}
