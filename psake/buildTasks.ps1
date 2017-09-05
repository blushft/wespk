Properties {
    $clientDir = "C:\Users\Tom\projects\go\src\github.com\blushft\wespk\client"
    $serverDir = "C:\Users\Tom\projects\go\src\github.com\blushft\wespk"
    $target = "dev"
}

Task Default -Depends BuildDev

Task BuildDev -Depends ClientBuild

Task ClientBuild {
    Exec { & cmd /c "cd $clientDir && npm run build" }
}

Task ServerBuild -depends ClientBuild {
    Exec { go build "$serverDir" }
}