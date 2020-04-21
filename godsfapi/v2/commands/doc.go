/*
Package commands implements all commands that can be sent to the server.
Most of the commands include a BaseCommand unnamed member. To obtain correctly
initialized command instances the user is strongly advised to use the prodived
NewCommandName() functions instead of creating a new instance of the according
struct.
*/
package commands
