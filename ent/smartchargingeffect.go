// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/cwmodel/ent/connector"
	"github.com/Kotodian/cwmodel/ent/equipment"
	"github.com/Kotodian/cwmodel/ent/orderinfo"
	"github.com/Kotodian/cwmodel/ent/smartchargingeffect"
	"github.com/Kotodian/cwmodel/ent/types"
	"github.com/Kotodian/gokit/datasource"
)

// SmartChargingEffect is the model entity for the SmartChargingEffect schema.
type SmartChargingEffect struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID datasource.UUID `json:"id,omitempty"`
	// 乐观锁
	Version int64 `json:"version,omitempty"`
	// 创建者
	CreatedBy datasource.UUID `json:"created_by,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"created_at,omitempty"`
	// 修改者
	UpdatedBy datasource.UUID `json:"updated_by,omitempty"`
	// 修改时间
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// 智慧id
	SmartID int64 `json:"smart_id,omitempty"`
	// 开始时间
	StartTime int64 `json:"start_time,omitempty"`
	// 父id
	Pid datasource.UUID `json:"pid,omitempty"`
	// 单位(W或者A)
	Unit string `json:"unit,omitempty"`
	// 桩sn号
	EquipmentSn string `json:"equipment_sn,omitempty"`
	// 有效开始时间
	ValidFrom *int64 `json:"valid_from,omitempty"`
	// 有效结束时间
	ValidTo *int64 `json:"valid_to,omitempty"`
	// 时间间隔
	Spec []types.ChargingSchedulePeriod `json:"spec,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SmartChargingEffectQuery when eager-loading is set.
	Edges        SmartChargingEffectEdges `json:"-"`
	connector_id *datasource.UUID
	equipment_id *datasource.UUID
	order_id     *datasource.UUID
}

// SmartChargingEffectEdges holds the relations/edges for other nodes in the graph.
type SmartChargingEffectEdges struct {
	// Equipment holds the value of the equipment edge.
	Equipment *Equipment `json:"equipment,omitempty"`
	// Connector holds the value of the connector edge.
	Connector *Connector `json:"connector,omitempty"`
	// OrderInfo holds the value of the order_info edge.
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SmartChargingEffectEdges) EquipmentOrErr() (*Equipment, error) {
	if e.loadedTypes[0] {
		if e.Equipment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// ConnectorOrErr returns the Connector value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SmartChargingEffectEdges) ConnectorOrErr() (*Connector, error) {
	if e.loadedTypes[1] {
		if e.Connector == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: connector.Label}
		}
		return e.Connector, nil
	}
	return nil, &NotLoadedError{edge: "connector"}
}

// OrderInfoOrErr returns the OrderInfo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SmartChargingEffectEdges) OrderInfoOrErr() (*OrderInfo, error) {
	if e.loadedTypes[2] {
		if e.OrderInfo == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: orderinfo.Label}
		}
		return e.OrderInfo, nil
	}
	return nil, &NotLoadedError{edge: "order_info"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SmartChargingEffect) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case smartchargingeffect.FieldSpec:
			values[i] = new([]byte)
		case smartchargingeffect.FieldID, smartchargingeffect.FieldVersion, smartchargingeffect.FieldCreatedBy, smartchargingeffect.FieldCreatedAt, smartchargingeffect.FieldUpdatedBy, smartchargingeffect.FieldUpdatedAt, smartchargingeffect.FieldSmartID, smartchargingeffect.FieldStartTime, smartchargingeffect.FieldPid, smartchargingeffect.FieldValidFrom, smartchargingeffect.FieldValidTo:
			values[i] = new(sql.NullInt64)
		case smartchargingeffect.FieldUnit, smartchargingeffect.FieldEquipmentSn:
			values[i] = new(sql.NullString)
		case smartchargingeffect.ForeignKeys[0]: // connector_id
			values[i] = new(sql.NullInt64)
		case smartchargingeffect.ForeignKeys[1]: // equipment_id
			values[i] = new(sql.NullInt64)
		case smartchargingeffect.ForeignKeys[2]: // order_id
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SmartChargingEffect", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SmartChargingEffect fields.
func (sce *SmartChargingEffect) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case smartchargingeffect.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sce.ID = datasource.UUID(value.Int64)
			}
		case smartchargingeffect.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				sce.Version = value.Int64
			}
		case smartchargingeffect.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				sce.CreatedBy = datasource.UUID(value.Int64)
			}
		case smartchargingeffect.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sce.CreatedAt = value.Int64
			}
		case smartchargingeffect.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				sce.UpdatedBy = datasource.UUID(value.Int64)
			}
		case smartchargingeffect.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sce.UpdatedAt = value.Int64
			}
		case smartchargingeffect.FieldSmartID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field smart_id", values[i])
			} else if value.Valid {
				sce.SmartID = value.Int64
			}
		case smartchargingeffect.FieldStartTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				sce.StartTime = value.Int64
			}
		case smartchargingeffect.FieldPid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pid", values[i])
			} else if value.Valid {
				sce.Pid = datasource.UUID(value.Int64)
			}
		case smartchargingeffect.FieldUnit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field unit", values[i])
			} else if value.Valid {
				sce.Unit = value.String
			}
		case smartchargingeffect.FieldEquipmentSn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_sn", values[i])
			} else if value.Valid {
				sce.EquipmentSn = value.String
			}
		case smartchargingeffect.FieldValidFrom:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field valid_from", values[i])
			} else if value.Valid {
				sce.ValidFrom = new(int64)
				*sce.ValidFrom = value.Int64
			}
		case smartchargingeffect.FieldValidTo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field valid_to", values[i])
			} else if value.Valid {
				sce.ValidTo = new(int64)
				*sce.ValidTo = value.Int64
			}
		case smartchargingeffect.FieldSpec:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field spec", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &sce.Spec); err != nil {
					return fmt.Errorf("unmarshal field spec: %w", err)
				}
			}
		case smartchargingeffect.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field connector_id", values[i])
			} else if value.Valid {
				sce.connector_id = new(datasource.UUID)
				*sce.connector_id = datasource.UUID(value.Int64)
			}
		case smartchargingeffect.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_id", values[i])
			} else if value.Valid {
				sce.equipment_id = new(datasource.UUID)
				*sce.equipment_id = datasource.UUID(value.Int64)
			}
		case smartchargingeffect.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				sce.order_id = new(datasource.UUID)
				*sce.order_id = datasource.UUID(value.Int64)
			}
		}
	}
	return nil
}

// QueryEquipment queries the "equipment" edge of the SmartChargingEffect entity.
func (sce *SmartChargingEffect) QueryEquipment() *EquipmentQuery {
	return (&SmartChargingEffectClient{config: sce.config}).QueryEquipment(sce)
}

// QueryConnector queries the "connector" edge of the SmartChargingEffect entity.
func (sce *SmartChargingEffect) QueryConnector() *ConnectorQuery {
	return (&SmartChargingEffectClient{config: sce.config}).QueryConnector(sce)
}

// QueryOrderInfo queries the "order_info" edge of the SmartChargingEffect entity.
func (sce *SmartChargingEffect) QueryOrderInfo() *OrderInfoQuery {
	return (&SmartChargingEffectClient{config: sce.config}).QueryOrderInfo(sce)
}

// Update returns a builder for updating this SmartChargingEffect.
// Note that you need to call SmartChargingEffect.Unwrap() before calling this method if this SmartChargingEffect
// was returned from a transaction, and the transaction was committed or rolled back.
func (sce *SmartChargingEffect) Update() *SmartChargingEffectUpdateOne {
	return (&SmartChargingEffectClient{config: sce.config}).UpdateOne(sce)
}

// Unwrap unwraps the SmartChargingEffect entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sce *SmartChargingEffect) Unwrap() *SmartChargingEffect {
	_tx, ok := sce.config.driver.(*txDriver)
	if !ok {
		panic("ent: SmartChargingEffect is not a transactional entity")
	}
	sce.config.driver = _tx.drv
	return sce
}

// String implements the fmt.Stringer.
func (sce *SmartChargingEffect) String() string {
	var builder strings.Builder
	builder.WriteString("SmartChargingEffect(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sce.ID))
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", sce.Version))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", sce.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", sce.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", sce.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", sce.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("smart_id=")
	builder.WriteString(fmt.Sprintf("%v", sce.SmartID))
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(fmt.Sprintf("%v", sce.StartTime))
	builder.WriteString(", ")
	builder.WriteString("pid=")
	builder.WriteString(fmt.Sprintf("%v", sce.Pid))
	builder.WriteString(", ")
	builder.WriteString("unit=")
	builder.WriteString(sce.Unit)
	builder.WriteString(", ")
	builder.WriteString("equipment_sn=")
	builder.WriteString(sce.EquipmentSn)
	builder.WriteString(", ")
	if v := sce.ValidFrom; v != nil {
		builder.WriteString("valid_from=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := sce.ValidTo; v != nil {
		builder.WriteString("valid_to=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("spec=")
	builder.WriteString(fmt.Sprintf("%v", sce.Spec))
	builder.WriteByte(')')
	return builder.String()
}

// SmartChargingEffects is a parsable slice of SmartChargingEffect.
type SmartChargingEffects []*SmartChargingEffect

func (sce SmartChargingEffects) config(cfg config) {
	for _i := range sce {
		sce[_i].config = cfg
	}
}
