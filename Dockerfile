FROM hdgigante/python-opencv:4.7.0-debian

WORKDIR /app

COPY requirements.txt requirements.txt
RUN #!bin/bash python -m pip install Flask

COPY . .

CMD [ "python", "app.py" ]
