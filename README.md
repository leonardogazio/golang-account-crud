# HTTP REST API CRUD OPERATIONS EXAMPLE

RESTful API example written using GIN framework;
https://github.com/gin-gonic/gin

# Unit Testing:

Unit-tests in order to guarantee code consistence were written. All tests implement mocked results using https://github.com/zhashkevych/go-sqlxmock and native net/http/httptest package.

To run unit-tests, just type <b>make unit-test</b> command in terminal at the repository root folder.

--------------------------------------------------------------------------------------------------------

It's possible to checkout the endpoints definition and test them in Swagger;
http://localhost:8080/swagger/index.html

--------------------------------------------------------------------------------------------------------

# Usage:

1- Run make file commands in the following  order;<br />
<b>$ make postgres</b><br />
<b>$ make migrate_up</b><br />
<b>$ make seed</b><br />
<b>$ make setup-swagger</b><br />
<br />
2- Then finally run;<br />
<b>$ source .env.example</b><br />
<b>$ make run</b><br />

--------------------------------------------------------------------------------------------------------

TODO:<br />
  :heavy_check_mark: Operation Types API.<br />
  :white_check_mark: Account API.<br />
  :white_check_mark: Transactions API.<br />
