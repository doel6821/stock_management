## Answer
1.  I think, in this case the inbound and outbound of goods is not recorded properly. Recording inbound and outbound of goods is very important, if not handled properly it will result in the actual stock of goods with stock in the system being inaccurate.
2. make records inbound and outbound of goods, and update stock every time there are inbound or outbound. 


## Concept
1. Each product is owned by a user, and only that user is allowed to change product data and make purchases to add stock
2. when the ordered goods have arrived, the user can receive the goods according to the order quantity and the system will automatically update the product stock
3. when the buyer makes a purchase, the stock will decrease automatically, if the stock is less than the purchase quantity, the purchase cannot be made.
4. All transactions will be recorded in the transaction history

## How to Run App
1. please update the environment according to the environment on your local computer
2. run queryTable.sql in your local database (postgreSQL)
3. run on your terminal "go run main.go"
4. you can open swagger documentation on web browser using url "http://localhost:8080/swagger/index.html" (dafault Port:8080)

