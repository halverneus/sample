package db

// GetVersion number for a particular schema.
func GetVersion(key string) (version int) {
	version = 0

	// Retrieve and place the schema version into the
	// version variable.

	result, err := Query(sqlVersionSelect, key)
	if nil != err {
		return
	}
	defer result.Free()

	result.ScanNextRow(&version)
	return
}

// SetVersion number for a particular schema.
func SetVersion(key string, version int) error {
	return Execute(sqlVersionUpdate, key, version)
}
