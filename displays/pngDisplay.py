class pngDisplay(object):
    config = None
    brightness = 255
    fileNumber = 0
    size = 0

    def __init__(self, size = 8):
        from config import getConfig

        self.config = getConfig("display")
        self.size = size

    def getEmptyFrame(self):
        x = self.config["width"]
        y = self.config["height"]

        return [[(0, 0, 0) for i in range(y)] for j in range(x)]

    def renderFrame(self, frame):
        import png

        imageWidth = self.config["width"] * self.size
        imageHeight = self.config["height"] * self.size

        p = [[0 for i in range(imageWidth * 3)] for j in range(imageHeight)]

        for frameX, row in enumerate(frame):
            for frameY, col in enumerate(row):
                for k in range(self.size):
                    for m in range(self.size):
                        for c in range(3):
                            imageX = (frameX * self.size + m) * 3 + c
                            imageY = frameY * self.size + k

                            p[imageY][imageX] = col[c]

        f = open('output/{}.png'.format(self.fileNumber), 'wb')
        w = png.Writer(imageWidth, imageHeight, greyscale=False)
        w.write(f, p)
        f.close()

        self.fileNumber += 1

    def setBrightness(self, newBrightness):
        self.brightness = newBrightness