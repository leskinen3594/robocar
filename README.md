### ไม่สามารถ deploy API ขึ้น Docker ได้

<br />

**ปัญหาที่เกิด**
- เนื่องจาก container MySQL เป็น IP 192.168.X.X ทำให้ API ไม่สามารถเชื่อมต่อฐานข้อมูล MySQL ได้

<br />

**แก้ปัญหาโดย run บน local** <br />
- run container ที่ต้องใช้ ได้แก่ MySQL, Redis, MQTT โดยใช้คำสั่ง <br />
  `docker-compose up -d` <br />
  โดยรายละเอียดอื่น ๆ สามารถดูได้ภายในโฟลเดอร์ของแต่ละ container <br />
  หรือดู config ในไฟล์ docker-compose.yaml
- เข้าไปที่โฟลเดอร์ goapi จากนั้น run ไฟล์ main.go โดยใช้คำสั่ง <br />
  `go run main.go`