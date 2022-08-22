// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsMoonbaseOrderRequestConf struct {
	CorrelationId int64
	EventType     OrderEventTypeEnum
	RejectReason  NewOrderRejectReasonEnum
	RejectDetails []byte
}

func (o *OmsMoonbaseOrderRequestConf) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.CorrelationId); err != nil {
		return err
	}
	if err := o.EventType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.RejectReason.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(o.RejectDetails))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.RejectDetails); err != nil {
		return err
	}
	return nil
}

func (o *OmsMoonbaseOrderRequestConf) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.CorrelationIdInActingVersion(actingVersion) {
		o.CorrelationId = o.CorrelationIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.CorrelationId); err != nil {
			return err
		}
	}
	if o.EventTypeInActingVersion(actingVersion) {
		if err := o.EventType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.RejectReasonInActingVersion(actingVersion) {
		if err := o.RejectReason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}

	if o.RejectDetailsInActingVersion(actingVersion) {
		var RejectDetailsLength uint16
		if err := _m.ReadUint16(_r, &RejectDetailsLength); err != nil {
			return err
		}
		if cap(o.RejectDetails) < int(RejectDetailsLength) {
			o.RejectDetails = make([]byte, RejectDetailsLength)
		}
		o.RejectDetails = o.RejectDetails[:RejectDetailsLength]
		if err := _m.ReadBytes(_r, o.RejectDetails); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OmsMoonbaseOrderRequestConf) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.CorrelationIdInActingVersion(actingVersion) {
		if o.CorrelationId < o.CorrelationIdMinValue() || o.CorrelationId > o.CorrelationIdMaxValue() {
			return fmt.Errorf("Range check failed on o.CorrelationId (%v < %v > %v)", o.CorrelationIdMinValue(), o.CorrelationId, o.CorrelationIdMaxValue())
		}
	}
	if err := o.EventType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.RejectReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	for idx, ch := range o.RejectDetails {
		if ch > 127 {
			return fmt.Errorf("o.RejectDetails[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	return nil
}

func OmsMoonbaseOrderRequestConfInit(o *OmsMoonbaseOrderRequestConf) {
	return
}

func (*OmsMoonbaseOrderRequestConf) SbeBlockLength() (blockLength uint16) {
	return 11
}

func (*OmsMoonbaseOrderRequestConf) SbeTemplateId() (templateId uint16) {
	return 1003
}

func (*OmsMoonbaseOrderRequestConf) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsMoonbaseOrderRequestConf) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsMoonbaseOrderRequestConf) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsMoonbaseOrderRequestConf) CorrelationIdId() uint16 {
	return 1
}

func (*OmsMoonbaseOrderRequestConf) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseOrderRequestConf) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CorrelationIdSinceVersion()
}

func (*OmsMoonbaseOrderRequestConf) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseOrderRequestConf) CorrelationIdMetaAttribute(meta int) string {
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

func (*OmsMoonbaseOrderRequestConf) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsMoonbaseOrderRequestConf) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsMoonbaseOrderRequestConf) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsMoonbaseOrderRequestConf) EventTypeId() uint16 {
	return 2
}

func (*OmsMoonbaseOrderRequestConf) EventTypeSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseOrderRequestConf) EventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.EventTypeSinceVersion()
}

func (*OmsMoonbaseOrderRequestConf) EventTypeDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseOrderRequestConf) EventTypeMetaAttribute(meta int) string {
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

func (*OmsMoonbaseOrderRequestConf) RejectReasonId() uint16 {
	return 3
}

func (*OmsMoonbaseOrderRequestConf) RejectReasonSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseOrderRequestConf) RejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RejectReasonSinceVersion()
}

func (*OmsMoonbaseOrderRequestConf) RejectReasonDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseOrderRequestConf) RejectReasonMetaAttribute(meta int) string {
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

func (*OmsMoonbaseOrderRequestConf) RejectDetailsMetaAttribute(meta int) string {
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

func (*OmsMoonbaseOrderRequestConf) RejectDetailsSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseOrderRequestConf) RejectDetailsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RejectDetailsSinceVersion()
}

func (*OmsMoonbaseOrderRequestConf) RejectDetailsDeprecated() uint16 {
	return 0
}

func (OmsMoonbaseOrderRequestConf) RejectDetailsCharacterEncoding() string {
	return "US-ASCII"
}

func (OmsMoonbaseOrderRequestConf) RejectDetailsHeaderLength() uint64 {
	return 2
}
