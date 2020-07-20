from __future__ import print_function
import logging
import time

import grpc

import service_pb2
import service_pb2_grpc

serial = 0


def generate_messages():
    messages = [
        service_pb2.Header(serial=1),
        service_pb2.Test(serial=1),
        service_pb2.Header(serial=2),
        service_pb2.Test(serial=2),
        service_pb2.Header(serial=3),
        service_pb2.Test(serial=3),
    ]
    for msg in messages:
        yield msg


def normalize(req):
    pass


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('localhost:50051',
                               options=[('grpc.keepalive_time_ms', 60000)]) as channel:
        stub = service_pb2_grpc.NodeStub(channel)
        response = stub.Auth(service_pb2.AuthReq(cookie='you'))
        print("client received: ", service_pb2.ErrorCode.Name(response.code))
        print(stub.Ping(service_pb2.Heartbeat(timestamp=time.time_ns())))
        for response in stub.Flow(generate_messages()):
            print(response)
        print(stub.Ping(service_pb2.Heartbeat(timestamp=time.time_ns())))
        time.sleep(60)


if __name__ == '__main__':
    logging.basicConfig()
    run()
