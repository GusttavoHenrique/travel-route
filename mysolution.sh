echo "---- Build and run docker ---"
FILE_PATH="$1"
FILE_DIR=$(basename "$FILE_PATH")
FILE_NAME=$(basename "$FILE_PATH")

docker build --build-arg FILE_NAME="${FILE_NAME}" -t app .
mkdir -p /tmp/travel-route
cp -p "${FILE_DIR}" /tmp/travel-route
gedit /tmp/travel-route/"${FILE_NAME}" &
docker run -p 8080:8080 -v /tmp/travel-route/"${FILE_NAME}":/tmp/"${FILE_NAME}" -it app:latest /bin/bash ls -la
