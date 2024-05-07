# gotreesize

Simple CLI utility to preview file sizes of a directory

-----

## Installation
#### Build from source
```bash
git clone https://github.com/TeemuTallskog/gotreesize.git

cd gotreesize

go build -o gotreesize.exe
```
Or download executable directly from releases:   
`https://github.com/TeemuTallskog/gotreesize/releases/latest/download/gotreesize.exe`

-----
## Usage
```bash
gotreesize.exe 'C:\Program Files\Git'

Directory: C:\Program Files\Git
---------------------------------
mingw64            |████████████████████████████████████████████████████████████████████████████████| 207.7 MB
usr                |███████████████████████████████████████████████████████████████████████         | 185.0 MB
unins000.exe       |█                                                                               | 3.1 MB
etc                |                                                                                | 1.5 MB
unins000.dat       |                                                                                | 1.2 MB
cmd                |                                                                                | 514.4 KB
ReleaseNotes.ht... |                                                                                | 255.2 KB
git-bash.exe       |                                                                                | 134.0 KB
git-cmd.exe        |                                                                                | 133.5 KB
bin                |                                                                                | 133.0 KB
unins000.msg       |                                                                                | 23.6 KB
LICENSE.txt        |                                                                                | 18.3 KB
dev                |                                                                                | 172 B
tmp                |                                                                                | 0 B
```