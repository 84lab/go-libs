package database

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	// required by migrate.New... to parse migration files directory
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
)

// supported migrate operations
const (
	defaultFilesDir = "dbmigrations"
	operationUp     = "up"
	operationDown   = "down"
	operationForce  = "force"
)

var ErrUnsupportedMigrationOperation = errors.New("unsupported migration operation")

type operationFn func(*MigrationRunner, OperationData) error

// OperationData contains information about the migration operation to be performed.
type OperationData struct {
	ID           string
	ForceVersion int
}

// MigrationRunner is responsible for managing and running database migrations.
type MigrationRunner struct {
	mgr      *migrate.Migrate
	filesDir string
}

// Option represents a function that configures a MigrationRunner.
type Option func(runner *MigrationRunner)

// WithFilesDir sets a custom directory containing migration files for the MigrationRunner.
// If not provided, the default directory "dbmigrations" will be used.
func WithFilesDir(filesDir string) Option {
	return func(runner *MigrationRunner) {
		runner.filesDir = filesDir
	}
}

// NewMigrationRunner creates a new MigrationRunner with the given database connection string (dsn) and options.
func NewMigrationRunner(dsn string, opts ...Option) (*MigrationRunner, error) {
	runner := &MigrationRunner{
		filesDir: defaultFilesDir,
	}

	for _, opt := range opts {
		opt(runner)
	}

	mgr, err := migrate.New("file://"+runner.filesDir, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create Migrate object: %w", err)
	}

	runner.mgr = mgr

	return runner, nil
}

// Run executes the migration operation specified by the OperationData.
func (m *MigrationRunner) Run(operation OperationData) error {
	supportedOperations := map[string]operationFn{
		operationUp:    runUp,
		operationDown:  runDown,
		operationForce: runForce,
	}

	operationName := operation.ID
	operationFn, found := supportedOperations[operationName]
	if !found {
		return fmt.Errorf("%w: %s", ErrUnsupportedMigrationOperation, operationName)
	}

	if err := operationFn(m, operation); err != nil {
		return fmt.Errorf("operation %s failed: %w", operationName, err)
	}

	return nil
}

// Version returns the current migration version, a dirty flag, and an error if any.
func (m *MigrationRunner) Version() (version uint, dirty bool, err error) {
	version, dirty, err = m.mgr.Version()
	if err != nil {
		return version, dirty, fmt.Errorf("failed to get the currently active migration version: %w", err)
	}

	return
}

// runUp runs the "up" migration operation, applying new migrations to the database.
func runUp(m *MigrationRunner, _ OperationData) error {
	log.Info("Running migrate UP")

	if err := m.mgr.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.WithField("directory", m.filesDir).Info("No new migrations found in the directory")

			return nil
		}

		return fmt.Errorf("failed to run migrations UP: %w", err)
	}

	return nil
}

// runDown runs the "down" migration operation, rolling back the latest applied migration.
func runDown(m *MigrationRunner, _ OperationData) error {
	log.Infof("Running migrate DOWN with STEPS=%d", 1)

	// always rollback the latest applied migration only
	if err := m.mgr.Steps(-1); err != nil {
		return fmt.Errorf("failed to run migrations DOWN: %w", err)
	}

	return nil
}

// runForce runs the "force" migration operation,
// forcibly setting the migration version without running the actual migrations.
func runForce(m *MigrationRunner, op OperationData) error {
	log.Infof("Running FORCE with VERSION %d", op.ForceVersion)

	if err := m.mgr.Force(op.ForceVersion); err != nil {
		return fmt.Errorf("failed to run migrations FORCE with VERSION %d failed: %w", op.ForceVersion, err)
	}

	return nil
}
