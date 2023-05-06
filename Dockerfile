FROM python:3.8-slim

WORKDIR /app

RUN apt update
RUN apt-get install build-essential -y

COPY requirements.txt requirements.txt
RUN python3 -m pip install -r requirements.txt

COPY . .

CMD ["python3", "main.py"]