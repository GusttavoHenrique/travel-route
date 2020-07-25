echo "---- Build and run docker ---"
FILE_PATH="$1"
FILE_NAME=$(basename "$FILE_PATH")

docker build --build-arg FILE_NAME="${FILE_NAME}" -t app .
docker run -p 8080:8080 -v "${FILE_PATH}":/app/data/"${FILE_NAME}" -it app:latest /bin/bash ls -la
