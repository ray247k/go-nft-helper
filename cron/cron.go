package cron

import (
	"log"
	"strconv"

	"go-nft-helper/database/model"
	"go-nft-helper/pkg/get_nft_owner_of"
	"go-nft-helper/pkg/get_nft_token_uri"
	"go-nft-helper/pkg/get_nft_total_supply"

	"github.com/robfig/cron/v3"
)

func Cronjob() {
	log.Println("Cron Starting...")

	c := cron.New(cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))

	// TODO 因為是送到 goroutine 所以 SkipIfStillRunning 要自己判斷才行
	c.AddFunc("* * * * *", func() {
		contracts := model.GetAllContracts()
		for _, contract := range contracts {
			go getItemsOwnerOfByContract(contract)
		}

	})
	c.Start()
	defer c.Stop()
	select {}
}

func getItemsOwnerOfByContract(contract string) {
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
