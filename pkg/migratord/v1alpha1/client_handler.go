package v1alpha1

import "fmt"

type clientMigationHandler struct {
	parent *Migratord
}

func (h *clientMigationHandler) PerformMigration(migration *Migration) {
	fmt.Println("Migration complete!")
}
