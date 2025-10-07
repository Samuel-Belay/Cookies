import requests

cookies = {
    'session_id': 'iiv0krhcgeksxaeq2mc3n8vybgimdb96',
    'csrftoken': '4Btngu6EjuEw8BbCtTLw9zw5IDsPOVyB'
}

url = 'https://support.mozilla.org/'

headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0 Safari/537.36',
    'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8'
}

response = requests.get(url, cookies=cookies, headers=headers)

print(f"Status code: {response.status_code}")

html = response.text.lower()

# Check for common logged-in indicators
indicators = ['logout', 'sign out', 'my account', 'profile', 'welcome', 'username']

found = [word for word in indicators if word in html]

if found:
    print("Possible logged-in session detected. Found keywords:", found)
else:
    print("No logged-in indicators found. Session might be invalid or expired.")

# Optional: save full HTML to a file for manual inspection
with open('response.html', 'w', encoding='utf-8') as f:
    f.write(response.text)
