// Code generated by ent, DO NOT EDIT.

package orderinfo

const (
	// Label holds the string label denoting the orderinfo type in the database.
	Label = "order_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRemoteStartID holds the string denoting the remote_start_id field in the database.
	FieldRemoteStartID = "remote_start_id"
	// FieldTransactionID holds the string denoting the transaction_id field in the database.
	FieldTransactionID = "transaction_id"
	// FieldAuthorizationID holds the string denoting the authorization_id field in the database.
	FieldAuthorizationID = "authorization_id"
	// FieldCustomerID holds the string denoting the customer_id field in the database.
	FieldCustomerID = "customer_id"
	// FieldCallerOrderID holds the string denoting the caller_order_id field in the database.
	FieldCallerOrderID = "caller_order_id"
	// FieldTotalElectricity holds the string denoting the total_electricity field in the database.
	FieldTotalElectricity = "total_electricity"
	// FieldChargeStartElectricity holds the string denoting the charge_start_electricity field in the database.
	FieldChargeStartElectricity = "charge_start_electricity"
	// FieldChargeStopElectricity holds the string denoting the charge_stop_electricity field in the database.
	FieldChargeStopElectricity = "charge_stop_electricity"
	// FieldSharpElectricity holds the string denoting the sharp_electricity field in the database.
	FieldSharpElectricity = "sharp_electricity"
	// FieldPeakElectricity holds the string denoting the peak_electricity field in the database.
	FieldPeakElectricity = "peak_electricity"
	// FieldFlatElectricity holds the string denoting the flat_electricity field in the database.
	FieldFlatElectricity = "flat_electricity"
	// FieldValleyElectricity holds the string denoting the valley_electricity field in the database.
	FieldValleyElectricity = "valley_electricity"
	// FieldStopReasonCode holds the string denoting the stop_reason_code field in the database.
	FieldStopReasonCode = "stop_reason_code"
	// FieldOffline holds the string denoting the offline field in the database.
	FieldOffline = "offline"
	// FieldPriceSchemeReleaseID holds the string denoting the price_scheme_release_id field in the database.
	FieldPriceSchemeReleaseID = "price_scheme_release_id"
	// FieldOrderStartTime holds the string denoting the order_start_time field in the database.
	FieldOrderStartTime = "order_start_time"
	// FieldOrderFinalTime holds the string denoting the order_final_time field in the database.
	FieldOrderFinalTime = "order_final_time"
	// FieldChargeStartTime holds the string denoting the charge_start_time field in the database.
	FieldChargeStartTime = "charge_start_time"
	// FieldChargeFinalTime holds the string denoting the charge_final_time field in the database.
	FieldChargeFinalTime = "charge_final_time"
	// FieldIntellectID holds the string denoting the intellect_id field in the database.
	FieldIntellectID = "intellect_id"
	// FieldStationID holds the string denoting the station_id field in the database.
	FieldStationID = "station_id"
	// FieldOperatorID holds the string denoting the operator_id field in the database.
	FieldOperatorID = "operator_id"
	// EdgeConnector holds the string denoting the connector edge name in mutations.
	EdgeConnector = "connector"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// EdgeOrderEvent holds the string denoting the order_event edge name in mutations.
	EdgeOrderEvent = "order_event"
	// Table holds the table name of the orderinfo in the database.
	Table = "order_info"
	// ConnectorTable is the table that holds the connector relation/edge.
	ConnectorTable = "order_info"
	// ConnectorInverseTable is the table name for the Connector entity.
	// It exists in this package in order to avoid circular dependency with the "connector" package.
	ConnectorInverseTable = "base_connector"
	// ConnectorColumn is the table column denoting the connector relation/edge.
	ConnectorColumn = "connector_id"
	// EquipmentTable is the table that holds the equipment relation/edge.
	EquipmentTable = "order_info"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "base_equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "equipment_id"
	// OrderEventTable is the table that holds the order_event relation/edge.
	OrderEventTable = "order_event"
	// OrderEventInverseTable is the table name for the OrderEvent entity.
	// It exists in this package in order to avoid circular dependency with the "orderevent" package.
	OrderEventInverseTable = "order_event"
	// OrderEventColumn is the table column denoting the order_event relation/edge.
	OrderEventColumn = "order_info_order_event"
)

// Columns holds all SQL columns for orderinfo fields.
var Columns = []string{
	FieldID,
	FieldRemoteStartID,
	FieldTransactionID,
	FieldAuthorizationID,
	FieldCustomerID,
	FieldCallerOrderID,
	FieldTotalElectricity,
	FieldChargeStartElectricity,
	FieldChargeStopElectricity,
	FieldSharpElectricity,
	FieldPeakElectricity,
	FieldFlatElectricity,
	FieldValleyElectricity,
	FieldStopReasonCode,
	FieldOffline,
	FieldPriceSchemeReleaseID,
	FieldOrderStartTime,
	FieldOrderFinalTime,
	FieldChargeStartTime,
	FieldChargeFinalTime,
	FieldIntellectID,
	FieldStationID,
	FieldOperatorID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "order_info"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"connector_id",
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
