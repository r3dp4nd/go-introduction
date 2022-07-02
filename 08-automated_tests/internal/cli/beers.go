package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Sraik25/golang-introduction/08-automated_tests/internal/fetching"
	"github.com/spf13/cobra"
)

const idFlag = "id"

// InitBeersCmd initialize beers command
func InitBeersCmd(service fetching.Service) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(service),
	}
	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beers")

	return beersCmd
}

func runBeersFn(service fetching.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			i, _ := strconv.Atoi(id)
			beer, err := service.FetchByID(i)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(beer)
			return
		}

		beers, err := service.FetchBeers()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(beers)
	}
}
