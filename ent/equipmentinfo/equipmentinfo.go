// Code generated by ent, DO NOT EDIT.

package equipmentinfo

import (
	"github.com/Kotodian/gokit/datasource"
)

const (
	// Label holds the string label denoting the equipmentinfo type in the database.
	Label = "equipment_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEquipmentSn holds the string denoting the equipment_sn field in the database.
	FieldEquipmentSn = "equipment_sn"
	// FieldModelID holds the string denoting the model_id field in the database.
	FieldModelID = "model_id"
	// FieldManufacturerID holds the string denoting the manufacturer_id field in the database.
	FieldManufacturerID = "manufacturer_id"
	// FieldFirmwareID holds the string denoting the firmware_id field in the database.
	FieldFirmwareID = "firmware_id"
	// FieldAccessPod holds the string denoting the access_pod field in the database.
	FieldAccessPod = "access_pod"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldEvseNumber holds the string denoting the evse_number field in the database.
	FieldEvseNumber = "evse_number"
	// FieldAlarmNumber holds the string denoting the alarm_number field in the database.
	FieldAlarmNumber = "alarm_number"
	// FieldRegisterDatetime holds the string denoting the register_datetime field in the database.
	FieldRegisterDatetime = "register_datetime"
	// FieldRemoteAddress holds the string denoting the remote_address field in the database.
	FieldRemoteAddress = "remote_address"
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
	EquipmentColumn = "equipment_id"
)

// Columns holds all SQL columns for equipmentinfo fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldEquipmentSn,
	FieldModelID,
	FieldManufacturerID,
	FieldFirmwareID,
	FieldAccessPod,
	FieldState,
	FieldEvseNumber,
	FieldAlarmNumber,
	FieldRegisterDatetime,
	FieldRemoteAddress,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "base_equipment_extra"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"equipment_id",
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
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion int64
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy datasource.UUID
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt int64
	// DefaultUpdatedBy holds the default value on creation for the "updated_by" field.
	DefaultUpdatedBy datasource.UUID
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID datasource.UUID
)
