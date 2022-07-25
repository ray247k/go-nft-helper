package main

import (
	"go-nft-helper/database/model"
	"go-nft-helper/pkg/get_nft_owner_of"
	"go-nft-helper/pkg/get_nft_token_uri"
	"go-nft-helper/pkg/get_nft_total_supply"
	"log"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		contracts := model.GetAllContracts()
		for _, contract := range contracts {
			wg.Add(1)
			go getItemsOwnerOfByContract(contract)
		}

		wg.Wait()
	}
}

func getItemsOwnerOfByContract(contract string) {
	defer wg.Done()

	totalSupply, err := get_nft_total_supply.GetTotalSupply(contract)
	if err != nil {
		log.Printf("Get NFT total supply failed, err: %v\n", err)
	}

	log.Printf("Contract: %s total supply: %s", contract, totalSupply)

	totalAmount, _ := strconv.Atoi(totalSupply.String())

	for i := 1; i <= totalAmount; i++ {
		owner, err := get_nft_owner_of.GetOwnerOf(contract, i)
		if err != nil {
			log.Printf("Get NFT owner of failed, err: %v\n", err)
		}

		tokenURI, err := get_nft_token_uri.GetNftTokenUri(contract, i)

		if err != nil {
			log.Printf("Get NFT tokenURI failed, err: %v\n", err)
		}

		log.Printf("Token Id: %v owner: %s tokenURI: %s", i, owner, *tokenURI)
	}
}
