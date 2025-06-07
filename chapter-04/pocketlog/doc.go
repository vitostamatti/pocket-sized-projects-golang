/*
Package pocketlog expoes an API to log your work.

First, instantiate a logger with pocketlog.New giving it a threshold level.
Messages of lesser criticality won't be logged.

# Sharing the logger is the responsiblity of the caller

The logger can be called to log messages on three levels:

- Debug: mostly used for debugging purposes
- Info: used to log information deemed valuable
- Error: used to log errors

Each of these methods takes a format string and a variadic number of arguments.
*/
package pocketlog
