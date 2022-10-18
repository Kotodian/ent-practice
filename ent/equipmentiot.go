// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/equipmentiot"
)

// EquipmentIot is the model entity for the EquipmentIot schema.
type EquipmentIot struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// iccid
	Iccid string `json:"iccid,omitempty"`
	// imei
	Imei string `json:"imei,omitempty"`
	// remote_address
	RemoteAddress string `json:"remote_address,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EquipmentIotQuery when eager-loading is set.
	Edges        EquipmentIotEdges `json:"edges"`
	equipment_id *int
}

// EquipmentIotEdges holds the relations/edges for other nodes in the graph.
type EquipmentIotEdges struct {
	// Equipment holds the value of the equipment edge.
	Equipment *Equipment `json:"equipment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EquipmentIotEdges) EquipmentOrErr() (*Equipment, error) {
	if e.loadedTypes[0] {
		if e.Equipment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EquipmentIot) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case equipmentiot.FieldID:
			values[i] = new(sql.NullInt64)
		case equipmentiot.FieldIccid, equipmentiot.FieldImei, equipmentiot.FieldRemoteAddress:
			values[i] = new(sql.NullString)
		case equipmentiot.ForeignKeys[0]: // equipment_id
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EquipmentIot", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EquipmentIot fields.
func (ei *EquipmentIot) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case equipmentiot.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ei.ID = int(value.Int64)
		case equipmentiot.FieldIccid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field iccid", values[i])
			} else if value.Valid {
				ei.Iccid = value.String
			}
		case equipmentiot.FieldImei:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field imei", values[i])
			} else if value.Valid {
				ei.Imei = value.String
			}
		case equipmentiot.FieldRemoteAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remote_address", values[i])
			} else if value.Valid {
				ei.RemoteAddress = value.String
			}
		case equipmentiot.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field equipment_id", value)
			} else if value.Valid {
				ei.equipment_id = new(int)
				*ei.equipment_id = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryEquipment queries the "equipment" edge of the EquipmentIot entity.
func (ei *EquipmentIot) QueryEquipment() *EquipmentQuery {
	return (&EquipmentIotClient{config: ei.config}).QueryEquipment(ei)
}

// Update returns a builder for updating this EquipmentIot.
// Note that you need to call EquipmentIot.Unwrap() before calling this method if this EquipmentIot
// was returned from a transaction, and the transaction was committed or rolled back.
func (ei *EquipmentIot) Update() *EquipmentIotUpdateOne {
	return (&EquipmentIotClient{config: ei.config}).UpdateOne(ei)
}

// Unwrap unwraps the EquipmentIot entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ei *EquipmentIot) Unwrap() *EquipmentIot {
	_tx, ok := ei.config.driver.(*txDriver)
	if !ok {
		panic("ent: EquipmentIot is not a transactional entity")
	}
	ei.config.driver = _tx.drv
	return ei
}

// String implements the fmt.Stringer.
func (ei *EquipmentIot) String() string {
	var builder strings.Builder
	builder.WriteString("EquipmentIot(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ei.ID))
	builder.WriteString("iccid=")
	builder.WriteString(ei.Iccid)
	builder.WriteString(", ")
	builder.WriteString("imei=")
	builder.WriteString(ei.Imei)
	builder.WriteString(", ")
	builder.WriteString("remote_address=")
	builder.WriteString(ei.RemoteAddress)
	builder.WriteByte(')')
	return builder.String()
}

// EquipmentIots is a parsable slice of EquipmentIot.
type EquipmentIots []*EquipmentIot

func (ei EquipmentIots) config(cfg config) {
	for _i := range ei {
		ei[_i].config = cfg
	}
}
