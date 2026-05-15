package app

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Application holds the core application dependencies
type Application struct {
	Router *gin.Engine
	DB     *gorm.DB
	Config *Config
}

// Config holds application configuration
type Config struct {
	// Server settings
	AppAddr    string `mapstructure:"app_addr"`
	AppPort    int    `mapstructure:"app_port"`
	AppVersion string `mapstructure:"app_version"`

	// Database settings
	DBPath string `mapstructure:"db_path"`

	// JWT settings
	JWTSecret     string        `mapstructure:"jwt_secret"`
	JWTExpiration time.Duration `mapstructure:"jwt_expiration"`

	// Log settings
	LogPath  string `mapstructure:"log_path"`
	LogLevel string `mapstructure:"log_level"`

	// Panel settings
	BaseDir     string `mapstructure:"base_dir"`
	DataDir     string `mapstructure:"data_dir"`
	TmpDir      string `mapstructure:"tmp_dir"`
	BackupDir   string `mapstructure:"backup_dir"`
	RunDir      string `mapstructure:"run_dir"`
	DockerSocket string `mapstructure:"docker_socket"`
}

// New creates a new Application instance with the provided configuration
func New(cfg *Config) *Application {
	if cfg.AppAddr == "" {
		cfg.AppAddr = "0.0.0.0"
	}
	if cfg.AppPort == 0 {
		cfg.AppPort = 9999
	}
	if cfg.JWTExpiration == 0 {
		cfg.JWTExpiration = 24 * time.Hour
	}
	if cfg.BaseDir == "" {
		cfg.BaseDir = "/opt/1panel"
	}

	return &Application{
		Config: cfg,
	}
}

// Addr returns the full listen address for the HTTP server
func (a *Application) Addr() string {
	return fmt.Sprintf("%s:%d", a.Config.AppAddr, a.Config.AppPort)
}

// DefaultConfig returns a Config populated with sensible defaults
func DefaultConfig() *Config {
	return &Config{
		AppAddr:      "0.0.0.0",
		AppPort:      9999,
		AppVersion:   "v1.0.0",
		DBPath:       "/opt/1panel/db/1panel.db",
		JWTExpiration: 24 * time.Hour,
		LogPath:      "/opt/1panel/log",
		LogLevel:     "info",
		BaseDir:      "/opt/1panel",
		DataDir:      "/opt/1panel/data",
		TmpDir:       "/opt/1panel/tmp",
		BackupDir:    "/opt/1panel/backup",
		RunDir:       "/opt/1panel/run",
		DockerSocket: "/var/run/docker.sock",
	}
}
