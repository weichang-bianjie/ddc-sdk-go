package handler

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/config"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/pkg/contracts"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/pkg/log"
)

// GetConn 获取连接实体
func GetConn() *ethclient.Client {
	var ctx context.Context
	ctx = context.Background()
	//设置请求参数
	ctx = context.WithValue(ctx, config.Info.HeaderKey(), config.Info.HeaderValue())
	conn, err := ethclient.DialContext(ctx, config.Info.OpbGatewayAddress())
	if err != nil {
		log.Error.Printf("failed to connect to the client :%v", err.Error())
	}
	return conn
}

// GetAuthority 获取Authority合约实体
func GetAuthority() *contracts.Authority {
	authority, err := contracts.NewAuthority(config.Info.AuthorityAddress(), GetConn())
	if err != nil {
		log.Error.Printf("failed to get contract Authority:%v", err)
		return nil
	}
	return authority
}

// GetCharge 获取Charge合约实体
func GetCharge() *contracts.Charge {
	charge, err := contracts.NewCharge(config.Info.ChargeAddress(), GetConn())
	if err != nil {
		log.Error.Printf("failed to get contract Charge:%v", err)
		return nil
	}
	return charge
}

// GetDDC721 获取DDC721合约实体
func GetDDC721() *contracts.DDC721 {
	ddc721, err := contracts.NewDDC721(config.Info.Ddc721Address(), GetConn())
	if err != nil {
		log.Error.Printf("failed to get contract DDC721:%v", err)
	}
	return ddc721
}

// GetDDC1155 获取DDC1155合约实体
func GetDDC1155() *contracts.DDC1155 {
	ddc1155, err := contracts.NewDDC1155(config.Info.Ddc1155Address(), GetConn())
	if err != nil {
		log.Error.Printf("failed to get contract DDC1155:%v", err)
	}
	return ddc1155
}
