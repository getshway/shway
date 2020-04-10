package homebrew

func IsAreadyUpdate(err error) bool {
	return err.Error() == "Already up-to-date."
}
