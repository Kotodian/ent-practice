// Code generated by entc, DO NOT EDIT.

package equipment

import (
	"fmt"

	"github.com/Kotodian/ent-practice/ent/enums"
	"github.com/Kotodian/gokit/datasource"
)

const (
	// Label holds the string label denoting the equipment type in the database.
	Label = "equipment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSn holds the string denoting the sn field in the database.
	FieldSn = "sn"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldOperatorID holds the string denoting the operator_id field in the database.
	FieldOperatorID = "operator_id"
	// FieldStationID holds the string denoting the station_id field in the database.
	FieldStationID = "station_id"
	// EdgeEquipmentInfo holds the string denoting the equipment_info edge name in mutations.
	EdgeEquipmentInfo = "equipment_info"
	// EdgeEvses holds the string denoting the evses edge name in mutations.
	EdgeEvses = "evses"
	// EdgeConnectors holds the string denoting the connectors edge name in mutations.
	EdgeConnectors = "connectors"
	// Table holds the table name of the equipment in the database.
	Table = "base_equipment"
	// EquipmentInfoTable is the table that holds the equipment_info relation/edge.
	EquipmentInfoTable = "base_equipment_extra"
	// EquipmentInfoInverseTable is the table name for the EquipmentInfo entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentinfo" package.
	EquipmentInfoInverseTable = "base_equipment_extra"
	// EquipmentInfoColumn is the table column denoting the equipment_info relation/edge.
	EquipmentInfoColumn = "equipment_equipment_info"
	// EvsesTable is the table that holds the evses relation/edge.
	EvsesTable = "base_evse"
	// EvsesInverseTable is the table name for the Evse entity.
	// It exists in this package in order to avoid circular dependency with the "evse" package.
	EvsesInverseTable = "base_evse"
	// EvsesColumn is the table column denoting the evses relation/edge.
	EvsesColumn = "equipment_evses"
	// ConnectorsTable is the table that holds the connectors relation/edge.
	ConnectorsTable = "base_connector"
	// ConnectorsInverseTable is the table name for the Connector entity.
	// It exists in this package in order to avoid circular dependency with the "connector" package.
	ConnectorsInverseTable = "base_connector"
	// ConnectorsColumn is the table column denoting the connectors relation/edge.
	ConnectorsColumn = "equipment_connectors"
)

// Columns holds all SQL columns for equipment fields.
var Columns = []string{
	FieldID,
	FieldSn,
	FieldCategory,
	FieldOperatorID,
	FieldStationID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// SnValidator is a validator for the "sn" field. It is called by the builders before save.
	SnValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID datasource.UUID
)

// CategoryValidator is a validator for the "category" field enum values. It is called by the builders before save.
func CategoryValidator(c enums.EquipmentCategory) error {
	switch c.String() {
	case "private", "public":
		return nil
	default:
		return fmt.Errorf("equipment: invalid enum value for category field: %q", c)
	}
}