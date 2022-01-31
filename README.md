# FreeVaxSessionMonitor
Exe runs on windows on command line. For Linux/Mac you need to build by installing GO first.(I might upload exe for Linux tomorrow, not sure of MAC though. If anybody could help in doing that for Mac?) Rohan

**Please make sure your are connected to internet, otherwise this command line crashes as certian error handling is missing**

This program monitors the free sessions available by pincode or district wise. More details on how to run, are below. District code need to be taken. This program generates beep only for

         1) Vaccine session is available for given age. Default is 18+ but you can give option as 45 for 45+ search.
         2) session must have capacity of doses greater than 9. (observed some sessions with just 1 dose capacity and no point in generating alarm for the same)

District codes
pune : 363
Kolhapur: 371

## To get district code other than Pune or Kolhapur:
     1) First get state code by running: https://api.demo.co-vin.in/api/v2/admin/location/states in any browser. Note down state id. For Maharshtra it is 21
     2) Then get district code using https://api.demo.co-vin.in/api/v2/admin/location/districts/21. This is for maharshtra. Change state code accordingly in thus URL

For Maharshtra:
{"districts":[{"district_id":388,"district_name":"a b c d"},{"district_id":391,"district_name":"Ahmednagar"},{"district_id":364,"district_name":"Akola"},{"district_id":366,"district_name":"Amravati"},{"district_id":397,"district_name":"Aurangabad "},{"district_id":384,"district_name":"Beed"},{"district_id":370,"district_name":"Bhandara"},{"district_id":367,"district_name":"Buldhana"},{"district_id":380,"district_name":"Chandrapur"},{"district_id":379,"district_name":"Gadchiroli"},{"district_id":378,"district_name":"Gondia"},{"district_id":386,"district_name":"Hingoli"},{"district_id":390,"district_name":"Jalgaon"},{"district_id":396,"district_name":"Jalna"},{"district_id":371,"district_name":"Kolhapur"},{"district_id":383,"district_name":"Latur"},{"district_id":395,"district_name":"Mumbai"},{"district_id":365,"district_name":"Nagpur"},{"district_id":382,"district_name":"Nanded"},{"district_id":387,"district_name":"Nandurbar"},{"district_id":389,"district_name":"Nashik"},{"district_id":381,"district_name":"Osmanabad"},{"district_id":394,"district_name":"Palghar"},{"district_id":385,"district_name":"Parbhani"},{"district_id":363,"district_name":"Pune"},{"district_id":393,"district_name":"Raigad"},{"district_id":372,"district_name":"Ratnagiri"},{"district_id":373,"district_name":"Sangli"},{"district_id":788,"district_name":"Sanglii"},{"district_id":376,"district_name":"Satara"},{"district_id":374,"district_name":"Sindhudurg"},{"district_id":375,"district_name":"Solapur"},{"district_id":783,"district_name":"South Mumbai"},{"district_id":392,"district_name":"Thane"},{"district_id":377,"district_name":"Wardha"},{"district_id":369,"district_name":"Washim"},{"district_id":368,"district_name":"Yavatmal"}],"ttl":24}


## Build
On windows/Linux:
1) go get -v github.com/hajimehoshi/oto
2) go build program.go

## Usage
1) First download program.exe anywhere on your windows machine(for other OS you need to build the program). I downloaded it into "C:\Users\rohan_pednekar\cowin".
2) Go Windows->"Type CMD". Start CMD(Command Prompt App), and got to directory where you downloaded

     C:>**cd C:\Users\rohan_pednekar\cowin**
3) Get Program usage details as

     C:\Users\rohan_pednekar\cowin>**program.exe -help**
     
     === Get Vaccination Sessions available ===
     
     Usage of program.exe:
     -age int
     
          Age limit, For 45+ please give 45 and 18+ give 18. (default 18)
        
     -beep int
     
          Time to Beep in Seconds (default 10 second)
          
     -date string
     
          Appointment Date in format DD-MM-YYYY. (default "tomorrow date")
          
     -district string
     
          District code, default is Pune's code. Check Readme.md for district code of your choice. (default "363")
          
     -pin string
     
          Optional Area Pincode. If District code is given then this will be ignored. (default "411027")
          
     -sleep int
     
          Optional Area Pincode. If District code is given then this will be ignored. (default "411027")
          

    -help prints this usage message.
4) Below some examples:
e.g. 
         
1) For Pune with 18+, simply run


        C:\Users\rohan_pednekar\cowin>program.exe
          
          
2) For Pune with 45+, run
     

         C:\Users\rohan_pednekar\cowin>program.exe -age 45
         
          
3) Search by district every 2 seconds for **tomorrows date**. For district code please see above. 363 is Pune district code.
     
 
         C:\Users\rohan_pednekar\cowin>program.exe -district "363" -sleep 2
          
  
4) Search for **specific date**
     

         C:\Users\rohan_pednekar\cowin>program.exe -district "363" -date "11-05-2021"
          
  
5) Search by pincode for **tomorrows date**
     

         C:\Users\rohan_pednekar\cowin>program.exe -district "0" -pin "411027"
   
  
  
  


