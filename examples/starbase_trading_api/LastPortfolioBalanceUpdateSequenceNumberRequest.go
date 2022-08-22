// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"io"
	"io/ioutil"
)

type LastPortfolioBalanceUpdateSequenceNumberRequest struct {
	RequestHeader OmsAdminRequestHeader
}

func (l *LastPortfolioBalanceUpdateSequenceNumberRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := l.RangeCheck(l.SbeSchemaVersion(), l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := l.RequestHeader.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (l *LastPortfolioBalanceUpdateSequenceNumberRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if l.RequestHeaderInActingVersion(actingVersion) {
		if err := l.RequestHeader.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := l.RangeCheck(actingVersion, l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (l *LastPortfolioBalanceUpdateSequenceNumberRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func LastPortfolioBalanceUpdateSequenceNumberRequestInit(l *LastPortfolioBalanceUpdateSequenceNumberRequest) {
	return
}

func (*LastPortfolioBalanceUpdateSequenceNumberRequest) SbeBlockLength() (blockLength uint16) {
	return 24
}

func (*LastPortfolioBalanceUpdateSequenceNumberRequest) SbeTemplateId() (templateId uint16) {
	return 1703
}

func (*LastPortfolioBalanceUpdateSequenceNumberRequest) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*LastPortfolioBalanceUpdateSequenceNumberRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*LastPortfolioBalanceUpdateSequenceNumberRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*LastPortfolioBalanceUpdateSequenceNumberRequest) RequestHeaderId() uint16 {
	return 1
}

func (*LastPortfolioBalanceUpdateSequenceNumberRequest) RequestHeaderSinceVersion() uint16 {
	return 0
}

func (l *LastPortfolioBalanceUpdateSequenceNumberRequest) RequestHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.RequestHeaderSinceVersion()
}

func (*LastPortfolioBalanceUpdateSequenceNumberRequest) RequestHeaderDeprecated() uint16 {
	return 0
}

func (*LastPortfolioBalanceUpdateSequenceNumberRequest) RequestHeaderMetaAttribute(meta int) string {
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
