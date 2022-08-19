from time import sleep


'''
total jiffies is fairly low
but individual proc jifs is fairly low

what exactly is going on with the proc%?
'''

sl = [0, 0, 0, 0]
prev = []
LCI = 0 # last cpu idle %, needed for idle deltas
lastSum = 0
cpuLast = 0
while True:
    with open("/proc/stat", "r") as f:
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
        cpuLast = 100.0 * (float(used) / float(delta))

    with open('/proc/%s/stat', 'r') as f:
        s = f.read(200)
        subs = ""
        for c in s:
            if c == "\n":
                break
            else:
                subs += c
        s = subs.split()
        s = s[13:15]
        s[0] = int(s[0])
        s[1] = int(s[1])
        print(s[0] - sl[0], s[1] - sl[1], end=' ')
        sl[2] = s[0] - sl[0]
        sl[3] = s[1] - sl[1]
        sl[0] = s[0]
        sl[1] = s[1]
    
    with open('/proc/stat', 'r') as f:
        head = f.read(100).split('\n')[0]
        wds = head.split()
        wds = wds[1:]

        ls = [int(i) for i in wds if len(i) > 0]
        ls = [n for i,n in enumerate(ls) if i != 3]

        if len(prev) == 0:
            prev = ls

        else:
            for i,v in enumerate(ls):
                v -= prev[i]
                prev[i] = ls[i]
                ls[i] = v

            sstr = sum(ls)
        
            ttl = sum(sl[2:])
            pc = ttl / sstr
            print(" "*6, f"{pc * cpuLast:.2f}%  -> top percent: {pc * cpuPC * 8:.2f}%")

    
    sleep(0.40)

'''
[pid]/stat
1 pid
2 comm
3 stt
4 ppid
5 pgrp
6 session
7 tty_nr
8 tpgid
9 flags
...

14 utime
15 stime
10984 (code) S 2631 2631 2631 0 -1 1077936128 2490242 0 0 0 34378 2338 0 0 20 0 20 0 305665 41215471


/proc/stat
1. user
2. nice
3. system
4. idle
5. iowait
6. irq
7. softirq
8. steal
9. guest
10. gnice
'''