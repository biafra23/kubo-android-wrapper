package wrapper

import (
	"fmt"
	logging "github.com/ipfs/go-log/v2"
	"github.com/ipfs/kubo/cmd/ipfs/kubo"
	"os"
	"strings"
)

// Run ipfs with command line options. For example: "daemon --init".
// Set IPFS_PATH to repoPath.
// Set GOLOG_LOG_LEVEL to gologLogLevel (case-insensitive) (DEBUG, INFO, WARN, ERROR, DPANIC, PANIC or FATAL)
func RunCli(repoPath string, gologLogLevel string, args string) error {
	//os.Setenv("GOLOG_LOG_LEVEL", gologLogLevel)
	err := logging.SetLogLevel("*", strings.ToUpper(gologLogLevel))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Invalid log level '%s': %v. Using default logging level.\n", gologLogLevel, err)
		return fmt.Errorf("failed to set log level: %w", err)
	}
	os.Args = append([]string{"ipfs"}, strings.Fields(args)...)
	err = os.Setenv("IPFS_PATH", repoPath)
	if err != nil {
		return fmt.Errorf("Failed to set IPFS_PATH: %v", err)
	}

	kubo.Start(kubo.BuildDefaultEnv)

	return nil
}
