Feature: Greeting

  @story-1
  Rule: A named request receives a personalized greeting

    Scenario: API consumer requests a greeting by name
      Given the greeter service is running
      When an API consumer sends a GET request to "/hello" with name "Ada"
      Then the response is JSON containing a greeting addressed to "Ada"

  @story-2
  Rule: A request with no name still receives a greeting

    Scenario: API consumer omits the name parameter
      Given the greeter service is running
      When an API consumer sends a GET request to "/hello" with no name parameter
      Then the response is JSON containing a generic greeting addressed to "World"
