package database

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Constants representing environment variable keys for migration configuration.
const (
	envKeyDsn          = "MIGRATION_DSN"
	envKeyOp           = "MIGRATION_OPERATION"
	envKeyForceVersion = "MIGRATION_FORCE_VERSION"
	envKeyFilesDir     = "MIGRATION_FILES_DIR"
)

var (
	ErrMissingDsnEnv       = errors.New("missing env: MIGRATION_DSN")
	ErrMissingOperationEnv = errors.New("missing env: MIGRATION_OPERATION")
)

func readForceVersion() (int, error) {
	forceVersionRaw, ok := os.LookupEnv(envKeyForceVersion)
	if !ok {
		return 0, nil
	}

	forceVersion, err := strconv.Atoi(forceVersionRaw)
	if err != nil {
		return 0, fmt.Errorf("convert forceVersion: %w", err)
	}

	return forceVersion, nil
}

// RunMigrationsFromEnv reads migration configuration from environment variables,
// creates a MigrationRunner, and runs the specified migration operation.
func RunMigrationsFromEnv() error {
	dsn, ok := os.LookupEnv(envKeyDsn)
	if !ok {
		return ErrMissingDsnEnv
	}

	operation, ok := os.LookupEnv(envKeyOp)
	if !ok {
		return ErrMissingOperationEnv
	}

	forceVersion, err := readForceVersion()
	if err != nil {
		return fmt.Errorf("failed to read forceVersion: %w", err)
	}

	opts := []Option{}
	if filesDir, ok := os.LookupEnv(envKeyFilesDir); ok {
		opts = append(opts, WithFilesDir(filesDir))
	}

	runner, err := NewMigrationRunner(dsn, opts...)
	if err != nil {
		return fmt.Errorf("failed to create a new migration runner: %w", err)
	}

	// Get the current migration version and log it.
	version, dirty, err := runner.Version()
	if err != nil {
		return fmt.Errorf("failed to get migration version before operation: %w", err)
	}

	log.WithFields(log.Fields{
		"version": version,
		"dirty":   dirty,
	}).Info("Migration version before operation")

	if err = runner.Run(OperationData{
		ID:           operation,
		ForceVersion: forceVersion,
	}); err != nil {
		return fmt.Errorf("run operation %s: %w", operation, err)
	}

	log.Info("Successfully finished migration")

	// Get the migration version after the operation and log it.
	version, dirty, err = runner.Version()
	if err != nil {
		return fmt.Errorf("failed to get migration version after operation: %w", err)
	}

	log.WithFields(log.Fields{
		"version": version,
		"dirty":   dirty,
	}).Info("Migration version after operation")

	return nil
}
