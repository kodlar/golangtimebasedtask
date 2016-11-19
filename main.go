package main

import(
		"fmt" 
	   "time"
	   "errors"
	   "os"
      )

func main(){
	
	hourOfDay := time.Now().Hour()
	
	greeting, err := getGreeting(hourOfDay)
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	fmt.Println(greeting)
	
}

func getGreeting(hour int)(string, error){
	
	var message string
	
	if hour < 7 {
	  err := errors.New("Selamlamak için çok erken bir saat")
	    return message, err
    }else if hour < 12 {
		message = "günaydın genç" 
	}else if hour < 18 {
		message = "iyi akşamlar genç"
	}else{
		message = "iyi geceler genç"
	}
	return message, nil
}