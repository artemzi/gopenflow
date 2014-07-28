package ofp4

import (
	"encoding"
	"encoding/binary"
)

type actionList []Action

func (obj actionList) MarshalBinary() ([]byte, error) {
	var data []byte
	for _, action := range []Action(obj) {
		if buf, err := action.MarshalBinary(); err != nil {
			return nil, err
		} else {
			data = append(data, buf...)
		}
	}
	return data, nil
}

func (obj *actionList) UnmarshalBinary(data []byte) (err error) {
	var actions []Action
	for cur := 0; cur < len(data); {
		atype := binary.BigEndian.Uint16(data[cur : 2+cur])
		alen := int(binary.BigEndian.Uint16(data[2+cur : 4+cur]))
		var action Action
		switch atype {
		default:
			return Error{OFPET_BAD_ACTION, OFPBAC_BAD_TYPE, nil}
		case OFPAT_COPY_TTL_OUT, OFPAT_COPY_TTL_IN, OFPAT_DEC_MPLS_TTL, OFPAT_POP_VLAN, OFPAT_DEC_NW_TTL, OFPAT_POP_PBB:
			action = new(ActionGeneric)
		case OFPAT_OUTPUT:
			action = new(ActionOutput)
		case OFPAT_SET_MPLS_TTL:
			action = new(ActionMplsTtl)
		case OFPAT_PUSH_VLAN, OFPAT_PUSH_MPLS, OFPAT_PUSH_PBB:
			action = new(ActionPush)
		case OFPAT_POP_MPLS:
			action = new(ActionPopMpls)
		case OFPAT_SET_QUEUE:
			action = new(ActionSetQueue)
		case OFPAT_GROUP:
			action = new(ActionGroup)
		case OFPAT_SET_NW_TTL:
			action = new(ActionNwTtl)
		case OFPAT_SET_FIELD:
			action = new(ActionSetField)
		case OFPAT_EXPERIMENTER:
			action = new(ActionExperimenter)
		}
		if err = action.(encoding.BinaryUnmarshaler).UnmarshalBinary(data[cur : cur+alen]); err != nil {
			return
		}
		actions = append(actions, action)
		cur += alen
	}
	*obj = actionList(actions)
	return
}

type actionIdList []Action

func (obj actionIdList) MarshalBinary() ([]byte, error) {
	var data []byte
	for _, action := range []Action(obj) {
		if buf, err := action.MarshalBinary(); err != nil {
			return nil, err
		} else {
			data = append(data, buf...)
		}
	}
	return data, nil
}

func (obj *actionIdList) UnmarshalBinary(data []byte) error {
	var actions []Action
	for cur := 0; cur < len(data); {
		atype := binary.BigEndian.Uint16(data[cur : 2+cur])
		alen := int(binary.BigEndian.Uint16(data[2+cur : 4+cur]))
		var action Action
		switch atype {
		default:
			return Error{OFPET_BAD_ACTION, OFPBAC_BAD_TYPE, nil}
		case OFPAT_OUTPUT,
			OFPAT_COPY_TTL_OUT,
			OFPAT_COPY_TTL_IN,
			OFPAT_SET_MPLS_TTL,
			OFPAT_DEC_MPLS_TTL,
			OFPAT_PUSH_VLAN,
			OFPAT_POP_VLAN,
			OFPAT_PUSH_MPLS,
			OFPAT_POP_MPLS,
			OFPAT_SET_QUEUE,
			OFPAT_GROUP,
			OFPAT_SET_NW_TTL,
			OFPAT_DEC_NW_TTL,
			OFPAT_SET_FIELD,
			OFPAT_PUSH_PBB,
			OFPAT_POP_PBB:
			action = new(ActionGeneric)
		case OFPAT_EXPERIMENTER:
			action = new(ActionExperimenter)
		}
		if err := action.(encoding.BinaryUnmarshaler).UnmarshalBinary(data[cur : cur+alen]); err != nil {
			return err
		}
		actions = append(actions, action)
		cur += alen
	}
	*obj = actionIdList(actions)
	return nil
}

type ActionGeneric struct {
	Type uint16
}

func (obj ActionGeneric) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8)
	binary.BigEndian.PutUint16(data[0:2], obj.Type)
	binary.BigEndian.PutUint16(data[2:4], 8)
	return data, nil
}

func (obj *ActionGeneric) UnmarshalBinary(data []byte) (err error) {
	obj.Type = binary.BigEndian.Uint16(data[0:2])
	return
}

type ActionOutput struct {
	Port   uint32
	MaxLen uint16
}

func (obj ActionOutput) MarshalBinary() ([]byte, error) {
	data := make([]byte, 16)
	binary.BigEndian.PutUint16(data[0:2], OFPAT_OUTPUT)
	binary.BigEndian.PutUint16(data[2:4], 16)
	binary.BigEndian.PutUint32(data[4:8], obj.Port)
	binary.BigEndian.PutUint16(data[8:10], obj.MaxLen)
	// 6 padding
	return data, nil
}

func (obj *ActionOutput) UnmarshalBinary(data []byte) (err error) {
	obj.Port = binary.BigEndian.Uint32(data[4:8])
	obj.MaxLen = binary.BigEndian.Uint16(data[8:10])
	return
}

type ActionMplsTtl struct {
	MplsTtl uint8
}

func (obj ActionMplsTtl) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8)
	binary.BigEndian.PutUint16(data[0:2], OFPAT_SET_MPLS_TTL)
	binary.BigEndian.PutUint16(data[2:4], 8)
	data[4] = obj.MplsTtl
	// 3 padding
	return data, nil
}

func (obj *ActionMplsTtl) UnmarshalBinary(data []byte) (err error) {
	obj.MplsTtl = data[4]
	return
}

type ActionPush struct {
	Type      uint16
	Ethertype uint16
}

func (obj ActionPush) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8)
	binary.BigEndian.PutUint16(data[0:2], obj.Type)
	binary.BigEndian.PutUint16(data[2:4], 8)
	binary.BigEndian.PutUint16(data[4:6], obj.Ethertype)
	// 2 padding
	return data, nil
}

func (obj *ActionPush) UnmarshalBinary(data []byte) (err error) {
	obj.Type = binary.BigEndian.Uint16(data[0:2])
	obj.Ethertype = binary.BigEndian.Uint16(data[4:6])
	return
}

type ActionPopMpls struct {
	Ethertype uint16
}

func (obj ActionPopMpls) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8)
	binary.BigEndian.PutUint16(data[0:2], OFPAT_POP_MPLS)
	binary.BigEndian.PutUint16(data[2:4], 8)
	binary.BigEndian.PutUint16(data[4:6], obj.Ethertype)
	// 2 padding
	return data, nil
}

func (obj *ActionPopMpls) UnmarshalBinary(data []byte) (err error) {
	obj.Ethertype = binary.BigEndian.Uint16(data[4:6])
	return
}

type ActionSetQueue struct {
	QueueId uint32
}

func (obj ActionSetQueue) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8)
	binary.BigEndian.PutUint16(data[0:2], OFPAT_SET_QUEUE)
	binary.BigEndian.PutUint16(data[2:4], 8)
	binary.BigEndian.PutUint32(data[4:8], obj.QueueId)
	return data, nil
}

func (obj *ActionSetQueue) UnmarshalBinary(data []byte) (err error) {
	obj.QueueId = binary.BigEndian.Uint32(data[4:8])
	return
}

type ActionGroup struct {
	GroupId uint32
}

func (obj ActionGroup) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8)
	binary.BigEndian.PutUint16(data[0:2], OFPAT_GROUP)
	binary.BigEndian.PutUint16(data[2:4], 8)
	binary.BigEndian.PutUint32(data[4:8], obj.GroupId)
	return data, nil
}

func (obj *ActionGroup) UnmarshalBinary(data []byte) (err error) {
	obj.GroupId = binary.BigEndian.Uint32(data[4:8])
	return
}

type ActionNwTtl struct {
	NwTtl uint8
}

func (obj ActionNwTtl) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8)
	binary.BigEndian.PutUint16(data[0:2], OFPAT_SET_NW_TTL)
	binary.BigEndian.PutUint16(data[2:4], 8)
	data[4] = obj.NwTtl
	// 3 padding
	return data, nil
}

func (obj *ActionNwTtl) UnmarshalBinary(data []byte) error {
	obj.NwTtl = data[4]
	return nil
}

type ActionSetField struct {
	Field []byte
}

func (obj ActionSetField) MarshalBinary() ([]byte, error) {
	length := align8(4 + len(obj.Field))
	data := make([]byte, length)
	binary.BigEndian.PutUint16(data[0:2], OFPAT_SET_FIELD)
	binary.BigEndian.PutUint16(data[2:4], uint16(length))
	copy(data[4:], obj.Field)
	return data, nil
}

func (obj *ActionSetField) UnmarshalBinary(data []byte) (err error) {
	oxm_length := int(data[7] & 0x7F)
	obj.Field = data[4 : 8+oxm_length]
	return
}

type ActionExperimenter struct {
	Experimenter uint32
	Data         []byte
}

func (obj ActionExperimenter) MarshalBinary() ([]byte, error) {
	length := 8 + align8(len(obj.Data))
	data := make([]byte, length)
	binary.BigEndian.PutUint16(data[0:2], OFPAT_EXPERIMENTER)
	binary.BigEndian.PutUint16(data[2:4], uint16(length))
	binary.BigEndian.PutUint32(data[4:8], obj.Experimenter)
	copy(data[8:], obj.Data)
	return data, nil
}

func (obj *ActionExperimenter) UnmarshalBinary(data []byte) (err error) {
	length := int(binary.BigEndian.Uint16(data[2:4]))
	obj.Experimenter = binary.BigEndian.Uint32(data[4:8])
	obj.Data = data[8:length]
	return
}
