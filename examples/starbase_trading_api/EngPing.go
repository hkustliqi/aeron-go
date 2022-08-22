// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngPing struct {
	CorrelationId int64
	Data          []byte
}

func (e *EngPing) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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

func (e *EngPing) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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

func (e *EngPing) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.CorrelationIdInActingVersion(actingVersion) {
		if e.CorrelationId < e.CorrelationIdMinValue() || e.CorrelationId > e.CorrelationIdMaxValue() {
			return fmt.Errorf("Range check failed on e.CorrelationId (%v < %v > %v)", e.CorrelationIdMinValue(), e.CorrelationId, e.CorrelationIdMaxValue())
		}
	}
	return nil
}

func EngPingInit(e *EngPing) {
	return
}

func (*EngPing) SbeBlockLength() (blockLength uint16) {
	return 8
}

func (*EngPing) SbeTemplateId() (templateId uint16) {
	return 2021
}

func (*EngPing) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngPing) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngPing) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngPing) CorrelationIdId() uint16 {
	return 1
}

func (*EngPing) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (e *EngPing) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CorrelationIdSinceVersion()
}

func (*EngPing) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*EngPing) CorrelationIdMetaAttribute(meta int) string {
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

func (*EngPing) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngPing) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngPing) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*EngPing) DataMetaAttribute(meta int) string {
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

func (*EngPing) DataSinceVersion() uint16 {
	return 0
}

func (e *EngPing) DataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DataSinceVersion()
}

func (*EngPing) DataDeprecated() uint16 {
	return 0
}

func (EngPing) DataCharacterEncoding() string {
	return "null"
}

func (EngPing) DataHeaderLength() uint64 {
	return 2
}
