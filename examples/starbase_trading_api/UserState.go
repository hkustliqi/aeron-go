// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type UserState struct {
	UserUuid    [2]int64
	PortfolioId int64
}

func (u *UserState) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := u.RangeCheck(u.SbeSchemaVersion(), u.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, u.UserUuid[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, u.PortfolioId); err != nil {
		return err
	}
	return nil
}

func (u *UserState) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !u.UserUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			u.UserUuid[idx] = u.UserUuidNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &u.UserUuid[idx]); err != nil {
				return err
			}
		}
	}
	if !u.PortfolioIdInActingVersion(actingVersion) {
		u.PortfolioId = u.PortfolioIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &u.PortfolioId); err != nil {
			return err
		}
	}
	if actingVersion > u.SbeSchemaVersion() && blockLength > u.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-u.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := u.RangeCheck(actingVersion, u.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (u *UserState) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if u.UserUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if u.UserUuid[idx] < u.UserUuidMinValue() || u.UserUuid[idx] > u.UserUuidMaxValue() {
				return fmt.Errorf("Range check failed on u.UserUuid[%d] (%v < %v > %v)", idx, u.UserUuidMinValue(), u.UserUuid[idx], u.UserUuidMaxValue())
			}
		}
	}
	if u.PortfolioIdInActingVersion(actingVersion) {
		if u.PortfolioId < u.PortfolioIdMinValue() || u.PortfolioId > u.PortfolioIdMaxValue() {
			return fmt.Errorf("Range check failed on u.PortfolioId (%v < %v > %v)", u.PortfolioIdMinValue(), u.PortfolioId, u.PortfolioIdMaxValue())
		}
	}
	return nil
}

func UserStateInit(u *UserState) {
	return
}

func (*UserState) SbeBlockLength() (blockLength uint16) {
	return 24
}

func (*UserState) SbeTemplateId() (templateId uint16) {
	return 14
}

func (*UserState) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*UserState) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*UserState) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*UserState) UserUuidId() uint16 {
	return 1
}

func (*UserState) UserUuidSinceVersion() uint16 {
	return 0
}

func (u *UserState) UserUuidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= u.UserUuidSinceVersion()
}

func (*UserState) UserUuidDeprecated() uint16 {
	return 0
}

func (*UserState) UserUuidMetaAttribute(meta int) string {
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

func (*UserState) UserUuidMinValue() int64 {
	return math.MinInt64 + 1
}

func (*UserState) UserUuidMaxValue() int64 {
	return math.MaxInt64
}

func (*UserState) UserUuidNullValue() int64 {
	return math.MinInt64
}

func (*UserState) PortfolioIdId() uint16 {
	return 2
}

func (*UserState) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (u *UserState) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= u.PortfolioIdSinceVersion()
}

func (*UserState) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*UserState) PortfolioIdMetaAttribute(meta int) string {
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

func (*UserState) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*UserState) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*UserState) PortfolioIdNullValue() int64 {
	return math.MinInt64
}
