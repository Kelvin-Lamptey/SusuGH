package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func testGroup(){
	setupApp()
	payments = nil
	groups["Susu"] = map[int]float64{1: 50}
	group_user["Ama"] = "Susu"
	group_user["Kofi"] = "Susu"
	group_user["Esi"] = "Susu"
	group_members["Susu"] = []string{"Ama", "Kofi", "Esi"}
	cur_pot_user["Susu"] = "Ama"
}

func withInput(input string, action func()){
	old := os.Stdin
	reader, writer, _ := os.Pipe()
	writer.WriteString(input)
	writer.Close()
	os.Stdin = reader
	action()
	os.Stdin = old
}

func captureOutput(action func()) string{
	old := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	action()
	writer.Close()
	os.Stdout = old
	result, _ := io.ReadAll(reader)
	return string(result)
}

func TestCurrentWeek(t *testing.T){
	testGroup()
	if getCurrentWeek("Susu") != 1{
		t.Fatal("expected current week to be 1")
	}
}

func TestMembersKeepJoinOrder(t *testing.T){
	testGroup()
	if getNextPotUser("Susu") != "Kofi"{
		t.Fatal("expected Kofi to receive the next payout")
	}
	cur_pot_user["Susu"] = "Esi"
	if getNextPotUser("Susu") != "Ama"{
		t.Fatal("expected payout order to return to Ama")
	}
}

func TestPartialPayment(t *testing.T){
	testGroup()
	cur_user = "Esi"
	withInput("30\n", func(){
		payAmountForGroup("Susu")
	})
	if payments[0]["Esi/|Susu"] != 30{
		t.Fatal("expected partial payment of 30")
	}
}

func TestPaymentAboveAmountIsRejected(t *testing.T){
	testGroup()
	cur_user = "Esi"
	output := captureOutput(func(){
		withInput("60\n", func(){
			payAmountForGroup("Susu")
		})
	})
	if len(payments) != 0{
		t.Fatal("expected payment above weekly amount to be rejected")
	}
	if !strings.Contains(output, "above the weekly amount"){
		t.Fatal("expected an overpayment message")
	}
}

func TestWeekCannotCloseWithPartialPayment(t *testing.T){
	testGroup()
	payments = []map[string]float64{
		{"Ama/|Susu": 50},
		{"Kofi/|Susu": 50},
		{"Esi/|Susu": 30},
	}
	output := captureOutput(func(){
		closeCurrentWeek("Susu")
	})
	if getCurrentWeek("Susu") != 1{
		t.Fatal("expected week to remain open")
	}
	if len(payout_history["Susu"]) != 0{
		t.Fatal("expected no payout when a member owes money")
	}
	if !strings.Contains(output, "Esi owes 20"){
		t.Fatal("expected the unpaid amount to be reported")
	}
}

func TestWeekClosesWhenEveryonePays(t *testing.T){
	testGroup()
	payments = []map[string]float64{
		{"Ama/|Susu": 50},
		{"Kofi/|Susu": 50},
		{"Esi/|Susu": 50},
	}
	closeCurrentWeek("Susu")
	if getCurrentWeek("Susu") != 2{
		t.Fatal("expected the next week to start")
	}
	if len(payout_history["Susu"]) != 1{
		t.Fatal("expected one payout in history")
	}
	if cur_pot_user["Susu"] != "Kofi"{
		t.Fatal("expected Kofi to receive the next payout")
	}
}

func TestReportShowsPaymentAndHistory(t *testing.T){
	testGroup()
	payments = []map[string]float64{
		{"Ama/|Susu": 50},
		{"Kofi/|Susu": 30},
		{"Esi/|Susu": 50},
	}
	payout_history["Susu"] = []string{"Week 0: Ama received 150.00"}
	output := captureOutput(func(){
		reportGroup("Susu")
	})
	if !strings.Contains(output, "Ama has paid in full") ||
		!strings.Contains(output, "Kofi owes 20.00") ||
		!strings.Contains(output, "Week 0: Ama received 150.00"){
		t.Fatal("expected report to show payments, amount owed, and payout history")
	}
}
