package main

import (
	"egox/pkg/get_nft_owner_of"
	"egox/pkg/get_nft_token_uri"
	"egox/pkg/get_nft_total_supply"
	"log"
	"strconv"
)

func main() {
	getItemsOwnerOfByContract("0x4e2Fbc9e3feB25999991C249217d8ee5608860AD")
}

func getItemsOwnerOfByContract(contract string) {
	totalSupply, err := get_nft_total_supply.GetTotalSupply(contract)
	if err != nil {
		log.Printf("Get NFT total supply failed, err: %v\n", err)
	}

	log.Printf("NFT total supply: %s", totalSupply)

	totalAmount, _ := strconv.Atoi(totalSupply.String())

	for i := 1; i <= totalAmount; i++ {
		owner, err := get_nft_owner_of.GetOwnerOf("0x4e2fbc9e3feb25999991c249217d8ee5608860ad", i)
		if err != nil {
			log.Printf("Get NFT owner of failed, err: %v\n", err)
		}

		tokenURI, err := get_nft_token_uri.GetNftTokenUri("0x4e2fbc9e3feb25999991c249217d8ee5608860ad", i)

		if err != nil {
			log.Printf("Get NFT tokenURI failed, err: %v\n", err)
		}

		log.Printf("Token Id: %v owner: %s tokenURI: %s", i, owner, *tokenURI)

	}
}
