import threading

def setInterval(func,time):
    e = threading.Event()
    count = 0
    while not e.wait(time):
        count+=1
        func(count)

def foo(count):
    print("Testando: {}".format(count))

# using
setInterval(foo,5)
