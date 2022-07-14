package get_nft_token_uri

import (
	"fmt"
	erc721 "go-nft-helper/assets/abi"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func GetNftTokenUri(hexAddress string, tokenId int) (*string, error) {
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
	tokenURI, err := instance.TokenURI(nil, tokenId_)
	if err != nil {
		return nil, err
	}

	return &tokenURI, nil
}
