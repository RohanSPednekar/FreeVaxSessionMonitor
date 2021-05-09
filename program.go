package main

import (
        "fmt"
        "net/http"
        "io/ioutil"
		"encoding/json"
		"flag"
		"os"
		//"errors"
		"io"
		"math"
		"sync"
		"time"
		"github.com/hajimehoshi/oto"
)

const (
	flagPinCode = "pin"
	flagDistrict = "district"
	flagDate = "date"
	flagSleepTime = "sleep"
	flagBeepTime = "beep"
	sampleRate int = 44100
	channelNum int = 2
	bitDepthInBytes int = 2
)

var (
	pinCode = flag.String(flagPinCode, "411027", "Area Pincode")
	distCode = flag.String(flagDistrict, "363", "District code")
	date = flag.String(flagDate, "09-05-2021", "Date to look at")
	timeToSleep = flag.Int(flagSleepTime, 4, "Sleep time in seconds before trying out")
	totalTimetoBeep = flag.Int(flagBeepTime, 10, "Time to Beep in Seconds")
)

type Center struct {
	Center_id int `json:"center_id"`
	Name string `json:"name"`
	State_name string `json:"state_name"`
	District_name string `json:"district_name"`
	Block_name string `json:"block_name"`
	Pincode int `json:"pincode"`
	From string `json:"from"`
	To string `json:"to"`
	Fee_type string `json:"fee_type"`
	Fee string `json:"fee"`
	Session_id string `json:"session_id"`
	Date string `json:"date"`
	Available_capacity int `json:"available_capacity"`
	Min_age_limit int `json:"min_age_limit"`
	Vaccine string `json:"vaccine"`
	Slots []string  `json:"slots"`
}

type ResponseStr struct {
	Sessions []Center `json:"sessions"`
}

type SineWave struct {
	freq   float64
	length int64
	pos    int64

	remaining []byte
}

func NewSineWave(freq float64, duration time.Duration) *SineWave {
	l := int64(channelNum) * int64(bitDepthInBytes) * int64(sampleRate) * int64(duration) / int64(time.Second)
	l = l / 4 * 4
	return &SineWave{
		freq:   freq,
		length: l,
	}
}

func (s *SineWave) Read(buf []byte) (int, error) {
	if len(s.remaining) > 0 {
		n := copy(buf, s.remaining)
		s.remaining = s.remaining[n:]
		return n, nil
	}

	if s.pos == s.length {
		return 0, io.EOF
	}

	eof := false
	if s.pos+int64(len(buf)) > s.length {
		buf = buf[:s.length-s.pos]
		eof = true
	}

	var origBuf []byte
	if len(buf)%4 > 0 {
		origBuf = buf
		buf = make([]byte, len(origBuf)+4-len(origBuf)%4)
	}

	length := float64(sampleRate) / float64(s.freq)

	num := (bitDepthInBytes) * (channelNum)
	p := s.pos / int64(num)
	switch bitDepthInBytes {
	case 1:
		for i := 0; i < len(buf)/num; i++ {
			const max = 127
			b := int(math.Sin(2*math.Pi*float64(p)/length) * 0.3 * max)
			for ch := 0; ch < channelNum; ch++ {
				buf[num*i+ch] = byte(b + 128)
			}
			p++
		}
	case 2:
		for i := 0; i < len(buf)/num; i++ {
			const max = 32767
			b := int16(math.Sin(2*math.Pi*float64(p)/length) * 0.3 * max)
			for ch := 0; ch < channelNum; ch++ {
				buf[num*i+2*ch] = byte(b)
				buf[num*i+1+2*ch] = byte(b >> 8)
			}
			p++
		}
	}

	s.pos += int64(len(buf))

	n := len(buf)
	if origBuf != nil {
		n = copy(origBuf, buf)
		s.remaining = buf[n:]
	}

	if eof {
		return n, io.EOF
	}
	return n, nil
}

func play(context *oto.Context, freq float64, duration time.Duration) error {
	p := context.NewPlayer()
	s := NewSineWave(freq, duration)
	if _, err := io.Copy(p, s); err != nil {
		return err
	}
	if err := p.Close(); err != nil {
		return err
	}
	return nil
}

func run() error {
	const (
		freqC = 523.3
		freqE = 659.3
		freqG = 784.0
	)

	c, err := oto.NewContext(sampleRate, channelNum, bitDepthInBytes, 4096)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := play(c, freqC, time.Duration(*totalTimetoBeep) * time.Second); err != nil {
			panic(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		if err := play(c, freqE, time.Duration(*totalTimetoBeep) * time.Second); err != nil {
			panic(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		if err := play(c, freqG, time.Duration(*totalTimetoBeep) * time.Second); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	c.Close()
	return nil
}

func runCommand(url string) (err error, done bool) {
        done = false
		req, _ := http.NewRequest("GET", url, nil)

        req.Header.Add("accept", "application/json")
        req.Header.Add("accept-language", "hi_IN")
		req.Header.Add("user-agent", "Rohan Pednekar")

        res, _ := http.DefaultClient.Do(req)

        defer res.Body.Close()
        body, _ := ioutil.ReadAll(res.Body)
		
		var sessions ResponseStr
		err = json.Unmarshal(body, &sessions)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("sessions:=", sessions)
		
		if len(sessions.Sessions) > 0 {
		    for _,session := range sessions.Sessions {
				if session.Min_age_limit == 18  && session.Available_capacity > 9 {
				    fmt.Println(string(body))
					fmt.Println("Registration is open for 18+..............")
					done = true
					return
				}
			}
			
			done = false
			fmt.Println("Registration is open for 45+..............")
			fmt.Println(string(body))
			return
		}
		done = false
		fmt.Println("No luck..............", string(body))
		return  
}

func main() {
		flag.Usage = func() {
			fmt.Fprintln(os.Stderr, "=== Get Vaccination Sessions available ===")
			fmt.Fprintf(os.Stderr, "\nUsage of %s:\n", os.Args[0])
			flag.PrintDefaults()

			fmt.Fprintf(os.Stderr, "\n-help prints this usage message.\n")
		}
		flag.Parse()
		
		if *date == "09-05-2021" {
			// Get tomorrow's date
			today := time.Now()
			tomorrow := today.AddDate(0, 0, 1)
			
			*date = fmt.Sprintf("%02d-%02d-%04d", tomorrow.Day(), tomorrow.Month(), tomorrow.Year())
		}
		
		var url string
		if *distCode == "0" {
			url = fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByPin?pincode=%s&date=%s", *pinCode, *date)
		} else {
			url = fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByDistrict?district_id=%s&date=%s", *distCode, *date)
		}
		
		fmt.Println("Area Pin Code: ", *pinCode)
		fmt.Println("Appointmet Date: ", *date)
		fmt.Println("Sleep time between 2 requests: ", *timeToSleep)
		fmt.Println("Time to Beep in Seconds: ", *totalTimetoBeep)
		fmt.Println("District code: ", *distCode)
		fmt.Println("Trying out at: ", url)
		done := false
		var err error
		
		for !done {
			err, done = runCommand(url)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				return
			}
			if done {
				if err := run(); err != nil {
					panic(err)
					return
				}
				done = false
			}
			time.Sleep(time.Duration(*timeToSleep) * time.Second)
		}
		
		fmt.Println("Registrion is Open. I am leaving now.....Happy Vaccination!")
}