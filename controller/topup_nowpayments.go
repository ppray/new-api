package controller

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting"
	"github.com/QuantumNous/new-api/setting/operation_setting"
	"github.com/QuantumNous/new-api/setting/system_setting"

	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
)

const NowPaymentsSignatureHeader = "x-nowpayments-sig"

var nowPaymentsAdaptor = &NowPaymentsAdaptor{}

type NowPaymentsPayRequest struct {
	Amount        int64  `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	PayCurrency   string `json:"pay_currency,omitempty"`
}

type NowPaymentsCreateRequest struct {
	PriceAmount      float64 `json:"price_amount"`
	PriceCurrency    string  `json:"price_currency"`
	PayCurrency      string  `json:"pay_currency"`
	OrderId          string  `json:"order_id"`
	OrderDescription string  `json:"order_description"`
	IpnCallbackUrl   string  `json:"ipn_callback_url"`
	SuccessUrl       string  `json:"success_url"`
	CancelUrl        string  `json:"cancel_url"`
}

type NowPaymentsCreateResponse struct {
	PaymentId   string  `json:"payment_id"`
	PaymentUrl  string  `json:"payment_url"`
	PayAddress  string  `json:"pay_address"`
	PayAmount   float64 `json:"pay_amount"`
	PayCurrency string  `json:"pay_currency"`
	ValidUntil  string  `json:"valid_until"`
}

type NowPaymentsWebhookData struct {
	PaymentStatus    string  `json:"payment_status"`
	PayAddress       string  `json:"pay_address"`
	PriceAmount      float64 `json:"price_amount"`
	PriceCurrency    string  `json:"price_currency"`
	PayAmount        float64 `json:"pay_amount"`
	PayCurrency      string  `json:"pay_currency"`
	OrderId          string  `json:"order_id"`
	OrderDescription string  `json:"order_description"`
}

type NowPaymentsAdaptor struct{}

func (*NowPaymentsAdaptor) RequestAmount(c *gin.Context, req *NowPaymentsPayRequest) {
	if req.Amount < int64(setting.NowPaymentsMinTopUp) {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": fmt.Sprintf("充值数量不能小于 %d", setting.NowPaymentsMinTopUp)})
		return
	}
	id := c.GetInt("id")
	group, err := model.GetUserGroup(id, true)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "获取用户分组失败"})
		return
	}
	payMoney := getNowPaymentsPayMoney(float64(req.Amount), group)
	if payMoney <= 0.01 {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "充值金额过低"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": strconv.FormatFloat(payMoney, 'f', 2, 64)})
}

func (*NowPaymentsAdaptor) RequestPay(c *gin.Context, req *NowPaymentsPayRequest) {
	if req.PaymentMethod != model.PaymentMethodNowPayments {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "不支持的支付渠道"})
		return
	}
	if !setting.NowPaymentsEnabled {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "NowPayments支付暂未开启"})
		return
	}
	if req.Amount < int64(setting.NowPaymentsMinTopUp) {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("充值数量不能小于 %d", setting.NowPaymentsMinTopUp), "data": 10})
		return
	}
	if req.Amount > 10000 {
		c.JSON(http.StatusOK, gin.H{"message": "充值数量不能大于 10000", "data": 10})
		return
	}

	id := c.GetInt("id")
	group, err := model.GetUserGroup(id, true)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "获取用户分组失败"})
		return
	}
	payMoney := getNowPaymentsPayMoney(float64(req.Amount), group)
	if payMoney <= 0.01 {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "充值金额过低"})
		return
	}

	referenceId := fmt.Sprintf("nowpay-%d-%d-%s", id, time.Now().UnixMilli(), randstr.String(6))

	payCurrency := req.PayCurrency
	if payCurrency == "" {
		payCurrency = setting.NowPaymentsCurrency
	}

	ipnUrl := system_setting.ServerAddress + "/api/nowpayments/webhook"
	successUrl := system_setting.ServerAddress + "/console/topup?success=1"
	cancelUrl := system_setting.ServerAddress + "/console/topup?cancel=1"

	apiReq := NowPaymentsCreateRequest{
		PriceAmount:      payMoney,
		PriceCurrency:    "usd",
		PayCurrency:      payCurrency,
		OrderId:          referenceId,
		OrderDescription: fmt.Sprintf("API额度充值 %.2f USD", payMoney),
		IpnCallbackUrl:   ipnUrl,
		SuccessUrl:       successUrl,
		CancelUrl:        cancelUrl,
	}

	payResp, err := createNowPaymentsOrder(apiReq)
	if err != nil {
		logger.LogError(c.Request.Context(), fmt.Sprintf("NowPayments创建支付订单失败 user_id=%d error=%v", id, err))
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "拉起支付失败"})
		return
	}

	topUp := &model.TopUp{
		UserId:        id,
		Amount:        req.Amount,
		Money:         payMoney,
		TradeNo:       referenceId,
		PaymentMethod: model.PaymentMethodNowPayments,
		PayAddress:    payResp.PayAddress,
		CreateTime:    time.Now().Unix(),
		Status:        common.TopUpStatusPending,
	}
	err = topUp.Insert()
	if err != nil {
		logger.LogError(c.Request.Context(), fmt.Sprintf("NowPayments创建充值订单失败 user_id=%d trade_no=%s error=%v", id, referenceId, err))
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "创建订单失败"})
		return
	}

	logger.LogInfo(c.Request.Context(), fmt.Sprintf("NowPayments充值订单创建成功 user_id=%d trade_no=%s amount=%d money=%.2f pay_url=%s",
		id, referenceId, req.Amount, payMoney, payResp.PaymentUrl))

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"trade_no":     referenceId,
			"pay_link":     payResp.PaymentUrl,
			"pay_address":  payResp.PayAddress,
			"pay_amount":   payResp.PayAmount,
			"pay_currency": payResp.PayCurrency,
			"valid_until":  payResp.ValidUntil,
		},
	})
}

func RequestNowPaymentsAmount(c *gin.Context) {
	var req NowPaymentsPayRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "参数错误"})
		return
	}
	nowPaymentsAdaptor.RequestAmount(c, &req)
}

func RequestNowPaymentsPay(c *gin.Context) {
	var req NowPaymentsPayRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "参数错误"})
		return
	}
	nowPaymentsAdaptor.RequestPay(c, &req)
}

func NowPaymentsWebhook(c *gin.Context) {
	ctx := c.Request.Context()
	if !isNowPaymentsWebhookEnabled() {
		logger.LogWarn(ctx, fmt.Sprintf("NowPayments webhook被拒绝 reason=webhook_disabled path=%q client_ip=%s", c.Request.RequestURI, c.ClientIP()))
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.LogError(ctx, fmt.Sprintf("NowPayments webhook读取请求体失败 path=%q client_ip=%s error=%q", c.Request.RequestURI, c.ClientIP(), err.Error()))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	signature := c.GetHeader(NowPaymentsSignatureHeader)
	if !verifyNowPaymentsSignature(body, signature, setting.NowPaymentsWebhookSecret) {
		logger.LogWarn(ctx, fmt.Sprintf("NowPayments webhook验签失败 path=%q client_ip=%s", c.Request.RequestURI, c.ClientIP()))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var webhookData NowPaymentsWebhookData
	err = common.Unmarshal(body, &webhookData)
	if err != nil {
		logger.LogError(ctx, fmt.Sprintf("NowPayments webhook解析数据失败 path=%q client_ip=%s error=%q", c.Request.RequestURI, c.ClientIP(), err.Error()))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	logger.LogInfo(ctx, fmt.Sprintf("NowPayments webhook收到通知 order_id=%s status=%s amount=%.2f",
		webhookData.OrderId, webhookData.PaymentStatus, webhookData.PriceAmount))

	tradeNo := webhookData.OrderId
	LockOrder(tradeNo)
	defer UnlockOrder(tradeNo)

	topUp := model.GetTopUpByTradeNo(tradeNo)
	if topUp == nil {
		logger.LogWarn(ctx, fmt.Sprintf("NowPayments订单不存在 trade_no=%s", tradeNo))
		c.Status(http.StatusOK)
		return
	}

	if topUp.PaymentMethod != model.PaymentMethodNowPayments {
		logger.LogWarn(ctx, fmt.Sprintf("NowPayments订单支付方式不匹配 trade_no=%s method=%s", tradeNo, topUp.PaymentMethod))
		c.Status(http.StatusOK)
		return
	}

	switch webhookData.PaymentStatus {
	case "finished":
		if topUp.Status == common.TopUpStatusSuccess {
			logger.LogInfo(ctx, fmt.Sprintf("NowPayments订单已处理 trade_no=%s", tradeNo))
			c.Status(http.StatusOK)
			return
		}
		if topUp.Status != common.TopUpStatusPending {
			logger.LogWarn(ctx, fmt.Sprintf("NowPayments订单状态异常 trade_no=%s status=%s", tradeNo, topUp.Status))
			c.Status(http.StatusOK)
			return
		}

		err = model.RechargeNowPayments(tradeNo, c.ClientIP())
		if err != nil {
			logger.LogError(ctx, fmt.Sprintf("NowPayments充值处理失败 trade_no=%s error=%v", tradeNo, err))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		logger.LogInfo(ctx, fmt.Sprintf("NowPayments充值成功 user_id=%d trade_no=%s amount=%d pay_address=%s", topUp.UserId, tradeNo, topUp.Amount, topUp.PayAddress))

	case "failed", "expired", "refunded":
		if topUp.Status == common.TopUpStatusPending {
			topUp.Status = common.TopUpStatusFailed
			if err := topUp.Update(); err != nil {
				logger.LogError(ctx, fmt.Sprintf("NowPayments更新订单状态失败 trade_no=%s error=%v", tradeNo, err))
			}
		}
		logger.LogInfo(ctx, fmt.Sprintf("NowPayments订单 %s trade_no=%s", webhookData.PaymentStatus, tradeNo))
	}

	c.Status(http.StatusOK)
}

func createNowPaymentsOrder(req NowPaymentsCreateRequest) (*NowPaymentsCreateResponse, error) {
	jsonData, err := common.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", "https://api.nowpayments.io/v1/payment", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-api-key", setting.NowPaymentsApiKey)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API请求失败: %d, %s", resp.StatusCode, string(body))
	}

	var payResp NowPaymentsCreateResponse
	err = common.Unmarshal(body, &payResp)
	if err != nil {
		return nil, err
	}

	return &payResp, nil
}

func verifyNowPaymentsSignature(body []byte, signature string, secret string) bool {
	if signature == "" || secret == "" {
		return false
	}
	// NowPayments requires sorted JSON fields before signing.
	// Use json.RawMessage to preserve original bytes and avoid float precision loss.
	var rawMap map[string]json.RawMessage
	if err := common.Unmarshal(body, &rawMap); err != nil {
		return false
	}

	keys := make([]string, 0, len(rawMap))
	for k := range rawMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, k := range keys {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(strconv.Quote(k))
		buf.WriteByte(':')
		buf.Write(rawMap[k])
	}
	buf.WriteByte('}')

	h := hmac.New(sha512.New, []byte(secret))
	h.Write(buf.Bytes())
	expectedSig := hex.EncodeToString(h.Sum(nil))
	return hmac.Equal([]byte(signature), []byte(expectedSig))
}

func getNowPaymentsPayMoney(amount float64, group string) float64 {
	originalAmount := amount
	if operation_setting.GetQuotaDisplayType() == operation_setting.QuotaDisplayTypeTokens {
		amount = amount / common.QuotaPerUnit
	}
	topupGroupRatio := common.GetTopupGroupRatio(group)
	if topupGroupRatio == 0 {
		topupGroupRatio = 1
	}
	discount := 1.0
	if ds, ok := operation_setting.GetPaymentSetting().AmountDiscount[int(originalAmount)]; ok {
		if ds > 0 {
			discount = ds
		}
	}
	payMoney := amount * setting.NowPaymentsUnitPrice * topupGroupRatio * discount
	return payMoney
}
