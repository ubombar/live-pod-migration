package v1alpha1

type clientMigationHandler struct {
	parent *Migratord
}

// Creates different adapters depending on the migration job
func (h *clientMigationHandler) PerformMigration(migration *MigrationJob) {
	switch migration.Method {
	case Basic:
		handleBasicMigrationJob(h.parent, migration)
	case Postcopy:
		// Not implemented yet
		handleUnknownMigrationJob(h.parent, migration)
	case Precopy:
		// Not implemented yet
		handleUnknownMigrationJob(h.parent, migration)
	default:
		// Error
		handleUnknownMigrationJob(h.parent, migration)
	}
}

// Handles unknown or not implemented migration job
func handleUnknownMigrationJob(m *Migratord, job *MigrationJob) {

}

// Do the basic migration job. This consists of Checkpointing, Transfering and Restoring.
func handleBasicMigrationJob(m *Migratord, job *MigrationJob) {

}
