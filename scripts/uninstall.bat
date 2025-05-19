@echo off
SETLOCAL

set SERVICE_NAME=GoCleanerService

echo.
echo 🛑 Parando o serviço %SERVICE_NAME%...

sc stop %SERVICE_NAME%
timeout /t 3 /nobreak >nul

echo 🧹 Removendo o serviço %SERVICE_NAME%...
sc delete %SERVICE_NAME%

echo.
echo ✅ Serviço %SERVICE_NAME% removido com sucesso.
echo.
pause
