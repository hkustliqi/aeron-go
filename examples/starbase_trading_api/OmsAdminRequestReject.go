// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
)

type OmsAdminRequestReject struct {
	ReplyHeader   OmsAdminReplyHeader
	ErrorMessage  [80]byte
	DetailsHeader MessageHeader
}

func (o *OmsAdminRequestReject) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := o.ReplyHeader.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.ErrorMessage[:]); err != nil {
		return err
	}
	if err := o.DetailsHeader.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (o *OmsAdminRequestReject) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if o.ReplyHeaderInActingVersion(actingVersion) {
		if err := o.ReplyHeader.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.ErrorMessageInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			o.ErrorMessage[idx] = o.ErrorMessageNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.ErrorMessage[:]); err != nil {
			return err
		}
	}
	if o.DetailsHeaderInActingVersion(actingVersion) {
		if err := o.DetailsHeader.Decode(_m, _r, actingVersion); err != nil {
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

func (o *OmsAdminRequestReject) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.ErrorMessageInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if o.ErrorMessage[idx] < o.ErrorMessageMinValue() || o.ErrorMessage[idx] > o.ErrorMessageMaxValue() {
				return fmt.Errorf("Range check failed on o.ErrorMessage[%d] (%v < %v > %v)", idx, o.ErrorMessageMinValue(), o.ErrorMessage[idx], o.ErrorMessageMaxValue())
			}
		}
	}
	for idx, ch := range o.ErrorMessage {
		if ch > 127 {
			return fmt.Errorf("o.ErrorMessage[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	return nil
}

func OmsAdminRequestRejectInit(o *OmsAdminRequestReject) {
	return
}

func (*OmsAdminRequestReject) SbeBlockLength() (blockLength uint16) {
	return 112
}

func (*OmsAdminRequestReject) SbeTemplateId() (templateId uint16) {
	return 1202
}

func (*OmsAdminRequestReject) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsAdminRequestReject) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsAdminRequestReject) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsAdminRequestReject) ReplyHeaderId() uint16 {
	return 1
}

func (*OmsAdminRequestReject) ReplyHeaderSinceVersion() uint16 {
	return 0
}

func (o *OmsAdminRequestReject) ReplyHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ReplyHeaderSinceVersion()
}

func (*OmsAdminRequestReject) ReplyHeaderDeprecated() uint16 {
	return 0
}

func (*OmsAdminRequestReject) ReplyHeaderMetaAttribute(meta int) string {
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

func (*OmsAdminRequestReject) ErrorMessageId() uint16 {
	return 2
}

func (*OmsAdminRequestReject) ErrorMessageSinceVersion() uint16 {
	return 0
}

func (o *OmsAdminRequestReject) ErrorMessageInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ErrorMessageSinceVersion()
}

func (*OmsAdminRequestReject) ErrorMessageDeprecated() uint16 {
	return 0
}

func (*OmsAdminRequestReject) ErrorMessageMetaAttribute(meta int) string {
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

func (*OmsAdminRequestReject) ErrorMessageMinValue() byte {
	return byte(32)
}

func (*OmsAdminRequestReject) ErrorMessageMaxValue() byte {
	return byte(126)
}

func (*OmsAdminRequestReject) ErrorMessageNullValue() byte {
	return 0
}

func (o *OmsAdminRequestReject) ErrorMessageCharacterEncoding() string {
	return "US-ASCII"
}

func (*OmsAdminRequestReject) DetailsHeaderId() uint16 {
	return 3
}

func (*OmsAdminRequestReject) DetailsHeaderSinceVersion() uint16 {
	return 0
}

func (o *OmsAdminRequestReject) DetailsHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DetailsHeaderSinceVersion()
}

func (*OmsAdminRequestReject) DetailsHeaderDeprecated() uint16 {
	return 0
}

func (*OmsAdminRequestReject) DetailsHeaderMetaAttribute(meta int) string {
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
