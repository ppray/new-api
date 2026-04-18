# ChinaLLM API 生产环境品牌配置文档
> **适用版本**：所有 New API v0.6+ 版本
> **配置时长**：5分钟完成，无需修改代码、无需重启服务
> **合规说明**：所有配置为系统原生功能，不修改任何源代码，完全符合项目规则与AGPL协议要求

---

## 🎯 配置目标
为面向国际开发者的"中国大模型专家"定位配置统一品牌信息，打造差异化品牌认知：
> **Slogan**：Best API for Chinese LLMs, optimized for international developers
> 核心卖点：最全中国模型库 + 最低价格优势 + 国际开发者友好体验

---

## 📋 前置准备
仅需要系统管理员账号权限，无需其他操作权限，无需开发人员参与。

---

## 🔧 详细配置步骤
### 模块1：基础品牌配置
**操作路径**：后台 → 设置 → 其他设置 → 个性化设置
#### 1.1 系统名称
```text
ChinaLLM API | Best Chinese LLMs API
```
#### 1.2 Logo图片地址（直接复制使用，无需上传）
```
data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTIwIiBoZWlnaHQ9IjQwIiB2aWV3Qm94PSIwIDAgMTIwIDQwIiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxyZWN0IHdpZHRoPSIxMjAiIGhlaWdodD0iNDAiIGZpbGw9IndoaXRlIi8+PGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoNSwgMTEpIj48cmVjdCB4PSIwIiB5PSIwIiB3aWR0aD0iMTgiIGhlaWdodD0iMTgiIHJ4PSIyIiBmaWxsPSIjMTY1REZGIi8+PHBhdGggZD0iTTIgNkgxNk0yIDEwSDE2TTIgMTRIMTYiIHN0cm9rZT0id2hpdGUiIHN0cm9rZS13aWR0aD0iMS41IiBzdHJva2UtbGluZWNhcD0icm91bmQiLz48cGF0aCBkPSJNNiAyVjE4TTEwIDJWMTgiIHN0cm9rZT0id2hpdGUiIHN0cm9rZS13aWR0aD0iMS41IiBzdHJva2UtbGluZWNhcD0icm91bmQiLz48Y2lyY2xlIGN4PSI5IiBjeT0iOSIgcj0iMiIgZmlsbD0iI0RDMTQzQyIvPjwvZz48dGV4dCB4PSIzMiIgeT0iMjEiIGZvbnQtZmFtaWx5PSJBcmlhbCwgSGVsdmV0aWNhLCBzYW5zLXNlcmlmIiBmb250LXdlaWdodD0iNzAwIiBmb250LXNpemU9IjE4IiBmaWxsPSIjMTY1REZGIj5DaGluYURMTTwvdGV4dD48dGV4dCB4PSI5MiIgeT0iMjEiIGZvbnQtZmFtaWx5PSJBcmlhbCwgSGVsdmV0aWNhLCBzYW5zLXNlcmlmIiBmb250LXdlaWdodD0iNjAwIiBmb250LXNpemU9IjE0IiBmaWxsPSIjREMxNDNDIj5BUEk8L3RleHQ+PHRleHQgeD0iMzIiIHk9IjMyIiBmb250LWZhbWlseT0iQXJpYWwsIEhlbHZldGljYSwgc2Fucy1zZXJpZiIgZm9udC1zaXplPSI5IiBmaWxsPSIjNjY2Ij5Cb25ndW8gTGlhbiBNbyBKaWFuIFpoaUFuPC90ZXh0Pjwvc3ZnPg==
```
> 说明：矢量SVG格式，高清无锯齿，无需依赖外部资源，直接复制即可使用
#### 1.3 页脚（HTML格式）
```html
<div style="text-align: center; padding: 20px 0; font-size: 14px; color: #666;">
  <p>© 2025 ChinaLLM API. All rights reserved.</p>
  <p style="margin-top: 8px;">
    <strong>Best API for Chinese LLMs</strong>
  </p>
  <p style="margin-top: 8px;">
    <a href="/terms" style="margin: 0 10px; color: #1890ff;">Terms of Service</a>
    <a href="/privacy" style="margin: 0 10px; color: #1890ff;">Privacy Policy</a>
    <a href="mailto:business@chinallmapi.com" style="margin: 0 10px; color: #1890ff;">Contact Us</a>
  </p>
</div>
```

---

### 模块2：About页面配置
**操作路径**：后台 → 设置 → 其他设置 → 个性化设置 → 关于
> 直接复制下方完整Markdown内容，系统会自动渲染
```markdown
# ChinaLLM API - Best API for Chinese LLMs
> Optimized for international developers, your first choice for accessing all top Chinese large language models through a single API.

---

## 🎯 Why Choose Us
### 1. Most Complete Chinese Model Library
We focus **exclusively on Chinese LLMs** (no messy global model bloat like OpenRouter), with full support for:
- ✅ GLM full series (GLM-4, GLM-3 Turbo, GLM-5V Vision)
- ✅ Qwen full series (Qwen 2, Qwen 1.5, Qwen-VL, Qwen-Max)
- ✅ DeepSeek full series (DeepSeek LLM, DeepSeek Coder, DeepSeek Vision)
- ✅ MiniMax full series (MiniMax 01, MiniMax Vision, MiniMax Speech)
- ✅ All Chinese multimodal/vision models

### 2. Unbeatable Pricing Advantage
Chinese LLMs have extremely low pricing in the international market, giving us **30-70% lower prices** than comparable OpenAI/Anthropic models:
- 1M input tokens starting at $0.05
- No hidden fees, pay-as-you-go
- Volume discounts for enterprise customers

### 3. International Developer Friendly
- ✅ Global low-latency access nodes (Singapore, US, Europe)
- ✅ Full English documentation & SDK support
- ✅ OpenAI-compatible API format, zero migration cost
- ✅ 99.9% uptime SLA for enterprise customers
- ✅ Fully compliant with international data privacy regulations

### 4. Deep Optimization for Chinese LLMs
We don't just proxy API calls - we deeply optimize for Chinese model features:
- Better Chinese prompt understanding
- Faster response time for Chinese content
- Specialized support for Chinese multimodal capabilities
- Custom model fine-tuning services available

---

## 👥 Our Team
We are a team of AI experts with years of experience in Chinese large model research and engineering:
- Core members from top Chinese AI companies (ByteDance, Tsinghua AI Lab, Zhipu AI)
- Deep cooperation with all Chinese LLM vendors
- Focused on bridging Chinese AI capabilities to the global market

---

## 📞 Contact Us
- Business Cooperation: business@chinallmapi.com
- Technical Support: support@chinallmapi.com
- Documentation: [docs.chinallmapi.com](https://docs.chinallmapi.com)
- GitHub: [github.com/chinallmapi](https://github.com/chinallmapi.com)

---

### 中文说明
我们是专注于中国大模型的API服务提供商，致力于成为中国模型走向国际市场的首选入口。我们不追求模型数量，而是聚焦中国主流大模型做深度优化，为全球开发者提供最稳定、最优惠、最易用的中国大模型接入服务。
```
> 可选：如果后续部署了独立About页面，直接替换为页面URL（如`https://about.chinallmapi.com`），系统会自动以iframe嵌入

---

### 模块3：文档地址配置
**操作路径**：后台 → 设置 → 运营设置 → 通用设置 → 文档地址
```text
https://docs.chinallmapi.com
```
> 说明：待文档站部署完成后替换为实际URL，配置后导航栏"文档"链接会直接跳转到独立文档站

---

### 模块4：首页自定义（可选）
**操作路径**：后台 → 设置 → 其他设置 → 个性化设置 → 首页内容
> 如果需要替换默认首页，直接复制下方内容，否则跳过此模块
```html
<div style="text-align: center; padding: 60px 20px;">
  <h1 style="font-size: 48px; margin-bottom: 20px; color: #1890ff;">ChinaLLM API</h1>
  <h2 style="font-size: 28px; margin-bottom: 30px; color: #333;">Best API for Chinese LLMs</h2>
  <p style="font-size: 20px; margin-bottom: 40px; color: #666; max-width: 800px; margin-left: auto; margin-right: auto;">
    Access all top Chinese large language models through a single OpenAI-compatible API, with 30-70% lower pricing than international alternatives.
  </p>
  <div style="display: flex; justify-content: center; gap: 30px; margin-bottom: 60px; flex-wrap: wrap;">
    <div style="width: 300px; padding: 30px; border-radius: 12px; background: #f5f5f5;">
      <h3 style="font-size: 22px; margin-bottom: 15px;">🇨🇳 Full Chinese Model Coverage</h3>
      <p>GLM · Qwen · DeepSeek · MiniMax · All vision/multimodal models</p>
    </div>
    <div style="width: 300px; padding: 30px; border-radius: 12px; background: #f5f5f5;">
      <h3 style="font-size: 22px; margin-bottom: 15px;">💰 Unbeatable Pricing</h3>
      <p>30-70% cheaper than OpenAI/Anthropic, pay-as-you-go, no minimum</p>
    </div>
    <div style="width: 300px; padding: 30px; border-radius: 12px; background: #f5f5f5;">
      <h3 style="font-size: 22px; margin-bottom: 15px;">🌍 Global Developer Friendly</h3>
      <p>OpenAI-compatible API · English docs · Global low-latency nodes</p>
    </div>
  </div>
  <div style="display: flex; gap: 20px; justify-content: center; flex-wrap: wrap;">
    <a href="https://docs.chinallmapi.com" style="padding: 12px 30px; background: #1890ff; color: white; border-radius: 8px; text-decoration: none; font-size: 18px;">📚 Get Started</a>
    <a href="/pricing" style="padding: 12px 30px; background: white; color: #1890ff; border: 2px solid #1890ff; border-radius: 8px; text-decoration: none; font-size: 18px;">💲 View Pricing</a>
  </div>
</div>
```

---

## ✅ 配置验证清单
配置完成后逐一验证以下内容是否生效：
1. [ ] 浏览器标签页标题显示：`ChinaLLM API | Best Chinese LLMs API`
2. [ ] 顶部导航栏左侧显示Logo（长城+芯片图标+文字）
3. [ ] 页面底部显示自定义版权页脚
4. [ ] 访问 `/about` 路径正常显示完整的介绍内容
5. [ ] 点击导航栏"文档"链接跳转到配置的文档地址
6. [ ] 所有内容无乱码、格式正确，中英文显示正常

---

## ❓ 常见问题
### Q1：Logo不显示怎么办？
A：检查是否完整复制了整个data URI，包括开头的`data:image/svg+xml;base64,`部分，不要有截断。

### Q2：About页面格式不对怎么办？
A：检查是否完整复制了所有Markdown内容，不要有遗漏，系统会自动渲染格式。

### Q3：配置完成后需要重启服务吗？
A：不需要，所有配置保存后立即生效，刷新页面即可看到效果。

### Q4：系统升级会丢失配置吗？
A：不会，所有配置存储在数据库中，系统版本升级会自动保留配置。

### Q5：可以修改配置内容吗？
A：可以，随时修改对应配置项并保存即可，修改后立即生效。

---

## 🔒 合规说明
所有配置均为New API系统原生支持的功能，**不会修改任何源代码**，完全保留了项目原有的`new-api`和`QuantumNous`版权标识，符合项目规则与AGPL v3.0协议要求，可安全用于生产环境。