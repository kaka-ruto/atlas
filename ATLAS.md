# Atlas - The AI Pair Programmer

## What Atlas Is

Atlas is a **100x productivity amplifier for software developers** that preserves the joy of coding. It's a terminal-native AI pair programmer built on top of lazygit, designed to make developers dramatically more productive without replacing their creative decision-making.

**Core Philosophy:** *"The AI thinks, you create."*

Atlas handles all the cognitive overhead that slows developers down - information gathering, system analysis, debugging investigations, and expert guidance - leaving developers free to focus purely on creative problem-solving and implementation.

## Why Atlas Exists

### The Problem with Current AI Coding Tools

**Existing tools like Claude Code try to replace developers:**
- They write code for you
- They make architectural decisions
- They reduce developers to code reviewers
- They take away the satisfaction of problem-solving

**The Real Developer Experience:**
- 80% of development time is spent gathering information, not writing code
- Developers constantly context-switch between files, logs, databases, documentation
- Complex debugging requires correlating data across multiple systems
- Best practices and patterns are hard to remember and apply consistently

### Atlas's Different Approach

**Atlas amplifies human capabilities instead of replacing them:**
- ✅ Provides expert guidance and step-by-step instructions
- ✅ Eliminates information gathering overhead
- ✅ Offers intelligent analysis and insights
- ✅ Teaches best practices through practical application
- ❌ Never writes code automatically
- ❌ Never makes decisions for the developer
- ❌ Never takes control away from the human

## How Atlas Works

### Architecture Overview

**Dual-Mode System:**
- **Git Mode**: Pure lazygit experience for version control
- **AI Mode**: Dedicated pair programming interface with AI assistance

**Built on Proven Foundation:**
- Forked from lazygit (preserves 100% of git functionality)
- Go backend for performance and reliability
- Terminal-native interface (works anywhere)
- BYOK (Bring Your Own Key) for AI models

### The Atlas Experience

#### Git Mode (Press 'G')
```
┌─ Status ─────────┐┌─ Unstaged Changes ────────────────┐
│ On branch main   ││ @@ -1,3 +1,6 @@                  │
│ 2 files changed  ││  class User < ApplicationRecord   │
├─ Files ──────────┤│ +  validates :email               │
│ M user.rb        ││                                   │
│ M routes.rb      ││                                   │
├─ Branches ───────┤├─ Command Log ─────────────────────┤
│ * main           ││ git add user.rb                   │
│   feature/auth   ││ git status --porcelain            │
├─ Commits ────────┤│                                   │
│ abc123 Add user  ││                                   │
└──────────────────┘└───────────────────────────────────┘
┌─ Keybindings ─────────────────────────────────────────┐
│ space: stage, c: commit, p: push, A: AI mode          │
└───────────────────────────────────────────────────────┘
```

#### AI Mode (Press 'A')
```
┌─ Project Status ──────────┐┌─ Current File Context ─────────────────┐
│ 🎯 Task: Add validations  ││ app/models/user.rb                     │
│ ⏱️  1h 23m active        ││ class User < ApplicationRecord         │
│ 📊 Progress: ████████░░ 80%││   validates :email, presence: true    │
├─ Smart File List ─────────┤├─ Step-by-Step Guide ───────────────────┤
│ 🎯 Current context:       ││ 📋 Add Email Validation:               │
│   • user.rb          ←─── ││ 1. Add format validation:              │
│   • users_controller.rb   ││    format: { with: URI::MailTo::EMAIL...│
│   • user_spec.rb          ││ 2. Add uniqueness constraint:          │
├─ System Health ───────────┤│    uniqueness: { case_sensitive: false}│
│ 🧪 Tests: 12/15 passing   ││ 3. Update database migration:          │
│ 📊 DB: Connected          ││    add_index :users, :email, unique:...│
└───────────────────────────┘└─────────────────────────────────────────┘
┌─ AI Teaching Session ──────────────────────────────────┐
│ 💬 You: How do I add proper email validation?          │
│ 🤖 Atlas: I'll walk you through Rails email validation │
│          step by step...                              │
└─────────────────────────────────────────────────────────┘
```

### Key Features

#### 1. Intelligent File Discovery
- **No file trees or manual searching**
- Smart file lists based on current context
- Git-aware file relationships
- Task-focused file suggestions
- MVC pattern recognition

#### 2. Read-Only System Intelligence
- **Database querying** (SELECT only)
- **Log analysis** and pattern recognition
- **Process monitoring** and system health
- **Dependency source code** reading and explanation
- **Code path tracing** and execution flow analysis

#### 3. Instructional AI Guidance
- **Step-by-step problem solving**
- **Teaching explanations** with every solution
- **Best practice guidance** specific to your codebase
- **Security and performance** recommendations
- **Debugging workflows** with clear instructions

#### 4. Multi-Model AI Support
- Anthropic Claude (primary)
- OpenAI GPT models
- Local models (Ollama, etc.)
- Custom/fine-tuned models
- BYOK - your keys, your privacy

## Technical Implementation

### Core Architecture

**Language:** Go (for performance, cross-platform distribution)
**UI Framework:** gocui (proven by lazygit)
**AI Integration:** REST APIs to various providers
**File Analysis:** Background goroutines with caching
**Database Integration:** Direct SQL connections (read-only)

### Permission System

**Always Allowed (No permission needed):**
- File reading (`cat`, `grep`, `find`)
- Database SELECT queries
- Process listing (`ps`, `top`)
- Log file analysis
- Dependency source code reading

**Never Allowed:**
- File modifications or creation
- Database writes (INSERT, UPDATE, DELETE)
- Process management
- System configuration changes
- Command execution that modifies state

### Project Structure
```
atlas/
├── pkg/
│   ├── lazygit/     # Pure lazygit fork (minimal changes)
│   ├── atlas/          # All Atlas AI functionality
│   │   ├── client/  # Multi-provider AI client
│   │   ├── context/ # Context analysis and management
│   │   ├── tools/   # Read-only system tools
│   │   └── guidance/# Instructional response generation
│   └── bridge/      # Integration between git and AI modes
├── main.go          # Entry point and mode switching
└── configs/         # AI model and tool configurations
```

## Development Phases

### Phase 1: Foundation
- **Goal:** Working AI mode with basic functionality
- Fork lazygit with clean separation
- Implement mode switching (G ↔ A)
- Basic AI client with Claude integration
- Smart file discovery system
- Simple conversation interface

### Phase 2: Intelligence Layer
- **Goal:** Read-only system analysis tools
- Database connection and query system
- Log file analysis and monitoring
- Process and system health checking
- Code path tracing and analysis
- Dependency source code reading

### Phase 3: Instructional AI
- **Goal:** Teaching-focused AI responses
- Step-by-step guidance system
- Context-aware instruction generation
- Rails-specific best practices
- Security and performance analysis
- Debugging workflow assistance

### Phase 4: Multi-Model & Polish
- **Goal:** Production-ready tool
- Multiple AI provider support
- Configuration and customization
- Performance optimization
- Documentation and onboarding
- IDE integration plugins (optional)

## Success Metrics

### Developer Experience Goals
- **10x faster** information gathering (no more grep/find cycles)
- **5x faster** debugging (systematic guidance vs trial-and-error)
- **3x better** code quality (real-time best practice guidance)
- **100% preserved** coding satisfaction (human writes all code)

### Technical Benchmarks
- **< 100ms** context switching between modes
- **< 200ms** file analysis and suggestions
- **< 2MB** memory usage for AI context
- **Works offline** for basic functionality (cached analysis)

## Competitive Positioning

### vs Claude Code / Cursor / Copilot
**They:** Replace developer decision-making, write code automatically
**Atlas:** Amplifies developer capabilities, provides expert guidance

### vs Traditional IDEs
**They:** Focus on editing features and extensions
**Atlas:** Terminal-native, works with any editor, focuses on system understanding

### vs Lazygit
**They:** Pure git interface
**Atlas:** Git interface + AI pair programmer (when needed)

## Business Model

### Open Source Foundation
Atlas is **free and open source software** with a BYOK (Bring Your Own Key) approach:
- Use your own API keys for Claude, OpenAI, or other AI providers
- No markup on AI costs - you pay the providers directly
- Complete control over your data and AI interactions
- Full access to all Atlas features and functionality

### Future Premium Plans
#### Atlas Pro ($20/month per developer)
- **Unified AI Access**: Access to all major AI models without managing multiple API keys
- **Custom Atlas Models**: Specialized coding models trained specifically for Atlas workflows
- **Enhanced Performance**: Optimized model routing and caching
- **Priority Support**: Direct support and feature requests
- **Enterprise Features**: Team management, usage analytics, SSO integration

### Market Positioning
- **Primary**: Senior developers who want to move faster while maintaining coding joy
- **Secondary**: Development teams wanting to level up junior developers
- **Enterprise**: Engineering organizations focused on productivity without compromising quality
- **Open Source Community**: All developers who want AI assistance without vendor lock-in

## Why This Will Succeed

### 1. Developer-First Design
- Preserves what developers love about coding
- Respects developer decision-making
- Works with existing workflows
- No vendor lock-in (BYOK)

### 2. Proven Foundation
- Built on lazygit's proven terminal UI
- Go's performance and reliability
- Terminal-native approach (works everywhere)
- Clean architecture enabling rapid iteration

### 3. Real Problem Solving
- Addresses actual developer pain points
- Focuses on information gathering overhead
- Provides expert guidance at the right time
- Teaches while solving problems

### 4. Sustainable Architecture
- Read-only operations eliminate security concerns
- Multi-model support prevents AI vendor lock-in
- Terminal-native works in any environment
- Extensible design for future capabilities

---

*Atlas: Your AI pair programmer that makes you 100x more productive while keeping the joy of coding alive.*