<p align="center">
  <img src="ChatGPT Image 23 de jun. de 2025, 00_14_30.png" alt="Spoofer Banner" width="800"/>
</p>




# WhatsApp Spoofed Message Tool

A Go-based tool to send spoofed messages on WhatsApp. This tool is for educational purposes only. Use it responsibly and only on systems you own or have explicit permission to test.

---

## **Installation**

### **Prerequisites**
- [Termux](https://termux.com/) (for Android) or a Linux terminal.
- Go (Golang) installed.

### **Steps**
1. **Install repository**:
   ```bash
   git clone https://github.com/ARESHAmohanad/spoofer.git
   cd whatsapp

   
2. **Install Go**:
   ```bash
   apt-get update && apt-get upgrade
   apt-get install golang
   apt-get upgrade golang
   go clean -modcache
   go get github.com/Rhymen/go-whatsapp
   go get google.golang.org/protobuf/proto
   
3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

4. **Build the tool**:
   ```bash
   chmod +x *
   go build spoofed_whatsapp.go
   ```

5. **Grant execute permissions**:
   ```bash
   chmod +x *
   ```


## **Usage**

1. **Run the tool**:
   ```bash
   ./spoofed_whatsapp
   ```

2. **Scan the QR code**:
   - Open WhatsApp on your phone.
   - Go to **Linked Devices** and scan the QR code displayed .

3. **Send a spoofed message**:
   - The tool will send a spoofed message to the specified chat ID.
