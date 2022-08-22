// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type PortfolioPermissionState struct {
	UserId      int64
	PortfolioId int64
	Permission  PortfolioPermissionTypeEnum
}

func (p *PortfolioPermissionState) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := p.RangeCheck(p.SbeSchemaVersion(), p.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, p.UserId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, p.PortfolioId); err != nil {
		return err
	}
	if err := p.Permission.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (p *PortfolioPermissionState) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !p.UserIdInActingVersion(actingVersion) {
		p.UserId = p.UserIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &p.UserId); err != nil {
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
	if p.PermissionInActingVersion(actingVersion) {
		if err := p.Permission.Decode(_m, _r, actingVersion); err != nil {
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

func (p *PortfolioPermissionState) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.UserIdInActingVersion(actingVersion) {
		if p.UserId < p.UserIdMinValue() || p.UserId > p.UserIdMaxValue() {
			return fmt.Errorf("Range check failed on p.UserId (%v < %v > %v)", p.UserIdMinValue(), p.UserId, p.UserIdMaxValue())
		}
	}
	if p.PortfolioIdInActingVersion(actingVersion) {
		if p.PortfolioId < p.PortfolioIdMinValue() || p.PortfolioId > p.PortfolioIdMaxValue() {
			return fmt.Errorf("Range check failed on p.PortfolioId (%v < %v > %v)", p.PortfolioIdMinValue(), p.PortfolioId, p.PortfolioIdMaxValue())
		}
	}
	if err := p.Permission.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func PortfolioPermissionStateInit(p *PortfolioPermissionState) {
	return
}

func (*PortfolioPermissionState) SbeBlockLength() (blockLength uint16) {
	return 17
}

func (*PortfolioPermissionState) SbeTemplateId() (templateId uint16) {
	return 15
}

func (*PortfolioPermissionState) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*PortfolioPermissionState) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*PortfolioPermissionState) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*PortfolioPermissionState) UserIdId() uint16 {
	return 1
}

func (*PortfolioPermissionState) UserIdSinceVersion() uint16 {
	return 0
}

func (p *PortfolioPermissionState) UserIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.UserIdSinceVersion()
}

func (*PortfolioPermissionState) UserIdDeprecated() uint16 {
	return 0
}

func (*PortfolioPermissionState) UserIdMetaAttribute(meta int) string {
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

func (*PortfolioPermissionState) UserIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioPermissionState) UserIdMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioPermissionState) UserIdNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioPermissionState) PortfolioIdId() uint16 {
	return 2
}

func (*PortfolioPermissionState) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (p *PortfolioPermissionState) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PortfolioIdSinceVersion()
}

func (*PortfolioPermissionState) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*PortfolioPermissionState) PortfolioIdMetaAttribute(meta int) string {
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

func (*PortfolioPermissionState) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PortfolioPermissionState) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*PortfolioPermissionState) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*PortfolioPermissionState) PermissionId() uint16 {
	return 3
}

func (*PortfolioPermissionState) PermissionSinceVersion() uint16 {
	return 0
}

func (p *PortfolioPermissionState) PermissionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PermissionSinceVersion()
}

func (*PortfolioPermissionState) PermissionDeprecated() uint16 {
	return 0
}

func (*PortfolioPermissionState) PermissionMetaAttribute(meta int) string {
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
