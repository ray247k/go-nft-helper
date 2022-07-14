package get_nft_total_supply

import (
	"fmt"
	"math/big"

	erc721 "go-nft-helper/assets/abi"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func GetTotalSupply(hexAddress string) (*big.Int, error) {
	viper.SetConfigFile("configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
	}

	infuraUrl := viper.GetString("infura.endpoint")

	client, err := ethclient.Dial(infuraUrl)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	address := common.HexToAddress(hexAddress)
	instance, err := erc721.NewErc721(address, client)
	if err != nil {
		return nil, err
	}

	totalSupply, err := instance.TotalSupply(nil)
	if err != nil {
		return nil, err
	}
	return totalSupply, nil
}
