// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"io"
	"io/ioutil"
)

type OmsAdminRequestReplyPart struct {
	ReplyHeader   OmsAdminReplyHeader
	DetailsHeader MessageHeader
}

func (o *OmsAdminRequestReplyPart) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := o.ReplyHeader.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.DetailsHeader.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (o *OmsAdminRequestReplyPart) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if o.ReplyHeaderInActingVersion(actingVersion) {
		if err := o.ReplyHeader.Decode(_m, _r, actingVersion); err != nil {
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

func (o *OmsAdminRequestReplyPart) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func OmsAdminRequestReplyPartInit(o *OmsAdminRequestReplyPart) {
	return
}

func (*OmsAdminRequestReplyPart) SbeBlockLength() (blockLength uint16) {
	return 32
}

func (*OmsAdminRequestReplyPart) SbeTemplateId() (templateId uint16) {
	return 1201
}

func (*OmsAdminRequestReplyPart) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsAdminRequestReplyPart) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsAdminRequestReplyPart) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsAdminRequestReplyPart) ReplyHeaderId() uint16 {
	return 1
}

func (*OmsAdminRequestReplyPart) ReplyHeaderSinceVersion() uint16 {
	return 0
}

func (o *OmsAdminRequestReplyPart) ReplyHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ReplyHeaderSinceVersion()
}

func (*OmsAdminRequestReplyPart) ReplyHeaderDeprecated() uint16 {
	return 0
}

func (*OmsAdminRequestReplyPart) ReplyHeaderMetaAttribute(meta int) string {
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

func (*OmsAdminRequestReplyPart) DetailsHeaderId() uint16 {
	return 2
}

func (*OmsAdminRequestReplyPart) DetailsHeaderSinceVersion() uint16 {
	return 0
}

func (o *OmsAdminRequestReplyPart) DetailsHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DetailsHeaderSinceVersion()
}

func (*OmsAdminRequestReplyPart) DetailsHeaderDeprecated() uint16 {
	return 0
}

func (*OmsAdminRequestReplyPart) DetailsHeaderMetaAttribute(meta int) string {
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
