from connect import doConnect
import ledDisplay
from layers.text import TextDrawer

display = ledDisplay()

frame = display.getEmptyFrame()

bigFont = TextDrawer('big')

bigFont.drawString('123', frame, (255,255,255))

print(frame)

display.renderFrame(frame)

# doConnect()