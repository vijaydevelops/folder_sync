# Folder Sync

```
To copy folder contents recursively 
from left folder
to right folder comparing right folder contents 
```

## Tech stack
Go lang <br>
+ 
<b>Fyne</b> (for native cross platform support) <br>


## Pre-requisites
```bash
sudo apt update

sudo apt install -y libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libxxf86vm-dev

sudo snap install go --classic
go version

go install fyne.io/fyne/v2/cmd/fyne@latest

echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
```

## Project setup

```bash
go mod init folder-sync

go mod tidy

go run main.go
go build -o folder-sync main.go

./folder-sync
```