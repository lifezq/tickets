/**
* PACKAGE service
* Name analysis
* Description TODO
* Author lifezqy@126.com
* Date 2023/8/10
 */
package service

import (
	"log"
	"strings"

	"github.com/lifezq/tickets/internal/gormgen"
)

func Analysis() {
	find, err := gormgen.MyDatum.Find()
	if err != nil {
		panic(err)
		return
	}

	for _, datum := range find {
		match(datum.Numbers)
		numbers := datum.Numbers[0:strings.LastIndex(datum.Numbers, ",")]
		match(numbers + "%")
		numbers = numbers[0:strings.LastIndex(numbers, ",")]
		match(numbers + "%")
		numbers = numbers[0:strings.LastIndex(numbers, ",")]
		match(numbers + "%")
	}
}

func match(numbers string) {
	matchResult, err := gormgen.Ticket.Where(gormgen.Ticket.Numbers.Like(numbers)).Find()
	if err != nil {
		return
	}
	log.Printf("analyzing datum: %s matched: %v", numbers, matchResult)
}
