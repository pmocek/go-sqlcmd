IF EXIST sqlcmd.exe (
  del sqlcmd.exe
)

IF EXIST output (
  rmdir output /s /q
)

IF EXIST release\output (
  rmdir release\output /s /q
)