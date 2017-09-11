# Signalize #

Dead simple OS-signal handling library

## Get The Library

```bash
go get github.com/avid/signalize
```

## Use The Library

```go
package main

import (
	"github.com/avid/signalize"
	"syscall"
)

func main() {
	signalize.Catch(func(signal os.Signal) {
		// catch SIGHUP and do something, ex. reload
	}, syscall.SIGHUP)

	signalize.Catch(func(signal os.Signal) {
		// catches SIGUSR1 & SIGUSR2 and run same function for both of them
	}, syscall.SIGUSR1, syscall.SIGUSR2)

	signalize.Catch(func(signal os.Signal) {
		// catch SIGINT and stop gracefully for example
	}, syscall.SIGINT)

	signalize.Catch(func(signal os.Signal) {
		// catch SIGQUIT and stop hardy for example
	}, syscall.SIGQUIT)

	// these signals will  stop listening signals
	signalize.Stops(syscall.SIGINT, syscall.SIGQUIT)

	// this starts listening OS signals
	// TRUE to block context and FALSE for non-blocking hook
	signalize.Hook(true)
}
```

