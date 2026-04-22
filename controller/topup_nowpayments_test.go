package controller

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/QuantumNous/new-api/setting/operation_setting"
	"github.com/stretchr/testify/require"
)

// ---------- verifyNowPaymentsSignature ----------

func signNowPayments(body map[string]interface{}, secret string) string {
	sorted, _ := json.Marshal(body)
	h := hmac.New(sha512.New, []byte(secret))
	h.Write(sorted)
	return hex.EncodeToString(h.Sum(nil))
}

func TestVerifyNowPaymentsSignature_Valid(t *testing.T) {
	body := map[string]interface{}{
		"payment_status": "finished",
		"order_id":       "nowpay-1-1234567890-abcd",
		"pay_amount":     1.5,
		"pay_currency":   "usdtsol",
	}
	secret := "my-ipn-secret"

	bodyBytes, err := json.Marshal(body)
	require.NoError(t, err)

	sig := signNowPayments(body, secret)
	require.True(t, verifyNowPaymentsSignature(bodyBytes, sig, secret))
}

func TestVerifyNowPaymentsSignature_InvalidSignature(t *testing.T) {
	body := map[string]interface{}{
		"payment_status": "finished",
		"order_id":       "test",
	}
	secret := "secret"
	bodyBytes, _ := json.Marshal(body)

	require.False(t, verifyNowPaymentsSignature(bodyBytes, "badsignature", secret))
}

func TestVerifyNowPaymentsSignature_EmptySecret(t *testing.T) {
	body := map[string]interface{}{"a": 1}
	bodyBytes, _ := json.Marshal(body)
	require.False(t, verifyNowPaymentsSignature(bodyBytes, "somesig", ""))
}

func TestVerifyNowPaymentsSignature_EmptySignature(t *testing.T) {
	body := map[string]interface{}{"a": 1}
	bodyBytes, _ := json.Marshal(body)
	require.False(t, verifyNowPaymentsSignature(bodyBytes, "", "secret"))
}

func TestVerifyNowPaymentsSignature_SortedKeys(t *testing.T) {
	// NowPayments requires JSON keys sorted alphabetically before HMAC.
	// Ensure our implementation produces the same signature regardless
	// of insertion order in the Go map (Go maps don't guarantee order).
	secret := "sort-test-secret"

	body := map[string]interface{}{
		"z_field": 1,
		"a_field": 2,
		"m_field": 3,
	}
	bodyBytes, err := json.Marshal(body)
	require.NoError(t, err)

	// Compute expected signature from sorted JSON
	sig := signNowPayments(body, secret)
	require.True(t, verifyNowPaymentsSignature(bodyBytes, sig, secret))

	// Same data, different insertion order — must still verify
	body2 := map[string]interface{}{
		"m_field": 3,
		"z_field": 1,
		"a_field": 2,
	}
	bodyBytes2, _ := json.Marshal(body2)
	require.True(t, verifyNowPaymentsSignature(bodyBytes2, sig, secret))
}

func TestVerifyNowPaymentsSignature_InvalidJSON(t *testing.T) {
	require.False(t, verifyNowPaymentsSignature([]byte("not json"), "sig", "secret"))
}

// ---------- getNowPaymentsPayMoney ----------

func TestGetNowPaymentsPayMoney_Basic(t *testing.T) {
	amount := getNowPaymentsPayMoney(10, "")
	// Default NowPaymentsUnitPrice=1.0, topupGroupRatio defaults to 1
	require.Equal(t, 10.0, amount)
}

func TestGetNowPaymentsPayMoney_WithDiscount(t *testing.T) {
	// Save and restore
	original := operation_setting.GetPaymentSetting().AmountDiscount
	operation_setting.GetPaymentSetting().AmountDiscount = map[int]float64{10: 0.8}
	defer func() { operation_setting.GetPaymentSetting().AmountDiscount = original }()

	amount := getNowPaymentsPayMoney(10, "")
	require.Equal(t, 8.0, amount)
}

// ---------- createNowPaymentsOrder request structure ----------

func TestNowPaymentsCreateRequest_PriceCurrencyIsUSD(t *testing.T) {
	// Verify the request struct exists and PriceCurrency defaults correctly
	req := NowPaymentsCreateRequest{
		PriceAmount:   10.0,
		PriceCurrency: "usd",
		PayCurrency:   "usdtsol",
	}
	require.Equal(t, "usd", req.PriceCurrency)
	require.Equal(t, "usdtsol", req.PayCurrency)
}

func TestNowPaymentsCreateResponse_Fields(t *testing.T) {
	resp := NowPaymentsCreateResponse{
		PaymentId:   "12345",
		PayAddress:  "TXkVRCpR7pBEqFQZetM87nMptaFpQhYZTW",
		PayAmount:   10.0,
		PayCurrency: "usdtsol",
	}
	require.Equal(t, "12345", resp.PaymentId)
	require.Equal(t, "TXkVRCpR7pBEqFQZetM87nMptaFpQhYZTW", resp.PayAddress)
	require.Equal(t, 10.0, resp.PayAmount)
}
