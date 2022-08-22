// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngCreateInstrumentReply struct {
	CorrelationId int64
	InstrumentId  int64
	Result        InstrumentCreateResultEnum
}

func (e *EngCreateInstrumentReply) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.CorrelationId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.InstrumentId); err != nil {
		return err
	}
	if err := e.Result.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (e *EngCreateInstrumentReply) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.CorrelationIdInActingVersion(actingVersion) {
		e.CorrelationId = e.CorrelationIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.CorrelationId); err != nil {
			return err
		}
	}
	if !e.InstrumentIdInActingVersion(actingVersion) {
		e.InstrumentId = e.InstrumentIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.InstrumentId); err != nil {
			return err
		}
	}
	if e.ResultInActingVersion(actingVersion) {
		if err := e.Result.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := e.RangeCheck(actingVersion, e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (e *EngCreateInstrumentReply) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.CorrelationIdInActingVersion(actingVersion) {
		if e.CorrelationId < e.CorrelationIdMinValue() || e.CorrelationId > e.CorrelationIdMaxValue() {
			return fmt.Errorf("Range check failed on e.CorrelationId (%v < %v > %v)", e.CorrelationIdMinValue(), e.CorrelationId, e.CorrelationIdMaxValue())
		}
	}
	if e.InstrumentIdInActingVersion(actingVersion) {
		if e.InstrumentId != e.InstrumentIdNullValue() && (e.InstrumentId < e.InstrumentIdMinValue() || e.InstrumentId > e.InstrumentIdMaxValue()) {
			return fmt.Errorf("Range check failed on e.InstrumentId (%v < %v > %v)", e.InstrumentIdMinValue(), e.InstrumentId, e.InstrumentIdMaxValue())
		}
	}
	if err := e.Result.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func EngCreateInstrumentReplyInit(e *EngCreateInstrumentReply) {
	e.InstrumentId = math.MinInt64
	return
}

func (*EngCreateInstrumentReply) SbeBlockLength() (blockLength uint16) {
	return 17
}

func (*EngCreateInstrumentReply) SbeTemplateId() (templateId uint16) {
	return 2013
}

func (*EngCreateInstrumentReply) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngCreateInstrumentReply) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngCreateInstrumentReply) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngCreateInstrumentReply) CorrelationIdId() uint16 {
	return 1
}

func (*EngCreateInstrumentReply) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentReply) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CorrelationIdSinceVersion()
}

func (*EngCreateInstrumentReply) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentReply) CorrelationIdMetaAttribute(meta int) string {
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

func (*EngCreateInstrumentReply) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrumentReply) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrumentReply) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrumentReply) InstrumentIdId() uint16 {
	return 2
}

func (*EngCreateInstrumentReply) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentReply) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentIdSinceVersion()
}

func (*EngCreateInstrumentReply) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentReply) InstrumentIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*EngCreateInstrumentReply) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrumentReply) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrumentReply) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrumentReply) ResultId() uint16 {
	return 3
}

func (*EngCreateInstrumentReply) ResultSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentReply) ResultInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ResultSinceVersion()
}

func (*EngCreateInstrumentReply) ResultDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentReply) ResultMetaAttribute(meta int) string {
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
