from config import getConfig

def getInterface():
    import network

    sta_if = network.WLAN(network.STA_IF)
    sta_if.active(True)
    
    return sta_if

def doConnect():
    wifi = getInterface()
    config = getConfig('wifi')

    if not wifi.isconnected():
        print('connecting to network...')

        wifi.connect(config['ssid'], config['password'])
        while not wifi.isconnected():
            pass
    print('network config:', wifi.ifconfig())

def listNetworks():
    wifi = getInterface()

    if wifi.isconnected():
        wifi.disconnect()
    print(wifi.scan())

# list_networks()
doConnect()
