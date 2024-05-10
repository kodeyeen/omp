# omp
omp is a Go package that allows you to write open.mp gamemodes.

## Installation

```shell
go get github.com/kodeyeen/omp
```

## Requirements

- `Go`
- `GCC/G++` go build will require you to have C and C++ 32 bit compilers available on your system.

Type `gcc -v` and `g++ -v` and you should see this line:

```
Target: i686-w64-mingw32
```

Otherwise it won't build

## Building

On Windows:

```powershell
$env:GOARCH=386; $env:CGO_ENABLED=1; go build -buildmode=c-shared -o build/gmname.dll
```

On Linux:

```bash
GOARCH=386 CGO_ENABLED=1 go build -buildmode=c-shared -o build/gmname.so
```

## Credits

* [Sreyas-Sreelal](https://github.com/Sreyas-Sreelal)
* [Hual](https://github.com/Hual)
* [AmyrAhmady](https://github.com/AmyrAhmady)
