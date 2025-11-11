FROM python:3.12-slim

WORKDIR /app

COPY . /app/

COPY .env /app/

RUN pip install -r requirements.txt

ENTRYPOINT [ "python3", "main.py" ]