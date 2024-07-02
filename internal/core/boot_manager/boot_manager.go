package bootmanager

import (
	"os"
	"os/signal"
	"syscall"
)

var shutdownSignals = []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM}

type Daemon func()

func Bootstrap(daemon func() Daemon) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, shutdownSignals...)

	cleanUp := daemon()
	// Wait for shutdown signal
	<-sigs

	// Execute shutdown clean up
	cleanUp()
	close(sigs)
}
