param([switch]$SoloCompilar)

$ErrorActionPreference = 'Stop'

# Go no resuelve los paquetes de este módulo desde algunas carpetas de OneDrive.
# Se compila una copia temporal de las fuentes, sin copiar configuración ni credenciales.
$temporal = (Resolve-Path -LiteralPath ([System.IO.Path]::GetTempPath())).ProviderPath
$nombre = 'agrocampo-noticias-' + [guid]::NewGuid().ToString('N')
$carpeta = Join-Path $temporal $nombre
New-Item -ItemType Directory -Path $carpeta | Out-Null

try {
    $archivos = @('go.mod', 'go.sum', 'controllers', 'models', 'routers')
    $archivos += @(Get-ChildItem -LiteralPath $PSScriptRoot -Filter '*.go' -File | Where-Object Name -NotLike '*_test.go' | Select-Object -ExpandProperty Name)
    foreach ($archivo in $archivos) {
        Copy-Item -LiteralPath (Join-Path $PSScriptRoot $archivo) -Destination $carpeta -Recurse
    }

    $ejecutable = Join-Path $carpeta 'noticias.exe'
    Write-Host 'Compilando Noticias en una carpeta temporal local…'
    Push-Location -LiteralPath $carpeta
    try {
        & go build -o $ejecutable .
        if ($LASTEXITCODE -ne 0) {
            throw "La compilación de Noticias falló (código $LASTEXITCODE)."
        }
    } finally {
        Pop-Location
    }

    if ($SoloCompilar) {
        Write-Host 'Compilación de Noticias correcta.'
        return
    }

    # Desde la carpeta original Beego encuentra conf/app.conf, Swagger y el .env padre.
    Push-Location -LiteralPath $PSScriptRoot
    try {
        & $ejecutable
        if ($LASTEXITCODE -ne 0) {
            throw "El CRUD de Noticias terminó con código $LASTEXITCODE."
        }
    } finally {
        Pop-Location
    }
} finally {
    $rutaVerificada = (Resolve-Path -LiteralPath $carpeta -ErrorAction SilentlyContinue).ProviderPath
    if ($rutaVerificada -eq $carpeta -and (Split-Path -Path $rutaVerificada -Parent) -eq $temporal) {
        Remove-Item -LiteralPath $rutaVerificada -Recurse -Force
    }
}
