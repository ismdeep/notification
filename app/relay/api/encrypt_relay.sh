openssl aes-256-cbc -e -salt -pbkdf2 -iter 1024 -k "%v" | base64 | curl -fsSL -X PUT -H 'X-Token: %v' -d @- %v/api/v1/msg/%v