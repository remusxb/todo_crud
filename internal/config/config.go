/*
Package config manages the loading and unmarshalling of configuration for the app.  it's meant
to be used pretty much only by the main program, and can be extended to auto-watch config, use remote
config repositories and some other stuff.

Don't use this for tests.
*/
package config

import (
	"log"
	"log/slog"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/remusxb/todo_crud/internal/featureflags"
	"github.com/remusxb/todo_crud/internal/server"
)

type (
	SvcConfig = server.Config
	FfConfig  = featureflags.Config
)

// Config lives here because how use init it in the main method is different from in tests.
// mapstructure - denotes key from config file
// env - denotes key from environment
type Config struct {
	SrvConfig SvcConfig `mapstructure:"server" env:"SERVER"`
	FfConfig  FfConfig  `mapstructure:"feature_flags" env:"FEATURE_FLAGS"`

	// Version toggles printing the version.
	Version bool `mapstructure:"version" env:"VERSION"`

	// Verbose toggles debug level logging.  If you want leveled logging, you can implement that here with an
	// enum and aliases or something
	Verbose bool `mapstructure:"verbose" env:"VERBOSE"`
}

// Parse the base config values using viper.  shove that into the provided cfg struct
// args are your command line args without the command name (`os.Args[:0]`) so the flags in the flag-set can be set.
func Parse(fs *flag.FlagSet, cfg any, args []string) {
	if err := fs.Parse(args); err != nil {
		log.Fatal("can't parse flags for config")
	}

	viper.SetConfigName("todo_crud")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.BindPFlags(fs); err != nil {
		slog.Error("Probably programmer error: can't bind pflags")
	}

	// Allow environmental variables to use _ delimiters
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			break
		default:
			slog.Error("fatal error loading config - ignoring", "error", err)
		}
	}

	if err := viper.Unmarshal(cfg); err != nil {
		slog.Error("Error unmarshaling the config")
		os.Exit(1)
	}
}
