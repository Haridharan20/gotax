# GoTax - GST Calculator

A Go application that calculates Goods and Services Tax (GST) for different Indian tax slabs.

## Project Structure

```
gotax/
├── README.md                 # This file
├── main.go                   # Entry point
├── go.mod                    # Go module definition
├── prices.txt                # Sample prices data
├── output/                   # Output directory (created at runtime)
│   ├── result_0.json         # 0% GST results
│   ├── result_5.json         # 5% GST results
│   ├── result_12.json        # 12% GST results
│   ├── result_18.json        # 18% GST results
│   └── result_28.json        # 28% GST results
├── prices/
│   └── prices.go             # GSTCalculation struct & IOManager interface
├── filemanager/
│   └── filemanager.go        # FileManager implementation (reads from file)
├── cmdmanager/
│   └── cmdmanager.go         # CMDManager implementation (reads from stdin)
└── utils/
    ├── filemanager.go        # File I/O utilities (ReadLines, WriteJson)
    └── conversion.go         # Type conversion utilities (StringToFloat)
```

## Key Components

### 1. **IOManager Interface** (`prices/prices.go`)
```go
type IOManager interface {
    LoadPrices() ([]float64, error)
    SaveResult(data any) error
}
```
Any manager that implements this interface can be used to load prices and save results.

### 2. **GSTCalculation** (`prices/prices.go`)
- Accepts a `Manager` instance
- Calculates GST-inclusive prices for each base price
- Formats output with rupee symbols (₹)

### 3. **FileManager** (`filemanager/filemanager.go`)
- Reads prices from a file (e.g., `prices.txt`)
- Writes results as JSON to `output/` directory
- Constructor: `filemanager.New(inputFile, outputFile)`

### 4. **CMDManager** (`cmdmanager/cmdmanager.go`)
- Reads prices from command-line input
- Prints results to console as JSON
- Constructor: `cmdmanager.New()`

### 5. **Utilities** (`utils/`)
- **ReadLines()** - Reads lines from a file
- **WriteJson()** - Writes data as JSON to file
- **StringToFloat()** - Converts string slice to float64 slice

## Usage

### Using FileManager (Default)
```go
gstRates := []float64{0.00, 0.05, 0.12, 0.18, 0.28}

for _, gstRate := range gstRates {
    fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", gstRate*100))
    calculation := prices.New(gstRate, fm)
    calculation.Process()
}
```

### Using CMDManager
```go
cm := cmdmanager.New()
pricesList, err := cm.LoadPrices()
if err != nil {
    return
}

// Create a manager wrapper to reuse prices
pm := PriceManager{prices: pricesList}

for _, gstRate := range gstRates {
    calculation := prices.New(gstRate, pm)
    calculation.Process()
}
```

## Running the Application

```bash
# Run with FileManager (default)
go run main.go

# Run with CMDManager (uncomment in main.go and run)
go run main.go
```

## GST Rates (India)

| Rate | Category |
|------|----------|
| 0% | Essentials (Food, Medicine) |
| 5% | Basic Items |
| 12% | Most Items |
| 18% | Services & Standard Items |
| 28% | Luxury Items |

## Example Input/Output

**Input (prices.txt):**
```
100
500
1000
48.96
```

**Output (result_18.json):**
```json
{
  "gst_rate": 0.18,
  "category": "Services & Standard Items",
  "base_prices": [100, 500, 1000, 48.96],
  "prices_with_gst": {
    "₹100.00": "₹118.00",
    "₹500.00": "₹590.00",
    "₹1000.00": "₹1180.00",
    "₹48.96": "₹57.77"
  }
}
```

## Architecture Pattern

The project uses the **Manager Pattern** (similar to Strategy Pattern):
- The `IOManager` interface defines how to load and save data
- Different implementations (`FileManager`, `CMDManager`) provide different I/O strategies
- `GSTCalculation` is decoupled from specific I/O implementation
- Easy to add new managers without modifying existing code

## File I/O Flow

```
FileManager           CMDManager
    ↓                     ↓
ReadLines()          User Input (stdin)
    ↓                     ↓
StringToFloat()      StringToFloat()
    ↓                     ↓
GSTCalculation.Process()
    ↓
SaveResult()
    ↓
WriteJson()    or    fmt.Println()
    ↓                     ↓
output/result_*.json    Console Output
```
