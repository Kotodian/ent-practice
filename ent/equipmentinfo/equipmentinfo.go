// Code generated by entc, DO NOT EDIT.

package equipmentinfo

import (
	"github.com/Kotodian/gokit/datasource"
)

const (
	// Label holds the string label denoting the equipmentinfo type in the database.
	Label = "equipment_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEquipmentSn holds the string denoting the equipment_sn field in the database.
	FieldEquipmentSn = "equipment_sn"
	// FieldAccessPod holds the string denoting the access_pod field in the database.
	FieldAccessPod = "access_pod"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// Table holds the table name of the equipmentinfo in the database.
	Table = "base_equipment_extra"
	// EquipmentTable is the table that holds the equipment relation/edge.
	EquipmentTable = "base_equipment_extra"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "base_equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "equipment_equipment_info"
)

// Columns holds all SQL columns for equipmentinfo fields.
var Columns = []string{
	FieldID,
	FieldEquipmentSn,
	FieldAccessPod,
	FieldState,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "base_equipment_extra"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"equipment_equipment_info",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID datasource.UUID
)
