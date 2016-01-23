package usersql

const (

	// Add a new user.
	// IN: [username] [password]
	// OUT: ID
	Add = "SELECT HEX(user_add(?,?))"

	// Authorize returns a user's ID upon successful authorization.
	Authorize = `
SELECT HEX(id) FROM user
WHERE name=LOWER(?) AND password=encrypt(LOWER(?), ?)
`

	// Delete an existing user.
	// IN: ID
	Delete = "DELETE FROM user WHERE id=UNHEX(?)"
)
