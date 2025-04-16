# Soup (Turtle Soup) - AI Riddle Game

LLM "Turtle Soup," a riddle game where players guess a story by asking "YES/NO" questions. The AI acts as the Dungeon Master (DM), answering only "YES," "NO," or "YES OR NO."

---

## Features

1. **Auto-Generated Riddles**  
   - Uses LLM  to generate riddles (for convenience, though AI-generated riddles are sick).  
   - Focus on gameplay rather than riddle quality.  

2. **AI Dungeon Master**  
   - The LLM strictly answers "YES/NO/YES OR NO" to player questions.  
   - Players deduce the story behind the riddle through logical questioning.  

3. **Deployment**  
   - Dockerized for easy setup.  
   - Configure `DEEPSEEK_API_KEY` and `DEEPSEEK_BASE_URI` in `.env`.  
   - Support for other LLMs is a work in progress (WIP).  

4. **CLI for Data Management**  
   - The CLI is designed for managing riddle data (e.g., adding, updating, or deleting riddles).  
   - Not intended for gameplay interaction.  

---

## Quick Start

1. Clone the repo:  
   ```bash
   git clone https://github.com/your-repo/soup.git
   cd soup
   ```

2. Configure `.env`:  
   ```bash
   cp .env.example .env
   # Edit .env with your DEEPSEEK_API_KEY and DEEPSEEK_BASE_URI
   ```

3. Run with Docker:  
   ```bash
   docker-compose up --watch
   ```

## Roadmap

- [ ] Support more LLMs.  
- [ ] Add multiplayer mode.  
- [ ] Implement full game logic.  
- [ ] i18n Prompt management
