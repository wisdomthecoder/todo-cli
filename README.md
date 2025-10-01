
# 📝 Go Todos CLI

A lightweight **command-line Todo app** built in **Go**.  
This project was created as a practice exercise to learn how to work with **Go flags**, **structs**, and **basic file I/O (optional extension)**.

With this CLI tool, you can easily manage todos from your terminal — add, edit, delete, toggle, and list tasks.

---

## ✨ Features
- ➕ **Add** a new todo  
- ✏️ **Edit** a todo by index with a new title  
- 🗑️ **Delete** a todo by index  
- ✅ **Toggle** a todo (mark as done/undone)  
- 📋 **List** all todos with their status  
- ❓ **Help** flag to display available commands  

---

## 📦 Requirements
- Go 1.21+ (recommended latest version)  
- A terminal or command-line environment  

---

## ⚙️ Installation

Clone the repository:

```bash
git clone https://github.com/wisdomthecoder/todos-cli.git
cd todos-cli
````

Build the binary:

```bash
go build -o todos
```

Run the program:

```bash
./todos [flags]
```

---

## 📌 Usage

### Add a new todo

```bash
./todos -add "Buy groceries"
```

### List todos

```bash
./todos -list
```

### Edit a todo

```bash
./todos -edit "1:Read Go documentation"
```

> **Format:** `index:new_title`

### Toggle a todo (done/undone)

```bash
./todos -toggle 1
```

### Delete a todo

```bash
./todos -del 1
```

### Help

```bash
./todos -h
```

---

## 🛠 Example Workflow

```bash
./todos -add "Learn Go"
./todos -add "Build a CLI app"
./todos -list
./todos -edit "1:Master Go CLI apps"
./todos -toggle 1
./todos -del 2
./todos -list
```

---

## 📂 Project Structure

```
.
├── main.go       # Entry point 
├── storage.go    # Handle Persistent using JSON 
├── command.go    # Commands and flags 
├── todos.go      # Todos logic (add, edit, delete, toggle, list)
├── todos.json    # Local Storage for Storing Todos
└── README.md     # Project documentation
```

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!
Feel free to fork this repo, submit a PR, or open an issue.

### Steps:

1. Fork the project
2. Create your feature branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -m "Add new feature"`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Open a Pull Request 🎉

---

## 🚀 Future Improvements

* Save todos to a file (persistent storage)
* Add due dates and priorities
* Support categories/tags
* Add colorized output for better UX
* Export todos to JSON/CSV

---

## 👨‍💻 Author

**Wisdom Dauda**

* GitHub: [@wisdomthecoder](https://github.com/wisdomthecoder)
* Twitter/X: [@wisdomthecoder](https://x.com/WisdomDauda4)

---

## 📜 License

This project is open source and available under the [MIT License](LICENSE).

