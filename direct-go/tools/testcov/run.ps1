param(
    [string]$CoverProfile = "coverage.out",
    [string]$Html = "coverage.html",
    [string[]]$Packages = @("./...")
)

$ErrorActionPreference = "Stop"

# direct-go ルートに移動（このスクリプトは tools/testcov/ 配下にある前提）
Push-Location (Join-Path $PSScriptRoot "..\\..")
try {
    Write-Host "Running go test with coverage..." -ForegroundColor Cyan
    go test -coverprofile=$CoverProfile @Packages

    Write-Host "`nFunction coverage summary:" -ForegroundColor Cyan
    go tool cover -func=$CoverProfile

    if ($Html) {
        go tool cover -html=$CoverProfile -o $Html
        Write-Host "`nHTML report generated at $Html" -ForegroundColor Green
    }
}
finally {
    Pop-Location
}
