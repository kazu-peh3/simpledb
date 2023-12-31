package storage

import (
	"github.com/gogo/protobuf/proto"
	"github.com/kazu-peh3/toydb/meta"
)

func NewTuple(minTxId uint64, values []interface{}) *Tuple {
	var t Tuple
	t.MinTxId = minTxId
	t.MaxTxId = minTxId

	var td *TupleData
	for _, v := range values {
		switch concrete := v.(type) {

		case int:
			td = &TupleData{
				Type:   TupleData_INT,
				Number: *proto.Int32(int32(concrete)),
			}

		case string:
			td = &TupleData{
				Type:    TupleData_STRING,
				String_: *proto.String(concrete),
			}
		}

		t.Data = append(t.Data, td)
	}

	return &t
}

func (m *Tuple) Less(than meta.Item) bool {
	t, ok := than.(*Tuple)
	if !ok {
		return false
	}

	// FIXME
	left := m.Data[0].Number
	right := t.Data[0].Number

	return left < right
}

func SerializeTuple(t *Tuple) ([128]byte, error) {
	out, err := proto.Marshal(t)

	if err != nil {
		return [128]byte{}, err
	}

	var b [128]byte
	copy(b[:], out)

	return b, nil
}

func DeserializeTuple(b [128]byte) (*Tuple, error) {
	var t Tuple

	err := proto.Unmarshal(b[:], &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (m *Tuple) Equal(order int, s string, n int) bool {
	tupleData := m.Data[order]

	if tupleData.Type == TupleData_STRING {
		return tupleData.String_ == s
	} else if tupleData.Type == TupleData_INT {
		return tupleData.Number == int32(n)
	}

	return false
}

func (m *Tuple) IsUnused() bool {
	// If minTxId is zero, it's an empty tuple.
	return m.MinTxId == 0
}

func (m *Tuple) CanSee(tran *Transaction) bool {
	if m.MinTxId == tran.txid {
		return true
	}

	if m.MaxTxId < tran.Txid() {
		return false
	}

	if m.MinTxId > tran.Txid() && tran.state != Commited {
		return false
	}

	return true
}
