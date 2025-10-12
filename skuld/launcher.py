import subprocess
import os

# Get the folder where this launcher is located
base_path = os.path.dirname(os.path.abspath(__file__))

# Paths to the other EXEs (assumes they are in the same folder)
app1_path = os.path.join(base_path, "app1.exe")
app2_path = os.path.join(base_path, "app2.exe")

subprocess.run([app1_path])  # waits for app1.exe to finish

subprocess.run([app2_path])  # runs after app1.exe finishes

print("All done!")

# pip install pyinstaller
# pyinstaller --onefile --noconsole --name Exela launcher.py