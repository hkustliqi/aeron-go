// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
)

type OrderRejectDetails struct {
	Reason  NewOrderRejectReasonEnum
	Details []byte
}

func (o *OrderRejectDetails) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := o.Reason.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(o.Details))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Details); err != nil {
		return err
	}
	return nil
}

func (o *OrderRejectDetails) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if o.ReasonInActingVersion(actingVersion) {
		if err := o.Reason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}

	if o.DetailsInActingVersion(actingVersion) {
		var DetailsLength uint16
		if err := _m.ReadUint16(_r, &DetailsLength); err != nil {
			return err
		}
		if cap(o.Details) < int(DetailsLength) {
			o.Details = make([]byte, DetailsLength)
		}
		o.Details = o.Details[:DetailsLength]
		if err := _m.ReadBytes(_r, o.Details); err != nil {
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

func (o *OrderRejectDetails) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := o.Reason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	for idx, ch := range o.Details {
		if ch > 127 {
			return fmt.Errorf("o.Details[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	return nil
}

func OrderRejectDetailsInit(o *OrderRejectDetails) {
	return
}

func (*OrderRejectDetails) SbeBlockLength() (blockLength uint16) {
	return 2
}

func (*OrderRejectDetails) SbeTemplateId() (templateId uint16) {
	return 211
}

func (*OrderRejectDetails) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OrderRejectDetails) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OrderRejectDetails) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OrderRejectDetails) ReasonId() uint16 {
	return 1
}

func (*OrderRejectDetails) ReasonSinceVersion() uint16 {
	return 0
}

func (o *OrderRejectDetails) ReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ReasonSinceVersion()
}

func (*OrderRejectDetails) ReasonDeprecated() uint16 {
	return 0
}

func (*OrderRejectDetails) ReasonMetaAttribute(meta int) string {
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

func (*OrderRejectDetails) DetailsMetaAttribute(meta int) string {
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

func (*OrderRejectDetails) DetailsSinceVersion() uint16 {
	return 0
}

func (o *OrderRejectDetails) DetailsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DetailsSinceVersion()
}

func (*OrderRejectDetails) DetailsDeprecated() uint16 {
	return 0
}

func (OrderRejectDetails) DetailsCharacterEncoding() string {
	return "US-ASCII"
}

func (OrderRejectDetails) DetailsHeaderLength() uint64 {
	return 2
}
