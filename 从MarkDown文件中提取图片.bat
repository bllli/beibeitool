for %%i in (%*) do (
    beibeitool.exe getAllBase64ImageFromMarkDownFile %%i
)
pause