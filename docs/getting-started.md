---
layout: default
title: Getting Started
---

# Getting Started

Set up OpenRouterA in under 2 minutes.

## Step 1: Create an Account

Go to [openroutera.com/console](https://openroutera.com/console) and sign up for a free account.

## Step 2: Get Your API Key

1. Navigate to **Tokens** in the dashboard
2. Click **Create Token**
3. Copy your API key (starts with `sk-`)

> **Important:** Keep your API key secret. Do not commit it to version control.

## Step 3: Configure Your Client

### OpenAI SDK (Python)

```python
from openai import OpenAI

client = OpenAI(
    api_key="sk-your-key-here",
    base_url="https://openroutera.com/v1"
)

response = client.chat.completions.create(
    model="glm-5.1",
    messages=[{"role": "user", "content": "Hello, how are you?"}]
)

print(response.choices[0].message.content)
```

### OpenAI SDK (Node.js)

```javascript
import OpenAI from 'openai';

const client = new OpenAI({
  apiKey: 'sk-your-key-here',
  baseURL: 'https://openroutera.com/v1',
});

const response = await client.chat.completions.create({
  model: 'glm-5.1',
  messages: [{ role: 'user', content: 'Hello, how are you?' }],
});

console.log(response.choices[0].message.content);
```

### Anthropic SDK (Python)

```python
import anthropic

client = anthropic.Anthropic(
    api_key="sk-your-key-here",
    base_url="https://openroutera.com"
)

response = client.messages.create(
    model="glm-5.1",
    max_tokens=1024,
    messages=[{"role": "user", "content": "Hello, how are you?"}]
)

print(response.content[0].text)
```

### cURL

```bash
curl https://openroutera.com/v1/chat/completions \
  -H "Authorization: Bearer sk-your-key-here" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "glm-5.1",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
```

## Step 4: Choose a Model

See [Supported Models](/#supported-models) for the full list. Popular choices:

| Model | Best For | Price |
|-------|----------|-------|
| `glm-5.1` | General purpose, Chinese tasks | $0.02/req |
| `glm-4.7` | General purpose, balanced cost | $0.01/req |
| `deepseek-v3.2` | Reasoning, code | $0.01/req |
| `doubao-seed-2.0-pro` | General purpose, cost-effective | $0.01/req |
| `kimi-k2.5` | Long context | $0.01/req |
| `MiniMax-M2.5` | General purpose | $0.01/req |

## Step 5: Add Credits

1. Go to **Top Up** in the dashboard
2. Add funds to your account
3. Start making API calls — each request deducts a flat fee

That's it! You're ready to go.

## Next Steps

- [API Reference](/api-reference.html) — Full endpoint documentation
- [Pricing](https://openroutera.com/pricing) — Detailed pricing information
