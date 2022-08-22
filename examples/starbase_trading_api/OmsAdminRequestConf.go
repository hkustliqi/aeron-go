// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"io"
	"io/ioutil"
)

type OmsAdminRequestConf struct {
	ReplyHeader   OmsAdminReplyHeader
	DetailsHeader MessageHeader
}

func (o *OmsAdminRequestConf) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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

func (o *OmsAdminRequestConf) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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

func (o *OmsAdminRequestConf) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func OmsAdminRequestConfInit(o *OmsAdminRequestConf) {
	return
}

func (*OmsAdminRequestConf) SbeBlockLength() (blockLength uint16) {
	return 32
}

func (*OmsAdminRequestConf) SbeTemplateId() (templateId uint16) {
	return 1200
}

func (*OmsAdminRequestConf) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsAdminRequestConf) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsAdminRequestConf) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsAdminRequestConf) ReplyHeaderId() uint16 {
	return 1
}

func (*OmsAdminRequestConf) ReplyHeaderSinceVersion() uint16 {
	return 0
}

func (o *OmsAdminRequestConf) ReplyHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ReplyHeaderSinceVersion()
}

func (*OmsAdminRequestConf) ReplyHeaderDeprecated() uint16 {
	return 0
}

func (*OmsAdminRequestConf) ReplyHeaderMetaAttribute(meta int) string {
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

func (*OmsAdminRequestConf) DetailsHeaderId() uint16 {
	return 2
}

func (*OmsAdminRequestConf) DetailsHeaderSinceVersion() uint16 {
	return 0
}

func (o *OmsAdminRequestConf) DetailsHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DetailsHeaderSinceVersion()
}

func (*OmsAdminRequestConf) DetailsHeaderDeprecated() uint16 {
	return 0
}

func (*OmsAdminRequestConf) DetailsHeaderMetaAttribute(meta int) string {
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
