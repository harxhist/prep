import sys
import requests
from bs4 import BeautifulSoup

def decode_secret(doc_url: str):
    try:
        response = requests.get(doc_url)
        response.raise_for_status()
    except requests.exceptions.RequestException as e:
        print(f"Error fetching the document: {e}")
        return

    soup = BeautifulSoup(response.text, "html.parser")

    char_data = {}
    max_x = 0
    max_y = 0

    table = soup.find("table")
    if not table:
        print("No table found in the document.")
        return

    rows = table.find_all("tr")
    for row in rows[1:]:
        cells = row.find_all("td")
        if len(cells) == 3:
            try:
                x = int(cells[0].get_text(strip=True))
                char = cells[1].get_text(strip=True)
                y = int(cells[2].get_text(strip=True))
                char_data[(x, y)] = char
                max_x = max(max_x, x)
                max_y = max(max_y, y)
            except ValueError:
                continue

    if not char_data:
        print("No valid character data found.")
        return

    grid = [[" " for _ in range(max_x + 1)] for _ in range(max_y + 1)]
    for (x, y), char in char_data.items():
        grid[y][x] = char

    for row in grid:
        print("".join(row))

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python3 a.py <google_doc_url>")
        sys.exit(1)
    decode_secret(sys.argv[1])