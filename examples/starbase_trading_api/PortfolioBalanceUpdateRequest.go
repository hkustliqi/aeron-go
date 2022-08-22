// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type PortfolioBalanceUpdateRequest struct {
	RequestHeader  OmsAdminRequestHeader
	PortfolioId    int64
	AssetId        int64
	Quantity       int64
	SequenceNumber int64
	Action         PortfolioBalanceUpdateActionEnum
}

func (p *PortfolioBalanceUpdateRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := p.RangeCheck(p.SbeSchemaVersion(), p.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := p.RequestHeader.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, p.PortfolioId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, p.AssetId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, p.Quantity); err != nil {
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

func (p *PortfolioBalanceUpdateRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if p.RequestHeaderInActingVersion(actingVersion) {
		if err := p.RequestHeader.Decode(_m, _r, actingVersion); err != nil {
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
	if !p.QuantityInActingVersion(actingVersion) {
		p.Quantity = p.QuantityNullValue()
	} else {
		if err := _m.ReadInt64(_r, &p.Quantity); err != nil {
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

func (p *PortfolioBalanceUpdateRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if p.QuantityInActingVersion(actingVersion) {
		if p.Quantity < p.QuantityMinValue() || p.Quantity > p.QuantityMaxValue() {
			return fmt.Errorf("Range check failed on p.Quantity (%v < %v > %v)", p.QuantityMinValue(), p.Quantity, p.QuantityMaxValue())
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

func PortfolioBalanceUpdateRequestInit(p *PortfolioBalanceUpdateRequest) {
	return
}

func (*PortfolioBalanceUpdateRequest) SbeBlockLength() (blockLength uint16) {
	return 57
}

func (*PortfolioBalanceUpdateRequest) SbeTemplateId() (templateId uint16) {
	return 1210
}

func (*PortfolioBalanceUpdateRequest) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*PortfolioBalanceUpdateRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*PortfolioBalanceUpdateRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*PortfolioBalanceUpdateRequest) RequestHeaderId() uint16 {
	return 1
}

func (*PortfolioBalanceUpdateRequest) RequestHeaderSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateRequest) RequestHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.RequestHeaderSinceVersion()
}

func (*PortfolioBalanceUpdateRequest) RequestHeaderDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateRequest) RequestHeaderMetaAttribute(meta int) string {
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

func (*PortfolioBalanceUpdateRequest) PortfolioIdId() uint16 {
	return 2
}

func (*PortfolioBalanceUpdateRequest) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateRequest) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PortfolioIdSinceVersion()
}

func (*PortfolioBalanceUpdateRequest) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateRequest) PortfolioIdMetaAttribute(meta int) string {
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

func (*PortfolioBalanceUpdateRequest) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioBalanceUpdateRequest) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioBalanceUpdateRequest) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioBalanceUpdateRequest) AssetIdId() uint16 {
	return 3
}

func (*PortfolioBalanceUpdateRequest) AssetIdSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateRequest) AssetIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.AssetIdSinceVersion()
}

func (*PortfolioBalanceUpdateRequest) AssetIdDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateRequest) AssetIdMetaAttribute(meta int) string {
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

func (*PortfolioBalanceUpdateRequest) AssetIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioBalanceUpdateRequest) AssetIdMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioBalanceUpdateRequest) AssetIdNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioBalanceUpdateRequest) QuantityId() uint16 {
	return 4
}

func (*PortfolioBalanceUpdateRequest) QuantitySinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateRequest) QuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.QuantitySinceVersion()
}

func (*PortfolioBalanceUpdateRequest) QuantityDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateRequest) QuantityMetaAttribute(meta int) string {
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

func (*PortfolioBalanceUpdateRequest) QuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioBalanceUpdateRequest) QuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioBalanceUpdateRequest) QuantityNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioBalanceUpdateRequest) SequenceNumberId() uint16 {
	return 5
}

func (*PortfolioBalanceUpdateRequest) SequenceNumberSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateRequest) SequenceNumberInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SequenceNumberSinceVersion()
}

func (*PortfolioBalanceUpdateRequest) SequenceNumberDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateRequest) SequenceNumberMetaAttribute(meta int) string {
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

func (*PortfolioBalanceUpdateRequest) SequenceNumberMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioBalanceUpdateRequest) SequenceNumberMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioBalanceUpdateRequest) SequenceNumberNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioBalanceUpdateRequest) ActionId() uint16 {
	return 6
}

func (*PortfolioBalanceUpdateRequest) ActionSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateRequest) ActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ActionSinceVersion()
}

func (*PortfolioBalanceUpdateRequest) ActionDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateRequest) ActionMetaAttribute(meta int) string {
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
