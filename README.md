# Launch My Project
A Go CLI to launch all services of a development project concurrently using goroutines.

Instead of manually starting frontend, backend, and other services one by one, run:

```launch <project_name>```

and your full dev environment boots instantly.

### Features
- Runs multiple project commands concurrently
- Config-based project management
- Lightweight and fast
- Built with Go and goroutines

### Installation

Clone the repository:

```
git clone https://github.com/<your-username>/launch-my-project.git
cd launch-my-project
go build -o <command_name>
```

Move the binary name <command_name> to your PATH and give permissions:
```
sudo cp <command_name> /usr/local/bin/<command_name>
sudo chmod +x /usr/local/bin/<command_name>
```

### Configuration
Create a config file at: ~/.config/launcher.json
```json
{
  "my_project": {
    "path": "/path/to/my_project",
    "commands": [
      "code .",
      "cd frontend && npm run dev",
      "cd backend && npx nodemon",
    ]
  }
}
```

### Usage
```bash
<command_name> my_project
```

### Why I Built This
While working on multiple projects, I found myself repeatedly starting several services manually.
This tool removes that friction and improves development workflow.
