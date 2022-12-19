package v1alpha1

import "github.com/ubombar/live-pod-migration/pkg/migratord/v1alpha1/handlers"

type clientMigationHandler struct {
	parent *Migratord
}

// Creates different adapters depending on the migration job
func (h *clientMigationHandler) PerformMigration(migration *MigrationJob) {
	switch migration.Method {
	case Basic:
		handlers.ClientMigrationHandlerBasic{}.Handle()
	}
}
