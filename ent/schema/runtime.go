package schema

import (
	"github.com/Kotodian/ent-practice/ent/connector"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/equipmentinfo"
	"github.com/Kotodian/ent-practice/ent/evse"
	"github.com/Kotodian/gokit/datasource"
)

func init() {
	connectorFields := Connector{}.Fields()
	_ = connectorFields
	// connectorDescID is the schema descriptor for id field.
	connectorDescID := connectorFields[0].Descriptor()
	// connector.DefaultID holds the default value on creation for the id field.
	connector.DefaultID = datasource.UUID(connectorDescID.Default.(uint64))
	equipmentFields := Equipment{}.Fields()
	_ = equipmentFields
	// equipmentDescSn is the schema descriptor for sn field.
	equipmentDescSn := equipmentFields[1].Descriptor()
	// equipment.SnValidator is a validator for the "sn" field. It is called by the builders before save.
	equipment.SnValidator = equipmentDescSn.Validators[0].(func(string) error)
	// equipmentDescID is the schema descriptor for id field.
	equipmentDescID := equipmentFields[0].Descriptor()
	// equipment.DefaultID holds the default value on creation for the id field.
	equipment.DefaultID = datasource.UUID(equipmentDescID.Default.(uint64))
	equipmentinfoFields := EquipmentInfo{}.Fields()
	_ = equipmentinfoFields
	// equipmentinfoDescID is the schema descriptor for id field.
	equipmentinfoDescID := equipmentinfoFields[0].Descriptor()
	// equipmentinfo.DefaultID holds the default value on creation for the id field.
	equipmentinfo.DefaultID = datasource.UUID(equipmentinfoDescID.Default.(uint64))
	evseFields := Evse{}.Fields()
	_ = evseFields
	// evseDescID is the schema descriptor for id field.
	evseDescID := evseFields[0].Descriptor()
	// evse.DefaultID holds the default value on creation for the id field.
	evse.DefaultID = datasource.UUID(evseDescID.Default.(uint64))
}
