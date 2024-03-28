FROM python:3.8

WORKDIR /app

RUN apt install -y \
    default-libmysqlclient-dev \
    build-essential \
    pkg-config

COPY requirements.txt .

RUN pip install -r requirements.txt

# TODO : Clean dependencies for python mysql

COPY . .

EXPOSE 8000

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
