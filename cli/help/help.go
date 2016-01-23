package help

const (

	// Overall containing help documentation.
	Overall = `
NAME
    sample

SYNOPSIS
    sample [ configuration_file ] command
    sample help

DESCRIPTION
    Sample provides a demo of some of the architectual and syntactic lessons
    learned while developing with Go. Some of the lessons had to be omitted due
    to the proprietary nature of the original code. The code included permits
    a user to add a user, authenticate and upload/download/list/remove files and
    file revisions.

DEPENDENCIES
    MariaDB 10.x or higher.

COMMAND OVERVIEW
    sample help
        USAGE: sample help
        DESCRIPTION: Prints application documentation to the console.

    sample [ configuration_file ] config [ noun ] [ verb ] [ options... ]
        USAGE: sample /etc/sample.yaml config user add username password
        DESCRIPTION: Allows the user to manipulate the configuration from the
        command-line. Additional documentation is available by executing:

            sample config help

    sample [ configuration_file ] run
        USAGE: sample /etc/sample.yaml run
        DESCRIPTION: Run Sample as a service.

    sample [ configuration_file ] upgrade [ noun ]
        USAGE: sample /etc/sample.yaml upgrade all
        DESCRIPTION: Performs upgrade actions. Available upgradeable targets
        include: all, database.

FILES
    configuration_file (ex: sample.yaml)
        This file contains all configuration settings required by the Sample
        application. File is in YAML format.

        EXAMPLE:
            database:
                host: 192.168.1.10  // Host name of the MariaDB 10.x database
                port: 3306          // Port of the MariaDB 10.x database
                name: sample        // Schema name created for Sample
                user: sampleuser    // User Sample uses for database access
                password: 12345     // Password Sample uses for database access
                max-connections: 8  // Max number of allowed database connections
            storage:
                folder: myfolder/   // Folder to store files.
            bind: :8080             // Bind to IP/Port.
`
)
