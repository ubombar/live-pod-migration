package v1alpha1

type clientMigationHandler struct {
	parent *Migratord
}

func (h *clientMigationHandler) PerformMigration(migration *Migration) {
	// Docker checkpoint

	// Checkpoint Transfer

	// Restore on server

	// Done!
}
