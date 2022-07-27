/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
)

// calcDeliveryCmd represents the calcDelivery command
var calcDeliveryCmd = &cobra.Command{
	Use:   "calcDelivery",
	Short: "Calculate total delivery cost with or without an offer code.",
	Long: `- input base cost and number of packages
- after that, loop through the packages to input data 
and print total cost
- offer codes are case sensitive`,
	Run: func(cmd *cobra.Command, args []string) {

		var baseDeliveryCost float64
		var noOfPkgs int

		//get input values
		fmt.Print("Enter base delivery cost:")
		fmt.Scanf("%f\n", &baseDeliveryCost)
		fmt.Print("Enter number of packages:")
		fmt.Scanf("%d\n", &noOfPkgs)

		if noOfPkgs > 0 {
			//continue
			getInputAndCalc(noOfPkgs, baseDeliveryCost)
		} else {
			// no packages
			log.Fatal("You entered an invalid number of packages")
		}

		//fmt.Println("calcDelivery called")

	},
}

func init() {
	rootCmd.AddCommand(calcDeliveryCmd)
}

type pkg struct {
	PkgID string `json:"ID" validate:"required"`
	// Greater than 0, Less Than or Equal to 500, required, numeric
	PkgWeightInKg float64 `json:"weight" validate:"required,numeric,gt=0,lte=500"`
	//Greater than 0, Less Than or Equal to 1000, required, numeric
	DistanceInKm float64 `json:"distance" validate:"required,numeric,gt=0,lte=1000"`
	//can be empty
	OfferCode string `json:"discCode"`
}

type OfferCrit struct {
	OfferCode string
	DistMin   float64
	DistMax   float64
	WeightMin float64
	WeightMax float64
	Disc      int
}

func getInputAndCalc(noPkgs int, baseDelvryCost float64) {

	var offer OfferCrit
	var offersList []OfferCrit

	//unmarshal the JSON file into a struct
	offers := getOffers()
	err := json.Unmarshal(offers, &offersList)
	if err != nil {
		panic(err)
	}

	pckg := make([]pkg, noPkgs)
	for i := 0; i < noPkgs; i++ {

		fmt.Print("Enter ID for package", i+1, ":")
		fmt.Scanf("%s", &pckg[i].PkgID)
		fmt.Print("Enter weight in kg for package", i+1, ":")
		fmt.Scanf("%f", &pckg[i].PkgWeightInKg)
		fmt.Print("Enter the distance in km for package", i+1, ":")
		fmt.Scanf("%f", &pckg[i].DistanceInKm)
		fmt.Print("Enter offer code for package", i+1, ":")
		fmt.Scanf("%s", &pckg[i].OfferCode)
		fmt.Println("You entered-> ID:", pckg[i].PkgID, "Weight:", pckg[i].PkgWeightInKg,
			"Distance:", pckg[i].DistanceInKm, "Offer Code:", pckg[i].OfferCode)

		//validate input
		isValid := validator.New()
		errs := isValid.Struct(pckg[i])
		if errs == nil {
			//return
		} else {
			fmt.Println("Errors found:")
			for _, e := range errs.(validator.ValidationErrors) {
				fmt.Println(e)
				// we can also handle the errors differently
				// not necessarily have to exit the program
				os.Exit(1)
			}
		}

		offer = OfferCrit{}
		if pckg[i].OfferCode != "" {
			for j := range offersList {
				//fmt.Println(j, "current offer:", offer)
				//fmt.Println(j, "offer:", offersList[j])
				if offersList[j].OfferCode == pckg[i].OfferCode {
					//found offer
					offer = offersList[j]
					fmt.Println("index:", j, "found offer:", offer)
					break
				}
			}
		}
		//total delivery cost
		//print delivery cost in the same loop,
		//otherwise we can print them later, and hold them in a variable
		fmt.Println("Delivery cost for", pckg[i].PkgID, "is:", calcDeliveryCost(baseDelvryCost, pckg[i], offer))
	}

}

func calcDeliveryCost(deliveryCost float64, pack pkg, discDetails OfferCrit) float64 {

	delCost := deliveryCost + pack.PkgWeightInKg*10 + pack.DistanceInKm*5

	if discDetails.DistMin <= pack.DistanceInKm && pack.DistanceInKm <= discDetails.DistMax &&
		discDetails.WeightMin <= pack.PkgWeightInKg && pack.PkgWeightInKg <= discDetails.WeightMax &&
		discDetails.Disc > 0 {

		delCost = delCost - (delCost / 100 * float64(discDetails.Disc))
	}

	return delCost
}

func getOffers() []byte {
	//read offers JSON file
	fileBytes, err := ioutil.ReadFile("./cmd/offercodes.json")

	if err != nil {
		panic(err)
	}

	return fileBytes
}
