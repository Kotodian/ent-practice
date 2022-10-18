// Code generated by ent, DO NOT EDIT.

package equipmentalarm

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Kotodian/ent-practice/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DtcCode applies equality check predicate on the "dtc_code" field. It's identical to DtcCodeEQ.
func DtcCode(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDtcCode), v))
	})
}

// RemoteAddress applies equality check predicate on the "remote_address" field. It's identical to RemoteAddressEQ.
func RemoteAddress(v string) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemoteAddress), v))
	})
}

// TriggerTime applies equality check predicate on the "trigger_time" field. It's identical to TriggerTimeEQ.
func TriggerTime(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTriggerTime), v))
	})
}

// FinalTime applies equality check predicate on the "final_time" field. It's identical to FinalTimeEQ.
func FinalTime(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFinalTime), v))
	})
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// DtcCodeEQ applies the EQ predicate on the "dtc_code" field.
func DtcCodeEQ(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDtcCode), v))
	})
}

// DtcCodeNEQ applies the NEQ predicate on the "dtc_code" field.
func DtcCodeNEQ(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDtcCode), v))
	})
}

// DtcCodeIn applies the In predicate on the "dtc_code" field.
func DtcCodeIn(vs ...int64) predicate.EquipmentAlarm {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDtcCode), v...))
	})
}

// DtcCodeNotIn applies the NotIn predicate on the "dtc_code" field.
func DtcCodeNotIn(vs ...int64) predicate.EquipmentAlarm {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDtcCode), v...))
	})
}

// DtcCodeGT applies the GT predicate on the "dtc_code" field.
func DtcCodeGT(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDtcCode), v))
	})
}

// DtcCodeGTE applies the GTE predicate on the "dtc_code" field.
func DtcCodeGTE(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDtcCode), v))
	})
}

// DtcCodeLT applies the LT predicate on the "dtc_code" field.
func DtcCodeLT(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDtcCode), v))
	})
}

// DtcCodeLTE applies the LTE predicate on the "dtc_code" field.
func DtcCodeLTE(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDtcCode), v))
	})
}

// RemoteAddressEQ applies the EQ predicate on the "remote_address" field.
func RemoteAddressEQ(v string) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemoteAddress), v))
	})
}

// RemoteAddressNEQ applies the NEQ predicate on the "remote_address" field.
func RemoteAddressNEQ(v string) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemoteAddress), v))
	})
}

// RemoteAddressIn applies the In predicate on the "remote_address" field.
func RemoteAddressIn(vs ...string) predicate.EquipmentAlarm {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRemoteAddress), v...))
	})
}

// RemoteAddressNotIn applies the NotIn predicate on the "remote_address" field.
func RemoteAddressNotIn(vs ...string) predicate.EquipmentAlarm {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRemoteAddress), v...))
	})
}

// RemoteAddressGT applies the GT predicate on the "remote_address" field.
func RemoteAddressGT(v string) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemoteAddress), v))
	})
}

// RemoteAddressGTE applies the GTE predicate on the "remote_address" field.
func RemoteAddressGTE(v string) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemoteAddress), v))
	})
}

// RemoteAddressLT applies the LT predicate on the "remote_address" field.
func RemoteAddressLT(v string) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemoteAddress), v))
	})
}

// RemoteAddressLTE applies the LTE predicate on the "remote_address" field.
func RemoteAddressLTE(v string) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemoteAddress), v))
	})
}

// RemoteAddressContains applies the Contains predicate on the "remote_address" field.
func RemoteAddressContains(v string) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemoteAddress), v))
	})
}

// RemoteAddressHasPrefix applies the HasPrefix predicate on the "remote_address" field.
func RemoteAddressHasPrefix(v string) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemoteAddress), v))
	})
}

// RemoteAddressHasSuffix applies the HasSuffix predicate on the "remote_address" field.
func RemoteAddressHasSuffix(v string) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemoteAddress), v))
	})
}

// RemoteAddressEqualFold applies the EqualFold predicate on the "remote_address" field.
func RemoteAddressEqualFold(v string) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemoteAddress), v))
	})
}

// RemoteAddressContainsFold applies the ContainsFold predicate on the "remote_address" field.
func RemoteAddressContainsFold(v string) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemoteAddress), v))
	})
}

// TriggerTimeEQ applies the EQ predicate on the "trigger_time" field.
func TriggerTimeEQ(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTriggerTime), v))
	})
}

// TriggerTimeNEQ applies the NEQ predicate on the "trigger_time" field.
func TriggerTimeNEQ(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTriggerTime), v))
	})
}

// TriggerTimeIn applies the In predicate on the "trigger_time" field.
func TriggerTimeIn(vs ...int64) predicate.EquipmentAlarm {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTriggerTime), v...))
	})
}

// TriggerTimeNotIn applies the NotIn predicate on the "trigger_time" field.
func TriggerTimeNotIn(vs ...int64) predicate.EquipmentAlarm {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTriggerTime), v...))
	})
}

// TriggerTimeGT applies the GT predicate on the "trigger_time" field.
func TriggerTimeGT(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTriggerTime), v))
	})
}

// TriggerTimeGTE applies the GTE predicate on the "trigger_time" field.
func TriggerTimeGTE(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTriggerTime), v))
	})
}

// TriggerTimeLT applies the LT predicate on the "trigger_time" field.
func TriggerTimeLT(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTriggerTime), v))
	})
}

// TriggerTimeLTE applies the LTE predicate on the "trigger_time" field.
func TriggerTimeLTE(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTriggerTime), v))
	})
}

// TriggerTimeIsNil applies the IsNil predicate on the "trigger_time" field.
func TriggerTimeIsNil() predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTriggerTime)))
	})
}

// TriggerTimeNotNil applies the NotNil predicate on the "trigger_time" field.
func TriggerTimeNotNil() predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTriggerTime)))
	})
}

// FinalTimeEQ applies the EQ predicate on the "final_time" field.
func FinalTimeEQ(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFinalTime), v))
	})
}

// FinalTimeNEQ applies the NEQ predicate on the "final_time" field.
func FinalTimeNEQ(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFinalTime), v))
	})
}

// FinalTimeIn applies the In predicate on the "final_time" field.
func FinalTimeIn(vs ...int64) predicate.EquipmentAlarm {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFinalTime), v...))
	})
}

// FinalTimeNotIn applies the NotIn predicate on the "final_time" field.
func FinalTimeNotIn(vs ...int64) predicate.EquipmentAlarm {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFinalTime), v...))
	})
}

// FinalTimeGT applies the GT predicate on the "final_time" field.
func FinalTimeGT(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFinalTime), v))
	})
}

// FinalTimeGTE applies the GTE predicate on the "final_time" field.
func FinalTimeGTE(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFinalTime), v))
	})
}

// FinalTimeLT applies the LT predicate on the "final_time" field.
func FinalTimeLT(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFinalTime), v))
	})
}

// FinalTimeLTE applies the LTE predicate on the "final_time" field.
func FinalTimeLTE(v int64) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFinalTime), v))
	})
}

// FinalTimeIsNil applies the IsNil predicate on the "final_time" field.
func FinalTimeIsNil() predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFinalTime)))
	})
}

// FinalTimeNotNil applies the NotNil predicate on the "final_time" field.
func FinalTimeNotNil() predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFinalTime)))
	})
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCount), v))
	})
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int) predicate.EquipmentAlarm {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCount), v...))
	})
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int) predicate.EquipmentAlarm {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCount), v...))
	})
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCount), v))
	})
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCount), v))
	})
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCount), v))
	})
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCount), v))
	})
}

// HasEquipment applies the HasEdge predicate on the "equipment" edge.
func HasEquipment() predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EquipmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EquipmentTable, EquipmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentWith applies the HasEdge predicate on the "equipment" edge with a given conditions (other predicates).
func HasEquipmentWith(preds ...predicate.Equipment) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EquipmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EquipmentTable, EquipmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EquipmentAlarm) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EquipmentAlarm) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EquipmentAlarm) predicate.EquipmentAlarm {
	return predicate.EquipmentAlarm(func(s *sql.Selector) {
		p(s.Not())
	})
}
