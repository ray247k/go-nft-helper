package get_nft_owner_of

import (
	"fmt"
	"math/big"

	erc721 "egox/assets/abi"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func GetOwnerOf(hexAddress string, tokenId int) (*common.Address, error) {
	viper.SetConfigFile("configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return nil, err
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

	tokenId_ := big.NewInt(int64(tokenId))
	owner, err := instance.OwnerOf(nil, tokenId_)
	if err != nil {
		return nil, err
	}

	return &owner, nil
}
