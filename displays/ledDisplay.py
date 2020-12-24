class ledDisplay(object):
    np = None
    config = None
    brightness = 255
    pixels = 0

    def __init__(self):
        import machine, neopixel
        from config import getConfig

        self.config = getConfig("display")

        self.pixels = self.config["width"] * self.config["height"]

        self.np = neopixel.NeoPixel(machine.Pin(self.config["pin"]), self.pixels)

    def __coordinatesToIndex(self, x, y):
        if x % 2 == 0:
            return x * self.config["height"] + y
        return (x + 1) * self.config["height"] - y - 1

    def getEmptyFrame(self):
        x = self.config["width"]
        y = self.config["height"]

        return [[(0, 0, 0) for i in range(y)] for j in range(x)]

    def renderFrame(self, frame):
        for i in range(self.config["width"]):
            for j in range(self.config["height"]):
                index = self.__coordinatesToIndex(i, j)

                self.np[index] = frame[i][j]
        self.np.write()

    def setBrightness(self, newBrightness):
        self.brightness = newBrightness