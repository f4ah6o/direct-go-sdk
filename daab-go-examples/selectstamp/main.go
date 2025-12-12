package main

import (
	"context"
	cryptorand "crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/f4ah6o/direct-go-sdk/daab-go/bot"
	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
)

const (
	menuQuestion = "機能メニュー"
	mirasapoBase = "https://mirasapo-plus.go.jp/jirei-api"
)

var menuOptions = []string{
	"uuid 占い",
	"ミラサポplus事例表示",
}

type menuTracker struct {
	mu             sync.Mutex
	lastQuestionID string
}

func (m *menuTracker) set(id string) {
	m.mu.Lock()
	m.lastQuestionID = id
	m.mu.Unlock()
}

func (m *menuTracker) matches(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return id != "" && id == m.lastQuestionID
}

type selectContent struct {
	Question  string
	Options   []string
	Response  *int
	InReplyTo string
}

type caseStudy struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Year     string `json:"year"`
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
}

type caseStudySearchResult struct {
	Total int         `json:"total"`
	Items []caseStudy `json:"items"`
}

type selectAnswerSummary struct {
	Option  string
	Count   int
	UserIDs []string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	debugServer := os.Getenv("DEBUG_SERVER")
	if debugServer == "" {
		debugServer = "http://localhost:9999"
	}
	direct.EnableDebugServer(debugServer)

	robot := bot.New(
		bot.WithName("selectbot"),
	)
	tracker := &menuTracker{}

	// Send the select menu when asked.
	robot.Respond("(menu|メニュー)$", func(ctx context.Context, res bot.Response) {
		id, err := res.SendSelect(menuQuestion, menuOptions)
		if err != nil {
			log.Printf("Error sending select stamp: %v", err)
			_ = res.Send("セレクトスタンプの送信に失敗しました。トークIDや権限を確認してください。")
			return
		}
		tracker.set(id)
		if err := res.Send("機能メニューを送信しました。選択して動作を試してみてください。"); err != nil {
			log.Printf("Error sending menu notice: %v", err)
		}

		// After a short delay, post the current answer summary.
		go func(roomID, questionID string) {
			time.Sleep(5 * time.Second)
			summaries, err := fetchSelectAnswerSummary(res.Robot, questionID)
			if err != nil {
				log.Printf("Failed to fetch answer summary: %v", err)
				return
			}
			summaryText := formatSelectSummary(summaries)
			if summaryText != "" {
				if err := res.Robot.SendText(roomID, summaryText); err != nil {
					log.Printf("Error sending summary: %v", err)
				}
			}
		}(res.RoomID(), id)
	})

	// Echo selection results and trigger features.
	robot.Hear(".*", func(ctx context.Context, res bot.Response) {
		handleSelectAction(ctx, res, tracker)
	})

	if err := robot.Run(context.Background()); err != nil {
		log.Fatalf("Bot error: %v", err)
	}
}

func handleSelectAction(ctx context.Context, res bot.Response, tracker *menuTracker) {
	// Debug: show message type
	log.Printf("[SELECT DEBUG] Message type: %d, Text: %q", res.Message.Type, res.Message.Text)

	content, err := extractSelectContent(res.Message)
	if err != nil {
		log.Printf("[SELECT DEBUG] extractSelectContent error: %v", err)
		return
	}
	if content == nil {
		log.Printf("[SELECT DEBUG] content is nil")
		return
	}
	if content.Response == nil {
		log.Printf("[SELECT DEBUG] content.Response is nil, Question=%q, Options=%v, InReplyTo=%q", content.Question, content.Options, content.InReplyTo)
		return
	}

	log.Printf("[SELECT DEBUG] Got response: idx=%d, Question=%q, Options=%v, InReplyTo=%q", *content.Response, content.Question, content.Options, content.InReplyTo)

	// Ensure this is the menu we sent.
	if content.Question != menuQuestion && !tracker.matches(content.InReplyTo) {
		log.Printf("[SELECT DEBUG] Skipping: Question mismatch and tracker doesn't match")
		return
	}

	idx := *content.Response
	choice := optionAt(content.Options, idx)
	switch choice {
	case menuOptions[0]:
		handleUUIDFortune(ctx, res)
	case menuOptions[1]:
		handleMirasapoCase(ctx, res)
	default:
		_ = res.Send(fmt.Sprintf("選択肢 %d を受信しました。", idx))
	}
}

func optionAt(options []string, idx int) string {
	if idx >= 0 && idx < len(options) {
		return options[idx]
	}
	return ""
}

func extractSelectContent(msg direct.ReceivedMessage) (*selectContent, error) {
	// Wire types: 502 = select stamp, 503 = select reply
	msgType := int(msg.Type)
	if msgType != direct.WireTypeSelect && msgType != direct.WireTypeSelectReply {
		return nil, fmt.Errorf("not a select message (type=%d)", msgType)
	}

	contentMap, err := pullContentMap(msg)
	if err != nil || contentMap == nil {
		return nil, err
	}

	// Debug: show all keys in contentMap
	log.Printf("[SELECT DEBUG] contentMap keys: %v", getMapKeys(contentMap))
	log.Printf("[SELECT DEBUG] contentMap full: %+v", contentMap)

	sc := &selectContent{
		Question: stringValue(contentMap["question"]),
		Options:  stringSlice(contentMap["options"]),
	}

	if resp, ok := intValue(contentMap["response"]); ok {
		sc.Response = &resp
	}
	if inReply, ok := contentMap["in_reply_to"]; ok {
		sc.InReplyTo = fmt.Sprintf("%v", inReply)
	}

	return sc, nil
}

func getMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func pullContentMap(msg direct.ReceivedMessage) (map[string]interface{}, error) {
	if content, ok := msg.Content.(map[string]interface{}); ok {
		return content, nil
	}

	if len(msg.Raw) == 0 {
		return nil, nil
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(msg.Raw, &raw); err != nil {
		return nil, err
	}

	if content, ok := raw["content"].(map[string]interface{}); ok {
		return content, nil
	}

	return nil, nil
}

func stringSlice(v interface{}) []string {
	raw, ok := v.([]interface{})
	if !ok {
		return nil
	}
	out := make([]string, 0, len(raw))
	for _, item := range raw {
		out = append(out, stringValue(item))
	}
	return out
}

func stringValue(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%v", v)
}

func intValue(v interface{}) (int, bool) {
	switch n := v.(type) {
	case int:
		return n, true
	case int8:
		return int(n), true
	case int16:
		return int(n), true
	case int32:
		return int(n), true
	case int64:
		return int(n), true
	case uint:
		return int(n), true
	case uint8:
		return int(n), true
	case uint16:
		return int(n), true
	case uint32:
		return int(n), true
	case uint64:
		return int(n), true
	case float32:
		return int(n), true
	case float64:
		return int(n), true
	case json.Number:
		i, err := n.Int64()
		if err != nil {
			return 0, false
		}
		return int(i), true
	default:
		return 0, false
	}
}

func handleUUIDFortune(ctx context.Context, res bot.Response) {
	uuidBytes, err := newUUIDv4()
	if err != nil {
		log.Printf("Failed to generate UUID: %v", err)
		return
	}

	fortunes := []string{
		"大吉",
		"中吉",
		"小吉",
		"吉",
		"末吉",
		"凶",
	}

	idx := int(uuidBytes[0]) % len(fortunes)
	lines := []string{
		"uuid 占いの結果です。",
		fmt.Sprintf("UUID: %s", formatUUID(uuidBytes)),
		fmt.Sprintf("運勢: %s", fortunes[idx]),
	}
	if err := res.Send(strings.Join(lines, "\n")); err != nil {
		log.Printf("Error sending fortune: %v", err)
	}
}

func newUUIDv4() ([]byte, error) {
	b := make([]byte, 16)
	if _, err := cryptorand.Read(b); err != nil {
		return nil, err
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return b, nil
}

func formatUUID(b []byte) string {
	if len(b) != 16 {
		return ""
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func handleMirasapoCase(ctx context.Context, res bot.Response) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client := &http.Client{Timeout: 10 * time.Second}

	cs, err := fetchRandomCaseStudy(ctx, client)
	if err != nil {
		if sendErr := res.Send(fmt.Sprintf("事例ナビの取得に失敗しました: %v", err)); sendErr != nil {
			log.Printf("Error sending failure message: %v", sendErr)
		}
		return
	}

	summary := summarizeCaseStudy(*cs)
	if err := res.Send(summary); err != nil {
		log.Printf("Error sending case study: %v", err)
	}
}

func fetchRandomCaseStudy(ctx context.Context, client *http.Client) (*caseStudy, error) {
	first, err := searchCaseStudies(ctx, client, 0)
	if err != nil {
		return nil, err
	}
	if len(first.Items) == 0 {
		return nil, fmt.Errorf("no case studies available")
	}
	if first.Total <= 1 {
		return &first.Items[0], nil
	}

	offset := rand.Intn(first.Total)
	if offset == 0 {
		return &first.Items[0], nil
	}

	result, err := searchCaseStudies(ctx, client, offset)
	if err != nil {
		return nil, err
	}
	if len(result.Items) == 0 {
		return &first.Items[0], nil
	}
	return &result.Items[0], nil
}

func searchCaseStudies(ctx context.Context, client *http.Client, offset int) (*caseStudySearchResult, error) {
	url := fmt.Sprintf("%s/case_studies?limit=1&offset=%d", mirasapoBase, offset)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API status %d", resp.StatusCode)
	}

	var result caseStudySearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func summarizeCaseStudy(cs caseStudy) string {
	lines := []string{
		"ミラサポplus 事例ナビからランダム紹介",
		fmt.Sprintf("タイトル: %s", cs.Title),
	}
	if cs.Year != "" {
		lines = append(lines, fmt.Sprintf("制度利用年: %s", cs.Year))
	}
	if cs.Location.Name != "" {
		lines = append(lines, fmt.Sprintf("地域: %s", cs.Location.Name))
	}
	if trimmed := trimSummary(cs.Summary, 140); trimmed != "" {
		lines = append(lines, "概要: "+trimmed)
	}
	lines = append(lines, fmt.Sprintf("データID: %s", cs.ID))
	lines = append(lines, fmt.Sprintf("API: %s/case_studies/%s", mirasapoBase, cs.ID))
	return strings.Join(lines, "\n")
}

func trimSummary(s string, limit int) string {
	s = strings.TrimSpace(s)
	if limit <= 0 || len([]rune(s)) <= limit {
		return s
	}
	runes := []rune(s)
	return string(runes[:limit]) + "..."
}

func fetchSelectAnswerSummary(robot *bot.Robot, questionID string) ([]selectAnswerSummary, error) {
	if questionID == "" {
		return nil, nil
	}

	// Convert questionID to uint64 as API expects numeric ID
	qid, err := strconv.ParseUint(questionID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid question ID: %w", err)
	}

	result, err := robot.Call("get_action", []interface{}{qid})
	if err != nil {
		return nil, err
	}

	data, ok := result.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T", result)
	}

	rawResponses, ok := data["responses"].([]interface{})
	if !ok {
		return nil, nil
	}

	summaries := make([]selectAnswerSummary, 0, len(rawResponses))
	for _, r := range rawResponses {
		respMap, ok := r.(map[string]interface{})
		if !ok {
			continue
		}
		option := stringValue(respMap["content"])
		count, _ := intValue(respMap["count"])
		userIDs := stringSlice(respMap["user_ids"])
		summaries = append(summaries, selectAnswerSummary{
			Option:  option,
			Count:   count,
			UserIDs: userIDs,
		})
	}

	return summaries, nil
}

func formatSelectSummary(summaries []selectAnswerSummary) string {
	if len(summaries) == 0 {
		return "アクションスタンプ回答状況: まだ回答がありません。"
	}

	lines := []string{"アクションスタンプ回答状況"}
	for _, s := range summaries {
		line := fmt.Sprintf("%s: %d件", s.Option, s.Count)
		if len(s.UserIDs) > 0 {
			line = fmt.Sprintf("%s (%s)", line, strings.Join(s.UserIDs, ", "))
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}
