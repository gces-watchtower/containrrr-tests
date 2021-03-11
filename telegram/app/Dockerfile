FROM python:3.6-slim

RUN mkdir /usr/src/app

WORKDIR /usr/src/app

COPY ./requirements.txt /usr/src/app/requirements.txt

RUN pip3 install -r requirements.txt

COPY . /usr/src/app

ENTRYPOINT ["python3"]

CMD ["main.py"]
