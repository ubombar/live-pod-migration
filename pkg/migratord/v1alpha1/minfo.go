package v1alpha1

type MigrationInfoBasic struct {
}

func (m MigrationInfoBasic) Type() MigrationMethod {
	return Basic
}
