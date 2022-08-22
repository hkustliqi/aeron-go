// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsCancelOrder struct {
	RequestHeader    OmsUserRequestHeader
	PortfolioId      int64
	ClientOrderId    [2]int64
	NewClientOrderId [2]int64
}

func (o *OmsCancelOrder) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := o.RequestHeader.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.PortfolioId); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, o.ClientOrderId[idx]); err != nil {
			return err
		}
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, o.NewClientOrderId[idx]); err != nil {
			return err
		}
	}
	return nil
}

func (o *OmsCancelOrder) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if o.RequestHeaderInActingVersion(actingVersion) {
		if err := o.RequestHeader.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.PortfolioIdInActingVersion(actingVersion) {
		o.PortfolioId = o.PortfolioIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.PortfolioId); err != nil {
			return err
		}
	}
	if !o.ClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			o.ClientOrderId[idx] = o.ClientOrderIdNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &o.ClientOrderId[idx]); err != nil {
				return err
			}
		}
	}
	if !o.NewClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			o.NewClientOrderId[idx] = o.NewClientOrderIdNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &o.NewClientOrderId[idx]); err != nil {
				return err
			}
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

func (o *OmsCancelOrder) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.PortfolioIdInActingVersion(actingVersion) {
		if o.PortfolioId < o.PortfolioIdMinValue() || o.PortfolioId > o.PortfolioIdMaxValue() {
			return fmt.Errorf("Range check failed on o.PortfolioId (%v < %v > %v)", o.PortfolioIdMinValue(), o.PortfolioId, o.PortfolioIdMaxValue())
		}
	}
	if o.ClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if o.ClientOrderId[idx] < o.ClientOrderIdMinValue() || o.ClientOrderId[idx] > o.ClientOrderIdMaxValue() {
				return fmt.Errorf("Range check failed on o.ClientOrderId[%d] (%v < %v > %v)", idx, o.ClientOrderIdMinValue(), o.ClientOrderId[idx], o.ClientOrderIdMaxValue())
			}
		}
	}
	if o.NewClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if o.NewClientOrderId[idx] < o.NewClientOrderIdMinValue() || o.NewClientOrderId[idx] > o.NewClientOrderIdMaxValue() {
				return fmt.Errorf("Range check failed on o.NewClientOrderId[%d] (%v < %v > %v)", idx, o.NewClientOrderIdMinValue(), o.NewClientOrderId[idx], o.NewClientOrderIdMaxValue())
			}
		}
	}
	return nil
}

func OmsCancelOrderInit(o *OmsCancelOrder) {
	return
}

func (*OmsCancelOrder) SbeBlockLength() (blockLength uint16) {
	return 72
}

func (*OmsCancelOrder) SbeTemplateId() (templateId uint16) {
	return 1001
}

func (*OmsCancelOrder) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsCancelOrder) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsCancelOrder) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsCancelOrder) RequestHeaderId() uint16 {
	return 1
}

func (*OmsCancelOrder) RequestHeaderSinceVersion() uint16 {
	return 0
}

func (o *OmsCancelOrder) RequestHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RequestHeaderSinceVersion()
}

func (*OmsCancelOrder) RequestHeaderDeprecated() uint16 {
	return 0
}

func (*OmsCancelOrder) RequestHeaderMetaAttribute(meta int) string {
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

func (*OmsCancelOrder) PortfolioIdId() uint16 {
	return 2
}

func (*OmsCancelOrder) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (o *OmsCancelOrder) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PortfolioIdSinceVersion()
}

func (*OmsCancelOrder) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*OmsCancelOrder) PortfolioIdMetaAttribute(meta int) string {
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

func (*OmsCancelOrder) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCancelOrder) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCancelOrder) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsCancelOrder) ClientOrderIdId() uint16 {
	return 3
}

func (*OmsCancelOrder) ClientOrderIdSinceVersion() uint16 {
	return 0
}

func (o *OmsCancelOrder) ClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClientOrderIdSinceVersion()
}

func (*OmsCancelOrder) ClientOrderIdDeprecated() uint16 {
	return 0
}

func (*OmsCancelOrder) ClientOrderIdMetaAttribute(meta int) string {
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

func (*OmsCancelOrder) ClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCancelOrder) ClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCancelOrder) ClientOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsCancelOrder) NewClientOrderIdId() uint16 {
	return 4
}

func (*OmsCancelOrder) NewClientOrderIdSinceVersion() uint16 {
	return 0
}

func (o *OmsCancelOrder) NewClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NewClientOrderIdSinceVersion()
}

func (*OmsCancelOrder) NewClientOrderIdDeprecated() uint16 {
	return 0
}

func (*OmsCancelOrder) NewClientOrderIdMetaAttribute(meta int) string {
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

func (*OmsCancelOrder) NewClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCancelOrder) NewClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCancelOrder) NewClientOrderIdNullValue() int64 {
	return math.MinInt64
}
