package app

import (
	"context"

	"github.com/spf13/cobra"
)

type App struct {
	name       string
	shortDesc  string
	descripton string
	run        RunFunc
	cmd        *cobra.Command
	args       cobra.PositionalArgs

	// +optional
	healthCheckFunc HealthCheckFunc

	// +optional
	options CliOptions

	// +optional
	silence bool

	// +optional
	noConfig bool

	// watching and re-reading config files
	// +optional
	watch bool

	contextExtractors map[string]func(context.Context) string
}

// RunFunc
type RunFunc func() error

type HealthCheckFunc func() error

// Option defines optional parameters for initializing the application
// structure.
type Option func(*App)
