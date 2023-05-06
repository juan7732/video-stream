FROM python:3.7-slim

WORKDIR /app

RUN apk update
RUN apk add make automake gcc g++ subversion python3-dev

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

COPY . .

CMD ["python", "main.py"]