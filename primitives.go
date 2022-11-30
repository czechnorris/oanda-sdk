package oanda_sdk

import "github.com/shopspring/decimal"

// Currency is a string containing an ISO 4217 currency (http://en.wikipedia.org/wiki/ISO_4217)
type Currency string

// Tag associated with an entity
type Tag struct {
	// The type of the tag.
	Type string `json:"type"`

	// The name of the tag.
	Name string `json:"name"`
}

// InstrumentType represents the type of Instrument
type InstrumentType string

const (
	InstrumentTypeCurrency = InstrumentType("CURRENCY")
	InstrumentTypeCFD      = InstrumentType("CFD")
	InstrumentTypeMetal    = InstrumentType("METAL")
)

// DayOfWeek provides a representation of the day of the week.
type DayOfWeek string

const (
	DayOfWeekSunday    = DayOfWeek("SUNDAY")
	DayOfWeekMonday    = DayOfWeek("MONDAY")
	DayOfWeekTuesday   = DayOfWeek("TUESDAY")
	DayOfWeekWednesday = DayOfWeek("WEDNESDAY")
	DayOfWeekThursday  = DayOfWeek("THURSDAY")
	DayOfWeekFriday    = DayOfWeek("FRIDAY")
	DayOfWeekSaturday  = DayOfWeek("SATURDAY")
)

// FinancingDayOfWeek message defines a day of the week when financing changes are debited or credited
type FinancingDayOfWeek struct {
	// The day of the week to charge the financing.
	DayOfWeek DayOfWeek `json:"dayOfWeek"`

	// The number of days worth of financing to be charged on dayOfWeek.
	DaysCharged int `json:"daysCharged"`
}

// InstrumentFinancing represents the financing data for the instrument
type InstrumentFinancing struct {
	// The financing rate to be used for a long position for the instrument. The
	// value is in decimal rather than percentage points, i.e. 5% is represented
	// as 0.05.
	LongRate decimal.Decimal `json:"longRate"`

	// The financing rate to be used for a short position for the instrument.
	// The value is in decimal rather than percentage points, i.e. 5% is
	// represented as 0.05.
	ShortRate decimal.Decimal `json:"shortRate"`

	// The days of the week to debit or credit financing charges; the exact time
	// of day at which to charge the financing is set in the
	// DivisionTradingGroup for the client’s account.
	FinancingDaysOfWeek []FinancingDayOfWeek `json:"financingDaysOfWeek"`
}

// Instrument represents a full specification of an instrument
type Instrument struct {
	// The name of the Instrument
	Name string `json:"name"`

	// The type of the Instrument
	Type InstrumentType `json:"type"`

	// The display name of the Instrument
	DisplayName string `json:"displayName"`

	// The location of the “pip” for this instrument. The decimal position of
	// the pip in this Instrument’s price can be found at 10 ^ pipLocation (e.g.
	// -4 pipLocation results in a decimal pip position of 10 ^ -4 = 0.0001).
	PipLocation int `json:"pipLocation"`

	// The number of decimal places that should be used to display prices for
	// this instrument. (e.g. a displayPrecision of 5 would result in a price of
	// “1” being displayed as “1.00000”)
	DisplayPrecision int `json:"displayPrecision"`

	// The amount of decimal places that may be provided when specifying the number of units traded for this instrument.
	TradeUnitsPrecision int `json:"tradeUnitsPrecision"`

	// The smallest number of units allowed to be traded for this instrument.
	MinimumTradeSize decimal.Decimal `json:"minimumTradeSize"`

	// The maximum trailing stop distance allowed for a trailing stop loss created for this instrument. Specified in
	// price units.
	MaximumTrailingStopDistance decimal.Decimal `json:"maximumTrailingStopDistance"`

	// The minimum distance allowed between the Trade’s fill price and the configured price for GuaranteedStopLossOrders
	// created for this instrument. Specified in price units.
	MinimumGuaranteedStopLossDistance decimal.Decimal `json:"minimumGuaranteedStopLossDistance"`

	// The minimum trailing stop distance allowed for a trailing stop loss created for this instrument. Specified in
	// price units.
	MinimumTrailingStopDistance decimal.Decimal `json:"minimumTrailingStopDistance"`

	// The maximum position size allowed for this instrument. Specified in units.
	MaximumPositionSize decimal.Decimal `json:"maximumPositionSize"`

	// The maximum units allowed for an Order placed for this instrument. Specified in units.
	MaximumOrderUnits decimal.Decimal `json:"maximumOrderUnits"`

	// The margin rate for this instrument.
	MarginRate decimal.Decimal `json:"marginRate"`

	// The commission structure for this instrument.
	Commission InstrumentCommission `json:"commission"`

	// The current Guaranteed Stop Loss Order mode of the Account for this Instrument.
	GuaranteedStopLossOrderMode GuaranteedStopLossOrderModeForInstrument `json:"guaranteedStopLossOrderMode"`

	// The amount that is charged to the account if a GuaranteedStopLossOrder
	// is triggered and filled. The value is in price units and is charged for
	// each unit of the Trade. This field will only be present if the Account’s
	// guaranteedStopLossOrderMode for this Instrument is not ‘DISABLED’.
	GuaranteedStopLosOrderExecutionPremium *decimal.Decimal `json:"guaranteedStopLosOrderExecutionPremium"`

	// The guaranteed Stop Loss Order level restriction for this instrument. This field will only be present if the
	// Account’s guaranteedStopLossOrderMode for this Instrument is not ‘DISABLED’.
	GuaranteedStopLossOrderLevelRestriction *GuaranteedStopLossOrderLevelRestriction `json:"guaranteedStopLossOrderLevelRestriction"`

	// Financing data for this instrument.
	Financing InstrumentFinancing `json:"financing"`

	// The tags associated with this instrument.
	Tags []Tag `json:"tags"`
}

// InstrumentCommission represents an instrument-specific commission
type InstrumentCommission struct {

	// The commission amount (in the Account’s home currency) charged per unitsTraded of the instrument
	Commission decimal.Decimal `json:"commission"`

	// The number of units traded that the commission amount is based on.
	UnitsTraded decimal.Decimal `json:"unitsTraded"`

	// The minimum commission amount (in the Account’s home currency) that is charged when an Order is filled for this
	// instrument.
	MinimumCommission decimal.Decimal `json:"minimumCommission"`
}

// GuaranteedStopLossOrderModeForInstrument represents the overall behaviour of the Account regarding
// [oanda_sdk.GuaranteedStopLossOrder]s for a specific Instrument
type GuaranteedStopLossOrderModeForInstrument string

const (
	// GuaranteedStopLossOrderModeForInstrumentDisabled represents the mode where the Account is not permitted to create
	// [oanda_sdk.GuaranteedStopLossOrder]s for this Instrument.
	GuaranteedStopLossOrderModeForInstrumentDisabled = GuaranteedStopLossOrderModeForInstrument("DISABLED")

	// GuaranteedStopLossOrderModeForInstrumentAllowed represents the mode where the Account is able, but not required
	// to have [oanda_sdk.GuaranteedStopLossOrder]s for open [oanda_sdk.Trade]s for this Instrument.
	GuaranteedStopLossOrderModeForInstrumentAllowed = GuaranteedStopLossOrderModeForInstrument("ALLOWED")

	// GuaranteedStopLossOrderModeForInstrumentRequired represents the mode where the Account is required to have
	// [oanda_sdk.GuaranteedStopLossOrder]s for all open [oanda_sdk.Trade]s for this Instrument.
	GuaranteedStopLossOrderModeForInstrumentRequired = GuaranteedStopLossOrderModeForInstrument("REQUIRED")
)

// GuaranteedStopLossOrderLevelRestriction represents the total position size that can exist within a given price window
// for [oanda_sdk.Trade]s with [oanda_sdk.GuaranteedStopLossOrder]s attached for a specific Instrument
type GuaranteedStopLossOrderLevelRestriction struct {
	// Applies to Trades with a GuaranteedStopLossOrder attached for the specified Instrument. This is the total allowed
	// Trade volume that can exist within the priceRange based on the trigger prices of the GuaranteedStopLossOrders.
	Volume decimal.Decimal `json:"volume"`

	// The price range the volume applies to. This value is in price units.
	PriceRange decimal.Decimal `json:"priceRange"`
}

// Direction in the context of an Order or a Trade defines whether the units are positive or negative.
type Direction string

const (
	// DirectionLong represents long direction. A long Order is used to buy units of an Instrument. A Trade is long when
	// it has bought units of an Instrument.
	DirectionLong = Direction("LONG")

	// DirectionShort represents short direction. A short Order is used to sell units of an Instrument. A Trade is short
	// when it has sold units of an Instrument.
	DirectionShort = Direction("SHORT")
)

// PricingComponent stands for the Price component(s) to get candlestick data for.
// Format: Can contain any combination of the characters “M” (midpoint candles) “B” (bid candles) and “A” (ask candles).
type PricingComponent string

// ConversionFactor contains information used to convert an amount, from an Instrument's base or quote currency, to the
// home currency of an Account
type ConversionFactor struct {
	// The factor by which to multiply the amount in the given currency to obtain the amount in the home currency of the
	// Account.
	Factor decimal.Decimal `json:"factor"`
}

// HomeConversionFactors message contains information used to convert amounts, from an Instrument's base or quote
// currency, to the home currency of an Account.
type HomeConversionFactors struct {
	// The ConversionFactor in effect for the Account for converting any gains realized in Instrument quote units into
	// units of the Account’s home currency.
	GainQuoteHome ConversionFactor `json:"gainQuoteHome"`

	// The ConversionFactor in effect for the Account for converting any losses realized in Instrument quote units into
	// units of the Account’s home currency.
	LossQuoteHome ConversionFactor `json:"lossQuoteHome"`

	// The ConversionFactor in effect for the Account for converting any gains realized in Instrument base units into
	// units of the Account’s home currency.
	GainBaseHome ConversionFactor `json:"gainBaseHome"`

	// The ConversionFactor in effect for the Account for converting any losses realized in Instrument base units into
	// units of the Account’s home currency.
	LossBaseHome ConversionFactor `json:"lossBaseHome"`
}
