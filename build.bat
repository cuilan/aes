@echo on

REM go build this application
REM @auther zhang.yan
REM @date 2022-04-19
TITLE build

rd /q /s dist
mkdir dist

go build -o dist/aes.exe

REM PAUSE
