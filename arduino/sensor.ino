
//#include "DHT.h"
#include "max6675.h"

//#define DHTPIN 2    
//#define DHTTYPE DHT22

//DHT dht(DHTPIN, DHTTYPE);

int thermoDO = 4;
int thermoCS = 5;
int thermoCLK = 6;

MAX6675 thermocouple(thermoCLK, thermoCS, thermoDO);

void setup() {
  Serial.begin(9600);
//  Serial.println("DHT22 test!");  
  Serial.println("MAX6675 test\n");

//  dht.begin();
  delay(300);
}

void loop() {
  // DHT22
//  float humid = dht.readHumidity();
//  float temp_celc = dht.readTemperature();
//  float temp_fahr = dht.readTemperature(true);
//
//  if (isnan(humid) || isnan(temp_celc) || isnan(temp_fahr)) {
//    Serial.println(F("Failed to read from DHT sensor!"));
//    return;
//  }
//
//  Serial.print("DHT22: \n");
//  Serial.print("\tHumidity\t: ");
//  Serial.print(humid);
//  Serial.print("%\n\tTemperature(C)\t: ");
//  Serial.print(temp_celc);
//  Serial.print("°C \n\tTemperature(F)\t: ");
//  Serial.print(temp_fahr);
//  Serial.print("°F \n");

//   MAX6675
  Serial.print("MAX6675: \n");
  Serial.print("\tTemperature(C)\t: "); 
  Serial.print(thermocouple.readCelsius());
  Serial.print("°C\n");
  Serial.print("\tTemperature(F)\t: ");
  Serial.print(thermocouple.readFahrenheit());
  Serial.print("°F\n\n");

  delay(5000);
}
