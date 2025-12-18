// Package webhook exposes a lightweight client and message schema for forwarding
// Direct4B events to external workflow engines such as n8n via HTTP webhooks.
// It pairs incoming chat data with bot metadata, posts it to a configured
// endpoint using Client.Send, and parses structured actions (reply, send,
// send_select, etc.) back from the workflow in WebhookResponse. Helper types
// like WebhookPayload and MessageTypeToName keep payloads consistent with the
// rest of daab-go while remaining framework-agnostic for custom integrations.
package webhook
