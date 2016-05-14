# automatic-carnival

This has been written in Golang.

* Download the binaries under bin directory. There are two binaries built. main_mac for mac and main_linux for linux

* Set GOPATH='directory where you have the binaries'. NOTE : Download the resources folder and put it in directory which is 
set as GOPATH. Because the code fetches the yaml file for the configuration

* Execute the binary. This is written as a webservice. 

# how to use
* http://localhost:8080/ is where the service runs.

* http://localhost:8080/getcost/{hours}/{cpus}/{price} . Here getcost is the method that calculates and returns the values. 

SCENARIOS

* http://localhost:8080/getcost/4/11/0 - minimum N CPUs for H hours - This will just calculate the amount. This find the best possible 
combinations of cpus and gives the reponse 

* http://localhost:8080/getcost/1/0/2.82  - maximum price they are willing to pay for H hours. This calculates the number of CPUS
it can provide with $10 for four hours

* http://localhost:8080/getcost/1/10/2.82 or a combination of both. - This calculates the best possible servers with the cpus user input
and within the price they have mentioned

* Output is in JSON Format
 
Test cases
* All parameters are valid.  http://localhost:8080/getcost/4/11/10. Should work fine
* Hours is 0 .  http://localhost:8080/getcost/0/11/0 . Throw an error . Hours should be valid
* When cpus is 0.  http://localhost:8080/getcost/4/0/10. Should calculate right counts and return values. Not for the number of CPUS
it allocates based on the cost efficiency
* When price is 0.  http://localhost:8080/getcost/4/11/0. Should calculate price based on the hours and number of cpus. 
* When price is 0 and cpus are 0 .  http://localhost:8080/getcost/4/0/0. Should throw error any one should be available 
* When the parameters are passed negative values.  http://localhost:8080/getcost/-4/-11/-1 . Hours cannot be valid

* When there cannot be any cpus allotted. http://localhost:8080/getcost/1/1/0.10. cost_details should be 0 and empty for the user to validate
* When the config file is not in place. rename the costconfig.yaml to costconfig1.yaml under resources directory and test. Should throw right 
error message without panicking
* When cpus are passed a float point -  http://localhost:8080/getcost/1/0.45/10. This truncates the cpu to be zero and calculates the number of cpus for the given 10 price.
* When cpus are passed a float point and price is zero -  http://localhost:8080/getcost/1/0.45/0. Throws the error message price or cpus should be valid
* when inputs are strings - http://localhost:8080/getcost/1/0.45/0 - This validates the inputs. Should throw error message inputs should be valid

Not done 
* when the inputs are really large number - Need to handle this case. Left it that practically there would be limited resources. If required, can add this







