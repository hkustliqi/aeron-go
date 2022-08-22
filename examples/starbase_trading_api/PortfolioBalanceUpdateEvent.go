// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type PortfolioBalanceUpdateEvent struct {
	TradingHeader  TradingEventHeader
	PortfolioId    int64
	AssetId        int64
	Balance        int64
	Hold           int64
	SequenceNumber int64
	Action         PortfolioBalanceUpdateActionEnum
}

func (p *PortfolioBalanceUpdateEvent) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := p.RangeCheck(p.SbeSchemaVersion(), p.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := p.TradingHeader.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, p.PortfolioId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, p.AssetId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, p.Balance); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, p.Hold); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, p.SequenceNumber); err != nil {
		return err
	}
	if err := p.Action.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (p *PortfolioBalanceUpdateEvent) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if p.TradingHeaderInActingVersion(actingVersion) {
		if err := p.TradingHeader.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !p.PortfolioIdInActingVersion(actingVersion) {
		p.PortfolioId = p.PortfolioIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &p.PortfolioId); err != nil {
			return err
		}
	}
	if !p.AssetIdInActingVersion(actingVersion) {
		p.AssetId = p.AssetIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &p.AssetId); err != nil {
			return err
		}
	}
	if !p.BalanceInActingVersion(actingVersion) {
		p.Balance = p.BalanceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &p.Balance); err != nil {
			return err
		}
	}
	if !p.HoldInActingVersion(actingVersion) {
		p.Hold = p.HoldNullValue()
	} else {
		if err := _m.ReadInt64(_r, &p.Hold); err != nil {
			return err
		}
	}
	if !p.SequenceNumberInActingVersion(actingVersion) {
		p.SequenceNumber = p.SequenceNumberNullValue()
	} else {
		if err := _m.ReadInt64(_r, &p.SequenceNumber); err != nil {
			return err
		}
	}
	if p.ActionInActingVersion(actingVersion) {
		if err := p.Action.Decode(_m, _r, actingVersion); err != nil {
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

func (p *PortfolioBalanceUpdateEvent) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.PortfolioIdInActingVersion(actingVersion) {
		if p.PortfolioId < p.PortfolioIdMinValue() || p.PortfolioId > p.PortfolioIdMaxValue() {
			return fmt.Errorf("Range check failed on p.PortfolioId (%v < %v > %v)", p.PortfolioIdMinValue(), p.PortfolioId, p.PortfolioIdMaxValue())
		}
	}
	if p.AssetIdInActingVersion(actingVersion) {
		if p.AssetId < p.AssetIdMinValue() || p.AssetId > p.AssetIdMaxValue() {
			return fmt.Errorf("Range check failed on p.AssetId (%v < %v > %v)", p.AssetIdMinValue(), p.AssetId, p.AssetIdMaxValue())
		}
	}
	if p.BalanceInActingVersion(actingVersion) {
		if p.Balance < p.BalanceMinValue() || p.Balance > p.BalanceMaxValue() {
			return fmt.Errorf("Range check failed on p.Balance (%v < %v > %v)", p.BalanceMinValue(), p.Balance, p.BalanceMaxValue())
		}
	}
	if p.HoldInActingVersion(actingVersion) {
		if p.Hold < p.HoldMinValue() || p.Hold > p.HoldMaxValue() {
			return fmt.Errorf("Range check failed on p.Hold (%v < %v > %v)", p.HoldMinValue(), p.Hold, p.HoldMaxValue())
		}
	}
	if p.SequenceNumberInActingVersion(actingVersion) {
		if p.SequenceNumber < p.SequenceNumberMinValue() || p.SequenceNumber > p.SequenceNumberMaxValue() {
			return fmt.Errorf("Range check failed on p.SequenceNumber (%v < %v > %v)", p.SequenceNumberMinValue(), p.SequenceNumber, p.SequenceNumberMaxValue())
		}
	}
	if err := p.Action.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func PortfolioBalanceUpdateEventInit(p *PortfolioBalanceUpdateEvent) {
	return
}

func (*PortfolioBalanceUpdateEvent) SbeBlockLength() (blockLength uint16) {
	return 73
}

func (*PortfolioBalanceUpdateEvent) SbeTemplateId() (templateId uint16) {
	return 1706
}

func (*PortfolioBalanceUpdateEvent) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*PortfolioBalanceUpdateEvent) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*PortfolioBalanceUpdateEvent) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*PortfolioBalanceUpdateEvent) TradingHeaderId() uint16 {
	return 1
}

func (*PortfolioBalanceUpdateEvent) TradingHeaderSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateEvent) TradingHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.TradingHeaderSinceVersion()
}

func (*PortfolioBalanceUpdateEvent) TradingHeaderDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateEvent) TradingHeaderMetaAttribute(meta int) string {
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

func (*PortfolioBalanceUpdateEvent) PortfolioIdId() uint16 {
	return 2
}

func (*PortfolioBalanceUpdateEvent) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateEvent) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PortfolioIdSinceVersion()
}

func (*PortfolioBalanceUpdateEvent) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateEvent) PortfolioIdMetaAttribute(meta int) string {
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

func (*PortfolioBalanceUpdateEvent) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioBalanceUpdateEvent) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioBalanceUpdateEvent) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioBalanceUpdateEvent) AssetIdId() uint16 {
	return 3
}

func (*PortfolioBalanceUpdateEvent) AssetIdSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateEvent) AssetIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.AssetIdSinceVersion()
}

func (*PortfolioBalanceUpdateEvent) AssetIdDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateEvent) AssetIdMetaAttribute(meta int) string {
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

func (*PortfolioBalanceUpdateEvent) AssetIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioBalanceUpdateEvent) AssetIdMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioBalanceUpdateEvent) AssetIdNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioBalanceUpdateEvent) BalanceId() uint16 {
	return 4
}

func (*PortfolioBalanceUpdateEvent) BalanceSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateEvent) BalanceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.BalanceSinceVersion()
}

func (*PortfolioBalanceUpdateEvent) BalanceDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateEvent) BalanceMetaAttribute(meta int) string {
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

func (*PortfolioBalanceUpdateEvent) BalanceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioBalanceUpdateEvent) BalanceMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioBalanceUpdateEvent) BalanceNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioBalanceUpdateEvent) HoldId() uint16 {
	return 5
}

func (*PortfolioBalanceUpdateEvent) HoldSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateEvent) HoldInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.HoldSinceVersion()
}

func (*PortfolioBalanceUpdateEvent) HoldDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateEvent) HoldMetaAttribute(meta int) string {
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

func (*PortfolioBalanceUpdateEvent) HoldMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioBalanceUpdateEvent) HoldMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioBalanceUpdateEvent) HoldNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioBalanceUpdateEvent) SequenceNumberId() uint16 {
	return 6
}

func (*PortfolioBalanceUpdateEvent) SequenceNumberSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateEvent) SequenceNumberInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SequenceNumberSinceVersion()
}

func (*PortfolioBalanceUpdateEvent) SequenceNumberDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateEvent) SequenceNumberMetaAttribute(meta int) string {
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

func (*PortfolioBalanceUpdateEvent) SequenceNumberMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioBalanceUpdateEvent) SequenceNumberMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioBalanceUpdateEvent) SequenceNumberNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioBalanceUpdateEvent) ActionId() uint16 {
	return 7
}

func (*PortfolioBalanceUpdateEvent) ActionSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateEvent) ActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ActionSinceVersion()
}

func (*PortfolioBalanceUpdateEvent) ActionDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateEvent) ActionMetaAttribute(meta int) string {
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
