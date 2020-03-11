# godsfapi
Implementation of DuetAPIClient in Go

## Differences
* A few functionalities had to be left out since there was no good representation in Go
* Since Go has no implicit type conversion there will be As<Type>() methods provided instead
* Currently there is no notification mechanism for object model updates
* In some cases zero values were chosen instead of nil that would be used by upstream
* Geometry was renamed to Kinematics
