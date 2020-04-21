/*
Package initmessages contains all init messages that can be used to initiate
a certain type of connection with DuetControlServer.

These are
- CommandInitMessage
- InterceptInitMessage
- SubscribeInitMessage

Even though all types are public it is strongly advised to use the corresponding NewXYZInitMessage()
functions to get a valid instance of an init message.
*/
package initmessages
