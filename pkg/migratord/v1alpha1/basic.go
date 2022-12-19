package v1alpha1

type BasicHandler struct {
}

func (h *BasicHandler) TransferCheckpointFile(job *MigrationJob) error {
	return nil
}
