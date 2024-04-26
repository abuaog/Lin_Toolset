from .db_handler import *
from influxdb_client import InfluxDBClient
from influxdb_client.client.write_api import ASYNCHRONOUS
# from colorlog import ColoredFormatter
import colorlog
import logging


log_format = "%(asctime)s: %(message)s"
log_format2 = "%(asctime)s — %(levelname)s — %(name)s — %(message)s"


def init(config):
    # formatter = colorlog.ColoredFormatter(
    #     "%(log_color)s%(levelname)-8s%(reset)s %(blue)s%(message)s",
    #     datefmt=None,
    #     reset=True,
    #     log_colors={
    #         'DEBUG': 'cyan',
    #         'INFO': 'green',
    #         'WARNING': 'yellow',
    #         'ERROR': 'red',
    #         'CRITICAL': 'red,bg_white',
    #     },
    #     secondary_log_colors={},
    #     style='%'
    # )
    # colorlog.basicConfig(format=formatter)
    logging.basicConfig(format=log_format2, level=logging.DEBUG)
    client = InfluxDBClient(url=config.influx_server, token=config.influx_token, org=config.influx_org, debug=False)
    write_api = client.write_api(write_options=ASYNCHRONOUS)
    return Log(write_api, config)


class Log:
    def __init__(self, write_api, config):
        self.write_api = write_api
        self.config = config

    def get_scope(self, scope_name):
        logger = logging.getLogger(scope_name)
        logger.setLevel(logging.DEBUG)
        logger.addHandler(self._get_db_handler())
        return logger

    def _get_db_handler(self):
        handler = LogDBHandler(self.write_api, self.config)
        return handler



