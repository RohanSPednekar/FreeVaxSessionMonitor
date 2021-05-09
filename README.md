# FreeVaxSessionMonitor
This program monitors the free sessions available by pincode or district wise. District code need to be taken

pune : 363
Kolhapur: 371

## To get district code:
     1) First get state code by running: https://api.demo.co-vin.in/api/v2/admin/location/states in any browser. Note down state id. For Maharshtra it is 21
     2) Then get district code using https://api.demo.co-vin.in/api/v2/admin/location/districts/21. This is for maharshtra. Change state code accordingly in thus URL

For Maharshtra:
{"districts":[{"district_id":388,"district_name":"a b c d"},{"district_id":391,"district_name":"Ahmednagar"},{"district_id":364,"district_name":"Akola"},{"district_id":366,"district_name":"Amravati"},{"district_id":397,"district_name":"Aurangabad "},{"district_id":384,"district_name":"Beed"},{"district_id":370,"district_name":"Bhandara"},{"district_id":367,"district_name":"Buldhana"},{"district_id":380,"district_name":"Chandrapur"},{"district_id":379,"district_name":"Gadchiroli"},{"district_id":378,"district_name":"Gondia"},{"district_id":386,"district_name":"Hingoli"},{"district_id":390,"district_name":"Jalgaon"},{"district_id":396,"district_name":"Jalna"},{"district_id":371,"district_name":"Kolhapur"},{"district_id":383,"district_name":"Latur"},{"district_id":395,"district_name":"Mumbai"},{"district_id":365,"district_name":"Nagpur"},{"district_id":382,"district_name":"Nanded"},{"district_id":387,"district_name":"Nandurbar"},{"district_id":389,"district_name":"Nashik"},{"district_id":381,"district_name":"Osmanabad"},{"district_id":394,"district_name":"Palghar"},{"district_id":385,"district_name":"Parbhani"},{"district_id":363,"district_name":"Pune"},{"district_id":393,"district_name":"Raigad"},{"district_id":372,"district_name":"Ratnagiri"},{"district_id":373,"district_name":"Sangli"},{"district_id":788,"district_name":"Sanglii"},{"district_id":376,"district_name":"Satara"},{"district_id":374,"district_name":"Sindhudurg"},{"district_id":375,"district_name":"Solapur"},{"district_id":783,"district_name":"South Mumbai"},{"district_id":392,"district_name":"Thane"},{"district_id":377,"district_name":"Wardha"},{"district_id":369,"district_name":"Washim"},{"district_id":368,"district_name":"Yavatmal"}],"ttl":24}


## Build
On windows/Linux:
1) go get -v github.com/hajimehoshi/oto
2) go build program.go

## Usage
C:\Users\rohan_pednekar\cowin>program.exe --help
flag provided but not defined: -helppro
=== Get Vaccination Sessions available ===

Usage of program.exe:
  -beep int
        Time to Beep in Seconds (default 10 second)
  -date string
        Date to look at (default tomorrows date inf "09-05-2021")
  -district string
        District code (default "363")
  -pin string
        Area Pincode (default "411027") if district code is given as ) then pin code is honoured
  -sleep int
        Sleep time in seconds before trying out (default 4)

-help prints this usage message.

e.g. 

1) Search by district every 2 seconds for tomorrows date
 
  program.exe -district "363" -sleep 2
  
2) Search for specific date

  program.exe -district "363" -date "11-05-2021"
  
3) Seach by pincode for tomorrows date

  program.exe -district "0" -pin "411027"
  
  
  


