package types

import (
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramstypes.ParamSet = (*Params)(nil)

var (
	KeyFeeRate             = []byte("FeeRate")
	KeyMinInitialLiquidity = []byte("MinInitialLiquidity")

	DefaultFeeRate             = sdk.NewDecWithPrec(3, 3)
	DefaultMinInitialLiquidity = math.NewInt(1000)
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(feeRate sdk.Dec, minInitialLiquidity math.Int) Params {
	return Params{
		FeeRate:             feeRate,
		MinInitialLiquidity: minInitialLiquidity,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultFeeRate,
		DefaultMinInitialLiquidity,
	)
}

// ParamSetPairs get the params.ParamSet
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyFeeRate, &params.FeeRate, validateFeeRate),
		paramstypes.NewParamSetPair(KeyMinInitialLiquidity, &params.MinInitialLiquidity, validateMinInitialLiquidity),
	}
}

func (params *Params) Validate() error {
	if err := validateFeeRate(params.FeeRate); err != nil {
		return err
	}
	if err := validateMinInitialLiquidity(params.MinInitialLiquidity); err != nil {
		return err
	}
	return nil
}
func (params *Params) String() string {
	out, _ := yaml.Marshal(params)
	return string(out)
}

func validateFeeRate(v interface{}) error {
	feeRate, ok := v.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	if feeRate.IsNegative() {
		return fmt.Errorf("invalid parameter value: %v", feeRate)
	}

	return nil
}

func validateMinInitialLiquidity(v interface{}) error {
	minInitialLiquidity, ok := v.(math.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	if minInitialLiquidity.IsNegative() {
		return fmt.Errorf("invalid parameter value: %v", minInitialLiquidity)
	}

	return nil
}
