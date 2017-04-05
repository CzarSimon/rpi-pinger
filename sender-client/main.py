import requests
import json
import schedule
import os
import time
import config


def get_ip():
    return os.environ["REOMOTE_ADDR"]


def send_ip(ip):
    data = json.dumps({"Addres": ip})
    try:
        requests.post(url=config.endpoint, data=data, header=config.header, timeout=config.timeout)
    except Exception as e:
        print e

def main():
    send_ip(get_ip())
    schedule.every().hour.do(send_ip, get_ip())
    while True:
        schedule.run_pending()
        time.sleep(config.sleep)


if __name__ == '__main__':
    main()
