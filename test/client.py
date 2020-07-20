from __future__ import print_function
import logging

import grpc

import admin_pb2
import service_pb2_grpc


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = service_pb2_grpc.NodeStub(channel)
        response = stub.Auth(admin_pb2.AuthReq(cookie='you'))
        response = stub.AlwaysAuth(admin_pb2.AuthReq(cookie='you'))
    print("client received: ", admin_pb2.ErrorCode.Name(response.code))


if __name__ == '__main__':
    logging.basicConfig()
    run()
