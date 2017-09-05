Import-Module psake

Invoke-psake ./psake/buildTasks.ps1

& go build
.\wespk.exe