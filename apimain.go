package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"log"

	"github.com/gin-contrib/cors"

	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql" // _ is basically an alias for the library...meaning another name
)

type Transaction struct { //Format in which data should be there.
	Lender   string `json:"lender"`
	Reciever string `json:"reciever"`
	Amount   int    `json:"amount"`
	Date     string `json:"date"`
}

var Payment = []Transaction{ //Array of transactions basically stores all the Transactions.
	{Lender: "Krish Sorathia", Reciever: "Me", Amount: 10000, Date: "2023-06-21"},
	{Lender: "Krish2", Reciever: "Me2", Amount: 20000, Date: "2023-06-21"},
}

func addTransactions(context *gin.Context) {
	var newTransactions Transaction

	if Err := context.BindJSON(&newTransactions); Err != nil { // contect.BindJSON means it is going to take JSON inside the body and it will bind it to the newTransactions variable which we just created of TRANSACTION type.
		// IF our newTransaction is not of Transaction type we will get an error.
		//The condition checks whether we have an error or not so if there is no error we will bind it without and issue.
		//But if error is not nil i.e it is not 0 then we will simply return and not Bind it.
		return
	}

	db, err := sql.Open("mysql", "root:Nevermind@22@tcp(localhost:3306)/payment") //Used to open a database
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Datbase Connected Succesfully")
	defer db.Close()

	_, err = db.Exec("INSERT INTO transaction (Lender, Reciever, Amount, Date) VALUES (?,?,?,?)",
		newTransactions.Lender, newTransactions.Reciever, newTransactions.Amount, newTransactions.Date)
	if err != nil {
		fmt.Println("Error", err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	fmt.Println("Database updated succesfully!")

	Payment = append(Payment, newTransactions) // Added the new Transaction in the array.

	context.IndentedJSON(http.StatusCreated, newTransactions)
	//Used to return the new Payment.
}

func getTransactions(context *gin.Context) { //It is used to transform data from Payment Array to json format.
	context.IndentedJSON(http.StatusOK, Payment) //Status of incoming request and second parameter is the data structure or array.
}

func main() {

	Router := gin.Default() //Created a server.

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:19006"} // Update with your client's origin
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	Router.Use(cors.New(config))

	Router.GET("/Payment", getTransactions)  //Endpoint and get request.
	Router.POST("/Payment", addTransactions) //Add a transaction in array.
	Router.Run("localhost:9000")             //To run the server on a particular path.
}
