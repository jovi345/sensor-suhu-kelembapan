#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <ArduinoJson.h>

#include "DHT.h"
#include "max6675.h"

// untuk arduino uno
//#define DHTPIN 2   
//int thermoSO = 4;
//int thermoCS = 5;
//int thermoSCK = 6;

// untuk esp8266
#define DHTPIN 5
int thermoSO = 13; // D7
int thermoCS = 15; // D8
int thermoSCK = 14; // D5

#define DHTTYPE DHT22

DHT dht(DHTPIN, DHTTYPE);
MAX6675 thermocouple(thermoSCK, thermoCS, thermoSO);

const char* ssid = "<ganti dengan nama wifi>";
const char* password = "<ganti dengan password wifi>";
// ganti dengan ip backend
const char* serverUrl = "http://192.168.24.182:8080/api/v1/data/add"; 

WiFiClient client;

void setup() {
  Serial.begin(9600);
  Serial.println("DHT22 test!");  
  Serial.println("MAX6675 test\n");

  WiFi.begin(ssid, password);
  Serial.print("Menghubungkan ke WiFi");
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.print(".");
  }
  Serial.print("\nTerhubung ke WiFi");
  
  dht.begin();
  delay(300);
}

void loop() {
  if (WiFi.status() == WL_CONNECTED) {
    float room_humid = dht.readHumidity();
    float room_temp_celc = dht.readTemperature();
    float room_temp_fahr = dht.readTemperature(true);
    float thermocouple_temp_c = thermocouple.readCelsius();
    float thermocouple_temp_f = thermocouple.readFahrenheit();

    if (isnan(room_humid) || isnan(room_temp_celc) || isnan(room_temp_fahr)) {
      Serial.println("Gagal membaca dari DHT sensor!");
      return;
    }

    Serial.println("DHT22:");
    Serial.printf("\tTemperature(C): %.2f°C\n", room_temp_celc);
    Serial.printf("\tTemperature(F): %.2f°F\n", room_temp_fahr);
    Serial.printf("\tHumidity: %.2f%%\n", room_humid);
    
    Serial.println("MAX6675:");
    Serial.printf("\tTemperature(C): %.2f°C\n", thermocouple_temp_c);
    Serial.printf("\tTemperature(F): %.2f°F\n\n", thermocouple_temp_f);

    // Mengirim data ke server
    HTTPClient http;
    http.begin(client, serverUrl);
    http.addHeader("Content-Type", "application/json");
    http.addHeader("Origin", "http://192.168.24.188"); // ganti dengan ip wifi
 
    // JSON untuk dikirim
    StaticJsonDocument<200> jsonDoc;
    jsonDoc["humidity"] = room_humid;
    jsonDoc["room_temperature"] = room_temp_celc;
    jsonDoc["object_temperature"] = thermocouple_temp_c;

    String jsonString;
    serializeJson(jsonDoc, jsonString);

    int httpResponseCode = http.POST(jsonString);

    if (httpResponseCode == 201) {
      Serial.print("Data terkirim! Response: ");
      Serial.println(httpResponseCode);
    } else {
      Serial.print("Gagal mengirim data. Error: ");
      Serial.println(httpResponseCode);
    }

    http.end();
  } else {
    Serial.println("WiFi tidak terhubung!");
  }

  delay(5000); // Kirim data setiap 5 detik
}