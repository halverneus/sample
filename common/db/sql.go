package db

const (

	// sqlVersionCreate is used to create a new version tracking table.
	sqlVersionCreate = `
CREATE TABLE IF NOT EXISTS dbversion (
	dbkey VARCHAR(40) UNIQUE NOT NULL PRIMARY KEY,
	dbversion INT
)
`

	// sqlVersionUpdate is used to insert or update a version key.
	// IN: [Key], [Version]
	sqlVersionUpdate = `
REPLACE INTO dbversion (
	dbkey, dbversion
) VALUES (
	?, ?
)
`

	// sqlVersionSelect is used to retrieve a version for a given key.
	// IN: [Key]
	// OUT: [Version]
	sqlVersionSelect = "SELECT dbversion FROM dbversion WHERE dbkey=?"
)
