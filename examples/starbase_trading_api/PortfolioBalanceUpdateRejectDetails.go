// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"io"
	"io/ioutil"
)

type PortfolioBalanceUpdateRejectDetails struct {
	RejectReason PortfolioBalanceUpdateRejectReasonEnum
}

func (p *PortfolioBalanceUpdateRejectDetails) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := p.RangeCheck(p.SbeSchemaVersion(), p.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := p.RejectReason.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (p *PortfolioBalanceUpdateRejectDetails) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if p.RejectReasonInActingVersion(actingVersion) {
		if err := p.RejectReason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > p.SbeSchemaVersion() && blockLength > p.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-p.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := p.RangeCheck(actingVersion, p.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (p *PortfolioBalanceUpdateRejectDetails) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := p.RejectReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func PortfolioBalanceUpdateRejectDetailsInit(p *PortfolioBalanceUpdateRejectDetails) {
	return
}

func (*PortfolioBalanceUpdateRejectDetails) SbeBlockLength() (blockLength uint16) {
	return 1
}

func (*PortfolioBalanceUpdateRejectDetails) SbeTemplateId() (templateId uint16) {
	return 1211
}

func (*PortfolioBalanceUpdateRejectDetails) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*PortfolioBalanceUpdateRejectDetails) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*PortfolioBalanceUpdateRejectDetails) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*PortfolioBalanceUpdateRejectDetails) RejectReasonId() uint16 {
	return 2
}

func (*PortfolioBalanceUpdateRejectDetails) RejectReasonSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateRejectDetails) RejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.RejectReasonSinceVersion()
}

func (*PortfolioBalanceUpdateRejectDetails) RejectReasonDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateRejectDetails) RejectReasonMetaAttribute(meta int) string {
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
