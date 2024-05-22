package constant

type MigrateType string

const (
	MigrateTypeUp    MigrateType = "up"
	MigrateTypeDown  MigrateType = "down"
	MigrateTypeFresh MigrateType = "fresh"
)
