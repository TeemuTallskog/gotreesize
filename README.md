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

```Bash
.\gotreesize.exe 'C:\Program Files\Git\mingw64'

Directory: C:\Program Files\Git\mingw64
-----------------------------------------
bin                |█████████████████████████████████████████████████████████| 94.1 MB
libexec            |██████████████████████████████████████████████████████   | 89.2 MB
share              |██████████                                               | 17.3 MB
lib                |███                                                      | 5.6 MB
etc                |                                                         | 1.5 MB
doc                |                                                         | 10.2 KB
```