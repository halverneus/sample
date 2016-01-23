package help

const (

	// ConfigNoun contains the help documentation for editing the configuration
	// from the command-line.
	ConfigNoun = `
NAME
    sample config

SYNOPSIS
    sample [ configuration_file ] config [ noun ] [ verb ] [ options... ]
    sample config help

DESCRIPTION
    Allows the user to manipulate the configuration from the command-line.

COMMAND OVERVIEW
    sample config help
        USAGE: sample config help
        DESCRIPTION: Prints available configuration noun documentation to the
        console.

    sample [ configuration_file ] config user [ verb ] [ options... ]
        USAGE: sample /etc/sample.yaml config user add username password
        DESCRIPTION: Allows the user to manipulate user configurations from the
        command-line. Additional documentation is available by executing:

            sample config user help
`

	// ConfigUser contains the help documentation for editing user configurations
	// from the command-line.
	ConfigUser = `
NAME
    sample config user

SYNOPSIS
    sample [ configuration_file ] config user [ verb ] [ options... ]
    sample config user help

DESCRIPTION
    Allows the user to manipulate user configurations from the command-line.

COMMAND OVERVIEW
    sample config user help
        USAGE: sample config user help
        DESCRIPTION: Prints available configuration documentation for
        manipulating users to the console.

    sample [ configuration_file ] config user add [ user_name ] [ password ]
        USAGE sample /etc/sample.yaml config user add username password
        DESCRIPTION: Allows the administrator to add a user to the system and to
        set their password.
`
)
