package ofp4

import (
	"encoding/binary"
	"fmt"
)

type ErrorMsg []byte

func (self ErrorMsg) Type() uint16 {
	return binary.BigEndian.Uint16(self[8:])
}

func (self ErrorMsg) Code() uint16 {
	return binary.BigEndian.Uint16(self[10:])
}

func (self ErrorMsg) Data() []byte {
	return self[16:]
}

func (self ErrorMsg) Error() string {
	s := "ofp unknown error"
	switch self.Type() {
	case OFPET_HELLO_FAILED:
		switch self.Code() {
		case OFPHFC_INCOMPATIBLE:
			s = "OFPET_HELLO_FAILED OFPHFC_INCOMPATIBLE"
		case OFPHFC_EPERM:
			s = "OFPET_HELLO_FAILED OFPHFC_EPERM"
		}
	case OFPET_BAD_REQUEST:
		switch self.Code() {
		case OFPBRC_BAD_VERSION:
			s = "OFPET_BAD_REQUEST OFPBRC_BAD_VERSION"
		case OFPBRC_BAD_TYPE:
			s = "OFPET_BAD_REQUEST OFPBRC_BAD_TYPE"
		case OFPBRC_BAD_MULTIPART:
			s = "OFPET_BAD_REQUEST OFPBRC_BAD_MULTIPART"
		case OFPBRC_BAD_EXPERIMENTER:
			s = "OFPET_BAD_REQUEST OFPBRC_BAD_EXPERIMENTER"
		case OFPBRC_BAD_EXP_TYPE:
			s = "OFPET_BAD_REQUEST OFPBRC_BAD_EXP_TYPE"
		case OFPBRC_EPERM:
			s = "OFPET_BAD_REQUEST OFPBRC_EPERM"
		case OFPBRC_BAD_LEN:
			s = "OFPET_BAD_REQUEST OFPBRC_BAD_LEN"
		case OFPBRC_BUFFER_EMPTY:
			s = "OFPET_BAD_REQUEST OFPBRC_BUFFER_EMPTY"
		case OFPBRC_BUFFER_UNKNOWN:
			s = "OFPET_BAD_REQUEST OFPBRC_BUFFER_UNKNOWN"
		case OFPBRC_BAD_TABLE_ID:
			s = "OFPET_BAD_REQUEST OFPBRC_BAD_TABLE_ID"
		case OFPBRC_IS_SLAVE:
			s = "OFPET_BAD_REQUEST OFPBRC_IS_SLAVE"
		case OFPBRC_BAD_PORT:
			s = "OFPET_BAD_REQUEST OFPBRC_BAD_PORT"
		case OFPBRC_BAD_PACKET:
			s = "OFPET_BAD_REQUEST OFPBRC_BAD_PACKET"
		case OFPBRC_MULTIPART_BUFFER_OVERFLOW:
			s = "OFPET_BAD_REQUEST OFPBRC_MULTIPART_BUFFER_OVERFLOW"
		}
	case OFPET_BAD_ACTION:
		switch self.Code() {
		case OFPBAC_BAD_TYPE:
			s = "OFPET_BAD_ACTION OFPBAC_BAD_TYPE"
		case OFPBAC_BAD_LEN:
			s = "OFPET_BAD_ACTION OFPBAC_BAD_LEN"
		case OFPBAC_BAD_EXPERIMENTER:
			s = "OFPET_BAD_ACTION OFPBAC_BAD_EXPERIMENTER"
		case OFPBAC_BAD_EXP_TYPE:
			s = "OFPET_BAD_ACTION OFPBAC_BAD_EXP_TYPE"
		case OFPBAC_BAD_OUT_PORT:
			s = "OFPET_BAD_ACTION OFPBAC_BAD_OUT_PORT"
		case OFPBAC_BAD_ARGUMENT:
			s = "OFPET_BAD_ACTION OFPBAC_BAD_ARGUMENT"
		case OFPBAC_EPERM:
			s = "OFPET_BAD_ACTION OFPBAC_EPERM"
		case OFPBAC_TOO_MANY:
			s = "OFPET_BAD_ACTION OFPBAC_TOO_MANY"
		case OFPBAC_BAD_QUEUE:
			s = "OFPET_BAD_ACTION OFPBAC_BAD_QUEUE"
		case OFPBAC_BAD_OUT_GROUP:
			s = "OFPET_BAD_ACTION OFPBAC_BAD_OUT_GROUP"
		case OFPBAC_MATCH_INCONSISTENT:
			s = "OFPET_BAD_ACTION OFPBAC_MATCH_INCONSISTENT"
		case OFPBAC_UNSUPPORTED_ORDER:
			s = "OFPET_BAD_ACTION OFPBAC_UNSUPPORTED_ORDER"
		case OFPBAC_BAD_TAG:
			s = "OFPET_BAD_ACTION OFPBAC_BAD_TAG"
		case OFPBAC_BAD_SET_TYPE:
			s = "OFPET_BAD_ACTION OFPBAC_BAD_SET_TYPE"
		case OFPBAC_BAD_SET_LEN:
			s = "OFPET_BAD_ACTION OFPBAC_BAD_SET_LEN"
		case OFPBAC_BAD_SET_ARGUMENT:
			s = "OFPET_BAD_ACTION OFPBAC_BAD_SET_ARGUMENT"
		}
	case OFPET_BAD_INSTRUCTION:
		switch self.Code() {
		case OFPBIC_UNKNOWN_INST:
			s = "OFPET_BAD_INSTRUCTION OFPBIC_UNKNOWN_INST"
		case OFPBIC_UNSUP_INST:
			s = "OFPET_BAD_INSTRUCTION OFPBIC_UNSUP_INST"
		case OFPBIC_BAD_TABLE_ID:
			s = "OFPET_BAD_INSTRUCTION OFPBIC_BAD_TABLE_ID"
		case OFPBIC_UNSUP_METADATA:
			s = "OFPET_BAD_INSTRUCTION OFPBIC_UNSUP_METADATA"
		case OFPBIC_UNSUP_METADATA_MASK:
			s = "OFPET_BAD_INSTRUCTION OFPBIC_UNSUP_METADATA_MASK"
		case OFPBIC_BAD_EXPERIMENTER:
			s = "OFPET_BAD_INSTRUCTION OFPBIC_BAD_EXPERIMENTER"
		case OFPBIC_BAD_EXP_TYPE:
			s = "OFPET_BAD_INSTRUCTION OFPBIC_BAD_EXP_TYPE"
		case OFPBIC_BAD_LEN:
			s = "OFPET_BAD_INSTRUCTION OFPBIC_BAD_LEN"
		case OFPBIC_EPERM:
			s = "OFPET_BAD_INSTRUCTION OFPBIC_EPERM"
		}
	case OFPET_BAD_MATCH:
		switch self.Code() {
		case OFPBMC_BAD_TYPE:
			s = "OFPET_BAD_MATCH OFPBMC_BAD_TYPE"
		case OFPBMC_BAD_LEN:
			s = "OFPET_BAD_MATCH OFPBMC_BAD_LEN"
		case OFPBMC_BAD_TAG:
			s = "OFPET_BAD_MATCH OFPBMC_BAD_TAG"
		case OFPBMC_BAD_DL_ADDR_MASK:
			s = "OFPET_BAD_MATCH OFPBMC_BAD_DL_ADDR_MASK"
		case OFPBMC_BAD_NW_ADDR_MASK:
			s = "OFPET_BAD_MATCH OFPBMC_BAD_NW_ADDR_MASK"
		case OFPBMC_BAD_WILDCARDS:
			s = "OFPET_BAD_MATCH OFPBMC_BAD_WILDCARDS"
		case OFPBMC_BAD_FIELD:
			s = "OFPET_BAD_MATCH OFPBMC_BAD_FIELD"
		case OFPBMC_BAD_VALUE:
			s = "OFPET_BAD_MATCH OFPBMC_BAD_VALUE"
		case OFPBMC_BAD_MASK:
			s = "OFPET_BAD_MATCH OFPBMC_BAD_MASK"
		case OFPBMC_BAD_PREREQ:
			s = "OFPET_BAD_MATCH OFPBMC_BAD_PREREQ"
		case OFPBMC_DUP_FIELD:
			s = "OFPET_BAD_MATCH OFPBMC_DUP_FIELD"
		case OFPBMC_EPERM:
			s = "OFPET_BAD_MATCH OFPBMC_EPERM"
		}
	case OFPET_FLOW_MOD_FAILED:
		switch self.Code() {
		case OFPFMFC_UNKNOWN:
			s = "OFPET_FLOW_MOD_FAILED OFPFMFC_UNKNOWN"
		case OFPFMFC_TABLE_FULL:
			s = "OFPET_FLOW_MOD_FAILED OFPFMFC_TABLE_FULL"
		case OFPFMFC_BAD_TABLE_ID:
			s = "OFPET_FLOW_MOD_FAILED OFPFMFC_BAD_TABLE_ID"
		case OFPFMFC_OVERLAP:
			s = "OFPET_FLOW_MOD_FAILED OFPFMFC_OVERLAP"
		case OFPFMFC_EPERM:
			s = "OFPET_FLOW_MOD_FAILED OFPFMFC_EPERM"
		case OFPFMFC_BAD_TIMEOUT:
			s = "OFPET_FLOW_MOD_FAILED OFPFMFC_BAD_TIMEOUT"
		case OFPFMFC_BAD_COMMAND:
			s = "OFPET_FLOW_MOD_FAILED OFPFMFC_BAD_COMMAND"
		case OFPFMFC_BAD_FLAGS:
			s = "OFPET_FLOW_MOD_FAILED OFPFMFC_BAD_FLAGS"
		}
	case OFPET_GROUP_MOD_FAILED:
		switch self.Code() {
		case OFPGMFC_GROUP_EXISTS:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_GROUP_EXISTS"
		case OFPGMFC_INVALID_GROUP:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_INVALID_GROUP"
		case OFPGMFC_WEIGHT_UNSUPPORTED:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_WEIGHT_UNSUPPORTED"
		case OFPGMFC_OUT_OF_GROUPS:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_OUT_OF_GROUPS"
		case OFPGMFC_OUT_OF_BUCKETS:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_OUT_OF_BUCKETS"
		case OFPGMFC_CHAINING_UNSUPPORTED:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_CHAINING_UNSUPPORTED"
		case OFPGMFC_WATCH_UNSUPPORTED:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_WATCH_UNSUPPORTED"
		case OFPGMFC_LOOP:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_LOOP"
		case OFPGMFC_UNKNOWN_GROUP:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_UNKNOWN_GROUP"
		case OFPGMFC_CHAINED_GROUP:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_CHAINED_GROUP"
		case OFPGMFC_BAD_TYPE:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_BAD_TYPE"
		case OFPGMFC_BAD_COMMAND:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_BAD_COMMAND"
		case OFPGMFC_BAD_BUCKET:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_BAD_BUCKET"
		case OFPGMFC_BAD_WATCH:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_BAD_WATCH"
		case OFPGMFC_EPERM:
			s = "OFPET_GROUP_MOD_FAILED OFPGMFC_EPERM"
		}
	case OFPET_PORT_MOD_FAILED:
		switch self.Code() {
		case OFPPMFC_BAD_PORT:
			s = "OFPET_PORT_MOD_FAILED OFPPMFC_BAD_PORT"
		case OFPPMFC_BAD_HW_ADDR:
			s = "OFPET_PORT_MOD_FAILED OFPPMFC_BAD_HW_ADDR"
		case OFPPMFC_BAD_CONFIG:
			s = "OFPET_PORT_MOD_FAILED OFPPMFC_BAD_CONFIG"
		case OFPPMFC_BAD_ADVERTISE:
			s = "OFPET_PORT_MOD_FAILED OFPPMFC_BAD_ADVERTISE"
		case OFPPMFC_EPERM:
			s = "OFPET_PORT_MOD_FAILED OFPPMFC_EPERM"
		}
	case OFPET_TABLE_MOD_FAILED:
		switch self.Code() {
		case OFPTMFC_BAD_TABLE:
			s = "OFPET_TABLE_MOD_FAILED OFPTMFC_BAD_TABLE"
		case OFPTMFC_BAD_CONFIG:
			s = "OFPET_TABLE_MOD_FAILED OFPTMFC_BAD_CONFIG"
		case OFPTMFC_EPERM:
			s = "OFPET_TABLE_MOD_FAILED OFPTMFC_EPERM"
		}
	case OFPET_QUEUE_OP_FAILED:
		switch self.Code() {
		case OFPQOFC_BAD_PORT:
			s = "OFPET_QUEUE_OP_FAILED OFPQOFC_BAD_PORT"
		case OFPQOFC_BAD_QUEUE:
			s = "OFPET_QUEUE_OP_FAILED OFPQOFC_BAD_QUEUE"
		case OFPQOFC_EPERM:
			s = "OFPET_QUEUE_OP_FAILED OFPQOFC_EPERM"
		}
	case OFPET_SWITCH_CONFIG_FAILED:
		switch self.Code() {
		case OFPSCFC_BAD_FLAGS:
			s = "OFPET_SWITCH_CONFIG_FAILED OFPSCFC_BAD_FLAGS"
		case OFPSCFC_BAD_LEN:
			s = "OFPET_SWITCH_CONFIG_FAILED OFPSCFC_BAD_LEN"
		case OFPSCFC_EPERM:
			s = "OFPET_SWITCH_CONFIG_FAILED OFPSCFC_EPERM"
		}
	case OFPET_ROLE_REQUEST_FAILED:
		switch self.Code() {
		case OFPRRFC_STALE:
			s = "OFPET_ROLE_REQUEST_FAILED OFPRRFC_STALE"
		case OFPRRFC_UNSUP:
			s = "OFPET_ROLE_REQUEST_FAILED OFPRRFC_UNSUP"
		case OFPRRFC_BAD_ROLE:
			s = "OFPET_ROLE_REQUEST_FAILED OFPRRFC_BAD_ROLE"
		}
	case OFPET_METER_MOD_FAILED:
		switch self.Code() {
		case OFPMMFC_UNKNOWN:
			s = "OFPET_METER_MOD_FAILED OFPMMFC_UNKNOWN"
		case OFPMMFC_METER_EXISTS:
			s = "OFPET_METER_MOD_FAILED OFPMMFC_METER_EXISTS"
		case OFPMMFC_INVALID_METER:
			s = "OFPET_METER_MOD_FAILED OFPMMFC_INVALID_METER"
		case OFPMMFC_UNKNOWN_METER:
			s = "OFPET_METER_MOD_FAILED OFPMMFC_UNKNOWN_METER"
		case OFPMMFC_BAD_COMMAND:
			s = "OFPET_METER_MOD_FAILED OFPMMFC_BAD_COMMAND"
		case OFPMMFC_BAD_FLAGS:
			s = "OFPET_METER_MOD_FAILED OFPMMFC_BAD_FLAGS"
		case OFPMMFC_BAD_RATE:
			s = "OFPET_METER_MOD_FAILED OFPMMFC_BAD_RATE"
		case OFPMMFC_BAD_BURST:
			s = "OFPET_METER_MOD_FAILED OFPMMFC_BAD_BURST"
		case OFPMMFC_BAD_BAND:
			s = "OFPET_METER_MOD_FAILED OFPMMFC_BAD_BAND"
		case OFPMMFC_BAD_BAND_VALUE:
			s = "OFPET_METER_MOD_FAILED OFPMMFC_BAD_BAND_VALUE"
		case OFPMMFC_OUT_OF_METERS:
			s = "OFPET_METER_MOD_FAILED OFPMMFC_OUT_OF_METERS"
		case OFPMMFC_OUT_OF_BANDS:
			s = "OFPET_METER_MOD_FAILED OFPMMFC_OUT_OF_BANDS"
		}
	case OFPET_TABLE_FEATURES_FAILED:
		switch self.Code() {
		case OFPTFFC_BAD_TABLE:
			s = "OFPET_TABLE_FEATURES_FAILED OFPTFFC_BAD_TABLE"
		case OFPTFFC_BAD_METADATA:
			s = "OFPET_TABLE_FEATURES_FAILED OFPTFFC_BAD_METADATA"
		case OFPTFFC_BAD_TYPE:
			s = "OFPET_TABLE_FEATURES_FAILED OFPTFFC_BAD_TYPE"
		case OFPTFFC_BAD_LEN:
			s = "OFPET_TABLE_FEATURES_FAILED OFPTFFC_BAD_LEN"
		case OFPTFFC_BAD_ARGUMENT:
			s = "OFPET_TABLE_FEATURES_FAILED OFPTFFC_BAD_ARGUMENT"
		case OFPTFFC_EPERM:
			s = "OFPET_TABLE_FEATURES_FAILED OFPTFFC_EPERM"
		}
	case OFPET_EXPERIMENTER:
		exp := ErrorExperimenterMsg(self)
		s = fmt.Sprintf("OFPET_EXPERIMENTER Experimenter=%x ExpType=%x", exp.Experimenter(), exp.ExpType())
	}
	return s
}

func MakeErrorMsg(etype, ecode uint16) ErrorMsg {
	self := make([]byte, 12)
	self[0] = 4
	self[1] = OFPT_ERROR
	binary.BigEndian.PutUint16(self[2:], 12)
	binary.BigEndian.PutUint16(self[8:], etype)
	binary.BigEndian.PutUint16(self[10:], ecode)
	return ErrorMsg(self)
}

type ErrorExperimenterMsg []byte

func (self ErrorExperimenterMsg) ExpType() uint16 {
	return binary.BigEndian.Uint16(self[10:])
}

func (self ErrorExperimenterMsg) Experimenter() uint32 {
	return binary.BigEndian.Uint32(self[12:])
}

func (self ErrorExperimenterMsg) Data() []uint8 {
	return self[16:]
}
