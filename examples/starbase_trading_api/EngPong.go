// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngPong struct {
	CorrelationId int64
	Data          []byte
}

func (e *EngPong) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.CorrelationId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(e.Data))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.Data); err != nil {
		return err
	}
	return nil
}

func (e *EngPong) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.CorrelationIdInActingVersion(actingVersion) {
		e.CorrelationId = e.CorrelationIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.CorrelationId); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}

	if e.DataInActingVersion(actingVersion) {
		var DataLength uint16
		if err := _m.ReadUint16(_r, &DataLength); err != nil {
			return err
		}
		if cap(e.Data) < int(DataLength) {
			e.Data = make([]byte, DataLength)
		}
		e.Data = e.Data[:DataLength]
		if err := _m.ReadBytes(_r, e.Data); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := e.RangeCheck(actingVersion, e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (e *EngPong) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.CorrelationIdInActingVersion(actingVersion) {
		if e.CorrelationId < e.CorrelationIdMinValue() || e.CorrelationId > e.CorrelationIdMaxValue() {
			return fmt.Errorf("Range check failed on e.CorrelationId (%v < %v > %v)", e.CorrelationIdMinValue(), e.CorrelationId, e.CorrelationIdMaxValue())
		}
	}
	return nil
}

func EngPongInit(e *EngPong) {
	return
}

func (*EngPong) SbeBlockLength() (blockLength uint16) {
	return 8
}

func (*EngPong) SbeTemplateId() (templateId uint16) {
	return 2022
}

func (*EngPong) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngPong) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngPong) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngPong) CorrelationIdId() uint16 {
	return 1
}

func (*EngPong) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (e *EngPong) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CorrelationIdSinceVersion()
}

func (*EngPong) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*EngPong) CorrelationIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*EngPong) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngPong) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngPong) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*EngPong) DataMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*EngPong) DataSinceVersion() uint16 {
	return 0
}

func (e *EngPong) DataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DataSinceVersion()
}

func (*EngPong) DataDeprecated() uint16 {
	return 0
}

func (EngPong) DataCharacterEncoding() string {
	return "null"
}

func (EngPong) DataHeaderLength() uint64 {
	return 2
}
