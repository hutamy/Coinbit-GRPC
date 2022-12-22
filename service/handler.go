package service

import (
	d "coinbit/deposit"
	"coinbit/proto/pb"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lovoo/goka"
)

func deposit(emitter *goka.Emitter, stream goka.Stream) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DepositRequest
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondWithError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		err = json.Unmarshal(b, &req)
		if err != nil {
			respondWithError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if !(req.Amount > 0) {
			respondWithError(w, http.StatusUnprocessableEntity, "Amount must be more than 0")
			return
		}
		deposit := &pb.Deposit{
			WalletId: req.WalletId,
			Amount:   req.Amount,
		}
		if stream == d.DepositStream {
			err = emitter.EmitSync(req.WalletId, deposit)
		} else {
			deposit.Amount = -1 * deposit.Amount
			err = emitter.EmitSync(req.WalletId, deposit)
		}

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, nil)
	}
}

func walletDetail(view *goka.View, flagView *goka.View) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		walletId := mux.Vars(r)["id"]
		var totalBalance float64
		var aboveThreshold bool

		response := CheckResponse{
			WalletId:       walletId,
			Balance:        totalBalance,
			AboveThreshold: aboveThreshold,
		}

		val, _ := view.Get(walletId)
		if val == nil {
			respondWithJSON(w, http.StatusOK, response)
			return
		}

		messages := val.(*pb.DepositHistory)
		for _, m := range messages.Deposits {
			totalBalance += m.Amount
		}

		flagVal, _ := flagView.Get(walletId)
		if flagVal != nil {
			b := flagVal.(*pb.FlagValue)
			aboveThreshold = b.Flagged
		}

		response.Balance = totalBalance
		response.AboveThreshold = aboveThreshold
		respondWithJSON(w, http.StatusOK, response)
	}
}
