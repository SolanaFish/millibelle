class TextDrawer:
    font = None

    def __init__(self, fontName):
        import json

        fontPath = 'fonts/{}.json'.format(fontName)

        file = open(fontPath, 'r')
        self.font = json.load(file)
        file.close()

    def drawString(self, text, frame, color, startX = 0, startY = 0):
        textHeight = self.font['height']
        textWidth = 0
        for sign in text:
            textWidth += self.font[sign]['width'] + 1

        frameWidth = len(frame)
        frameHeight = len(frame[0])

        if textWidth > startX + frameWidth:
            textWidth = startX + frameWidth

        if textHeight > startY + frameHeight:
            textHeight = startY + frameHeight

        x = startX

        for sign in text:
            signWidth = self.font[sign]['width']
            signPixels = self.font[sign]['pixels']

            if signWidth > textWidth - x:
                signWidth = textWidth - x;

            print("asdf", textHeight, textWidth, signWidth, x)

            for i in range(textHeight):
                for j in range(signWidth):
                    if signPixels[i][j] == 1:
                        frame[x + j][startY + i] = color

            x += signWidth + 1

        return frame

