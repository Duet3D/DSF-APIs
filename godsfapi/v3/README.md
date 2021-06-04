# godsfapi
Implementation of DuetAPIClient in Go

## Differences
* A few functionalities had to be left out since there was no good representation in Go
* Since Go has no implicit type conversion there will be As<Type>() methods provided instead
* Currently there is no notification mechanism for object model updates
* In some cases zero values were chosen instead of nil that would be used by upstream
* Geometry was renamed to Kinematics


Go:
Workflow (neovim + vim-go Plugin)

    Port der Änderungen von C# nach Go
    Bei Änderungen am ObjectModel noch die JSON Tags hinzufügen (:GoAddTags mit Cursor innerhalb des struct)
    Compile mit <Plug>(go-build), um zu sehen, ob alles passt, ggf. Bugs fixen
    Compile der Testanwendung in cmd/examples mit go build main.go und dann die (aktuell) drei Modi testen und schauen, ob der Output passt.

In cmd/examples/main.go kann der SockPath geändert werden.

