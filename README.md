# HTTP REST API CRUD OPERATIONS EXAMPLE

RESTful API example written using GIN framework;
github.com/gin-gonic/gin

# Unit Testing:

Unit-tests in order to guarantee code consistence were written. All tests implement mocked results using github.com/zhashkevych/go-sqlxmock and native net/http/httptest package.

To run unit-tests, just type <b>make unit-test</b> command in terminal at the repository root folder.

--------------------------------------------------------------------------------------------------------

It's possible to checkout the endpoints definition and test them in Swagger;
http://localhost:8080/swagger/index.html

--------------------------------------------------------------------------------------------------------

# Usage:

1- Run make file commands in the following  order;
<b>$ make postgres</b>
<b>$ make migrate_up</b>
<b>$ make seed</b>
<b>$ make setup-swagger</b>

Then finally run;
<b>$ source .env.example</b>
<b>$ make run</b>

--------------------------------------------------------------------------------------------------------

TODO:
  [x] Operation Types API.
  [] Account API.
  [] Transactions API.
