@echo off
SETLOCAL ENABLEEXTENSIONS ENABLEDELAYEDEXPANSION

REM === CONFIGURAÇÕES ===
set SERVICE_NAME=GoCleanerService
set DISPLAY_NAME=GoCleaner - Limpeza Automática
set DESCRIPTION=Serviço de limpeza automática de arquivos com agendamento embutido via gocron
set EXE_PATH=C:\gocleaner\gocleaner.exe

REM === VERIFICAR BINÁRIO ===
IF NOT EXIST "%EXE_PATH%" (
    echo [ERRO] Arquivo %EXE_PATH% n\u00e3o encontrado.
    echo Verifique se gocleaner.exe est\u00e1 no caminho correto.
    EXIT /B 1
)

REM === CRIAR SERVI\u00c7O ===
echo Criando o servi\u00e7o %SERVICE_NAME%...

sc create %SERVICE_NAME% binPath= "%EXE_PATH%" start= auto DisplayName= "%DISPLAY_NAME%"
sc description %SERVICE_NAME% "%DESCRIPTION%"

REM === CONFIGURAR RECUPERA\u00c7\u00c3O AUTOM\u00c1TICA ===
sc failure %SERVICE_NAME% reset= 60 actions= restart/5000/restart/5000/restart/5000

REM === INICIAR SERVI\u00c7O ===
echo Iniciando o servi\u00e7o...
sc start %SERVICE_NAME%

echo.
echo ✅ Servi\u00e7o %SERVICE_NAME% instalado com sucesso!
echo 🔁 O gocleaner.exe ser\u00e1 executado automaticamente e reiniciado em caso de falha.
echo.
pause
