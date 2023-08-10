/**
* PACKAGE service
* Name tickets
* Description TODO
* Author lifezqy@126.com
* Date 2023/8/10
 */
package service

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	"github.com/lifezq/tickets/internal/gormgen"
	"github.com/lifezq/tickets/internal/model"
)

type DataResult struct {
	Data []struct {
		Name        string `json:"name"`
		Code        string `json:"code"`
		DetailsLink string `json:"detailsLink"`
		VideoLink   string `json:"videoLink"`
		Date        string `json:"date"`
		Week        string `json:"week"`
		Red         string `json:"red"`
		Blue        string `json:"blue"`
		Blue2       string `json:"blue2"`
		Sales       string `json:"sales"`
		Poolmoney   string `json:"poolmoney"`
		Content     string `json:"content"`
		Addmoney    string `json:"addmoney"`
		Addmoney2   string `json:"addmoney2"`
		Msg         string `json:"msg"`
		Z2Add       string `json:"z2add"`
		M2Add       string `json:"m2add"`
		Prizegrades []struct {
			Type      int    `json:"type"`
			Typenum   string `json:"typenum"`
			Typemoney string `json:"typemoney"`
		} `json:"prizegrades"`
	} `json:"data"`
}

type Data2Result struct {
	Data []struct {
		Issue              string `json:"issue"`
		OpenTime           string `json:"openTime"`
		FrontWinningNum    string `json:"frontWinningNum"`
		BackWinningNum     string `json:"backWinningNum"`
		SeqFrontWinningNum string `json:"seqFrontWinningNum"`
		SeqBackWinningNum  string `json:"seqBackWinningNum"`
		SaleMoney          string `json:"saleMoney"`
		R9SaleMoney        string `json:"r9SaleMoney"`
		PrizePoolMoney     string `json:"prizePoolMoney"`
		Week               string `json:"week"`
		WinnerDetails      []struct {
			AwardEtc      string `json:"awardEtc"`
			BaseBetWinner struct {
				Remark     string `json:"remark"`
				AwardNum   string `json:"awardNum"`
				AwardMoney string `json:"awardMoney"`
				TotalMoney string `json:"totalMoney"`
			} `json:"baseBetWinner"`
			AddToBetWinner  string `json:"addToBetWinner"`
			AddToBetWinner2 string `json:"addToBetWinner2"`
			AddToBetWinner3 string `json:"addToBetWinner3"`
		} `json:"winnerDetails"`
	} `json:"data"`
}

func DataSync() {
	fp, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}

	var value DataResult
	err = json.Unmarshal(body, &value)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range value.Data {
		gormgen.Ticket.Create(&model.Ticket{
			Numbers: s.Red + "," + s.Blue,
			Ymd:     s.Code,
		})
	}
}

func Data2Sync() {
	fp, err := os.Open("data2.json")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}

	var value Data2Result
	err = json.Unmarshal(body, &value)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range value.Data {
		gormgen.Ticket.Create(&model.Ticket{
			Numbers: strings.Join(strings.Split(s.FrontWinningNum, " "),
				",") + "," + s.BackWinningNum,
			Ymd: s.Issue,
		})
	}
}
