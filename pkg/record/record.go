package record

import (
	"bytes"
	"encoding/binary"
	"time"
)

type Record struct {
	Timestamp time.Time `json:"timestamp"`
	Line      string    `json:"line"`
}

func (r *Record) ToBinary() ([]byte, error) {
	buf := new(bytes.Buffer)

	timeBuf := make([]byte, 8)
	binary.LittleEndian.PutUint64(timeBuf, uint64(r.Timestamp.UnixNano()))
	if _, err := buf.Write(timeBuf); err != nil {
		return nil, err
	}

	if _, err := buf.WriteString(r.Line); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func FromBinary(v []byte) *Record {
	entry := new(Record)
	entry.Timestamp = time.Unix(0, int64(binary.LittleEndian.Uint64(v[:8])))
	entry.Line = string(v[8:])
	return entry
}
