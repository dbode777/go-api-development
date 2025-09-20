# go-api-development

This is a repository that contains examples used in learning Go for a CourseCareers course. 

The GO-API-TUTORIAL is more of a mini-project that serves as a rudimentary library/bookstore system. There are a few components I added to get a bit more familiar with the language. 
In particular, I added a validator and error handler so there was a dedicated component for each of those cases. I added manual test files for regression testing. Still need to add
unit tests and integration tests.

To run the project locally, use _go run main.go_ in the terminal to run the program, and then in another terminal use _curl localhost:8080/{endpoint} --request "{request_type}"_ to make requests. 
If the endpoint requires input JSON, such as for updating a book or creating a book, use 
  _curl -d @testData/{filename}.json localhost:8080/{endpoint} --request "{request_type}"_ 
to use the corresponding file as input. 
