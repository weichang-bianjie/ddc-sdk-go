package service

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/handler"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/log"
	types2 "github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/types"
)

type ChargeService struct {
	Base
}

func NewChargeService() *ChargeService {
	return &ChargeService{}
}

// Recharge
// @Description:运营方、平台方调用该接口为所属同一方的同一级别账户或者下级账户充值
// @receiver c
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param to 接收者账户
// @param amount 充值金额
// @return signedTx 签名好的交易
// @return error
func (c *ChargeService) Recharge(opts *bind.TransactOpts, to string, amount int64) (signedTx *types.Transaction, err error) {
	if common.HexToAddress(to) == common.HexToAddress("0") || !common.IsHexAddress(to) {
		return nil, types2.ToAccountError
	}
	if amount <= 0 {
		return nil, types2.AmountError
	}

	c.setOpts(opts)
	signedTx, err = handler.GetCharge().Recharge(opts, common.HexToAddress(to), big.NewInt(amount))
	if err != nil {
		log.Error.Printf("failed to execute Recharge: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// BalanceOf
// @Description: 查询指定账户的余额
// @receiver c
// @param accAddr 要查询的账户地址
// @return uint64 账户所对应的业务费余额
// @return error
func (c *ChargeService) BalanceOf(accAddr string) (uint64, error) {
	if common.HexToAddress(accAddr) == common.HexToAddress("0") || !common.IsHexAddress(accAddr) {
		return 0, types2.AccountError
	}

	balance, err := handler.GetCharge().BalanceOf(nil, common.HexToAddress(accAddr))
	if err != nil {
		log.Error.Printf("failed to execute BalanceOf: %v", err.Error())
	}

	return balance.Uint64(), nil
}

// QueryFee
// @Description: 查询指定的DDC业务方法所对应的业务费用
// @receiver c
// @param contrAddr DDC业务合约地址
// @param sig DDC业务方法对应的sig编码
// @return uint32 业务费用
// @return error
func (c *ChargeService) QueryFee(contrAddr string, sig [4]byte) (uint32, error) {

	if common.HexToAddress(contrAddr) == common.HexToAddress("0") || !common.IsHexAddress(contrAddr) {
		return 0, types2.ContractAddressError
	}

	fee, err := handler.GetCharge().QueryFee(nil, common.HexToAddress(contrAddr), sig)
	if err != nil {
		log.Error.Printf("failed to execute QueryFee: %v", err.Error())
		return 0, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return fee, nil
}
