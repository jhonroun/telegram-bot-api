package tgbotapi

import (
	"encoding/json"
	"testing"
)

// Chat flags
func Test61_Chat_JoinFlags_JSON(t *testing.T) {
	const js = `{"join_to_send_messages":true,"join_by_request":true}`
	var ch Chat
	if err := json.Unmarshal([]byte(js), &ch); err != nil {
		t.Fatal(err)
	}
	if !ch.JoinToSendMessages || !ch.JoinByRequest {
		t.Fatal("join flags not populated")
	}
}

// User.added_to_attachment_menu
func Test61_User_AddedToAttachmentMenu_JSON(t *testing.T) {
	const js = `{"added_to_attachment_menu":true}`
	var u User
	if err := json.Unmarshal([]byte(js), &u); err != nil {
		t.Fatal(err)
	}
	if !u.AddedToAttachmentMenu {
		t.Fatal("flag not populated")
	}
}

// setWebhook: secret_token
func Test61_SetWebhook_SecretToken_Param(t *testing.T) {
	cfg := WebhookConfig{SecretToken: "s3cr3t"}
	p, err := cfg.params()
	if err != nil {
		t.Fatal(err)
	}
	if p["secret_token"] != "s3cr3t" {
		t.Fatalf("secret_token=%q", p["secret_token"])
	}
}

// createInvoiceLink: no chat_id, has prices
func Test61_CreateInvoiceLink_Params(t *testing.T) {
	cfg := CreateInvoiceLinkConfig{
		Title: "t", Description: "d", Payload: "p",
		ProviderToken: "prov", Currency: CurrencyRUB,
		Prices: []LabeledPrice{{Label: "X", Amount: 100}},
	}
	p, err := cfg.params()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := cfg.method(), "createInvoiceLink"; got != want {
		t.Fatal("method mismatch")
	}
	if _, ok := p["chat_id"]; ok {
		t.Fatal("chat_id must not be set")
	}
	if _, ok := p["prices"]; !ok {
		t.Fatal("prices missing")
	}
	if p["title"] != "t" || p["provider_token"] != "prov" {
		t.Fatal("required fields missing")
	}
}

func Test61_ThemeParams_SecondaryBG_JSON(t *testing.T) {
	const js = `{"secondary_bg_color":"#112233"}`
	var tp ThemeParams
	if err := json.Unmarshal([]byte(js), &tp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if tp.SecondaryBGColor != "#112233" {
		t.Fatalf("secondary_bg_color=%q", tp.SecondaryBGColor)
	}
}

func Test61_WebAppInitData_ChatAndCanSendAfter_JSON(t *testing.T) {
	const js = `{
		"query_id":"q",
		"user":{"id":1},
		"chat":{"id":-100123,"type":"supergroup","title":"X"},
		"can_send_after":5,
		"auth_date":1660000000,
		"hash":"h"
	}`
	var d WebAppInitData
	if err := json.Unmarshal([]byte(js), &d); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if d.Chat == nil || d.Chat.ID != -100123 {
		t.Fatalf("chat not parsed: %+v", d.Chat)
	}
	if d.CanSendAfter != 5 {
		t.Fatalf("can_send_after=%d", d.CanSendAfter)
	}
}

type apiResponse[T any] struct {
	Ok          bool   `json:"ok"`
	Result      T      `json:"result,omitempty"`
	Description string `json:"description,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
}

// Успешный ответ: ok=true и в result — ссылка (string).
func Test61_CreateInvoiceLink_ResponseOK(t *testing.T) {
	raw := []byte(`{
		"ok": true,
		"result": "https://t.me/invoice?start=AAABBBCCC"
	}`)

	var resp apiResponse[string]
	if err := json.Unmarshal(raw, &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if !resp.Ok {
		t.Fatal("expected ok=true")
	}
	if resp.Result == "" {
		t.Fatal("empty result link")
	}
	if want := "https://t.me/invoice?start=AAABBBCCC"; resp.Result != want {
		t.Fatalf("result mismatch: got %q want %q", resp.Result, want)
	}
}

// Ошибка API: ok=false + error_code + description.
func Test61_CreateInvoiceLink_ResponseError(t *testing.T) {
	raw := []byte(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: prices are required"
	}`)

	var resp apiResponse[string]
	if err := json.Unmarshal(raw, &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if resp.Ok {
		t.Fatal("expected ok=false")
	}
	if resp.ErrorCode == 0 {
		t.Fatal("expected non-zero error_code")
	}
	if resp.Description == "" {
		t.Fatal("expected description")
	}
}

// Кросс‑проверка сериализации/десериализации prices.
func Test61_CreateInvoiceLink_PricesJSONRoundTrip(t *testing.T) {
	cfg := NewCreateInvoiceLinkConfig(ChatID, "t", "d", "p", "prov", []LabeledPrice{{Label: "X", Amount: 100}, {Label: "Y", Amount: 2500}}, CurrencyRUB)

	p, err := cfg.params()
	if err != nil {
		t.Fatalf("params error: %v", err)
	}
	raw, ok := p["prices"]
	if !ok {
		t.Fatal("prices missing in params")
	}

	var got []LabeledPrice
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal prices failed: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("prices len mismatch: got %d want 2", len(got))
	}
	if got[0].Label != "X" || got[0].Amount != 100 {
		t.Fatalf("first price mismatch: %+v", got[0])
	}
	if got[1].Label != "Y" || got[1].Amount != 2500 {
		t.Fatalf("second price mismatch: %+v", got[1])
	}
}

// Бонус: проверяем, что метод соответствует createInvoiceLink (как в твоём первом тесте).
func Test61_CreateInvoiceLink_Method(t *testing.T) {
	cfg := CreateInvoiceLinkConfig{}
	if got, want := cfg.method(), "createInvoiceLink"; got != want {
		t.Fatalf("method mismatch: got %q want %q", got, want)
	}
}

// Проверяем разбор успешного ответа Telegram для sendInvoice.
func Test61_SendInvoice_ResponseOK(t *testing.T) {
	raw := []byte(`{
		"ok": true,
		"result": {
			"message_id": 42,
			"chat": { "id": 1234567890, "type": "private" },
			"date": 1700000000,
			"invoice": {
				"title": "t",
				"description": "d",
				"start_parameter": "start-abc",
				"currency": "RUB",
				"total_amount": 100
			}
		}
	}`)

	var resp apiResponse[Message]
	if err := json.Unmarshal(raw, &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if !resp.Ok {
		t.Fatal("expected ok=true")
	}
	if resp.Result.MessageID != 42 {
		t.Fatalf("message_id mismatch: %d", resp.Result.MessageID)
	}
	if resp.Result.Chat.ID != 1234567890 || resp.Result.Chat.Type != "private" {
		t.Fatalf("chat mismatch: %+v", resp.Result.Chat)
	}
	if resp.Result.Invoice == nil {
		t.Fatal("invoice missing")
	}
	inv := resp.Result.Invoice
	if inv.Title != "t" || inv.Currency != "RUB" || inv.TotalAmount != 100 {
		t.Fatalf("invoice fields mismatch: %+v", inv)
	}
}

// Проверяем, что параметры собираются корректно для sendInvoice.
func Test61_SendInvoice_Params(t *testing.T) {
	cfg := InvoiceConfig{
		BaseChat:      BaseChat{ChatID: (int64)(1)},
		Title:         "t",
		Description:   "d",
		Payload:       "p",
		ProviderToken: "prov",
		Currency:      CurrencyRUB,
		Prices:        []LabeledPrice{{Label: "X", Amount: 100}},
	}

	p, err := cfg.params()
	if err != nil {
		t.Fatalf("params error: %v", err)
	}

	if got, want := cfg.method(), "sendInvoice"; got != want {
		t.Fatalf("method mismatch: got %q want %q", got, want)
	}

	// chat_id обязателен
	if _, ok := p["chat_id"]; !ok {
		t.Fatal("chat_id missing")
	}

	// prices обязателен и должен быть JSON-массивом
	rawPrices, ok := p["prices"]
	if !ok {
		t.Fatal("prices missing")
	}
	var prices []LabeledPrice
	if err := json.Unmarshal([]byte(rawPrices), &prices); err != nil {
		t.Fatalf("prices not valid JSON: %v", err)
	}
	if len(prices) != 1 || prices[0].Label != "X" || prices[0].Amount != 100 {
		t.Fatalf("prices content mismatch: %+v", prices)
	}

	// обязательные строковые поля + валюта
	if p["title"] != "t" || p["provider_token"] != "prov" || p["currency"] != "RUB" {
		t.Fatal("required fields missing or mismatched")
	}
}

// createInvoiceLink
func Test61_Live_CreateInvoiceLink(t *testing.T) {

	bot := getBot(t)

	cfg := NewCreateInvoiceLinkConfig(ChatID, "Test Item", "Test description", "test-payload-123", "", []LabeledPrice{{Label: "Test item", Amount: 100}}, CurrencyXTR)

	link, err := bot.CreateInvoiceLink(cfg)
	if err != nil {
		t.Fatalf("CreateInvoiceLink failed: %v", err)
	}
	if link == "" {
		t.Fatal("empty link returned")
	}
	t.Logf("Invoice link: %s", link)
	bot.Send(NewMessage(ChatID, "Invoice link: "+link))
}

// sendInvoice
func Test61_Live_SendInvoice(t *testing.T) {

	bot := getBot(t)

	cfg := NewInvoice(ChatID, "Test Item", "Test description", "test-payload-123", "", "", CurrencyXTR, []LabeledPrice{{Label: "Test item", Amount: 100}})

	msg, err := bot.Send(cfg)
	if err != nil {
		t.Fatalf("SendInvoice failed: %v", err)
	}
	if msg.MessageID == 0 {
		t.Fatal("empty message_id")
	}
	if msg.Invoice == nil {
		t.Fatal("invoice missing in message")
	}
	if msg.Invoice.Currency != "XTR" {
		t.Fatalf("unexpected currency: %s", msg.Invoice.Currency)
	}
}

func Test61_WebAppInitData_Unmarshal_NewFields(t *testing.T) {
	raw := []byte(`{
        "chat": {"id": 42, "type": "private"},
        "can_send_after": 15
    }`)
	var v WebAppInitData
	if err := json.Unmarshal(raw, &v); err != nil {
		t.Fatal(err)
	}
	if v.Chat == nil || v.Chat.ID != 42 || v.Chat.Type != "private" {
		t.Fatalf("chat parsed wrong: %+v", v.Chat)
	}
	if v.CanSendAfter != 15 {
		t.Fatalf("can_send_after mismatch: %d", v.CanSendAfter)
	}
}

func Test61_ThemeParams_Unmarshal_SecondaryBG(t *testing.T) {
	raw := []byte(`{"secondary_bg_color":"#0f0f0f"}`)
	var tp ThemeParams
	if err := json.Unmarshal(raw, &tp); err != nil {
		t.Fatal(err)
	}
	if tp.SecondaryBGColor != "#0f0f0f" {
		t.Fatalf("secondary_bg_color mismatch: %s", tp.SecondaryBGColor)
	}
}
