from flask import Flask, send_file  

app = Flask(__name__)  

@app.route('/proxy.pac')  
def serve_pac():  
    return send_file('proxy.pac', mimetype='application/x-ns-proxy-autoconfig')  

if __name__ == '__main__':  
    app.run(host='0.0.0.0', port=80)  
