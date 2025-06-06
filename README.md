# UsernameAlgo

🔍 **UsernameAlgo** is a lightweight command-line utility written in Go that scans large text files line by line and extracts common prefixes found in **groups of three or more consecutive lines**.

---

## 📦 Features

- ✅ Reads input file from CLI argument or prompts via `stdin`
- 🚀 Efficient prefix matching on massive files (e.g., `usernames.txt`)
- ✂️ Outputs only **significant common prefixes**
- 📝 Creates an `.out` file containing the results
- 🧪 Clean and simple Go implementation

---

## 📂 Usage

```bash
go run app.go <path-to-input-file>
```

Or if no argument is passed:

```bash
$ go run app.go
Enter file path: usernames.txt
```

### Example Input (`usernames.txt`):

```
alpha_apple
alpha_apricot
alpha_apollo
beta_box
beta_banana
beta_bounce
gamma_grape
```

### Example Output (`usernames.txt.out`):

```
alpha_ap
beta_b
```

## 🧠 How it Works

- Reads all lines from the file
- Compares each line with the next one to find their longest common prefix
- Continues as long as lines share a non-empty prefix
- Only groups of 3 or more are considered valid
- Writes only the final common prefix of each valid group into a .out file

## 🛠 Build

To compile it as a standalone executable:

```bash
go build -o usernamealgo.exe app.go
```

## 📁 Project Structure

```
.
├── app.go           # Main logic
├── go.mod           # Go module definition
├── LICENSE          # MIT License
├── README.md        # You're here!
├── test.txt         # Sample input
├── test.txt.out     # Sample output
├── usernamealgo.exe # Built binary
```

## 📄 License

This project is licensed under the MIT License.

© 2025 Max Base

## 💡 Author

Created with ❤️ by Max Base
