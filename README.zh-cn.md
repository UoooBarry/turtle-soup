# Soup (海龟汤) - AI 谜题游戏

[English](./README.md) [中文](./README.zh-cn.md)

LLM“海龟汤”谜题游戏。玩家通过提问“是/否”问题来推理故事。

---

## 功能简介

1. **自动生成谜题**  
   - LLM生成谜题（仅为方便，AI 生成的谜题质量sucks）。  

2. **AI 主持人**  
   - 大语言模型严格回答“是/否/是或否”来引导玩家。  
   - 玩家通过逻辑提问推理出谜题背后的故事。  

3. **部署**  
   - 使用 Docker 快速部署。  
   - 需在 `.env` 中配置 `DEEPSEEK_API_KEY` 和 `DEEPSEEK_BASE_URI`。  
   - 其他大语言模型的支持正在开发中（WIP）。  

---

## 快速开始

1. 克隆仓库：  
   ```bash
   git clone https://github.com/your-repo/soup.git
   cd soup
   ```

2. 配置 `.env`：  
   ```bash
   cp .env.example .env
   # 在 .env 中填写你的 DEEPSEEK_API_KEY 和 DEEPSEEK_BASE_URI
   ```

3. 使用 Docker 运行：  
   ```bash
   docker-compose up
   cd vue && bun dev
   ```

4. 网页版`http://localhost:5173`。  

---

## Roeadmap

- [ ] 支持更多大语言模型。  
- [ ] 添加多人游戏模式。  
- [ ] 完成游戏逻辑。  
- [ X ] i18n提示词管理。
