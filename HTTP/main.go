package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// fmt.Println(resp) this will not give us the body of response

	/*
		//body is of type interface
		//field inside a struct can be of type interface
		Response Struct:
		1. Status - string
		2. StatusCode - int
		3. Body - io.ReadCloser (interface)

		//We can take different interfaces and assemble them together to form another interface
		io.ReadCloser interface
		1. Reader - (interface)
		2. Closer - (interface)

		//Reader interface is used to provide a common output ie []byte for different sources of input data (eg text file on hard drive, image file, http req body, text from cmd line, data from analog sensor)

		io.Reader Interface {
			Read(p []byte)(n int, err error) //function
			We create a slice of byte and then pass it to Reader func as para. Since slice is passed by reference thus it modifies the original slice which was passed and injects response body into it.
			return values: n is no of bytes in res body and err is error occured

			//source of data -> Reader-> []byte
		}

		io.Closer Interface{
			Close()() //function
		}




	*/

	/*

		//read function is not set up to automatically resize the slice. Thus we make an arbitary guess to accodomate all data
		bs := make([]byte, 99999) //make a slice where no of empty spaces_>99999
		resp.Body.Read(bs)      //pass our byte slice to Read func
		fmt.Println(string(bs)) //prints the response body
	*/

	//Writer interface describes something that can take info and send it outside of our program
	// []byte -> Writer -> Some form of output
	//we  pass []byte to a value that implements writer to log out all the info

	io.Copy(os.Stdout, resp.Body) //shorthand for the above three lines
	//copies from src to dest
	//os.Stdout (dest) is a value that implemets writer interface
	//resp.Body(src) is a value that implemets reader interface

}
