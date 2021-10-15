## เรียนรู้การออกแบบ micro service ด้วย Hexagonal Architecture

<br />
<br />

### **Repository (domain model)** <br />
```
เป็นส่วนที่ใช้ติดต่อกับฐานข้อมูล เช่น RDBMS, NoSQL, In memory database เป็นต้น
```

### **Service** <br />
```
เป็นส่วนของ Business Logic กำหนดว่ามี Service อะไรให้บริการบ้าง
ควร Response อะไรบ้าง หรือควรบอก error อย่างไรให้ user รู้
เป็นมุมมองของคนทำ Business Analysts หรือ SA
```

### **Handler** <br />
```
เป็นส่วนของการจัดการ Operation ต่าง ๆ ตามที่ Request เข้ามา
แล้ว Response กลับไป
```

<br />

### **Why Hexagonal Architecture?**
- สามารถแบ่งงานกับคนในทีมได้ง่ายขึ้น
- สามารถโฟกัสงานของตัวเองได้ดีขึ้น เช่น <br />
  คนทำ BL สามารถทำงานได้โดยไม่ต้องกังวลเกี่ยวกับ database มากจนเกินไป, <br />
  คนทำด้านติดต่อกับ database แค่ทำงานของตัวเองแล้วส่งต่อให้ BL
- สามารถแก้ไข/ต่อเติมได้ง่าย ไม่ต้องไล่แก้โค้ดเยอะ

<br />

### **Deploy with Docker** <br />
```
docker build -t goapi-robocar:v1.0 -f Dockerfile.multistage .
docker run --name goapi --restart=always -d -p 8088:8088 goapi-robocar:v1.0
```
