import qrcode  
import argparse  

def generate_wifi_qr(ssid, password, auth_type="WPA", hidden=False, output_file="wifi_qr.png"):  
    hidden_str = "true" if hidden else "false"  
    qr_data = f"WIFI:S:{ssid};T:{auth_type};P:{password};H:{hidden_str};;"  
    qr = qrcode.QRCode(  
        version=1,  
        error_correction=qrcode.constants.ERROR_CORRECT_L,  
        box_size=10,  
        border=4,  
    )  
    qr.add_data(qr_data)  
    qr.make(fit=True)  
    img = qr.make_image(fill_color="black", back_color="white")  
    img.save(output_file)  
    print(f"Wi-Fi QR code saved to {output_file}")  
    print(f"QR code data: {qr_data}")  

if __name__ == "__main__":  
    parser = argparse.ArgumentParser(description="Generate Wi-Fi QR code")  
    parser.add_argument("--ssid", required=True, help="Wi-Fi SSID")  
    parser.add_argument("--password", required=True, help="Wi-Fi password")  
    parser.add_argument("--auth_type", default="WPA", choices=["WPA", "WEP", "nopass"], help="Wi-Fi authentication type")  
    parser.add_argument("--hidden", action="store_true", help="Set if Wi-Fi network is hidden")  
    parser.add_argument("--output", default="wifi_qr.png", help="Output QR code image filename")  
    args = parser.parse_args()  
    generate_wifi_qr(args.ssid, args.password, args.auth_type, args.hidden, args.output)  