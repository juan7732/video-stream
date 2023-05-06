FROM python:3.7-slim

WORKDIR /app

RUN apt update
RUN apt add make automake gcc g++ subversion python3-dev

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

COPY . .

CMD ["python", "main.py"]