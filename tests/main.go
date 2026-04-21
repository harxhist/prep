package main

import (
    "bufio"
    "fmt"
    "io"
    "net/http"
    "regexp"
    "strconv"
    "strings"
)

// decodeSecretMessage fetches a Google Doc, parses character coordinates,
// constructs a grid, and prints the secret message.
// It takes the raw Google Doc URL (e.g., ending in /pub or /edit) as input.
func decodeSecretMessage(docURL string) {
    // 1. Construct the plain text export URL from the given Google Doc URL.
    // This handles common /edit or /pub suffixes to get the raw text content.
    txtURL := docURL
    if strings.HasSuffix(txtURL, "/edit") {
        txtURL = strings.TrimSuffix(txtURL, "/edit") + "/export?format=txt"
    } else if strings.HasSuffix(txtURL, "/pub") {
        txtURL += "/export?format=txt"
    } else if !strings.HasSuffix(txtURL, "/export?format=txt") {
        // Fallback for other formats, try appending. This assumes the base URL
        // without /edit or /pub can also be appended with /export?format=txt.
        txtURL += "/export?format=txt"
    }

    // 2. Fetch the content from the modified URL
    resp, err := http.Get(txtURL)
    if err != nil {
        fmt.Printf("Error fetching document from %s: %v\n", txtURL, err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("HTTP error fetching document from %s: %s\n", txtURL, resp.Status)
        // Try reading body for more details if available
        bodyBytes, _ := io.ReadAll(resp.Body)
        if len(bodyBytes) > 0 {
            fmt.Printf("Response body: %s\n", string(bodyBytes))
        }
        return
    }

    // 3. Parse coordinates and characters
    // charMap stores characters: Y-coordinate -> X-coordinate -> Character
    charMap := make(map[int]map[int]rune)
    maxX, maxY := 0, 0 // To track the maximum dimensions of the grid

    // Regex to match lines like "(X, Y, Char)" where Char is an uppercase letter.
    // It handles optional whitespace around commas.
    re := regexp.MustCompile(`^\((\d+)\s*,\s*(\d+)\s*,\s*([A-Z])\)$`)

    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text()) // Trim leading/trailing whitespace from the line
        matches := re.FindStringSubmatch(line)
        if len(matches) == 4 { // Full match plus 3 captured groups (X, Y, Char)
            x, err := strconv.Atoi(matches[1])
            if err != nil {
                // Should not occur if regex correctly identifies digits, but good practice.
                continue
            }
            y, err := strconv.Atoi(matches[2])
            if err != nil {
                continue
            }
            char := rune(matches[3][0]) // Get the first character (the uppercase letter)

            // Update maximum coordinates encountered
            if x > maxX {
                maxX = x
            }
            if y > maxY {
                maxY = y
            }

            // Store character in the map
            if _, ok := charMap[y]; !ok {
                charMap[y] = make(map[int]rune)
            }
            charMap[y][x] = char
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading document content: %v\n", err)
        return
    }

    // If no valid character coordinates were found
    if len(charMap) == 0 {
        fmt.Println("No character coordinates found in the document to decode a secret message.")
        return
    }

    // 4. Construct the grid and build the secret message string
    // Grid dimensions are (maxY+1) rows and (maxX+1) columns (since coordinates start from 0)
    grid := make([][]rune, maxY+1)
    for i := range grid {
        grid[i] = make([]rune, maxX+1)
        for j := range grid[i] {
            grid[i][j] = ' ' // Initialize all cells with a space character for empty spots
        }
    }

    var secretMessageBuilder strings.Builder
    // Populate the grid and build the secret message by iterating in row-major order
    // This ensures the secret message is constructed correctly from left-to-right, top-to-bottom.
    for y := 0; y <= maxY; y++ {
        for x := 0; x <= maxX; x++ {
            if char, ok := charMap[y][x]; ok {
                grid[y][x] = char
                secretMessageBuilder.WriteRune(char) // Add character to the secret message
            }
        }
    }

    // 5. Print the decoded grid
    fmt.Println("--- Decoded Grid ---")
    for _, row := range grid {
        fmt.Println(string(row))
    }
    fmt.Println("--------------------")

    // 6. Print the extracted secret message
    fmt.Printf("The secret message is: %s\n", secretMessageBuilder.String())
}

func main() {
    // The URL provided in the problem statement for verification
    testURL := "https://docs.google.com/document/d/e/2PACX-1vSvM5gDlNvt7npYHhp_XfsJvuntUhq184By5xO_pA4b_gCWeXb6dM6ZxwN8rE6S4ghUsCj2VKR21oEP/pub/export?format=txt"
    decodeSecretMessage(testURL)
}