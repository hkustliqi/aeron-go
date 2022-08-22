// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsHeartbeat struct {
	CorrelationId int64
}

func (o *OmsHeartbeat) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.CorrelationId); err != nil {
		return err
	}
	return nil
}

func (o *OmsHeartbeat) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.CorrelationIdInActingVersion(actingVersion) {
		o.CorrelationId = o.CorrelationIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.CorrelationId); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OmsHeartbeat) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.CorrelationIdInActingVersion(actingVersion) {
		if o.CorrelationId < o.CorrelationIdMinValue() || o.CorrelationId > o.CorrelationIdMaxValue() {
			return fmt.Errorf("Range check failed on o.CorrelationId (%v < %v > %v)", o.CorrelationIdMinValue(), o.CorrelationId, o.CorrelationIdMaxValue())
		}
	}
	return nil
}

func OmsHeartbeatInit(o *OmsHeartbeat) {
	return
}

func (*OmsHeartbeat) SbeBlockLength() (blockLength uint16) {
	return 8
}

func (*OmsHeartbeat) SbeTemplateId() (templateId uint16) {
	return 1009
}

func (*OmsHeartbeat) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsHeartbeat) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsHeartbeat) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsHeartbeat) CorrelationIdId() uint16 {
	return 1
}

func (*OmsHeartbeat) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (o *OmsHeartbeat) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CorrelationIdSinceVersion()
}

func (*OmsHeartbeat) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*OmsHeartbeat) CorrelationIdMetaAttribute(meta int) string {
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

func (*OmsHeartbeat) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsHeartbeat) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsHeartbeat) CorrelationIdNullValue() int64 {
	return math.MinInt64
}
