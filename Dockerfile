FROM python:3.7-slim

WORKDIR /app

RUN apt update
RUN apt add make automake gcc g++ subversion python3-dev

COPY requirements.txt requirements.txt
RUN python3 pip install -r requirements.txt

COPY . .

CMD ["python3", "main.py"]