// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/ent-practice/ent/connector"
	"github.com/Kotodian/ent-practice/ent/enums"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/evse"
	"github.com/Kotodian/gokit/datasource"
)

// Connector is the model entity for the Connector schema.
type Connector struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID datasource.UUID `json:"id,omitempty"`
	// EquipmentSn holds the value of the "equipment_sn" field.
	// 桩序列号
	EquipmentSn string `json:"equipment_sn,omitempty"`
	// EvseSerial holds the value of the "evse_serial" field.
	// 设备序列号
	EvseSerial string `json:"evse_serial,omitempty"`
	// Serial holds the value of the "serial" field.
	// 枪序列号
	Serial string `json:"serial,omitempty"`
	// CurrentState holds the value of the "current_state" field.
	// 当前状态
	CurrentState enums.ConnectorState `json:"current_state,omitempty"`
	// BeforeState holds the value of the "before_state" field.
	// 之前状态
	BeforeState enums.ConnectorState `json:"before_state,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ConnectorQuery when eager-loading is set.
	Edges                ConnectorEdges `json:"edges"`
	equipment_connectors *datasource.UUID
	evse_connectors      *datasource.UUID
}

// ConnectorEdges holds the relations/edges for other nodes in the graph.
type ConnectorEdges struct {
	// Evse holds the value of the evse edge.
	Evse *Evse `json:"evse,omitempty"`
	// Equipment holds the value of the equipment edge.
	Equipment *Equipment `json:"equipment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EvseOrErr returns the Evse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConnectorEdges) EvseOrErr() (*Evse, error) {
	if e.loadedTypes[0] {
		if e.Evse == nil {
			// The edge evse was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: evse.Label}
		}
		return e.Evse, nil
	}
	return nil, &NotLoadedError{edge: "evse"}
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConnectorEdges) EquipmentOrErr() (*Equipment, error) {
	if e.loadedTypes[1] {
		if e.Equipment == nil {
			// The edge equipment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Connector) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case connector.FieldCurrentState, connector.FieldBeforeState:
			values[i] = new(enums.ConnectorState)
		case connector.FieldID:
			values[i] = new(sql.NullInt64)
		case connector.FieldEquipmentSn, connector.FieldEvseSerial, connector.FieldSerial:
			values[i] = new(sql.NullString)
		case connector.ForeignKeys[0]: // equipment_connectors
			values[i] = new(sql.NullInt64)
		case connector.ForeignKeys[1]: // evse_connectors
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Connector", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Connector fields.
func (c *Connector) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case connector.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = datasource.UUID(value.Int64)
		case connector.FieldEquipmentSn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_sn", values[i])
			} else if value.Valid {
				c.EquipmentSn = value.String
			}
		case connector.FieldEvseSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field evse_serial", values[i])
			} else if value.Valid {
				c.EvseSerial = value.String
			}
		case connector.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				c.Serial = value.String
			}
		case connector.FieldCurrentState:
			if value, ok := values[i].(*enums.ConnectorState); !ok {
				return fmt.Errorf("unexpected type %T for field current_state", values[i])
			} else if value != nil {
				c.CurrentState = *value
			}
		case connector.FieldBeforeState:
			if value, ok := values[i].(*enums.ConnectorState); !ok {
				return fmt.Errorf("unexpected type %T for field before_state", values[i])
			} else if value != nil {
				c.BeforeState = *value
			}
		case connector.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field equipment_connectors", value)
			} else if value.Valid {
				c.equipment_connectors = new(datasource.UUID)
				*c.equipment_connectors = datasource.UUID(value.Int64)
			}
		case connector.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field evse_connectors", value)
			} else if value.Valid {
				c.evse_connectors = new(datasource.UUID)
				*c.evse_connectors = datasource.UUID(value.Int64)
			}
		}
	}
	return nil
}

// QueryEvse queries the "evse" edge of the Connector entity.
func (c *Connector) QueryEvse() *EvseQuery {
	return (&ConnectorClient{config: c.config}).QueryEvse(c)
}

// QueryEquipment queries the "equipment" edge of the Connector entity.
func (c *Connector) QueryEquipment() *EquipmentQuery {
	return (&ConnectorClient{config: c.config}).QueryEquipment(c)
}

// Update returns a builder for updating this Connector.
// Note that you need to call Connector.Unwrap() before calling this method if this Connector
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Connector) Update() *ConnectorUpdateOne {
	return (&ConnectorClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Connector entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Connector) Unwrap() *Connector {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Connector is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Connector) String() string {
	var builder strings.Builder
	builder.WriteString("Connector(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", equipment_sn=")
	builder.WriteString(c.EquipmentSn)
	builder.WriteString(", evse_serial=")
	builder.WriteString(c.EvseSerial)
	builder.WriteString(", serial=")
	builder.WriteString(c.Serial)
	builder.WriteString(", current_state=")
	builder.WriteString(fmt.Sprintf("%v", c.CurrentState))
	builder.WriteString(", before_state=")
	builder.WriteString(fmt.Sprintf("%v", c.BeforeState))
	builder.WriteByte(')')
	return builder.String()
}

// Connectors is a parsable slice of Connector.
type Connectors []*Connector

func (c Connectors) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
