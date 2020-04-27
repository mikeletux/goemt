# GoEMT library
![Image of GoEMT](images/Logo.png )

This project aims to implement the Go bindings for the EMT Rest API.  
The EMT is the public transportation company from Madrid city. Using this API can be useful for getting live information from the service (waiting times, stop locations, etc). For full Rest API referece, please reffer to [HERE](https://apidocs.emtmadrid.es/).  

To start using it, we have to create a **APIClient** struct. To do so, first a **ClientConfig** struct needs to be created with your credentials:
```
config := goemt.ClientConfig{
		Enpoint:   "https://openapi.emtmadrid.es/v2",
		XClientID: "YOUR_XCLIENTID",
		PassKey:   "YOUR_PASSKEY",
```
*There are three kind of autentication. Basic, advanced and protected. The one used in the example is **protected**. Depending on which values from the **ClientConfig** struct are set, one or another will be used. Please reffer to [HERE](https://apidocs.emtmadrid.es/#api-Block_1_User_identity-login) to check which values your need to set for each login method.  
  
Once a config struct is created, the **Connect()** function must be used to log in:
```
api, err := goemt.Connect(config)
	if err != nil {
		panic(err)
	}

    //Do your thing

	err = api.Logout()
	if err != nil {
		panic(err)
	}
```
Once you're logged in the platform, you can use the different functions **(TO BE DONE)**. When you are done, use the Logout() method from the **APIClient** to close the session agains the EMT server.  

Known issues:
- GetTimeTableStartStop function doesn't work because platform returns Error code 90: Error managing internal services  

  
Miguel Sama 2020 (miguelsamamerino@gmail.com)  
GoEMT logo by Marta Recio (martarecio2011ti22@gmail.com)

