FROM python:3.7-slim

WORKDIR /app

RUN apt update


COPY requirements.txt requirements.txt
RUN python3 -m pip install -r requirements.txt

COPY . .

CMD ["python3", "main.py"]