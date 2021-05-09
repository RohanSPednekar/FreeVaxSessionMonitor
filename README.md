# FreeVaxSessionMonitor
This program monitors the free sessions available by pincode or district wise. District code need to be taken

pune : 363
Kolhpuar: 371


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
  
  
  


