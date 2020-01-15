# Back End Development Technical Challenge

## Notes

### Implementation

I have fully implemented the Pets ("/pet") endpoint per the specification, with the following exceptions/comments:

* No actual upload of images is processed; it simply updates the image data for the pet
* The image upload specifies an optional metadata field, but no storage exists for this information in the spec
* since the Pet struct contains full Tag and Category data (rather than just their IDs) using the "/pet" POST and PUT functions to add/update a Pet entry with arbitrary JSON data would allow inconsistent Tag and Category entries to be created within the Pet. I have not added code to fix this, since I figure in a real world situation the data would be structured differently anyway
* I made usage of the api-key ("special-key", per the spec page) mandatory for all interactions with the API. The spec only seems to require it for the DELETE method, but it seemed to me that any writing to the database should be protected. Possibly that requirement could be relaxed for GET methods?

I have included an additional "/reset4test" url which allows the test suite to reset the dummy database to a known state before each run of the tests

That said, I have created a dummydb package which models all the data; in a real world situation this would, of course, be linked to an actual database, but for this task it simply initalises in memory with a few sample data points.

Finally, I used the external golang.org/x/crypto/bcrypt library for password hashing. If you do not already have it on your system, "go get golang.org/x/crypto/bcrypt" should install it.

### Testing

I implemented a main_test package which tests the API itself, via calls to cURL, and examines the results. The curlEXE const points to the location of the executable; obviously I developed this on a Windows machine.

Perhaps it goes without saying, but since it is a server, the program must be compiled and run *before* the test suite is run.

There are no unit tests. Yet.

## Challenge Specification

### Email

Hi Peter

Thank you for your application for the role of Back End Developer with Jumbo Interactive,

Having reviewed your CV, we would like to progress your application in the recruitment process for this role. Recruitment is a two-step process, the first being an at-home Technical Challenge which may be followed by an interview.

Details of the technical challenge can be found here and in the attached document. We expect the challenge to take no more than 8 hours, therefore can you please submit your completed challenge within 7 days by 10am. If you are unable to submit the challenge by this deadline please let me know.

Please submit your technical challenge via <https://jumbointeractive.typeform.com/to/TtVa2R>

Below is a word document which further outlines the requirement of the technical challenge.

If you have any questions regarding the role or the recruitment process, feel free to contact me.

Regards,
Jumbo Interactive HR Team

### Word Doc

As part of Jumbos Programmer candidate process we would like you to complete a short development task.

The aim is to give you an opportunity to show us the way you solve problems and the development pattern you use.

There is no right or wrong way to approach this, there is only your way.

We would like you to create an API for a Pet store using a Microservices architecture style.

The API requirements have been set and can be found at <https://petstore.swagger.io.>

You should allow yourself 4 hours to complete your work; it is not expected that you will provide a full solution.

We will be looking at:

* One endpoint fully operational
* Your approach to Microservices architecture
* Models used
* Code style
* How testable your solution is

Once finished please publish your results to an online repository hosting service (GitHub for example) and please submit your technical challenge via <https://jumbointeractive.typeform.com/to/TtVa2R>
