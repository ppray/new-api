---
layout: default
title: API Reference
---

# API Reference

OpenRouterA provides two compatible API formats:

- **OpenAI format**: `/v1/chat/completions`, `/v1/embeddings`, `/v1/images/generations`
- **Anthropic format**: `/v1/messages`

All endpoints require authentication via API key.

---

## Authentication

Include your API key in the `Authorization` header:

```
Authorization: Bearer sk-your-key-here
```

## Base URL

```
https://openroutera.com
```

---

## Chat Completions

```
POST /v1/chat/completions
```

OpenAI-compatible chat completion endpoint.

### Request

```json
{
  "model": "glm-5.1",
  "messages": [
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "Hello!"}
  ],
  "temperature": 0.7,
  "max_tokens": 1024,
  "stream": false
}
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `model` | string | Yes | Model ID (e.g., `glm-5.1`, `deepseek-v3.2`, `glm-4.7`) |
| `messages` | array | Yes | Array of message objects with `role` and `content` |
| `temperature` | float | No | Sampling temperature (0-2). Default: 1 |
| `max_tokens` | integer | No | Maximum tokens to generate |
| `top_p` | float | No | Nucleus sampling parameter |
| `stream` | boolean | No | Enable streaming. Default: false |
| `stop` | string/array | No | Stop sequences |

### Response

```json
{
  "id": "chatcmpl-abc123",
  "object": "chat.completion",
  "created": 1713523200,
  "model": "glm-5.1",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": "Hello! How can I help you today?"
      },
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 10,
    "completion_tokens": 8,
    "total_tokens": 18
  }
}
```

### Streaming

Set `"stream": true` to receive Server-Sent Events:

```json
{
  "model": "glm-5.1",
  "messages": [{"role": "user", "content": "Hello!"}],
  "stream": true
}
```

Response is a stream of `data:` prefixed JSON chunks, ending with `data: [DONE]`.

---

## Messages (Anthropic Format)

```
POST /v1/messages
```

Anthropic-compatible messages endpoint.

### Request

```json
{
  "model": "glm-5.1",
  "max_tokens": 1024,
  "messages": [
    {"role": "user", "content": "Hello!"}
  ]
}
```

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `model` | string | Yes | Model ID |
| `messages` | array | Yes | Array of message objects |
| `max_tokens` | integer | Yes | Maximum tokens to generate |
| `temperature` | float | No | Sampling temperature |
| `system` | string | No | System prompt |
| `stream` | boolean | No | Enable streaming |

### Response

```json
{
  "id": "msg-abc123",
  "type": "message",
  "role": "assistant",
  "content": [
    {
      "type": "text",
      "text": "Hello! How can I help you today?"
    }
  ],
  "model": "glm-5.1",
  "stop_reason": "end_turn",
  "usage": {
    "input_tokens": 10,
    "output_tokens": 8
  }
}
```

---

## Embeddings

```
POST /v1/embeddings
```

### Request

```json
{
  "model": "<your-embedding-model>",
  "input": "Hello, world!"
}
```

### Response

```json
{
  "object": "list",
  "data": [
    {
      "object": "embedding",
      "index": 0,
      "embedding": [0.0023, -0.0094, ...]
    }
  ],
  "model": "<your-embedding-model>",
  "usage": {
    "prompt_tokens": 2,
    "total_tokens": 2
  }
}
```

---

## Image Generation

```
POST /v1/images/generations
```

### Request

```json
{
  "model": "<your-image-model>",
  "prompt": "A cute cat wearing a hat",
  "n": 1,
  "size": "1024x1024"
}
```

---

## Error Handling

All errors follow the OpenAI error format:

```json
{
  "error": {
    "message": "Invalid API key",
    "type": "authentication_error",
    "code": "invalid_api_key"
  }
}
```

### Common HTTP Status Codes

| Code | Meaning |
|------|---------|
| 200 | Success |
| 400 | Bad request (invalid parameters) |
| 401 | Invalid or missing API key |
| 402 | Insufficient credits |
| 404 | Model not found |
| 429 | Rate limit exceeded |
| 500 | Internal server error |

---

## Rate Limits

Rate limits vary by account tier. If you hit a rate limit, you'll receive a `429` status code with a `Retry-After` header.

---

## SDK Compatibility

Any OpenAI-compatible SDK works out of the box:

| SDK | Configuration |
|-----|--------------|
| Python `openai` | `base_url="https://openroutera.com/v1"` |
| Node.js `openai` | `baseURL: 'https://openroutera.com/v1'` |
| Python `anthropic` | `base_url="https://openroutera.com"` |
| LangChain | `ChatOpenAI(openai_api_base="https://openroutera.com/v1")` |
| LlamaIndex | `Settings.api_base = "https://openroutera.com/v1"` |

No code changes needed — just swap the base URL and API key.
