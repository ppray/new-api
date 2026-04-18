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

| Provider | Models | Price |
|----------|--------|-------|
| **Zhipu** | glm-5.1, glm-5-turbo, glm-4.7 | $0.010 - $0.02/req |
| **DeepSeek** | deepseek-v3.2 | $0.01/req |
| **Doubao (ByteDance)** | doubao-seed-2.0-pro, doubao-seed-2.0-code, doubao-seed-2.0-lite, doubao-seed-code | $0.01/req |
| **Moonshot (Kimi)** | kimi-k2.5 | $0.01/req |
| **MiniMax** | MiniMax-M2.5 | $0.01/req |

> Full list available at [openroutera.com/pricing](https://openroutera.com/pricing).

## Pricing

**Per-request flat rate.** No token counting, no surprises.

- Same price regardless of context length (1K or 200K+)
- Models starting at **$0.01/request**
- GLM 5.1 at **$0.02/request** — up to **12x cheaper** than per-token competitors at 200K+ context

Visit [openroutera.com/pricing](https://openroutera.com/pricing) for full pricing details.

## Why OpenRouterA?

- **Simplest pricing** — flat per-request, no token math
- **Best value for large context** — 12x cheaper at 200K+ tokens
- **Full Chinese model coverage** — every major Chinese LLM provider
- **Zero migration cost** — OpenAI & Anthropic compatible, just swap the URL
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
