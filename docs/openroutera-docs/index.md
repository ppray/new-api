---
layout: default
title: OpenRouterA - Best API for Chinese LLMs
---

# OpenRouterA API

**Best API for Chinese LLMs** — Access all top Chinese large language models through a single OpenAI & Anthropic compatible API.

Flat-rate pricing per request. No token math. GLM 5.1 at just **$0.02**/req — up to **12x cheaper** than competitors at 200K+ context.

## Quick Start

```bash
# 1. Get your API key at https://openroutera.com/console
# 2. Set the base URL
export OPENAI_API_BASE=https://openroutera.com/v1
export OPENAI_API_KEY=sk-your-key-here

# 3. Make a request
curl https://openroutera.com/v1/chat/completions \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "glm-5.1",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
```

That's it. No SDK changes needed — just swap the base URL.

## Supported Endpoints

| Endpoint | Description |
|----------|-------------|
| `/v1/chat/completions` | OpenAI-compatible chat completions |
| `/v1/messages` | Anthropic-compatible messages API |
| `/v1/embeddings` | Text embeddings |
| `/v1/images/generations` | Image generation |

## Supported Models

| Provider | Models |
|----------|--------|
| **GLM (Zhipu)** | GLM-5.1, GLM-4, GLM-4V, GLM-3 Turbo |
| **DeepSeek** | DeepSeek V3, DeepSeek R1, DeepSeek Coder |
| **Qwen (Alibaba)** | Qwen-Max, Qwen 2.5, Qwen-VL |
| **MiniMax** | MiniMax 01, MiniMax Vision, MiniMax Speech |
| **Moonshot** | Moonshot v1 (Kimi) |
| **Wenxin (Baidu)** | ERNIE 4.0, ERNIE 3.5 |
| **Hunyuan (Tencent)** | Hunyuan Turbo, Hunyuan Vision |
| **Spark (iFlytek)** | Spark 4.0 Ultra |
| **Volcengine (ByteDance)** | Doubao series |
| **Claude** | Claude 4 Opus, Claude 4 Sonnet |
| **OpenAI** | GPT-4o, GPT-4o-mini |
| **Gemini** | Gemini 2.5 Pro, Gemini 2.5 Flash |
| **Grok** | Grok 3 |
| **+ 30 more** | Llama, Mistral, Cohere, ... |

## Pricing

**Per-request flat rate.** No token counting, no surprises.

- Same price regardless of context length (1K or 200K+)
- Premium models starting at **$0.02/request**
- Up to **12x cheaper** than per-token competitors at large context sizes

Visit [openroutera.com/pricing](https://openroutera.com/pricing) for full pricing details.

## Why OpenRouterA?

- **Simplest pricing** — flat per-request, no token math
- **Best value for large context** — 12x cheaper at 200K+ tokens
- **Full Chinese model coverage** — every major Chinese LLM provider
- **Zero migration cost** — OpenAI & Anthropic compatible, just swap the URL
- **Global low-latency** — Singapore, US, Europe nodes
- **No vendor lock-in** — standard API format, switch anytime

## Documentation

- [Getting Started](/getting-started.html) — Setup guide and first API call
- [API Reference](/api-reference.html) — Full endpoint documentation
- [Pricing](https://openroutera.com/pricing) — Pricing details
- [Terms of Service](/terms.html)
- [Privacy Policy](/privacy.html)

## Contact

- Email: [openroutera@163.com](mailto:openroutera@163.com)
- Documentation: [docs.openroutera.com](https://docs.openroutera.com)
- Main site: [openroutera.com](https://openroutera.com)
