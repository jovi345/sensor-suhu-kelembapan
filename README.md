# Sistem Monitoring Suhu

## Langkah-Langkah menjalankan program ini

### 1. Clone project

```bash
git clone https://github.com/jovi345/sensor-suhu-kelembapan.git
```

### 2. Salin kode arduino ke Arduino IDE Anda

### 3. Jalankan MySQL DB

### 4. Buat tabel pada DB Anda

```bash
-- temp_humid.sensor_data definition

CREATE TABLE `sensor_data` (
  `id` int NOT NULL AUTO_INCREMENT,
  `room_temperature` double DEFAULT NULL COMMENT 'in celcius',
  `object_temperature` double DEFAULT NULL COMMENT 'in celcius',
  `humidity` double DEFAULT NULL COMMENT 'percentage',
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7409 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

## BACKEND

### Buka direktori backend pada IDE Anda dan jalankan command ini

```bash
go mod tidy
go run main.go
```

## FRONTEND

### Buka direktori frontend pada IDE Anda dan jalankan command ini

```bash
npm install
npm run dev
```

### Buka localhost:5173 di browser

## ARDUINO

### Pada kode baris 25 dan 26 ganti dengan SSID dan Password Wifi Anda

```bash
const char* ssid = "namawifisaya";
const char* password = "passwordwifisaya";
```

### Pada kode baris 28, ganti IP ini dengan IP Backend Anda

```bash
const char* serverUrl = "<ip_backend_anda>/api/v1/data/add";
```

Info: IP ini dapat Anda dapatkan dengan menjalankan server backend terlebih dahulu

### Pada Linux & macOS, IP dapat diperoleh melalui shell command berikut

```bash
ifconfig
```

### Pada kode baris 75, ganti IP ini dengan IP ESP8266 yang telah terhubung ke WiFi Anda

Info: Menggunakan Mobile Hotspot akan lebih mudah untuk mendapatkan IP ini

```bash
http.addHeader("Origin", "<ip_esp8266>");
```
