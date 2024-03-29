// Code generated by ent, DO NOT EDIT.

package smartchargingeffect

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Kotodian/cwmodel/ent/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// ID filters vertices based on their ID field.
func ID(id datasource.UUID) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id datasource.UUID) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id datasource.UUID) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...datasource.UUID) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...datasource.UUID) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id datasource.UUID) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id datasource.UUID) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id datasource.UUID) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id datasource.UUID) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), vc))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// SmartID applies equality check predicate on the "smart_id" field. It's identical to SmartIDEQ.
func SmartID(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSmartID), v))
	})
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// Pid applies equality check predicate on the "pid" field. It's identical to PidEQ.
func Pid(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPid), vc))
	})
}

// Unit applies equality check predicate on the "unit" field. It's identical to UnitEQ.
func Unit(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnit), v))
	})
}

// EquipmentSn applies equality check predicate on the "equipment_sn" field. It's identical to EquipmentSnEQ.
func EquipmentSn(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEquipmentSn), v))
	})
}

// ValidFrom applies equality check predicate on the "valid_from" field. It's identical to ValidFromEQ.
func ValidFrom(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValidFrom), v))
	})
}

// ValidTo applies equality check predicate on the "valid_to" field. It's identical to ValidToEQ.
func ValidTo(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValidTo), v))
	})
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVersion), v))
	})
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldVersion), v...))
	})
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldVersion), v...))
	})
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVersion), v))
	})
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVersion), v))
	})
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVersion), v))
	})
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVersion), v))
	})
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...datasource.UUID) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...datasource.UUID) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedBy), vc))
	})
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedBy), vc))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...datasource.UUID) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...datasource.UUID) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedBy), vc))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// SmartIDEQ applies the EQ predicate on the "smart_id" field.
func SmartIDEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSmartID), v))
	})
}

// SmartIDNEQ applies the NEQ predicate on the "smart_id" field.
func SmartIDNEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSmartID), v))
	})
}

// SmartIDIn applies the In predicate on the "smart_id" field.
func SmartIDIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSmartID), v...))
	})
}

// SmartIDNotIn applies the NotIn predicate on the "smart_id" field.
func SmartIDNotIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSmartID), v...))
	})
}

// SmartIDGT applies the GT predicate on the "smart_id" field.
func SmartIDGT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSmartID), v))
	})
}

// SmartIDGTE applies the GTE predicate on the "smart_id" field.
func SmartIDGTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSmartID), v))
	})
}

// SmartIDLT applies the LT predicate on the "smart_id" field.
func SmartIDLT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSmartID), v))
	})
}

// SmartIDLTE applies the LTE predicate on the "smart_id" field.
func SmartIDLTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSmartID), v))
	})
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartTime), v))
	})
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartTime), v...))
	})
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartTime), v...))
	})
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartTime), v))
	})
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartTime), v))
	})
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartTime), v))
	})
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartTime), v))
	})
}

// PidEQ applies the EQ predicate on the "pid" field.
func PidEQ(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPid), vc))
	})
}

// PidNEQ applies the NEQ predicate on the "pid" field.
func PidNEQ(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPid), vc))
	})
}

// PidIn applies the In predicate on the "pid" field.
func PidIn(vs ...datasource.UUID) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPid), v...))
	})
}

// PidNotIn applies the NotIn predicate on the "pid" field.
func PidNotIn(vs ...datasource.UUID) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPid), v...))
	})
}

// PidGT applies the GT predicate on the "pid" field.
func PidGT(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPid), vc))
	})
}

// PidGTE applies the GTE predicate on the "pid" field.
func PidGTE(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPid), vc))
	})
}

// PidLT applies the LT predicate on the "pid" field.
func PidLT(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPid), vc))
	})
}

// PidLTE applies the LTE predicate on the "pid" field.
func PidLTE(v datasource.UUID) predicate.SmartChargingEffect {
	vc := uint64(v)
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPid), vc))
	})
}

// UnitEQ applies the EQ predicate on the "unit" field.
func UnitEQ(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnit), v))
	})
}

// UnitNEQ applies the NEQ predicate on the "unit" field.
func UnitNEQ(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnit), v))
	})
}

// UnitIn applies the In predicate on the "unit" field.
func UnitIn(vs ...string) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUnit), v...))
	})
}

// UnitNotIn applies the NotIn predicate on the "unit" field.
func UnitNotIn(vs ...string) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUnit), v...))
	})
}

// UnitGT applies the GT predicate on the "unit" field.
func UnitGT(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnit), v))
	})
}

// UnitGTE applies the GTE predicate on the "unit" field.
func UnitGTE(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnit), v))
	})
}

// UnitLT applies the LT predicate on the "unit" field.
func UnitLT(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnit), v))
	})
}

// UnitLTE applies the LTE predicate on the "unit" field.
func UnitLTE(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnit), v))
	})
}

// UnitContains applies the Contains predicate on the "unit" field.
func UnitContains(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUnit), v))
	})
}

// UnitHasPrefix applies the HasPrefix predicate on the "unit" field.
func UnitHasPrefix(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUnit), v))
	})
}

// UnitHasSuffix applies the HasSuffix predicate on the "unit" field.
func UnitHasSuffix(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUnit), v))
	})
}

// UnitEqualFold applies the EqualFold predicate on the "unit" field.
func UnitEqualFold(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUnit), v))
	})
}

// UnitContainsFold applies the ContainsFold predicate on the "unit" field.
func UnitContainsFold(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUnit), v))
	})
}

// EquipmentSnEQ applies the EQ predicate on the "equipment_sn" field.
func EquipmentSnEQ(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnNEQ applies the NEQ predicate on the "equipment_sn" field.
func EquipmentSnNEQ(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnIn applies the In predicate on the "equipment_sn" field.
func EquipmentSnIn(vs ...string) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEquipmentSn), v...))
	})
}

// EquipmentSnNotIn applies the NotIn predicate on the "equipment_sn" field.
func EquipmentSnNotIn(vs ...string) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEquipmentSn), v...))
	})
}

// EquipmentSnGT applies the GT predicate on the "equipment_sn" field.
func EquipmentSnGT(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnGTE applies the GTE predicate on the "equipment_sn" field.
func EquipmentSnGTE(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnLT applies the LT predicate on the "equipment_sn" field.
func EquipmentSnLT(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnLTE applies the LTE predicate on the "equipment_sn" field.
func EquipmentSnLTE(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnContains applies the Contains predicate on the "equipment_sn" field.
func EquipmentSnContains(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnHasPrefix applies the HasPrefix predicate on the "equipment_sn" field.
func EquipmentSnHasPrefix(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnHasSuffix applies the HasSuffix predicate on the "equipment_sn" field.
func EquipmentSnHasSuffix(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnEqualFold applies the EqualFold predicate on the "equipment_sn" field.
func EquipmentSnEqualFold(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEquipmentSn), v))
	})
}

// EquipmentSnContainsFold applies the ContainsFold predicate on the "equipment_sn" field.
func EquipmentSnContainsFold(v string) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEquipmentSn), v))
	})
}

// ValidFromEQ applies the EQ predicate on the "valid_from" field.
func ValidFromEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValidFrom), v))
	})
}

// ValidFromNEQ applies the NEQ predicate on the "valid_from" field.
func ValidFromNEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValidFrom), v))
	})
}

// ValidFromIn applies the In predicate on the "valid_from" field.
func ValidFromIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldValidFrom), v...))
	})
}

// ValidFromNotIn applies the NotIn predicate on the "valid_from" field.
func ValidFromNotIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldValidFrom), v...))
	})
}

// ValidFromGT applies the GT predicate on the "valid_from" field.
func ValidFromGT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValidFrom), v))
	})
}

// ValidFromGTE applies the GTE predicate on the "valid_from" field.
func ValidFromGTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValidFrom), v))
	})
}

// ValidFromLT applies the LT predicate on the "valid_from" field.
func ValidFromLT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValidFrom), v))
	})
}

// ValidFromLTE applies the LTE predicate on the "valid_from" field.
func ValidFromLTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValidFrom), v))
	})
}

// ValidFromIsNil applies the IsNil predicate on the "valid_from" field.
func ValidFromIsNil() predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldValidFrom)))
	})
}

// ValidFromNotNil applies the NotNil predicate on the "valid_from" field.
func ValidFromNotNil() predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldValidFrom)))
	})
}

// ValidToEQ applies the EQ predicate on the "valid_to" field.
func ValidToEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValidTo), v))
	})
}

// ValidToNEQ applies the NEQ predicate on the "valid_to" field.
func ValidToNEQ(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValidTo), v))
	})
}

// ValidToIn applies the In predicate on the "valid_to" field.
func ValidToIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldValidTo), v...))
	})
}

// ValidToNotIn applies the NotIn predicate on the "valid_to" field.
func ValidToNotIn(vs ...int64) predicate.SmartChargingEffect {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldValidTo), v...))
	})
}

// ValidToGT applies the GT predicate on the "valid_to" field.
func ValidToGT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValidTo), v))
	})
}

// ValidToGTE applies the GTE predicate on the "valid_to" field.
func ValidToGTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValidTo), v))
	})
}

// ValidToLT applies the LT predicate on the "valid_to" field.
func ValidToLT(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValidTo), v))
	})
}

// ValidToLTE applies the LTE predicate on the "valid_to" field.
func ValidToLTE(v int64) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValidTo), v))
	})
}

// ValidToIsNil applies the IsNil predicate on the "valid_to" field.
func ValidToIsNil() predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldValidTo)))
	})
}

// ValidToNotNil applies the NotNil predicate on the "valid_to" field.
func ValidToNotNil() predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldValidTo)))
	})
}

// HasEquipment applies the HasEdge predicate on the "equipment" edge.
func HasEquipment() predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EquipmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EquipmentTable, EquipmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentWith applies the HasEdge predicate on the "equipment" edge with a given conditions (other predicates).
func HasEquipmentWith(preds ...predicate.Equipment) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
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

// HasConnector applies the HasEdge predicate on the "connector" edge.
func HasConnector() predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConnectorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ConnectorTable, ConnectorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConnectorWith applies the HasEdge predicate on the "connector" edge with a given conditions (other predicates).
func HasConnectorWith(preds ...predicate.Connector) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConnectorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ConnectorTable, ConnectorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrderInfo applies the HasEdge predicate on the "order_info" edge.
func HasOrderInfo() predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderInfoTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OrderInfoTable, OrderInfoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderInfoWith applies the HasEdge predicate on the "order_info" edge with a given conditions (other predicates).
func HasOrderInfoWith(preds ...predicate.OrderInfo) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderInfoInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OrderInfoTable, OrderInfoColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SmartChargingEffect) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SmartChargingEffect) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
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
func Not(p predicate.SmartChargingEffect) predicate.SmartChargingEffect {
	return predicate.SmartChargingEffect(func(s *sql.Selector) {
		p(s.Not())
	})
}
