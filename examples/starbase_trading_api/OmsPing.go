// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsPing struct {
	CorrelationId int64
	Data          []byte
}

func (o *OmsPing) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.CorrelationId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(o.Data))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Data); err != nil {
		return err
	}
	return nil
}

func (o *OmsPing) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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

	if o.DataInActingVersion(actingVersion) {
		var DataLength uint16
		if err := _m.ReadUint16(_r, &DataLength); err != nil {
			return err
		}
		if cap(o.Data) < int(DataLength) {
			o.Data = make([]byte, DataLength)
		}
		o.Data = o.Data[:DataLength]
		if err := _m.ReadBytes(_r, o.Data); err != nil {
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

func (o *OmsPing) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.CorrelationIdInActingVersion(actingVersion) {
		if o.CorrelationId < o.CorrelationIdMinValue() || o.CorrelationId > o.CorrelationIdMaxValue() {
			return fmt.Errorf("Range check failed on o.CorrelationId (%v < %v > %v)", o.CorrelationIdMinValue(), o.CorrelationId, o.CorrelationIdMaxValue())
		}
	}
	return nil
}

func OmsPingInit(o *OmsPing) {
	return
}

func (*OmsPing) SbeBlockLength() (blockLength uint16) {
	return 8
}

func (*OmsPing) SbeTemplateId() (templateId uint16) {
	return 1021
}

func (*OmsPing) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsPing) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsPing) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsPing) CorrelationIdId() uint16 {
	return 1
}

func (*OmsPing) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (o *OmsPing) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CorrelationIdSinceVersion()
}

func (*OmsPing) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*OmsPing) CorrelationIdMetaAttribute(meta int) string {
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

func (*OmsPing) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsPing) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsPing) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsPing) DataMetaAttribute(meta int) string {
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

func (*OmsPing) DataSinceVersion() uint16 {
	return 0
}

func (o *OmsPing) DataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DataSinceVersion()
}

func (*OmsPing) DataDeprecated() uint16 {
	return 0
}

func (OmsPing) DataCharacterEncoding() string {
	return "null"
}

func (OmsPing) DataHeaderLength() uint64 {
	return 2
}
