sample-golang-ports-adaptors-project
Controller layer does input validation
controller layer calls app layer or facade layer
there is a middleware layer too. for example : there is request id generation for request or there is login or authentication system.
app init layer
logger
ports are interfaces responsible for talking to other layers or external services .
incoming port : whatever functionalities that are being exposed . outgoing : wherever we are talking outside.
Domain will have core , application and infrastructure layer .
core - model and ports
