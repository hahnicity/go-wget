# Go-WGET
Perform a GET request and then save its contents to file

## Installation
It's Easy! Type in at the command line:

        go get github.com/hahnicity/go-wget

## Usage
Since there exists a builtin wget utility on Unix that is far more mature than this 
piece of code I have written, I am only allowing this code to be utilized via scripting.

### Interface
The user can utilize the Wget function

        Wget(url, fileName string)

A file name can be input if the user desires a custom file path/name to be used. If not
then the user can input an empty string "" and a file named as the last chunk of the url 
will be created in the current working directory. 

Example:

        Wget("https://google.com", "")

Will create a file google.com in the current working directory

More advanced usage can utilize the `MakeRequest` and `WriteToFile` functions

`MakeRequest` will return a http.Response object using the url we provided

        MakeRequest(url string) *http.Response

`WriteToFile` will write the content in the http.Response to file

        WriteToFile(fileName string, resp *http.Response)
