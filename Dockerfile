FROM python:3.8-slim

WORKDIR /app

RUN python3 -m pip install --upgrade pip
RUN apt update
RUN apt-get install build-essential libopencv-dev -y

COPY requirements.txt requirements.txt
RUN python3 -m pip install -r requirements.txt

COPY . .

CMD ["python3", "main.py"]