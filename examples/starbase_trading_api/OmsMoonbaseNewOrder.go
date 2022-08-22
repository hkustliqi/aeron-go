// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsMoonbaseNewOrder struct {
	GatewayLogPosition      int64
	RequestTime             int64
	CorrelationId           int64
	UserName                [80]byte
	InstrumentSymbol        [80]byte
	ClientOrderId           [2]int64
	PortfolioName           [80]byte
	BrokerageClientId       [2]int64
	LimitPrice              [80]byte
	StopPrice               [80]byte
	Quantity                [80]byte
	Funds                   [80]byte
	OrderFlags              OrderFlags
	ExpireTime              int32
	SelfMatchPreventionMode SelfMatchPreventionModeEnum
	Side                    SideEnum
	TimeInForce             TimeInForceEnum
}

func (o *OmsMoonbaseNewOrder) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.GatewayLogPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.RequestTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.CorrelationId); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.UserName[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.InstrumentSymbol[:]); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, o.ClientOrderId[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, o.PortfolioName[:]); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, o.BrokerageClientId[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, o.LimitPrice[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.StopPrice[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Quantity[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Funds[:]); err != nil {
		return err
	}
	if err := o.OrderFlags.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.ExpireTime); err != nil {
		return err
	}
	if err := o.SelfMatchPreventionMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (o *OmsMoonbaseNewOrder) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.GatewayLogPositionInActingVersion(actingVersion) {
		o.GatewayLogPosition = o.GatewayLogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.GatewayLogPosition); err != nil {
			return err
		}
	}
	if !o.RequestTimeInActingVersion(actingVersion) {
		o.RequestTime = o.RequestTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.RequestTime); err != nil {
			return err
		}
	}
	if !o.CorrelationIdInActingVersion(actingVersion) {
		o.CorrelationId = o.CorrelationIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.CorrelationId); err != nil {
			return err
		}
	}
	if !o.UserNameInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			o.UserName[idx] = o.UserNameNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.UserName[:]); err != nil {
			return err
		}
	}
	if !o.InstrumentSymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			o.InstrumentSymbol[idx] = o.InstrumentSymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.InstrumentSymbol[:]); err != nil {
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
	if !o.PortfolioNameInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			o.PortfolioName[idx] = o.PortfolioNameNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.PortfolioName[:]); err != nil {
			return err
		}
	}
	if !o.BrokerageClientIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			o.BrokerageClientId[idx] = o.BrokerageClientIdNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &o.BrokerageClientId[idx]); err != nil {
				return err
			}
		}
	}
	if !o.LimitPriceInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			o.LimitPrice[idx] = o.LimitPriceNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.LimitPrice[:]); err != nil {
			return err
		}
	}
	if !o.StopPriceInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			o.StopPrice[idx] = o.StopPriceNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.StopPrice[:]); err != nil {
			return err
		}
	}
	if !o.QuantityInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			o.Quantity[idx] = o.QuantityNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Quantity[:]); err != nil {
			return err
		}
	}
	if !o.FundsInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			o.Funds[idx] = o.FundsNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Funds[:]); err != nil {
			return err
		}
	}
	if o.OrderFlagsInActingVersion(actingVersion) {
		if err := o.OrderFlags.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.ExpireTimeInActingVersion(actingVersion) {
		o.ExpireTime = o.ExpireTimeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &o.ExpireTime); err != nil {
			return err
		}
	}
	if o.SelfMatchPreventionModeInActingVersion(actingVersion) {
		if err := o.SelfMatchPreventionMode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.SideInActingVersion(actingVersion) {
		if err := o.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.TimeInForceInActingVersion(actingVersion) {
		if err := o.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
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

func (o *OmsMoonbaseNewOrder) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.GatewayLogPositionInActingVersion(actingVersion) {
		if o.GatewayLogPosition < o.GatewayLogPositionMinValue() || o.GatewayLogPosition > o.GatewayLogPositionMaxValue() {
			return fmt.Errorf("Range check failed on o.GatewayLogPosition (%v < %v > %v)", o.GatewayLogPositionMinValue(), o.GatewayLogPosition, o.GatewayLogPositionMaxValue())
		}
	}
	if o.RequestTimeInActingVersion(actingVersion) {
		if o.RequestTime < o.RequestTimeMinValue() || o.RequestTime > o.RequestTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.RequestTime (%v < %v > %v)", o.RequestTimeMinValue(), o.RequestTime, o.RequestTimeMaxValue())
		}
	}
	if o.CorrelationIdInActingVersion(actingVersion) {
		if o.CorrelationId != o.CorrelationIdNullValue() && (o.CorrelationId < o.CorrelationIdMinValue() || o.CorrelationId > o.CorrelationIdMaxValue()) {
			return fmt.Errorf("Range check failed on o.CorrelationId (%v < %v > %v)", o.CorrelationIdMinValue(), o.CorrelationId, o.CorrelationIdMaxValue())
		}
	}
	if o.UserNameInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if o.UserName[idx] < o.UserNameMinValue() || o.UserName[idx] > o.UserNameMaxValue() {
				return fmt.Errorf("Range check failed on o.UserName[%d] (%v < %v > %v)", idx, o.UserNameMinValue(), o.UserName[idx], o.UserNameMaxValue())
			}
		}
	}
	for idx, ch := range o.UserName {
		if ch > 127 {
			return fmt.Errorf("o.UserName[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if o.InstrumentSymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if o.InstrumentSymbol[idx] < o.InstrumentSymbolMinValue() || o.InstrumentSymbol[idx] > o.InstrumentSymbolMaxValue() {
				return fmt.Errorf("Range check failed on o.InstrumentSymbol[%d] (%v < %v > %v)", idx, o.InstrumentSymbolMinValue(), o.InstrumentSymbol[idx], o.InstrumentSymbolMaxValue())
			}
		}
	}
	for idx, ch := range o.InstrumentSymbol {
		if ch > 127 {
			return fmt.Errorf("o.InstrumentSymbol[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if o.ClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if o.ClientOrderId[idx] < o.ClientOrderIdMinValue() || o.ClientOrderId[idx] > o.ClientOrderIdMaxValue() {
				return fmt.Errorf("Range check failed on o.ClientOrderId[%d] (%v < %v > %v)", idx, o.ClientOrderIdMinValue(), o.ClientOrderId[idx], o.ClientOrderIdMaxValue())
			}
		}
	}
	if o.PortfolioNameInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if o.PortfolioName[idx] < o.PortfolioNameMinValue() || o.PortfolioName[idx] > o.PortfolioNameMaxValue() {
				return fmt.Errorf("Range check failed on o.PortfolioName[%d] (%v < %v > %v)", idx, o.PortfolioNameMinValue(), o.PortfolioName[idx], o.PortfolioNameMaxValue())
			}
		}
	}
	for idx, ch := range o.PortfolioName {
		if ch > 127 {
			return fmt.Errorf("o.PortfolioName[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if o.BrokerageClientIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if o.BrokerageClientId[idx] < o.BrokerageClientIdMinValue() || o.BrokerageClientId[idx] > o.BrokerageClientIdMaxValue() {
				return fmt.Errorf("Range check failed on o.BrokerageClientId[%d] (%v < %v > %v)", idx, o.BrokerageClientIdMinValue(), o.BrokerageClientId[idx], o.BrokerageClientIdMaxValue())
			}
		}
	}
	if o.LimitPriceInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if o.LimitPrice[idx] < o.LimitPriceMinValue() || o.LimitPrice[idx] > o.LimitPriceMaxValue() {
				return fmt.Errorf("Range check failed on o.LimitPrice[%d] (%v < %v > %v)", idx, o.LimitPriceMinValue(), o.LimitPrice[idx], o.LimitPriceMaxValue())
			}
		}
	}
	for idx, ch := range o.LimitPrice {
		if ch > 127 {
			return fmt.Errorf("o.LimitPrice[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if o.StopPriceInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if o.StopPrice[idx] < o.StopPriceMinValue() || o.StopPrice[idx] > o.StopPriceMaxValue() {
				return fmt.Errorf("Range check failed on o.StopPrice[%d] (%v < %v > %v)", idx, o.StopPriceMinValue(), o.StopPrice[idx], o.StopPriceMaxValue())
			}
		}
	}
	for idx, ch := range o.StopPrice {
		if ch > 127 {
			return fmt.Errorf("o.StopPrice[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if o.QuantityInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if o.Quantity[idx] < o.QuantityMinValue() || o.Quantity[idx] > o.QuantityMaxValue() {
				return fmt.Errorf("Range check failed on o.Quantity[%d] (%v < %v > %v)", idx, o.QuantityMinValue(), o.Quantity[idx], o.QuantityMaxValue())
			}
		}
	}
	for idx, ch := range o.Quantity {
		if ch > 127 {
			return fmt.Errorf("o.Quantity[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if o.FundsInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if o.Funds[idx] < o.FundsMinValue() || o.Funds[idx] > o.FundsMaxValue() {
				return fmt.Errorf("Range check failed on o.Funds[%d] (%v < %v > %v)", idx, o.FundsMinValue(), o.Funds[idx], o.FundsMaxValue())
			}
		}
	}
	for idx, ch := range o.Funds {
		if ch > 127 {
			return fmt.Errorf("o.Funds[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if o.ExpireTimeInActingVersion(actingVersion) {
		if o.ExpireTime != o.ExpireTimeNullValue() && (o.ExpireTime < o.ExpireTimeMinValue() || o.ExpireTime > o.ExpireTimeMaxValue()) {
			return fmt.Errorf("Range check failed on o.ExpireTime (%v < %v > %v)", o.ExpireTimeMinValue(), o.ExpireTime, o.ExpireTimeMaxValue())
		}
	}
	if err := o.SelfMatchPreventionMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func OmsMoonbaseNewOrderInit(o *OmsMoonbaseNewOrder) {
	o.CorrelationId = math.MinInt64
	for idx := 0; idx < 80; idx++ {
		o.StopPrice[idx] = 0
	}
	for idx := 0; idx < 80; idx++ {
		o.Funds[idx] = 0
	}
	o.ExpireTime = math.MinInt32
	return
}

func (*OmsMoonbaseNewOrder) SbeBlockLength() (blockLength uint16) {
	return 631
}

func (*OmsMoonbaseNewOrder) SbeTemplateId() (templateId uint16) {
	return 1002
}

func (*OmsMoonbaseNewOrder) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsMoonbaseNewOrder) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsMoonbaseNewOrder) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsMoonbaseNewOrder) GatewayLogPositionId() uint16 {
	return 1
}

func (*OmsMoonbaseNewOrder) GatewayLogPositionSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) GatewayLogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.GatewayLogPositionSinceVersion()
}

func (*OmsMoonbaseNewOrder) GatewayLogPositionDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) GatewayLogPositionMetaAttribute(meta int) string {
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

func (*OmsMoonbaseNewOrder) GatewayLogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsMoonbaseNewOrder) GatewayLogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsMoonbaseNewOrder) GatewayLogPositionNullValue() int64 {
	return math.MinInt64
}

func (*OmsMoonbaseNewOrder) RequestTimeId() uint16 {
	return 2
}

func (*OmsMoonbaseNewOrder) RequestTimeSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) RequestTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RequestTimeSinceVersion()
}

func (*OmsMoonbaseNewOrder) RequestTimeDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) RequestTimeMetaAttribute(meta int) string {
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

func (*OmsMoonbaseNewOrder) RequestTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsMoonbaseNewOrder) RequestTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsMoonbaseNewOrder) RequestTimeNullValue() int64 {
	return math.MinInt64
}

func (*OmsMoonbaseNewOrder) CorrelationIdId() uint16 {
	return 3
}

func (*OmsMoonbaseNewOrder) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CorrelationIdSinceVersion()
}

func (*OmsMoonbaseNewOrder) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) CorrelationIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*OmsMoonbaseNewOrder) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsMoonbaseNewOrder) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsMoonbaseNewOrder) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsMoonbaseNewOrder) UserNameId() uint16 {
	return 4
}

func (*OmsMoonbaseNewOrder) UserNameSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) UserNameInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.UserNameSinceVersion()
}

func (*OmsMoonbaseNewOrder) UserNameDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) UserNameMetaAttribute(meta int) string {
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

func (*OmsMoonbaseNewOrder) UserNameMinValue() byte {
	return byte(32)
}

func (*OmsMoonbaseNewOrder) UserNameMaxValue() byte {
	return byte(126)
}

func (*OmsMoonbaseNewOrder) UserNameNullValue() byte {
	return 0
}

func (o *OmsMoonbaseNewOrder) UserNameCharacterEncoding() string {
	return "US-ASCII"
}

func (*OmsMoonbaseNewOrder) InstrumentSymbolId() uint16 {
	return 5
}

func (*OmsMoonbaseNewOrder) InstrumentSymbolSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) InstrumentSymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.InstrumentSymbolSinceVersion()
}

func (*OmsMoonbaseNewOrder) InstrumentSymbolDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) InstrumentSymbolMetaAttribute(meta int) string {
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

func (*OmsMoonbaseNewOrder) InstrumentSymbolMinValue() byte {
	return byte(32)
}

func (*OmsMoonbaseNewOrder) InstrumentSymbolMaxValue() byte {
	return byte(126)
}

func (*OmsMoonbaseNewOrder) InstrumentSymbolNullValue() byte {
	return 0
}

func (o *OmsMoonbaseNewOrder) InstrumentSymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*OmsMoonbaseNewOrder) ClientOrderIdId() uint16 {
	return 6
}

func (*OmsMoonbaseNewOrder) ClientOrderIdSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) ClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClientOrderIdSinceVersion()
}

func (*OmsMoonbaseNewOrder) ClientOrderIdDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) ClientOrderIdMetaAttribute(meta int) string {
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

func (*OmsMoonbaseNewOrder) ClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsMoonbaseNewOrder) ClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsMoonbaseNewOrder) ClientOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsMoonbaseNewOrder) PortfolioNameId() uint16 {
	return 7
}

func (*OmsMoonbaseNewOrder) PortfolioNameSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) PortfolioNameInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PortfolioNameSinceVersion()
}

func (*OmsMoonbaseNewOrder) PortfolioNameDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) PortfolioNameMetaAttribute(meta int) string {
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

func (*OmsMoonbaseNewOrder) PortfolioNameMinValue() byte {
	return byte(32)
}

func (*OmsMoonbaseNewOrder) PortfolioNameMaxValue() byte {
	return byte(126)
}

func (*OmsMoonbaseNewOrder) PortfolioNameNullValue() byte {
	return 0
}

func (o *OmsMoonbaseNewOrder) PortfolioNameCharacterEncoding() string {
	return "US-ASCII"
}

func (*OmsMoonbaseNewOrder) BrokerageClientIdId() uint16 {
	return 8
}

func (*OmsMoonbaseNewOrder) BrokerageClientIdSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) BrokerageClientIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.BrokerageClientIdSinceVersion()
}

func (*OmsMoonbaseNewOrder) BrokerageClientIdDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) BrokerageClientIdMetaAttribute(meta int) string {
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

func (*OmsMoonbaseNewOrder) BrokerageClientIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsMoonbaseNewOrder) BrokerageClientIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsMoonbaseNewOrder) BrokerageClientIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsMoonbaseNewOrder) LimitPriceId() uint16 {
	return 9
}

func (*OmsMoonbaseNewOrder) LimitPriceSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) LimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LimitPriceSinceVersion()
}

func (*OmsMoonbaseNewOrder) LimitPriceDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) LimitPriceMetaAttribute(meta int) string {
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

func (*OmsMoonbaseNewOrder) LimitPriceMinValue() byte {
	return byte(32)
}

func (*OmsMoonbaseNewOrder) LimitPriceMaxValue() byte {
	return byte(126)
}

func (*OmsMoonbaseNewOrder) LimitPriceNullValue() byte {
	return 0
}

func (o *OmsMoonbaseNewOrder) LimitPriceCharacterEncoding() string {
	return "US-ASCII"
}

func (*OmsMoonbaseNewOrder) StopPriceId() uint16 {
	return 10
}

func (*OmsMoonbaseNewOrder) StopPriceSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) StopPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopPriceSinceVersion()
}

func (*OmsMoonbaseNewOrder) StopPriceDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) StopPriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*OmsMoonbaseNewOrder) StopPriceMinValue() byte {
	return byte(32)
}

func (*OmsMoonbaseNewOrder) StopPriceMaxValue() byte {
	return byte(126)
}

func (*OmsMoonbaseNewOrder) StopPriceNullValue() byte {
	return 0
}

func (o *OmsMoonbaseNewOrder) StopPriceCharacterEncoding() string {
	return "US-ASCII"
}

func (*OmsMoonbaseNewOrder) QuantityId() uint16 {
	return 11
}

func (*OmsMoonbaseNewOrder) QuantitySinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) QuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.QuantitySinceVersion()
}

func (*OmsMoonbaseNewOrder) QuantityDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) QuantityMetaAttribute(meta int) string {
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

func (*OmsMoonbaseNewOrder) QuantityMinValue() byte {
	return byte(32)
}

func (*OmsMoonbaseNewOrder) QuantityMaxValue() byte {
	return byte(126)
}

func (*OmsMoonbaseNewOrder) QuantityNullValue() byte {
	return 0
}

func (o *OmsMoonbaseNewOrder) QuantityCharacterEncoding() string {
	return "US-ASCII"
}

func (*OmsMoonbaseNewOrder) FundsId() uint16 {
	return 12
}

func (*OmsMoonbaseNewOrder) FundsSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) FundsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.FundsSinceVersion()
}

func (*OmsMoonbaseNewOrder) FundsDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) FundsMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*OmsMoonbaseNewOrder) FundsMinValue() byte {
	return byte(32)
}

func (*OmsMoonbaseNewOrder) FundsMaxValue() byte {
	return byte(126)
}

func (*OmsMoonbaseNewOrder) FundsNullValue() byte {
	return 0
}

func (o *OmsMoonbaseNewOrder) FundsCharacterEncoding() string {
	return "US-ASCII"
}

func (*OmsMoonbaseNewOrder) OrderFlagsId() uint16 {
	return 13
}

func (*OmsMoonbaseNewOrder) OrderFlagsSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) OrderFlagsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderFlagsSinceVersion()
}

func (*OmsMoonbaseNewOrder) OrderFlagsDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) OrderFlagsMetaAttribute(meta int) string {
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

func (*OmsMoonbaseNewOrder) ExpireTimeId() uint16 {
	return 14
}

func (*OmsMoonbaseNewOrder) ExpireTimeSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) ExpireTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExpireTimeSinceVersion()
}

func (*OmsMoonbaseNewOrder) ExpireTimeDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) ExpireTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*OmsMoonbaseNewOrder) ExpireTimeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OmsMoonbaseNewOrder) ExpireTimeMaxValue() int32 {
	return math.MaxInt32
}

func (*OmsMoonbaseNewOrder) ExpireTimeNullValue() int32 {
	return math.MinInt32
}

func (*OmsMoonbaseNewOrder) SelfMatchPreventionModeId() uint16 {
	return 15
}

func (*OmsMoonbaseNewOrder) SelfMatchPreventionModeSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) SelfMatchPreventionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SelfMatchPreventionModeSinceVersion()
}

func (*OmsMoonbaseNewOrder) SelfMatchPreventionModeDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) SelfMatchPreventionModeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*OmsMoonbaseNewOrder) SideId() uint16 {
	return 16
}

func (*OmsMoonbaseNewOrder) SideSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OmsMoonbaseNewOrder) SideDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) SideMetaAttribute(meta int) string {
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

func (*OmsMoonbaseNewOrder) TimeInForceId() uint16 {
	return 17
}

func (*OmsMoonbaseNewOrder) TimeInForceSinceVersion() uint16 {
	return 0
}

func (o *OmsMoonbaseNewOrder) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeInForceSinceVersion()
}

func (*OmsMoonbaseNewOrder) TimeInForceDeprecated() uint16 {
	return 0
}

func (*OmsMoonbaseNewOrder) TimeInForceMetaAttribute(meta int) string {
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
