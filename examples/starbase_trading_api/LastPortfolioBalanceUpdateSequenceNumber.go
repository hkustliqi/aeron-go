// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type LastPortfolioBalanceUpdateSequenceNumber struct {
	ReplyHeader        OmsAdminReplyHeader
	LastSequenceNumber int64
}

func (l *LastPortfolioBalanceUpdateSequenceNumber) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := l.RangeCheck(l.SbeSchemaVersion(), l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := l.ReplyHeader.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, l.LastSequenceNumber); err != nil {
		return err
	}
	return nil
}

func (l *LastPortfolioBalanceUpdateSequenceNumber) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if l.ReplyHeaderInActingVersion(actingVersion) {
		if err := l.ReplyHeader.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !l.LastSequenceNumberInActingVersion(actingVersion) {
		l.LastSequenceNumber = l.LastSequenceNumberNullValue()
	} else {
		if err := _m.ReadInt64(_r, &l.LastSequenceNumber); err != nil {
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

func (l *LastPortfolioBalanceUpdateSequenceNumber) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if l.LastSequenceNumberInActingVersion(actingVersion) {
		if l.LastSequenceNumber < l.LastSequenceNumberMinValue() || l.LastSequenceNumber > l.LastSequenceNumberMaxValue() {
			return fmt.Errorf("Range check failed on l.LastSequenceNumber (%v < %v > %v)", l.LastSequenceNumberMinValue(), l.LastSequenceNumber, l.LastSequenceNumberMaxValue())
		}
	}
	return nil
}

func LastPortfolioBalanceUpdateSequenceNumberInit(l *LastPortfolioBalanceUpdateSequenceNumber) {
	return
}

func (*LastPortfolioBalanceUpdateSequenceNumber) SbeBlockLength() (blockLength uint16) {
	return 32
}

func (*LastPortfolioBalanceUpdateSequenceNumber) SbeTemplateId() (templateId uint16) {
	return 1704
}

func (*LastPortfolioBalanceUpdateSequenceNumber) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*LastPortfolioBalanceUpdateSequenceNumber) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*LastPortfolioBalanceUpdateSequenceNumber) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*LastPortfolioBalanceUpdateSequenceNumber) ReplyHeaderId() uint16 {
	return 1
}

func (*LastPortfolioBalanceUpdateSequenceNumber) ReplyHeaderSinceVersion() uint16 {
	return 0
}

func (l *LastPortfolioBalanceUpdateSequenceNumber) ReplyHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ReplyHeaderSinceVersion()
}

func (*LastPortfolioBalanceUpdateSequenceNumber) ReplyHeaderDeprecated() uint16 {
	return 0
}

func (*LastPortfolioBalanceUpdateSequenceNumber) ReplyHeaderMetaAttribute(meta int) string {
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

func (*LastPortfolioBalanceUpdateSequenceNumber) LastSequenceNumberId() uint16 {
	return 2
}

func (*LastPortfolioBalanceUpdateSequenceNumber) LastSequenceNumberSinceVersion() uint16 {
	return 0
}

func (l *LastPortfolioBalanceUpdateSequenceNumber) LastSequenceNumberInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.LastSequenceNumberSinceVersion()
}

func (*LastPortfolioBalanceUpdateSequenceNumber) LastSequenceNumberDeprecated() uint16 {
	return 0
}

func (*LastPortfolioBalanceUpdateSequenceNumber) LastSequenceNumberMetaAttribute(meta int) string {
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

func (*LastPortfolioBalanceUpdateSequenceNumber) LastSequenceNumberMinValue() int64 {
	return math.MinInt64 + 1
}

func (*LastPortfolioBalanceUpdateSequenceNumber) LastSequenceNumberMaxValue() int64 {
	return math.MaxInt64
}

func (*LastPortfolioBalanceUpdateSequenceNumber) LastSequenceNumberNullValue() int64 {
	return math.MinInt64
}
