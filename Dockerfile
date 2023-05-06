FROM hdgigante/python-opencv:4.7.0-debian

WORKDIR /app

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

COPY . .

CMD [ "python", "app.py" ]
