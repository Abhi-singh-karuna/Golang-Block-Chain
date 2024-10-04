# docker build -t abstore-server .
# docker run -p 8080:8080 abstore-server


postman Collection :

# curl -X POST "http://localhost:8080/init?A=battery1&Aval=100&B=battery2&Bval=200"
# curl "http://localhost:8080/query?A=battery1"
