# Visitor counter
Simple visitor counter which logs unique daily web visitors to your server and stores them into a .log file.

## Usage
Run `go build .` and copy the executable into a server you want to use it on `scp visitor-counter username@host:visitor-counter`.  
Connect into the server and create the directory to store log files `mkdir visitor-counter-logs` (optional).  
Go into newly created folder (if created in previous step) `cd visitor-counter-logs` and run the executable into the background `../visitor-counter &`, `./visitor-counter &` if directory was not created.  
Visit your site and you will see new .log file created in the directory from which the application was started.

## Extra info
Application stores .log files in the directory from which it was started, so they could be stored wherever you like.  
Only visitors to 443 port (https) are recorded.  
You can use `ssh username@host "ls visitor-counter-logs/* | xargs -n 1 wc -l"` to see the number of unique visitors from your local machine as a shorthand.
