// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type PortfolioState struct {
	PortfolioUuid [2]int64
}

func (p *PortfolioState) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := p.RangeCheck(p.SbeSchemaVersion(), p.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, p.PortfolioUuid[idx]); err != nil {
			return err
		}
	}
	return nil
}

func (p *PortfolioState) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !p.PortfolioUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			p.PortfolioUuid[idx] = p.PortfolioUuidNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &p.PortfolioUuid[idx]); err != nil {
				return err
			}
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

func (p *PortfolioState) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.PortfolioUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if p.PortfolioUuid[idx] < p.PortfolioUuidMinValue() || p.PortfolioUuid[idx] > p.PortfolioUuidMaxValue() {
				return fmt.Errorf("Range check failed on p.PortfolioUuid[%d] (%v < %v > %v)", idx, p.PortfolioUuidMinValue(), p.PortfolioUuid[idx], p.PortfolioUuidMaxValue())
			}
		}
	}
	return nil
}

func PortfolioStateInit(p *PortfolioState) {
	return
}

func (*PortfolioState) SbeBlockLength() (blockLength uint16) {
	return 16
}

func (*PortfolioState) SbeTemplateId() (templateId uint16) {
	return 12
}

func (*PortfolioState) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*PortfolioState) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*PortfolioState) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*PortfolioState) PortfolioUuidId() uint16 {
	return 1
}

func (*PortfolioState) PortfolioUuidSinceVersion() uint16 {
	return 0
}

func (p *PortfolioState) PortfolioUuidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PortfolioUuidSinceVersion()
}

func (*PortfolioState) PortfolioUuidDeprecated() uint16 {
	return 0
}

func (*PortfolioState) PortfolioUuidMetaAttribute(meta int) string {
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

func (*PortfolioState) PortfolioUuidMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioState) PortfolioUuidMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioState) PortfolioUuidNullValue() int64 {
	return math.MinInt64
}
