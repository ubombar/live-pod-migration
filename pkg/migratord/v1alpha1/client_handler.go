package v1alpha1

type clientMigationHandler struct {
	parent *Migratord
}

// Creates different adapters depending on the migration job
func (h *clientMigationHandler) PerformMigration(migration *MigrationJob) {

}
