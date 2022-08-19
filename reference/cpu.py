#!/bin/python3

from time import sleep

def cpuUsage() -> None:
    lastSum = 0
    LCI = 0 # last cpu idle %, needed for idle deltas
    with open("/proc/stat", "r") as f:
        while True:
            head = f.read(100) # only grab a few bytes
            f.seek(0)          # reset to beginning

            n1 = "" # grab up to the first newline
            for c in head:
                if c != "\n": 
                    n1 += c
                else: 
                    break

            # slice off 'cpu:' string and empty split
            nums = [int(i) for i in n1.split(" ")[1:] if len(i) > 0]
            cpuSum = sum(nums)

            delta = cpuSum - lastSum
            idle = nums[3] - LCI
            LCI = nums[3]

            lastSum = cpuSum

            used = delta - idle

            pc = 100.0 * (float(used) / float(delta))

            print(f" cpu: {pc:.2f}%")
            sleep(1)

try:             
    cpuUsage()
except KeyboardInterrupt:
    print()