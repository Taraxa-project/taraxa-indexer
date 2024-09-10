package lara

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

type StakerResponse struct {
	Stakers []Staker `json:"stakers"`
}

type Staker struct {
	ID string `json:"id"`
}

func GetStakedTaraHolders(lara *Lara, blockNumber uint64) []string {
	if lara.graphQLEndpoint == "" {
		log.Fatal("GraphQL endpoint is not set")
	}
	client := graphql.NewClient(lara.graphQLEndpoint)

	req := graphql.NewRequest(`
		query($blockNumber: Int!) {
			stakers(
				block: {number: $blockNumber},
				where: { stTaraBalance_gt: 0 }
			) {
				id
			}
		}
	`)

	req.Var("blockNumber", int(blockNumber))

	var resp StakerResponse

	if err := client.Run(context.Background(), req, &resp); err != nil {
		log.Printf("GraphQL query error: %v", err)
		return nil
	}

	if len(resp.Stakers) == 0 {
		fmt.Println("No stakers found in the response")
		return nil
	}

	// Extract staker IDs
	var stakerIDs []string
	for _, staker := range resp.Stakers {
		stakerIDs = append(stakerIDs, staker.ID)
	}

	return stakerIDs
}
