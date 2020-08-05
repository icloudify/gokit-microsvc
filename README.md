# gokit-microsvc
### Encrypt
curl -XPOST -d'{"key":"111023043350789514532147", "text": "I am A Message"}' http://localhost:8080/encrypt
###Decrypt
curl -XPOST -d'{"key":"111023043350789514532147", "message": "8/+JCfTb+ibIjzQtmCo="}' http://localhost:8080/decrypt