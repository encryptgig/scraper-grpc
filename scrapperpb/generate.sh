#!/bin/bash

protoc scrapper.proto --go_out=plugins=grpc:.
