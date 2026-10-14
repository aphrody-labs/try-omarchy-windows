param([string]$Root = 'C:\msys64')
$ErrorActionPreference = 'Stop'
$bash = Join-Path $Root 'usr\bin\bash.exe'
if (!(Test-Path -LiteralPath $bash)) { throw "MSYS2 bash missing: $bash" }
& $bash -lc 'pacman --noconfirm --needed -Syu mingw-w64-ucrt-x86_64-qemu mingw-w64-ucrt-x86_64-zstd'
if ($LASTEXITCODE -ne 0) { throw "MSYS2 package installation failed: $LASTEXITCODE" }
$qemu = Join-Path $Root 'ucrt64\bin\qemu-system-x86_64w.exe'
if (!(Test-Path -LiteralPath $qemu)) { throw "MSYS2 QEMU missing after installation: $qemu" }
[Environment]::SetEnvironmentVariable('TRYOMARCHY_MSYS2_ROOT', $Root, 'User')
Write-Output "Try Omarchy MSYS2 runtime ready: $qemu"
