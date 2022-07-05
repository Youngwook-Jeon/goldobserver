package main

import (
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var jsonToReturn = `{"ts":1657001160369,"tsj":1657001156052,"date":"Jul 5th 2022, 02:05:56 am NY","items":[{"curr":"USD","xauPrice":1811.445,"xagPrice":20.1408,"chgXau":1.825,"chgXag":0.1483,"pcXau":0.1008,"pcXag":0.7418,"xauClose":1809.62,"xagClose":19.9925}]}`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
