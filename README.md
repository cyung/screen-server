##Screen Server

A server for hosting images and creating URL aliases.

##API Documentations


###Adding a user
* URL: `/user`
* Method: `POST`

###Adding a picture
* URL: `/img`
* Method: `POST`
* Parameters:
  * `user` username string
  * `authToken` authorization token
  * `img` image
* Success response:
  * `path` URL for image uploaded and hosted on server
* Error response:
  * `error` error response

###Deleting a picture
* URL: `/img`
* Method: `DELETE`
* Parameters:
  * `user` username string
  * `authToken` authorization token
  * `imgHash` hash of image