docker ps -a || GOTO error

sqlcmd config view || GOTO error
IF EXIST "c:\users\stuartpa\.sqlcmd" (
    rmdir /s /q c:\users\stuartpa\.sqlcmd
)

sqlcmd install mssql get-tags || GOTO error
sqlcmd install mssql || GOTO error
sqlcmd install mssql --tag 2019-latest --encrypt-password || GOTO error
sqlcmd query "SELECT @@version" || GOTO error
sqlcmd uninstall --yes || GOTO error
sqlcmd create mssql-edge || GOTO error
sqlcmd query "SELECT @@SERVERNAME" || GOTO error
sqlcmd config use-context mssql || GOTO error
sqlcmd config use-context edge || GOTO error
sqlcmd query "SELECT @@SERVERNAME" || GOTO error
sqlcmd config view || GOTO error
sqlcmd config view --raw || GOTO error
sqlcmd config connection-strings || GOTO error
sqlcmd config get-contexts || GOTO error
sqlcmd config get-contexts -o yaml || GOTO error
sqlcmd config get-contexts -o xml || GOTO error
sqlcmd config get-contexts -o json || GOTO error
sqlcmd config get-endpoints || GOTO error
sqlcmd config get-endpoints -o json || GOTO error
sqlcmd config get-users || GOTO error
sqlcmd config get-users -o xml || GOTO error
sqlcmd uninstall --yes --force || GOTO error
sqlcmd install mssql -u foo || GOTO error
sqlcmd query "SELECT DB_NAME()" || GOTO error
sqlcmd config use-context mssql2 || GOTO error
sqlcmd config current-context || GOTO error
sqlcmd install mssql-edge -u foo || GOTO error
sqlcmd query "SELECT DB_NAME()" || GOTO error
sqlcmd install mssql get-tags || GOTO error
sqlcmd install mssql-edge get-tags || GOTO error
sqlcmd uninstall --yes --force  || GOTO error
sqlcmd delete --yes --force || GOTO error
sqlcmd drop --yes --force || GOTO error
sqlcmd drop

echo Tests completed successfully

ECHO NOTE: Confirm config should be empty
sqlcmd config view
docker ps -a

exit /B 0

:error

echo Tests FAILED

exit /B 1