package usersql

const (

	// CreateTable for managing users.
	CreateTable = `
CREATE TABLE IF NOT EXISTS user (
	id BINARY(16) PRIMARY KEY NOT NULL,
	name VARCHAR(64) UNIQUE,
	password VARCHAR(512)
)
`

	// CreateEncryptionFunction for simple password encryption.
	CreateEncryptionFunction = `
CREATE FUNCTION encrypt(
	name VARCHAR(64),
	password VARCHAR(64))
RETURNS VARCHAR(512)
NOT DETERMINISTIC
RETURN AES_ENCRYPT(email, SHA2(password, 512));
`

	// CreateIDFunction for generating random binary IDs.
	CreateIDFunction = `
CREATE FUNCTION make_id()
RETURNS BINARY(16)
NOT DETERMINISTIC
RETURN UNHEX(REPLACE(UUID(),'-',''));
`

	// CreateNullFunction returns binary '0' representing a failure.
	CreateNullFunction = `
CREATE FUNCTION make_null()
RETURNS BINARY(16)
RETURN UNHEX('00000000000000000000000000000000');
`

	// AddFunction used to add a new user.
	AddFunction = `
CREATE FUNCTION user_add(
	v_name VARCHAR(128),
	v_password VARCHAR(64))
RETURNS BINARY(16)
NOT DETERMINISTIC
BEGIN
	DECLARE v_id BINARY(16);
	IF EXISTS(SELECT 1 FROM user WHERE name=v_name) THEN
		SELECT make_null() INTO v_id;
	ELSE
		SELECT make_id() INTO v_id;
		INSERT INTO user (
			id,
			name,
			password
		) VALUES (
			v_id,
			LOWER(v_name),
			encrypt(LOWER(v_name), v_password)
		);
	END IF;
	RETURN v_id;
END
`
)
