# omp
omp is a Go package that allows you to write open.mp gamemodes.

## Installation

```shell
go get github.com/kodeyeen/omp
```

## Requirements

- `GCC/G++` go build tool will require you to have C and C++ 32 bit compilers available on your system.

Type `gcc -v` and `g++ -v` in your terminal and you should see something like this:

```
Target: i686-w64-mingw32
```

Note the i686. Otherwise it won't build

## Quickstart

1. Install the latest [Gomponent](https://github.com/kodeyeen/gomponent).
2. Initialize a go module with `go mod init <YOUR MODULE NAME>`.
3. Write some basic gamemode.

```go
package main

import (
	"github.com/kodeyeen/omp"
)

// Gamemode entry point
func init() {
	// Listen to some predefined event
	omp.On(omp.EventTypePlayerConnect, func(e *omp.PlayerConnectEvent) bool {
		// Send client message to the connected player
		e.Player.SendClientMessage("Hello, world!", 0xFFFF00FF)
		return true
	})
}

// You MUST declare the main function, otherwise it fails to build
// You shouldn't write any code here
func main() {}

```

4. Build it depending on your system (see the section below).
5. Add the compiled file to the `gamemodes` folder of your server.
6. Add this to your server's `config.json`:

```json
"go": {
    "gamemode": "<YOUR GAMEMODE NAME>"
}
```

Now if you run the server and connect to it you should see the message "Hello, world!"

## Building

On Windows:

```powershell
$env:GOARCH=386; $env:CGO_ENABLED=1; go build -buildmode=c-shared -o build/gamemode.dll
```

On Linux:

```bash
GOARCH=386 CGO_ENABLED=1 go build -buildmode=c-shared -o build/gamemode.so
```

## Credits

* [Sreyas-Sreelal](https://github.com/Sreyas-Sreelal)
* [Hual](https://github.com/Hual)
* [AmyrAhmady](https://github.com/AmyrAhmady)
