// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type PortfolioBalanceSnapshot struct {
	PortfolioId int64
	AssetId     int64
	Balance     int64
	Hold        int64
}

func (p *PortfolioBalanceSnapshot) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := p.RangeCheck(p.SbeSchemaVersion(), p.SbeSchemaVersion()); err != nil {
			return err
		}
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
	return nil
}

func (p *PortfolioBalanceSnapshot) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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

func (p *PortfolioBalanceSnapshot) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	return nil
}

func PortfolioBalanceSnapshotInit(p *PortfolioBalanceSnapshot) {
	return
}

func (*PortfolioBalanceSnapshot) SbeBlockLength() (blockLength uint16) {
	return 32
}

func (*PortfolioBalanceSnapshot) SbeTemplateId() (templateId uint16) {
	return 1705
}

func (*PortfolioBalanceSnapshot) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*PortfolioBalanceSnapshot) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*PortfolioBalanceSnapshot) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*PortfolioBalanceSnapshot) PortfolioIdId() uint16 {
	return 1
}

func (*PortfolioBalanceSnapshot) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceSnapshot) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PortfolioIdSinceVersion()
}

func (*PortfolioBalanceSnapshot) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceSnapshot) PortfolioIdMetaAttribute(meta int) string {
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

func (*PortfolioBalanceSnapshot) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioBalanceSnapshot) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioBalanceSnapshot) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioBalanceSnapshot) AssetIdId() uint16 {
	return 2
}

func (*PortfolioBalanceSnapshot) AssetIdSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceSnapshot) AssetIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.AssetIdSinceVersion()
}

func (*PortfolioBalanceSnapshot) AssetIdDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceSnapshot) AssetIdMetaAttribute(meta int) string {
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

func (*PortfolioBalanceSnapshot) AssetIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioBalanceSnapshot) AssetIdMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioBalanceSnapshot) AssetIdNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioBalanceSnapshot) BalanceId() uint16 {
	return 3
}

func (*PortfolioBalanceSnapshot) BalanceSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceSnapshot) BalanceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.BalanceSinceVersion()
}

func (*PortfolioBalanceSnapshot) BalanceDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceSnapshot) BalanceMetaAttribute(meta int) string {
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

func (*PortfolioBalanceSnapshot) BalanceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioBalanceSnapshot) BalanceMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioBalanceSnapshot) BalanceNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioBalanceSnapshot) HoldId() uint16 {
	return 4
}

func (*PortfolioBalanceSnapshot) HoldSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceSnapshot) HoldInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.HoldSinceVersion()
}

func (*PortfolioBalanceSnapshot) HoldDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceSnapshot) HoldMetaAttribute(meta int) string {
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

func (*PortfolioBalanceSnapshot) HoldMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioBalanceSnapshot) HoldMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioBalanceSnapshot) HoldNullValue() int64 {
	return math.MinInt64
}
