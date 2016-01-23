package upgrade

// All upgrades to be performed. Inputs ignored.
func All(...string) (err error) {
	if err = Database(); nil != err {
		return
	}

	return
}
